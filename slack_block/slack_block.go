package slackBlock

import (
	"github.com/slack-go/slack"
)

func Blocks(blocks ...slack.Block) slack.MsgOption {
	return slack.MsgOptionBlocks(blocks...)
}
