package discordbot

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/tombuente/tresor/rest"
)

var commands = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "snippet",
		Description: "Create a snippet",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "key",
				Description: "Base36 encoded key",
				Required:    true,
			},
		},
	},
}

func newCommandHandler(api rest.Client) *handler.Mux {
	h := handler.New()
	h.Command("/snippet", snippetCommandHandler(api))

	return h
}

func snippetCommandHandler(api rest.Client) handler.CommandHandler {
	return func(event *handler.CommandEvent) error {
		data := event.SlashCommandInteractionData()
		key := data.String("key")

		snippet, err := api.GetSnippet(key)
		if err != nil {
			return event.CreateMessage(CreateErrorMessage("Unable to retrieve snippet"))
		}

		return event.CreateMessage(CreateMessage(fmt.Sprintf("Snippet: %s", key), (fmt.Sprintf("```%s\n%s\n```", snippet.Language, snippet.Content))))
	}
}
