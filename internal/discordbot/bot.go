package discordbot

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/snowflake/v2"
	"github.com/tombuente/tresor/rest"
)

type Config struct {
	DevMode  bool
	DevGuild snowflake.ID
	Token    string
}

type DiscordBot struct {
	API    rest.Client
	Client bot.Client
}

func New(config Config, api rest.Client) (DiscordBot, error) {
	discordBot := DiscordBot{
		API: api,
	}

	var err error
	discordBot.Client, err = disgo.New(
		config.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessages,
			),
		),
		bot.WithEventListeners(
			bot.NewListenerFunc(discordBot.OnReady()),
			bot.NewListenerFunc(discordBot.OnMessageCreate()),
			newCommandHandler(discordBot.API),
		),
	)
	if err != nil {
		return DiscordBot{}, fmt.Errorf("error creating Discord bot client: %w", err)
	}

	if _, err = discordBot.Client.Rest().SetGuildCommands(discordBot.Client.ApplicationID(), config.DevGuild, commands); err != nil {
		log.Fatal("Error registering commands:", err)
	}

	return discordBot, nil
}

func NewConfigFromEnv() (Config, error) {
	token := os.Getenv("TOKEN")
	if token == "" {
		return Config{}, errors.New("no token provided")
	}

	guild := snowflake.GetEnv("GUILD")
	if guild == snowflake.ID(0) {
		return Config{}, errors.New("no guild provided")
	}

	return Config{
		DevMode:  true,
		DevGuild: guild,
		Token:    token,
	}, nil
}
