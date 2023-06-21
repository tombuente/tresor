package tresorbot

import (
	"errors"
	"os"

	"github.com/disgoorg/snowflake/v2"
)

type Config struct {
	DevMode  bool
	DevGuild snowflake.ID
	Token    string
}

func NewConfigFromEnv() (Config, error) {
	token := os.Getenv("TOKEN")
	guild := snowflake.GetEnv("GUILD")

	if token == "" {
		return Config{}, errors.New("no token provided")
	}

	return Config{
		DevMode:  true,
		DevGuild: guild,
		Token:    token,
	}, nil
}
