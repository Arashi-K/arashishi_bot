package slackAPI

import (
	config "arashishi_bot/config"

	"github.com/slack-go/slack"
)

var (
	api          = slack.New(config.SlackClientToken)
	fallbackText = slack.MsgOptionText("This client is not supported.", false)
)

func PostMessage(message slack.MsgOption, channelId string) (err error) {
	_, _, err = api.PostMessage(channelId, message)
	return
}

func PostMessagePrivate(message slack.MsgOption, channelId, userId string) (err error) {
	_, err = api.PostEphemeral(channelId, userId, fallbackText, message)
	return
}

func PostMessageReplace(message slack.MsgOption, url string) (err error) {
	replace := slack.MsgOptionReplaceOriginal(url)
	_, _, _, err = api.SendMessage("", replace, message)
	return
}

func DeleteMessage(url string) (err error) {
	delete := slack.MsgOptionDeleteOriginal(url)
	_, _, _, err = api.SendMessage("", delete)
	return
}
