package tresorbot

import (
	"context"
	"log"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
)

type Bot struct {
	Config Config
	Client bot.Client // Must be added after initialization with SetupClient
}

func New(cfg Config) *Bot {
	return &Bot{
		Config: cfg,
	}
}

func (b *Bot) SetupClient(listeners ...bot.EventListener) {
	var err error
	b.Client, err = disgo.New(
		b.Config.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessages,
			),
		),
		bot.WithEventListeners(listeners...),
	)
	if err != nil {
		log.Fatal("Error while building bot client: ", err)
	}
}

func (b *Bot) OnReady(_ *events.Ready) {
	log.Println("Ready!")
	if err := b.Client.SetPresence(context.TODO(), gateway.WithListeningActivity("you"), gateway.WithOnlineStatus(discord.OnlineStatusOnline)); err != nil {
		log.Printf("Failed to set presence: %s", err)
	}
}
