package tresorbot

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
)

const (
	PrimaryColor = 0x5c5fea
	DangerColor  = 0xd43535
)

func CreateMessage(title string, content string) discord.MessageCreate {
	return discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Title:       title,
				Description: content,
				Color:       PrimaryColor,
			},
		},
	}
}

func CreateMessagef(title string, format string, a ...any) discord.MessageCreate {
	return CreateMessage(title, fmt.Sprintf(format, a...))
}

func CreateErrorMessage(message string, a ...any) discord.MessageCreate {
	return discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Description: fmt.Sprintf(message, a...),
				Color:       DangerColor,
			},
		},
		Flags: discord.MessageFlagEphemeral,
	}
}

func CreateDerivedErrorMessage(message string, err error, a ...any) discord.MessageCreate {
	if message == "" {
		return CreateErrorMessage(err.Error())
	}
	return CreateErrorMessage(message + ": " + err.Error())
}
