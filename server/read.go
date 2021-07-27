package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/jhump/protoreflect/desc"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	skafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber/backends/kafka"
	"github.com/batchcorp/plumber/serializers"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/common"
	"github.com/batchcorp/plumber-schemas/build/go/protos/conns"
	"github.com/batchcorp/plumber-schemas/build/go/protos/encoding"
	"github.com/batchcorp/plumber-schemas/build/go/protos/records"
)

type Read struct {
	PlumberID     string
	ID            string
	ContextCxl    context.Context
	CancelFunc    context.CancelFunc
	Backend       *kafka.KafkaReader // TODO: have to genercize once backend refactor is done
	ConnectionID  string
	MessageCh     chan *records.Message // TODO: have to genercize this somehow
	DecodeOptions *encoding.Options
	MsgDesc       *desc.MessageDescriptor
	log           *logrus.Entry
}

func (p *PlumberServer) StartRead(req *protos.StartReadRequest, srv protos.PlumberServer_StartReadServer) error {
	if err := p.validateRequest(req.Auth); err != nil {
		return CustomError(common.Code_UNAUTHENTICATED, fmt.Sprintf("invalid auth: %s", err))
	}

	requestID := uuid.NewV4().String()

	// Reader needs a unique ID that frontend can reference
	readerID := uuid.NewV4().String()

	md, err := generateMD(req.DecodeOptions)
	if err != nil {
		return err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	// TODO: figure out what we want to do with background reader? should this only be for writing to disk?
	// For now, just close reader when client disconnects from this endpoint
	defer cancelFunc()

	backend, err := p.getBackendRead(req)
	if err != nil {
		return CustomError(common.Code_ABORTED, err.Error())
	}

	// Launch reader and record
	reader := &Read{
		PlumberID:     p.PersistentConfig.PlumberID,
		ID:            readerID,
		ContextCxl:    ctx,
		CancelFunc:    cancelFunc,
		Backend:       backend,
		ConnectionID:  req.ConnectionId,
		MessageCh:     make(chan *records.Message, 100),
		DecodeOptions: req.DecodeOptions,
		MsgDesc:       md,
		log:           p.Log.WithField("read_id", readerID),
	}

	p.setRead(readerID, reader)

	go reader.Read()

	for {
		select {
		case msg := <-reader.MessageCh:
			messages := make([]*records.Message, 0)

			// TODO: batch these up and send multiple per response?
			messages = append(messages, msg)

			srv.Send(&protos.StartReadResponse{
				Status: &common.Status{
					Code:      common.Code_OK,
					Message:   "Message read",
					RequestId: requestID,
				},
				ReadId:   readerID,
				Messages: messages,
			})
		}
	}

	// TODO: what should we do when connection to the endpoint is ended (UI is closed)?
	// TODO: should we continue reading and filling up a channel or should we pause the read?

	return nil
}

// generateKafkaPayload generates a records.Message protobuf struct from a kafka message struct
func (r *Read) generateKafkaPayload(msg *skafka.Message) (*records.Message, error) {
	var err error
	var data []byte

	switch r.DecodeOptions.Type {
	case encoding.Type_PROTOBUF:
		data, err = DecodeProtobuf(r.MsgDesc, msg.Value)
		if err != nil {
			return nil, errors.Wrap(err, "unable to decode protobuf")
		}
	case encoding.Type_AVRO:
		data, err = serializers.AvroDecode(r.DecodeOptions.GetAvro().Schema, msg.Value)
		if err != nil {
			return nil, errors.Wrap(err, "unable to decode AVRO message")
		}
		fallthrough
	case encoding.Type_JSON_SCHEMA:
		// TODO
		fallthrough
	default:
		data = msg.Value
	}

	return &records.Message{
		MessageId:        uuid.NewV4().String(),
		PlumberId:        r.PlumberID,
		UnixTimestampUtc: time.Now().UTC().UnixNano(),
		Message: &records.Message_Kafka{Kafka: &records.Kafka{
			Topic:     msg.Topic,
			Key:       msg.Key,
			Value:     data,
			Blob:      msg.Value,
			Timestamp: msg.Time.UTC().UnixNano(),
			Offset:    msg.Offset,
			Partition: int32(msg.Partition),
			Headers:   convertKafkaHeadersToProto(msg.Headers),
		}},
	}, nil
}

// convertKafkaHeadersToProto converts type of header slice from segmentio's to our protobuf type
func convertKafkaHeadersToProto(original []skafka.Header) []*records.KafkaHeader {
	converted := make([]*records.KafkaHeader, 0)

	for _, o := range original {
		converted = append(converted, &records.KafkaHeader{
			Key:   o.Key,
			Value: string(o.Value),
		})
	}

	return converted
}

// Read is a goroutine that is launched when a read is started. It will continue running until plumber exiits
// or a read is stopped via the API
func (r *Read) Read() {
	defer r.Backend.Reader.Close()

	for {
		select {
		case <-r.ContextCxl.Done():
			r.log.Info("Read stopped")
			return
		default:
			// noop
		}

		var err error

		msg, err := r.Backend.Reader.ReadMessage(r.ContextCxl)
		if err != nil && err != context.Canceled {
			r.log.Errorf("unable to read kafka message: %s", err)
			continue
		}

		payload, err := r.generateKafkaPayload(&msg)
		if err != nil {
			r.log.Errorf("unable to generate kafka payload: %s", err)
		}

		r.MessageCh <- payload
	}

}

func getKafkaAuthConfig(cfg *conns.Kafka) (sasl.Mechanism, error) {
	switch cfg.SaslType {
	case conns.SASLType_SCRAM:
		return scram.Mechanism(scram.SHA512, cfg.SaslUsername, cfg.SaslPassword)
	case conns.SASLType_PLAIN:
		return plain.Mechanism{
			Username: cfg.SaslUsername,
			Password: cfg.SaslPassword,
		}, nil
	}

	return nil, nil
}

// getBackendRead gets the backend message bus needed to read/write from
// TODO: genericize after backend refactor
func (p *PlumberServer) getBackendRead(req *protos.StartReadRequest) (*kafka.KafkaReader, error) {
	args := req.GetKafka()

	connCfg := p.getConn(req.ConnectionId)
	if connCfg == nil {
		return nil, errors.New("connection does not exist")
	}

	switch {
	case req.GetKafka() != nil:
		return p.getBackendReadKafka(connCfg, args)
	}

	return nil, errors.New("unknown message bus")
}

func (p *PlumberServer) getBackendReadKafka(connCfg *protos.Connection, args *args.Kafka) (*kafka.KafkaReader, error) {
	kafkaCfg := connCfg.GetKafka()

	dialer := &skafka.Dialer{
		DualStack: true,
		Timeout:   time.Second * 10,
	}

	if kafkaCfg.InsecureTls {
		dialer.TLS = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	auth, err := getKafkaAuthConfig(kafkaCfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not get authentication mechanism")
	}

	dialer.SASLMechanism = auth

	// Attempt to establish connection on startup
	ctxDeadline, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))

	kafkaConn, err := dialer.DialContext(ctxDeadline, "tcp", kafkaCfg.Address[0])
	if err != nil {
		logrus.Errorf("unable to create initial connection to broker '%s', trying next broker", kafkaCfg.Address[0])
	}
	if err != nil {
		return nil, err
	}

	commitInterval, err := time.ParseDuration(fmt.Sprintf("%ds", args.CommitIntervalSeconds))
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse CommitIntervalSeconds")
	}

	maxWait, err := time.ParseDuration(fmt.Sprintf("%ds", args.MaxWaitSeconds))
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse MaxWaitSeconds")
	}

	rebalanceTimeout, err := time.ParseDuration(fmt.Sprintf("%ds", args.RebalanceTimeoutSeconds))
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse RebalanceTimeoutSeconds")
	}

	rc := skafka.ReaderConfig{
		Brokers:          kafkaCfg.Address,
		CommitInterval:   commitInterval,
		Dialer:           dialer,
		MaxWait:          maxWait,
		MinBytes:         int(args.MinBytes),
		MaxBytes:         int(args.MaxBytes),
		QueueCapacity:    100, // TODO: add to protos?
		RebalanceTimeout: rebalanceTimeout,
	}

	if args.UseConsumerGroup {
		rc.GroupTopics = args.Topics
		rc.GroupID = args.ConsumerGroupName
	} else {
		rc.Topic = args.Topics[0]
	}

	r := skafka.NewReader(rc)

	if !args.UseConsumerGroup {
		if err := r.SetOffset(args.ReadOffset); err != nil {
			return nil, errors.Wrap(err, "unable to set read offset")
		}
	}

	return &kafka.KafkaReader{
		Reader: r,
		Conn:   kafkaConn,
	}, nil
}

func (p *PlumberServer) StopRead(_ context.Context, req *protos.StopReadRequest) (*protos.StopReadResponse, error) {
	if err := p.validateRequest(req.Auth); err != nil {
		return nil, CustomError(common.Code_UNAUTHENTICATED, fmt.Sprintf("invalid auth: %s", err))
	}

	requestID := uuid.NewV4().String()

	// Get reader and cancel
	read := p.getRead(req.ReadId)
	if read == nil {
		return nil, CustomError(common.Code_NOT_FOUND, "read does not exist or has already been stopped")
	}

	read.CancelFunc()

	p.Log.WithField("request_id", requestID).Infof("Read '%s' stopped", req.ReadId)

	return &protos.StopReadResponse{
		Status: &common.Status{
			Code:      common.Code_OK,
			Message:   "Message read",
			RequestId: requestID,
		},
	}, nil
}
