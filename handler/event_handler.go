package handler

import (
	action "arashishi_bot/api/action_type"
	slackAPI "arashishi_bot/api/slack_api"
	slackBlock "arashishi_bot/slack_block"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/slack-go/slack/slackevents"
)

func handleEvent() {
	slackVerificationMiddleware("/slack/events", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch eventsAPIEvent.Type {
		case slackevents.URLVerification:
			var res *slackevents.ChallengeResponse
			if err := json.Unmarshal(body, &res); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/plain")
			if _, err := w.Write([]byte(res.Challenge)); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		case slackevents.CallbackEvent:
			innerEvent := eventsAPIEvent.InnerEvent
			switch event := innerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				message := strings.Split(event.Text, " ")
				if len(message) < 2 {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				command := message[1]
				switch command {
				case "会議室":

					text := slackBlock.TextBlock("どれにする？")

					buttons := slackBlock.ButtonsBlock(
						action.SelectVersion,
						slackBlock.Button(slackBlock.Primary, "reserve", "予約する"),
						slackBlock.Button(slackBlock.Primary, "schedule", "今日の予約"),
					)

					blocks := slackBlock.Blocks(text, buttons)
					log.Println(blocks)
					if err := slackAPI.PostMessagePrivate(blocks, event.Channel, event.User); err != nil {
						log.Println(1)
						log.Println(err)
						w.WriteHeader(http.StatusInternalServerError)
						log.Println(2)
						return
					}
				}
			}
		}
	})
}
