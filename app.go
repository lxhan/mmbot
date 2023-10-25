package main

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/rs/zerolog"
)

type application struct {
	config          config
	logger          zerolog.Logger
	client          *model.Client4
	websocketClient *model.WebSocketClient
	user            *model.User
	channels        []*model.Channel
	team            *model.Team
}
