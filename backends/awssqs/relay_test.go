package awssqs

import (
	"context"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber-schemas/build/go/protos/records"

	"github.com/streamdal/plumber/backends/awssqs/sqsfakes"
	"github.com/streamdal/plumber/validate"
)

var _ = Describe("AWSSQS Backend", func() {
	var a *AWSSQS
	var relayOpts *opts.RelayOptions

	BeforeEach(func() {
		a = &AWSSQS{
			connArgs: &args.AWSSQSConn{
				AwsRegion:          "us-west-2",
				AwsSecretAccessKey: "testing",
			},
			client: &sqsfakes.FakeSQSAPI{},
			log:    logrus.NewEntry(&logrus.Logger{Out: ioutil.Discard}),
		}

		relayOpts = &opts.RelayOptions{
			AwsSqs: &opts.RelayGroupAWSSQSOptions{
				Args: &args.AWSSQSRelayArgs{
					QueueName:               "test",
					RemoteAccountId:         "test",
					MaxNumMessages:          1,
					ReceiveRequestAttemptId: "",
					AutoDelete:              true,
					WaitTimeSeconds:         1,
				},
			},
		}
	})

	Context("validateRelayOptions", func() {
		It("validates nil relay options", func() {
			err := validateRelayOptions(nil)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(validate.ErrEmptyRelayOpts))
		})
		It("validates missing backend group", func() {
			relayOpts.AwsSqs = nil
			err := validateRelayOptions(relayOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(validate.ErrEmptyBackendGroup))
		})
		It("validates missing backend args", func() {
			relayOpts.AwsSqs.Args = nil
			err := validateRelayOptions(relayOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(validate.ErrEmptyBackendArgs))
		})
		It("validates empty topic", func() {
			relayOpts.AwsSqs.Args.QueueName = ""
			err := validateRelayOptions(relayOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrMissingQueue))
		})

		It("validates wait time seconds", func() {
			relayOpts.AwsSqs.Args.WaitTimeSeconds = -1
			err := validateRelayOptions(relayOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrInvalidWaitTime))
		})

		It("validates max num messages", func() {
			relayOpts.AwsSqs.Args.MaxNumMessages = 0
			err := validateRelayOptions(relayOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrInvalidMaxNumMessages))
		})
	})

	Context("Relay", func() {
		It("validates relay options", func() {
			err := a.Relay(context.Background(), nil, nil, nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(validate.ErrEmptyRelayOpts.Error()))
		})

		It("relays a message", func() {
			fakeSQS := &sqsfakes.FakeSQSAPI{}
			fakeSQS.GetQueueUrlStub = func(*sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
				return &sqs.GetQueueUrlOutput{}, nil
			}

			ctx, cancel := context.WithCancel(context.Background())

			fakeSQS.ReceiveMessageStub = func(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
				// Cancel relay after first message
				defer cancel()

				return &sqs.ReceiveMessageOutput{
					Messages: []*sqs.Message{
						{
							MessageId: aws.String("test"),
							Body:      aws.String("test"),
						},
					},
				}, nil
			}

			a.client = fakeSQS

			relayCh := make(chan interface{}, 1)
			errorCh := make(chan *records.ErrorRecord, 1)

			err := a.Relay(ctx, relayOpts, relayCh, errorCh)

			Expect(err).ToNot(HaveOccurred())
			Expect(fakeSQS.ReceiveMessageCallCount()).To(Equal(1))
			Expect(errorCh).ToNot(Receive())
			Expect(relayCh).To(Receive())
		})
	})
})
