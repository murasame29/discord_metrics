package metrics

import (
	"context"
	"fmt"

	"github.com/murasame29/discord_metrics/receiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"

	"go.opentelemetry.io/collector/receiver"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		NewDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
	)
}

func createMetricsReceiver(
	_ context.Context,
	settings receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	if consumer == nil {
		return nil, fmt.Errorf("nil consumer")
	}

	c := cfg.(*Config)
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return newDiscordReceiver(c, settings, consumer)
}
