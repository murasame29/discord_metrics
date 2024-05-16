package metrics

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

func (dh *discordHandler) messageCreateToMetrics(e messageCreateEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	channelID := e.m.ChannelID
	matched := false
	switch {
	case dh.config.ServerWide:
		channelID = "@all@"
		matched = true
	case len(dh.config.Channels) == 0:
		matched = true
	default:
		for _, ch := range dh.config.Channels {
			if ch == channelID {
				matched = true
				break
			}
		}
	}
	if !matched {
		return
	}
	dh.mb.RecordDiscordMessagesCountDataPoint(now, 1, channelID)
}

func (dh *discordHandler) guildJoinToMetrics(e guildMemberAddEvent) {
	now := pcommon.NewTimestampFromTime(time.Now())
	// ignore bot
	if e.g.Member.User.Bot {
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

	dh.mb.RecordDiscordLeaveCountDataPoint(now, 1, e.g.Member.User.ID)
}
