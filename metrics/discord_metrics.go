package metrics

import (
	"log"
	"time"

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
	dh.mb.RecordDiscordMessagesCountDataPoint(
		now,
		1,
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

	dh.mb.RecordDiscordJoinCountDataPoint(now, 1, e.g.Member.User.GlobalName)
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

	dh.mb.RecordDiscordLeaveCountDataPoint(now, 1, e.g.Member.User.GlobalName)
}
