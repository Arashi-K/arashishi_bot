package slackBlock

import (
	"fmt"

	"github.com/slack-go/slack"
)

func text(message string) *slack.TextBlockObject {
	return slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprint(message), false, false)
}

func TextBlock(message string) *slack.SectionBlock {
	textBlock := text(message)
	return slack.NewSectionBlock(textBlock, nil, nil)
}
