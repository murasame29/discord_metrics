package metrics

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

type discordReceiver struct {
	consumer consumer.Metrics
	settings receiver.CreateSettings
	cancel   context.CancelFunc
	config   *Config
	dh       *discordHandler
	obsrecv  *receiverhelper.ObsReport
}

func newDiscordReceiver(config *Config, settings receiver.CreateSettings, consumer consumer.Metrics) (*discordReceiver, error) {
	obsrecv, err := receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             settings.ID,
		Transport:              "event",
		ReceiverCreateSettings: settings,
	})
	if err != nil {
		return nil, err
	}
	return &discordReceiver{
		consumer: consumer,
		settings: settings,
		config:   config,
		obsrecv:  obsrecv,
	}, nil
}

func (r *discordReceiver) Start(ctx context.Context, _ component.Host) error {
	ctx, r.cancel = context.WithCancel(ctx)

	var err error
	r.dh, err = newDiscordHandler(r.consumer, r.config, r.settings, r.obsrecv)
	if err != nil {
		return err
	}
	if r.dh == nil {
		return errors.New("failed to create discord handler")
	}
	if err := r.dh.run(ctx); err != nil {
		return err
	}
	return nil
}

func (r *discordReceiver) Shutdown(ctx context.Context) error {
	if r.cancel != nil {
		r.cancel()
	}
	return nil
}
