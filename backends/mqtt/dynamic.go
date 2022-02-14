package mqtt

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber-schemas/build/go/protos/records"
	"github.com/batchcorp/plumber/dynamic"
	"github.com/batchcorp/plumber/util"
	"github.com/batchcorp/plumber/validate"
)

func (m *MQTT) Dynamic(ctx context.Context, dynamicOpts *opts.DynamicOptions, dynamicSvc dynamic.IDynamic, errorCh chan<- *records.ErrorRecord) error {
	if err := validateDynamicOptions(dynamicOpts); err != nil {
		return errors.Wrap(err, "invalid dynamic options")
	}

	llog := m.log.WithField("pkg", "mqtt/dynamic")

	if err := dynamicSvc.Start(ctx, "MQTT", errorCh); err != nil {
		return errors.Wrap(err, "unable to create dynamic")
	}

	timeout := util.DurationSec(dynamicOpts.Mqtt.Args.WriteTimeoutSeconds)
	topic := dynamicOpts.Mqtt.Args.Topic

	outboundCh := dynamicSvc.Read()

	for {
		select {
		case outbound := <-outboundCh:
			token := m.client.Publish(topic, byte(int(m.connArgs.QosLevel)), false, outbound.Blob)

			if !token.WaitTimeout(timeout) {
				return fmt.Errorf("timed out attempting to publish message after %d seconds",
					dynamicOpts.Mqtt.Args.WriteTimeoutSeconds)
			}

			if token.Error() != nil {
				return errors.Wrap(token.Error(), "unable to replay message")
			}

			llog.Debugf("Replayed message to MQTT topic '%s' for replay '%s'", topic, outbound.ReplayId)
		case <-ctx.Done():
			m.log.Debug("context cancelled")
			return nil
		}
	}
}

func validateDynamicOptions(dynamicOpts *opts.DynamicOptions) error {
	if dynamicOpts == nil {
		return validate.ErrEmptyDynamicOpts
	}

	if dynamicOpts.Mqtt == nil {
		return validate.ErrEmptyBackendGroup
	}

	if dynamicOpts.Mqtt.Args == nil {
		return validate.ErrEmptyBackendArgs
	}

	if dynamicOpts.Mqtt.Args.Topic == "" {
		return ErrEmptyTopic
	}

	return nil
}
