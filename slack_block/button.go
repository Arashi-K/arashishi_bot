package slackBlock

import (
	action "arashishi_bot/api/action_type"

	"github.com/slack-go/slack"
)

type ButtonType slack.Style

const (
	Primary = ButtonType(slack.StylePrimary)
	Warning = ButtonType(slack.StyleDanger)
)

func (t ButtonType) Value() slack.Style {
	return slack.Style(t)
}

func Button(buttonType ButtonType, key, label string) *slack.ButtonBlockElement {
	text := text(label)
	button := slack.NewButtonBlockElement("", key, text)
	button.WithStyle(buttonType.Value())
	return button
}

func ButtonsBlock(actionType action.ActionType, buttons ...slack.BlockElement) *slack.ActionBlock {
	return slack.NewActionBlock(actionType.Value(), buttons...)
}
