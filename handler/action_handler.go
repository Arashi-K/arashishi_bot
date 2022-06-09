package handler

import (
	action "arashishi_bot/api/action_type"
	slackAPI "arashishi_bot/api/slack_api"
	slackBlock "arashishi_bot/slack_block"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

func handleAction() {
	slackVerificationMiddleware("/slack/actions", func(w http.ResponseWriter, r *http.Request) {
		var payload *slack.InteractionCallback
		if err := json.Unmarshal([]byte(r.FormValue("payload")), &payload); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch payload.Type {
		case slack.InteractionTypeBlockActions:
			if len(payload.ActionCallback.BlockActions) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			a := payload.ActionCallback.BlockActions[0]
			switch a.BlockID {
			case action.SelectVersion.Value():
				value := a.Value
				log.Println(value)
				switch value {
				case "reserve":
					text := slackBlock.TextBlock("いつ？")

					selectBlock := slackBlock.SelectBlock(
						action.SelectDay,
						"予約日",
						slackBlock.Option("today", "今日"),
						slackBlock.Option("tommorow", "明日"),
						slackBlock.Option("dayAfterTommorow", "明後日"),
					)

					blocks := slackBlock.Blocks(text, selectBlock)

					if err := slackAPI.PostMessageReplace(blocks, payload.ResponseURL); err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
				case "schedule":
					text := slackBlock.TextBlock(fmt.Sprintf("Could I deploy `%s`?", value))

					buttons := slackBlock.ButtonsBlock(
						action.ConfirmDeploy,
						slackBlock.Button(slackBlock.Primary, "yes", "デプロイする"),
						slackBlock.Button(slackBlock.Primary, "no", "やっぱしない"),
					)

					blocks := slackBlock.Blocks(text, buttons)

					if err := slackAPI.PostMessageReplace(blocks, payload.ResponseURL); err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
				}
			case action.ConfirmDeploy.Value():
				if strings.HasPrefix(a.Value, "v") {
					version := a.Value
					go func() {
						willDeployText := slackBlock.TextBlock(fmt.Sprintf("<@%s> OK, I will deploy `%s`.", payload.User.ID, version))
						willDeployBlocks := slackBlock.Blocks(willDeployText)

						if err := slackAPI.PostMessage(willDeployBlocks, payload.Channel.ID); err != nil {
							log.Println(err)
						}

						deploy(version)

						completeDeployText := slackBlock.TextBlock(fmt.Sprintf("`%s` deployment completed!", version))
						completeDeployBlocks := slackBlock.Blocks(completeDeployText)

						if err := slackAPI.PostMessage(completeDeployBlocks, payload.Channel.ID); err != nil {
							log.Println(err)
						}
					}()
				}

				if err := slackAPI.DeleteMessage(payload.ResponseURL); err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
	})
}

func deploy(version string) {
	log.Printf("deploy %s", version)
	time.Sleep(10 * time.Second)
}
