package command

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"

	tresorAPI "github.com/tombuente/tresor/api"
	tresorbot "github.com/tombuente/tresor/client/discord"
)

var Commands = []discord.ApplicationCommandCreate{
	snippet,
}

var snippet = discord.SlashCommandCreate{
	Name:        "snippet",
	Description: "Create a snippet",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:        "key",
			Description: "Base36 encoded key",
			Required:    true,
		},
	},
}

func SnippetHandler(tresor tresorAPI.Tresor) handler.CommandHandler {
	return func(event *handler.CommandEvent) error {
		data := event.SlashCommandInteractionData()
		key := data.String("key")

		snippet, err := tresor.GetSnippet(key)
		if err != nil {
			return event.CreateMessage(tresorbot.CreateErrorMessage("Unable to get snippet."))
		}

		return event.CreateMessage(tresorbot.CreateMessage(fmt.Sprintf("Snippet: %s", key), (fmt.Sprintf("```%s\n%s\n```", snippet.Language, snippet.Content))))
	}
}
