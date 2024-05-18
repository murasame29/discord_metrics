package metrics

import (
	"errors"

	"github.com/murasame29/discord_metrics/metrics/internal/metadata"
	"go.opentelemetry.io/collector/component"
)

var collectGuildIDs map[string]struct{}

type Config struct {
	Token string `mapstructure:"token"`
	// BufferInterval is the interval period to buffer messages from Discord.
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m" and "h", which are
	// supported by time.ParseDuration().
	// Default is "30s".
	BufferInterval string `mapstructure:"buffer_interval"`

	// ServerWide is an optional setting to collect statistics from all channels in the server.
	// If it is false, the receiver collects statistics per channel.
	// If it is true, the receiver ignores the Channels setting.
	// Default is false.
	ServerWide bool `mapstructure:"server_wide,omitempty"`

	// Channels is an optional setting to collect statistics from specific channels.
	// If it is empty, the receiver collects statistics from all channels.
	// The ServerWide setting is true, receiver ignores this setting.
	Guilds []string `mapstructure:"guilds,omitempty"`

	MetricsBuilderConfig metadata.MetricsBuilderConfig
}

func NewDefaultConfig() component.Config {
	return &Config{
		Token:                "",
		BufferInterval:       "30s",
		ServerWide:           true,
		Guilds:               []string{},
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
	}
}

var _ component.Config = (*Config)(nil)

func (c *Config) Validate() error {
	if c.Token == "" {
		return errors.New("token cannot be empty")
	}

	collectGuildIDs = make(map[string]struct{}, len(c.Guilds))
	for _, guild := range c.Guilds {
		collectGuildIDs[guild] = struct{}{}
	}

	return nil
}
