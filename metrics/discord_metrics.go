package metrics

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

func (dh *discordHandler) messageCreateToMetrics(e messageCreateEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	if _, ok := collectGuildIDs[e.m.GuildID]; !ok {
		return
	}

	dh.mb.RecordDiscordMessagesCountDataPoint(now, 1, e.m.ChannelID)
}

func (dh *discordHandler) guildJoinToMetrics(e guildMemberAddEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	// ignore bot
	if e.g.Member.User.Bot {
		return
	}

	if _, ok := collectGuildIDs[e.g.GuildID]; !ok {
		return
	}

	dh.mb.RecordDiscordJoinCountDataPoint(now, 1, e.g.Member.User.ID)
}

func (dh *discordHandler) guildLeaveToMetrics(e guildMemberRemoveEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	// ignore bot
	if e.g.Member.User.Bot {
		return
	}

	if _, ok := collectGuildIDs[e.g.GuildID]; !ok {
		return
	}

	dh.mb.RecordDiscordLeaveCountDataPoint(now, 1, e.g.Member.User.ID)
}
