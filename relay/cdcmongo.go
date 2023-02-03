package relay

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/batchcorp/collector-schemas/build/go/protos/records"
	"github.com/batchcorp/collector-schemas/build/go/protos/services"

	"github.com/batchcorp/plumber/backends/cdcmongo/types"
)

// handleCDCMongo sends a mongo relay message to the GRPC server
func (r *Relay) handleCDCMongo(ctx context.Context, conn *grpc.ClientConn, messages []interface{}) error {
	sinkRecords, err := r.convertMessagesToMongoRecords(messages)
	if err != nil {
		return fmt.Errorf("unable to convert messages to generic records: %s", err)
	}

	client := services.NewGRPCCollectorClient(conn)

	return r.CallWithRetry(ctx, "AddRecord", func(ctx context.Context) error {
		_, err := client.AddRecord(ctx, &services.GenericRecordRequest{
			Records: sinkRecords,
		}, grpc.MaxCallSendMsgSize(MaxGRPCMessageSize))
		return err
	})
}

// validateCDCMongoMessage ensures all necessary values are present for a mongo message
func (r *Relay) validateCDCMongoMessage(msg *types.RelayMessage) error {
	if msg == nil {
		return ErrMissingMessage
	}

	if msg.Value == nil {
		return ErrMissingMessageValue
	}

	return nil
}

// convertMessagesToMongoRecords creates a records.GenericRecord from a mongo change stream message which can then
// be sent to the GRPC server
func (r *Relay) convertMessagesToMongoRecords(messages []interface{}) ([]*records.GenericRecord, error) {
	sinkRecords := make([]*records.GenericRecord, 0)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	for i, v := range messages {
		relayMessage, ok := v.(*types.RelayMessage)
		if !ok {
			return nil, fmt.Errorf("unable to type assert incoming message as RelayMessage (index: %d)", i)
		}

		if err := r.validateCDCMongoMessage(relayMessage); err != nil {
			return nil, fmt.Errorf("unable to validate mongo relay message (index: %d): %s", i, err)
		}

		sinkRecords = append(sinkRecords, &records.GenericRecord{
			Source:          hostname,
			Body:            []byte(relayMessage.Value.String()),
			Timestamp:       time.Now().UTC().UnixNano(),
			ForceDeadLetter: r.DeadLetter,
		})
	}

	return sinkRecords, nil
}
