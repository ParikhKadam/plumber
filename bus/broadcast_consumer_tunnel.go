package bus

import (
	"context"
	"fmt"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber/validate"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func (b *Bus) doCreateTunnel(ctx context.Context, msg *Message) error {
	tunnelOptions := &opts.DynamicOptions{}
	if err := proto.Unmarshal(msg.Data, tunnelOptions); err != nil {
		return errors.Wrap(err, "unable to unmarshal message into protos.DynamicOptions")
	}

	if err := validate.TunnelOptionsForServer(tunnelOptions); err != nil {
		return errors.Wrap(err, "tunnel option validation failed")
	}

	if _, err := b.config.Actions.CreateTunnel(ctx, tunnelOptions); err != nil {
		return errors.Wrap(err, "unable to create tunnel")
	}

	b.log.Infof("Created tunnel '%s' (from broadcast msg)", tunnelOptions.XDynamicId)

	return nil
}

func (b *Bus) doUpdateTunnel(ctx context.Context, msg *Message) error {
	tunnelOptions := &opts.DynamicOptions{}
	if err := proto.Unmarshal(msg.Data, tunnelOptions); err != nil {
		return errors.Wrap(err, "unable to unmarshal message into protos.DynamicOptions")
	}

	if tunnelOptions.XDynamicId == "" {
		return errors.New("tunnel id in options cannot be empty")
	}

	if _, err := b.config.Actions.UpdateTunnel(ctx, tunnelOptions.XDynamicId, tunnelOptions); err != nil {
		return fmt.Errorf("unable to update tunnel '%s': %s", tunnelOptions.XDynamicId, err)
	}

	b.log.Infof("Updated tunnel '%s' (from broadcast msg)", tunnelOptions.XDynamicId)

	return nil
}

func (b *Bus) doStopTunnel(ctx context.Context, msg *Message) error {
	// Only unmarshalling to get XDynamicID - we'll be operating off of what's in
	// our cache.
	tunnelOptions := &opts.DynamicOptions{}
	if err := proto.Unmarshal(msg.Data, tunnelOptions); err != nil {
		return errors.Wrap(err, "unable to unmarshal message into protos.DynamicOptions")
	}

	if tunnelOptions.XDynamicId == "" {
		return errors.New("tunnel id in options cannot be empty")
	}

	if _, err := b.config.Actions.StopTunnel(ctx, tunnelOptions.XDynamicId); err != nil {
		return fmt.Errorf("unable to stop tunnel '%s': %s", tunnelOptions.XDynamicId, err)
	}

	b.log.Infof("Stopped tunnel '%s' (from broadcast msg)", tunnelOptions.XDynamicId)

	return nil
}

func (b *Bus) doResumeTunnel(ctx context.Context, msg *Message) error {
	// Only unmarshalling to get DynamicId - we'll be operating off of what's in
	// our cache.
	tunnelOptions := &opts.DynamicOptions{}
	if err := proto.Unmarshal(msg.Data, tunnelOptions); err != nil {
		return errors.Wrap(err, "unable to unmarshal message into protos.DynamicOptions")
	}

	if tunnelOptions.XDynamicId == "" {
		return errors.New("tunnel id in options cannot be empty")
	}

	if _, err := b.config.Actions.ResumeTunnel(ctx, tunnelOptions.XDynamicId); err != nil {
		return fmt.Errorf("unable to resume tunnel '%s': %s", tunnelOptions.XDynamicId, err)
	}

	b.log.Infof("Resumed tunnel '%s' (from broadcast msg)", tunnelOptions.XDynamicId)

	return nil
}

func (b *Bus) doDeleteTunnel(ctx context.Context, msg *Message) error {
	// Only unmarshalling to get XDynamicId - we'll be operating off of what's in
	// our cache.
	tunnelOptions := &opts.DynamicOptions{}
	if err := proto.Unmarshal(msg.Data, tunnelOptions); err != nil {
		return errors.Wrap(err, "unable to unmarshal message into protos.DynamicOptions")
	}

	if tunnelOptions.XDynamicId == "" {
		return errors.New("tunnel id in options cannot be empty")
	}

	if err := b.config.Actions.DeleteTunnel(ctx, tunnelOptions.XDynamicId); err != nil {
		return fmt.Errorf("unable to delete tunnel '%s': %s", tunnelOptions.XDynamicId, err)
	}

	b.log.Infof("Deleted tunnel '%s' (from broadcast msg)", tunnelOptions.XDynamicId)

	return nil
}
