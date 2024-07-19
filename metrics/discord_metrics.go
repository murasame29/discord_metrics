package metrics

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func (dh *discordHandler) messageCreateToMetrics(e messageCreateEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	if e.m.GuildID != dh.config.GuildID {
		log.Panicln("[INFO] ingore", e.m.Content)
		return
	}

	channel := ChannelMaps[e.m.ChannelID]
	parent := ChannelMaps[channel.ParentID]
	dh.mb.RecordDiscordMetricsMessagesCountDataPoint(
		now,
		1,
		e.m.GuildID,
		dh.config.Environment,
		parent.ID,
		parent.Name,
		channel.ID,
		channel.Name,
		e.m.Author.ID,
		e.m.Author.GlobalName,
	)
}

func (dh *discordHandler) guildJoinToMetrics(e guildMemberAddEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	// ignore bot
	if e.g.Member.User.Bot {
		return
	}

	if e.g.GuildID != dh.config.GuildID {
		return
	}

	dh.mb.RecordDiscordMetricsJoinCountDataPoint(
		now,
		1,
		e.g.GuildID,
		dh.config.Environment,
		e.g.Member.User.ID,
		e.g.Member.User.GlobalName,
	)
}

func (dh *discordHandler) guildLeaveToMetrics(e guildMemberRemoveEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	// ignore bot
	if e.g.Member.User.Bot {
		return
	}

	if e.g.GuildID != dh.config.GuildID {
		return
	}

	dh.mb.RecordDiscordMetricsLeaveCountDataPoint(
		now,
		1,
		e.g.GuildID,
		dh.config.Environment,
		e.g.Member.User.ID,
		e.g.Member.User.GlobalName,
	)
}

func (dh *discordHandler) voiceStateEventToMetrics(e voiceStateEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	if e.v.GuildID != dh.config.GuildID {
		return
	}

	channel := ChannelMaps[e.v.ChannelID]
	parent := ChannelMaps[channel.ParentID]

	dh.mb.RecordDiscordMetricsVcEventCountDataPoint(
		now,
		1,
		e.v.GuildID,
		dh.config.Environment,
		parent.ID,
		parent.Name,
		channel.ID,
		channel.Name,
		e.v.Member.User.ID,
		e.v.Member.User.GlobalName,
		string(checkVoiceChannelEvent(e.v)),
	)
}

func checkVoiceChannelEvent(v *discordgo.VoiceStateUpdate) VoiceChannelEvent {
	if v.BeforeUpdate == nil {
		return VoiceChannelEventJoin
	}

	if v.VoiceState.ChannelID == "" {
		return VoiceChannelEventLeave
	}

	if v.VoiceState.ChannelID != v.BeforeUpdate.ChannelID {
		return VoiceChannelEventMove
	}

	return VoiceChannelEventMute
}
