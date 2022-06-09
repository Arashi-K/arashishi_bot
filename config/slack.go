package config

import "os"

const (
	SlackChannel = "#dev"
)

var (
	SlackClientToken = os.Getenv("SLACK_CLIENT_TOKEN")
	SlackSecretKey   = os.Getenv("SLACK_SECRET_KEY")
)
