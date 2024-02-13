package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tombuente/tresor/internal/discordbot"
	"github.com/tombuente/tresor/rest"
)

func main() {
	config, err := discordbot.NewConfigFromEnv()
	if err != nil {
		log.Fatal("Error reading config from envionment: ", err)
	}

	discordBot, err := discordbot.New(config, rest.New())
	if err != nil {
		log.Fatal("Error creating Discord bot: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = discordBot.Client.OpenGateway(ctx); err != nil {
		log.Fatal("Error opening bot gateway: ", err)
	}
	defer discordBot.Client.Close(context.TODO())

	log.Println("Running. Press CTRL-C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-stop
}
