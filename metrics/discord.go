package metrics

import (
	"context"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/murasame29/discord_metrics/metrics/internal/metadata"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

// ChannelMaps は、channelIDをkey, channel_nameをvalueで保存する
var (
	ChannelMaps map[string]*discordgo.Channel
)

func initChannelKV(ss *discordgo.Session, guildID string) {
	channels, err := ss.GuildChannels(guildID)
	if err != nil {
		panic(err)
	}

	channelMaps := make(map[string]*discordgo.Channel)
	for _, channel := range channels {
		channelMaps[channel.ID] = channel
	}

	log.Println(channelMaps)
	ChannelMaps = channelMaps
}

type discordHandler struct {
	session  *discordgo.Session
	consumer consumer.Metrics
	cancel   context.CancelFunc
	config   *Config
	obsrecv  *receiverhelper.ObsReport
	mb       *metadata.MetricsBuilder
	eventChannels
}

func newDiscordHandler(consumer consumer.Metrics, cfg *Config, settings receiver.CreateSettings, obsrecv *receiverhelper.ObsReport) (*discordHandler, error) {
	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, err
	}

	dh := &discordHandler{
		session:       s,
		consumer:      consumer,
		config:        cfg,
		obsrecv:       obsrecv,
		mb:            metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
		eventChannels: newEventChannels(),
	}
	return dh, nil
}

const (
	dataFormat = "discord"
)

func (dh *discordHandler) run(ctx context.Context) error {
	dh.initEvents(ctx)

	if err := dh.session.Open(); err != nil {
		return err
	}
	defer dh.session.Close()

	d, err := time.ParseDuration(dh.config.BufferInterval)
	if err != nil {
		return err
	}
	initChannelKV(dh.session, dh.config.GuildID)
	ticker := time.NewTicker(d)
TICK:
	for {
		select {
		case e := <-dh.mcCh:
			dh.messageCreateToMetrics(e)
		case e := <-dh.gaCh:
			dh.guildJoinToMetrics(e)
		case e := <-dh.grCh:
			dh.guildLeaveToMetrics(e)
		case e := <-dh.vsCh:
			dh.voiceStateEventToMetrics(e)
		case <-ticker.C:
			metrics := dh.mb.Emit()
			dh.obsrecv.StartMetricsOp(ctx)
			err := dh.consumer.ConsumeMetrics(ctx, metrics)
			dh.obsrecv.EndMetricsOp(ctx, dataFormat, metrics.DataPointCount(), err)
		case <-ctx.Done():
			break TICK
		}
	}

	return nil
}
