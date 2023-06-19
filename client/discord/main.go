package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

const api = "http://0.0.0.0:8080/api"

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "get-code-snippet",
			Description: "Get code snippet",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "key",
					Description: "Code snippet key",
					Required:    true,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"get-code-snippet": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			var resMsg string
			var body []byte
			var snippet Snippet

			var key string
			if opt, ok := optionMap["key"]; ok {
				key = opt.StringValue()
			}

			res, err := http.Get(api + "/snippets/" + key)
			if err != nil {
				resMsg = "Error: unable to get snippet"
				goto respond
			}
			if res.StatusCode != 200 {
				resMsg = "res.StatusCode"
				goto respond
			}

			body, err = io.ReadAll(res.Body)
			if err != nil {
				resMsg = "failed to read response body"
				goto respond
			}

			err = json.Unmarshal(body, &snippet)
			if err != nil {
				resMsg = "received malformatted response body"
				goto respond
			}

			resMsg = fmt.Sprintf("```%s\n%s\n```", snippet.Language, snippet.Content)

		respond:
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: resMsg,
				},
			})
		},
	}
)

type Snippet struct {
	Key      string `json:"key"`
	Content  string `json:"content"`
	Language string `json:"language"`
}

func main() {
	token := os.Getenv("TOKEN")
	guild := os.Getenv("GUILD")

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = s.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, guild, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Shutting down...")
}
