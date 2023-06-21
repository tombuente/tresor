package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/handler"
	"github.com/tombuente/tresor/api"
	tresorbot "github.com/tombuente/tresor/client/discord"
	"github.com/tombuente/tresor/client/discord/command"
	"github.com/tombuente/tresor/client/discord/listener"
)

func main() {
	api := api.New()

	cfg, err := tresorbot.NewConfigFromEnv()
	if err != nil {
		log.Fatal("Error reading config from env:", err)
	}

	b := tresorbot.New(cfg)

	h := handler.New()
	h.Command("/snippet", command.SnippetHandler(api))

	b.SetupClient(h, bot.NewListenerFunc(listener.OnMessageCreate(api)))

	if _, err = b.Client.Rest().SetGuildCommands(b.Client.ApplicationID(), b.Config.DevGuild, command.Commands); err != nil {
		log.Fatal("Error while registering commands: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = b.Client.OpenGateway(ctx); err != nil {
		log.Fatal("Failed to connect to gateway:", err)
	}
	defer b.Client.Close(context.TODO())

	log.Println("Running. Press CTRL-C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-stop
}
