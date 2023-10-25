package main

import (
	"net/url"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	teamName         string
	token            string
	server           *url.URL
	broascastMessage string
}

func loadConfig() config {
	var settings config

	settings.teamName = os.Getenv("MM_TEAMNAME")
	settings.token = os.Getenv("MM_TOKEN")
	settings.server, _ = url.Parse(os.Getenv("MM_SERVER"))
	settings.broascastMessage = os.Getenv("BROADCAST_MESSAGE")

	return settings
}
