package listener

import (
	"fmt"
	"regexp"

	"github.com/disgoorg/disgo/events"
	tresorbot "github.com/tombuente/tresor/client/discord"
	"github.com/tombuente/tresor/spec/snippetspec"

	tresorAPI "github.com/tombuente/tresor/api"
)

func OnMessageCreate(tresor tresorAPI.Tresor) func(event *events.MessageCreate) {
	return func(event *events.MessageCreate) {
		if event.Message.Author.Bot {
			return
		}

		id := event.Client().ApplicationID().String() // ID of this bot

		snippetExp := regexp.MustCompile(fmt.Sprintf(`(?sU)<@%s>\s*\n\s*\x60\x60\x60(?P<language>\w*)\s*\n(?P<code>.+)\n\s*\x60\x60\x60`, id))

		snippets := []snippetspec.SnippetReq{}

		for _, match := range snippetExp.FindAllStringSubmatch(event.Message.Content, -1) {
			snippet := snippetspec.SnippetReq{
				Language: match[1],
				Content:  match[2],
			}
			snippets = append(snippets, snippet)
		}

		for _, snippet := range snippets {
			createdSnippet, err := tresor.PostSnippet(snippet)
			if err != nil {
				event.Client().Rest().CreateMessage(event.ChannelID, tresorbot.CreateErrorMessage("Error creating snippet."))
				return
			}

			event.Client().Rest().CreateMessage(event.ChannelID, tresorbot.CreateMessagef("Snippet", "Key: %s", createdSnippet.Key))
		}
	}
}
