package metrics

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (dh *discordHandler) initEvents(ctx context.Context) {
	dh.session.AddHandler(dh.messageCreateFunc(ctx))
	dh.session.AddHandler(dh.joinFunc(ctx))
	dh.session.AddHandler(dh.leaveFunc(ctx))
	dh.session.AddHandler(dh.createChannelFunc(ctx))
	dh.session.AddHandler(dh.deletechannelFunc(ctx))
	dh.session.AddHandler(dh.voiceStateFunc(ctx))
}

type eventChannels struct {
	mcCh chan messageCreateEvent
	gaCh chan guildMemberAddEvent
	grCh chan guildMemberRemoveEvent
	vsCh chan voiceStateEvent
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
		log.Println("messageCreateFunc", m)
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
		log.Println("joinFunc", g)
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
		log.Println("leaveFunc", g)
		dh.grCh <- guildMemberRemoveEvent{
			s: s,
			g: g,
		}

	}
}

type channelCreateEvent struct {
	s *discordgo.Session
	c *discordgo.ChannelCreate
}

func (dh *discordHandler) createChannelFunc(_ context.Context) func(s *discordgo.Session, c *discordgo.ChannelCreate) {
	return func(s *discordgo.Session, c *discordgo.ChannelCreate) {
		log.Println("createChannelFunc", c)
		ChannelMaps[c.ID] = c.Channel
	}
}

type channelDeleteEvent struct {
	s *discordgo.Session
	c *discordgo.ChannelDelete
}

func (dh *discordHandler) deletechannelFunc(_ context.Context) func(s *discordgo.Session, c *discordgo.ChannelDelete) {
	return func(s *discordgo.Session, c *discordgo.ChannelDelete) {
		log.Println("deletechannelFunc", c)
		delete(ChannelMaps, c.ID)
	}
}

type voiceStateEvent struct {
	s *discordgo.Session
	v *discordgo.VoiceStateUpdate
}

type VoiceChannelEvent string

const (
	VoiceChannelEventJoin  VoiceChannelEvent = "join"
	VoiceChannelEventLeave VoiceChannelEvent = "leave"
	VoiceChannelEventMove  VoiceChannelEvent = "mute"
	VoiceChannelEventMute  VoiceChannelEvent = "mute"
)

func (dh *discordHandler) voiceStateFunc(_ context.Context) func(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	return func(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
		log.Println("voiceStateFunc", v)
		dh.vsCh <- voiceStateEvent{
			s,
			v,
		}
	}
}
