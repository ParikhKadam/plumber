package awssqs

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber-schemas/build/go/protos/records"

	"github.com/batchcorp/plumber/dynamic"
	"github.com/batchcorp/plumber/validate"
)

func (a *AWSSQS) Dynamic(ctx context.Context, opts *opts.DynamicOptions, dynamicSvc dynamic.IDynamic, errorCh chan<- *records.ErrorRecord) error {
	if err := validateDynamicOptions(opts); err != nil {
		return errors.Wrap(err, "unable to validate dynamic options")
	}

	llog := a.log.WithField("pkg", "awssqs/dynamic")

	args := opts.AwsSqs.Args

	queueURL, err := a.getQueueURL(args.QueueName, args.RemoteAccountId)
	if err != nil {
		return errors.Wrap(err, "unable to get queue url")
	}

	if err := dynamicSvc.Start(ctx, "AWS SQS"); err != nil {
		return errors.Wrap(err, "unable to create dynamic")
	}

	outboundCh := dynamicSvc.Read()

	for {
		select {
		case outbound := <-outboundCh:
			// write
			if err := a.writeMsg(args, string(outbound.Blob), queueURL); err != nil {
				err = fmt.Errorf("unable to replay message: %s", err)
				llog.Error(err)
				return err
			}
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

	if dynamicOpts.AwsSqs == nil {
		return validate.ErrEmptyBackendGroup
	}

	if dynamicOpts.AwsSqs.Args == nil {
		return validate.ErrEmptyBackendArgs
	}

	if dynamicOpts.AwsSqs.Args.QueueName == "" {
		return ErrMissingQueue
	}

	return nil
}
