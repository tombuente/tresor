package discordbot

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

func CreateErrorMessage(format string, a ...any) discord.MessageCreate {
	return discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Description: fmt.Sprintf(format, a...),
				Color:       DangerColor,
			},
		},
		Flags: discord.MessageFlagEphemeral,
	}
}
