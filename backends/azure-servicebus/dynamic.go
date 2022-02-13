package azure_servicebus

import (
	"context"

	serviceBus "github.com/Azure/azure-service-bus-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber/dynamic"
	"github.com/batchcorp/plumber/validate"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber-schemas/build/go/protos/records"
)

func (a *AzureServiceBus) Dynamic(ctx context.Context, dynamicOpts *opts.DynamicOptions, dynamicSvc dynamic.IDynamic, errorCh chan<- *records.ErrorRecord) error {
	if err := validateDynamicOptions(dynamicOpts); err != nil {
		return errors.Wrap(err, "invalid dynamic options")
	}

	llog := logrus.WithField("pkg", "azure/dynamic")

	var queue *serviceBus.Queue
	var topic *serviceBus.Topic
	var err error

	if dynamicOpts.AzureServiceBus.Args.Queue != "" {
		queue, err = a.client.NewQueue(dynamicOpts.AzureServiceBus.Args.Queue)
		if err != nil {
			return errors.Wrap(err, "unable to create new azure service bus queue client")
		}

		defer queue.Close(ctx)
	} else {
		topic, err = a.client.NewTopic(dynamicOpts.AzureServiceBus.Args.Topic)
		if err != nil {
			return errors.Wrap(err, "unable to create new azure service bus topic client")
		}

		defer topic.Close(ctx)
	}

	if err := dynamicSvc.Start(ctx, "Azure Service Bus"); err != nil {
		return errors.Wrap(err, "unable to create dynamic")
	}

	outboundCh := dynamicSvc.Read()

	// Continually loop looking for messages on the channel.
	for {
		select {
		case outbound := <-outboundCh:
			msg := serviceBus.NewMessage(outbound.Blob)

			if queue != nil {
				// Publishing to queue
				if err := queue.Send(ctx, msg); err != nil {
					llog.Errorf("Unable to replay message: %s", err)
					break
				}
			} else {
				// Publishing to topic
				if err := topic.Send(ctx, msg); err != nil {
					llog.Errorf("Unable to replay message: %s", err)
					break
				}
			}

			llog.Debugf("Replayed message to Azure Service Bus for replay '%s'", outbound.ReplayId)
		case <-ctx.Done():
			llog.Debug("context cancelled")
			return nil
		}
	}

	return nil
}

func validateDynamicOptions(dynamicOpts *opts.DynamicOptions) error {
	if dynamicOpts == nil {
		return validate.ErrEmptyDynamicOpts
	}

	if dynamicOpts.AzureServiceBus == nil {
		return validate.ErrEmptyBackendGroup
	}

	args := dynamicOpts.AzureServiceBus.Args
	if dynamicOpts.AzureServiceBus.Args == nil {
		return validate.ErrEmptyBackendArgs
	}

	if args.Queue == "" && args.Topic == "" {
		return ErrQueueOrTopic
	}

	if args.Queue != "" && args.Topic != "" {
		return ErrQueueAndTopic
	}

	return nil
}
