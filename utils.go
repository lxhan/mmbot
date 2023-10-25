package main

import (
	"os"
	"os/signal"

	"github.com/mattermost/mattermost-server/v6/model"
)

func setupGracefulShutdown(app *application) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			if app.websocketClient != nil {
				app.logger.Info().Msg("Closing websocket connection")
				app.websocketClient.Close()
			}
			app.logger.Info().Msg("Shutting down")
			os.Exit(0)
		}
	}()
}

func sendMessageToChannels(app *application, msg string) {
	for _, channel := range app.channels {
		if channel.Name == "town-square" {
			continue
		}
		app.logger.Info().Str("channel", channel.Name).Msg("Sending message")
		if _, _, err := app.client.CreatePost(&model.Post{
			ChannelId: channel.Id,
			Message:   msg,
		}); err != nil {
			app.logger.Error().Err(err).Msg("Failed to send message")
		}
	}
}
