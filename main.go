package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/rs/zerolog"
)

func main() {
	app := &application{
		logger: zerolog.New(
			zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC822,
			},
		).With().Timestamp().Logger(),
	}

	app.config = loadConfig()
	app.logger.Info().Str("config", fmt.Sprint(app.config)).Msg("")

	setupGracefulShutdown(app)

	app.client = model.NewAPIv4Client(app.config.server.String())

	app.client.SetToken(app.config.token)

	if user, res, err := app.client.GetUser("me", ""); err != nil {
		app.logger.Fatal().Err(err).Msg("Login failed")
	} else {

		app.logger.Debug().Interface("user", user).Interface("res", res).Msg("")
		app.logger.Info().Msg("Logged in")
		app.user = user
	}

	if team, res, err := app.client.GetTeamByName(app.config.teamName, ""); err != nil {
		app.logger.Fatal().Err(err).Msg("Team not found")
	} else {
		app.logger.Debug().Interface("team", team).Interface("res", res).Msg("")
		app.team = team
	}

	if channels, res, err := app.client.GetChannelsForTeamForUser(app.team.Id, app.user.Id, false, ""); err != nil {
		app.logger.Fatal().Err(err).Msg("Channels not found")
	} else {
		app.logger.Debug().Interface("channels", channels).Interface("res", res).Msg("")
		app.channels = channels
	}

	sendMessageToChannels(app, app.config.broascastMessage)
}
