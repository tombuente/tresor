package discordbot

import (
	"fmt"
	"log"
	"regexp"

	"github.com/disgoorg/disgo/events"
	"github.com/tombuente/tresor/rest"
)

func (discordBot DiscordBot) OnReady() func(event *events.Ready) {
	return func(event *events.Ready) {
		log.Println("Discord Bot is ready!")
	}
}

func (discordBot DiscordBot) OnMessageCreate() func(event *events.MessageCreate) {
	return func(event *events.MessageCreate) {
		if event.Message.Author.Bot {
			return
		}

		id := event.Client().ApplicationID().String()
		snippetExp := regexp.MustCompile(fmt.Sprintf(`(?sU)<@%s>\s*\n\s*\x60\x60\x60(?P<language>\w*)\s*\n(?P<code>.+)\n\s*\x60\x60\x60`, id))

		snippets := []rest.SnippetRequest{}
		for _, match := range snippetExp.FindAllStringSubmatch(event.Message.Content, -1) {
			snippet := rest.SnippetRequest{
				Language: match[1],
				Content:  match[2],
			}
			snippets = append(snippets, snippet)
		}

		for _, snippet := range snippets {
			createdSnippet, err := discordBot.API.PostSnippet(snippet)
			if err != nil {
				event.Client().Rest().CreateMessage(event.ChannelID, CreateErrorMessage("Error creating snippet"))
				return
			}

			event.Client().Rest().CreateMessage(event.ChannelID, CreateMessagef("Snippet", "Key: %s", createdSnippet.Key))
		}
	}
}
