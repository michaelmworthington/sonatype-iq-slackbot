package main

import (
	"encoding/json"
	"bytes"
	"net/http"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
	"log"
)

//https://github.com/nlopes/slack/blob/master/examples/eventsapi/events.go
func slackEventSubscriptions(w http.ResponseWriter, r *http.Request) {
	slackAPI := slack.New(config.BotToken)

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: config.VerificationToken}))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("[ERROR] %s\n", err)
		return
	}

	//todo: switch this...maybe
	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("[ERROR] %s\n", err)
			return
			}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		//log.Printf("[DEBUG] Processing Callback %s\n", innerEvent.Data)

		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			slackAPI.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
		}
	}
}
