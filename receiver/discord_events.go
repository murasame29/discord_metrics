package metrics

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (dh *discordHandler) initEvents(ctx context.Context) {
	dh.session.AddHandler(dh.messageCreateFunc(ctx))
	dh.session.AddHandler(dh.joinFunc(ctx))
	dh.session.AddHandler(dh.leaveFunc(ctx))
}

type eventChannels struct {
	mcCh chan messageCreateEvent
	gaCh chan guildMemberAddEvent
	grCh chan guildMemberRemoveEvent
}

func newEventChannels() eventChannels {
	return eventChannels{
		mcCh: make(chan messageCreateEvent),
		gaCh: make(chan guildMemberAddEvent),
		grCh: make(chan guildMemberRemoveEvent),
	}
}

type messageCreateEvent struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

func (dh *discordHandler) messageCreateFunc(_ context.Context) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		dh.mcCh <- messageCreateEvent{
			s: s,
			m: m,
		}
	}
}

type guildMemberAddEvent struct {
	s *discordgo.Session
	g *discordgo.GuildMemberAdd
}

func (dh *discordHandler) joinFunc(_ context.Context) func(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
	return func(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
		dh.gaCh <- guildMemberAddEvent{
			s: s,
			g: g,
		}
	}
}

type guildMemberRemoveEvent struct {
	s *discordgo.Session
	g *discordgo.GuildMemberRemove
}

func (dh *discordHandler) leaveFunc(_ context.Context) func(s *discordgo.Session, g *discordgo.GuildMemberRemove) {
	return func(s *discordgo.Session, g *discordgo.GuildMemberRemove) {
		dh.grCh <- guildMemberRemoveEvent{
			s: s,
			g: g,
		}

	}
}
