package slackBlock

import (
	action "arashishi_bot/api/action_type"

	"github.com/slack-go/slack"
)

func Option(key, value string) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(value, text(key), text(key))
}

func SelectBlock(
	actionType action.ActionType,
	placeholder string,
	options ...*slack.OptionBlockObject,
) *slack.ActionBlock {
	selectBlock := slack.NewOptionsSelectBlockElement(
		slack.OptTypeStatic,
		text(placeholder),
		"",
		options...,
	)
	return slack.NewActionBlock(actionType.Value(), selectBlock)
}
