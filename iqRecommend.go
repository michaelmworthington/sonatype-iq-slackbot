package main

import (
	"github.com/nlopes/slack"
	"net/http"
	"log"
)

// Sample from https://github.com/nlopes/slack/blob/master/examples/slash/slash.go
// TODO: look at https://glitch.com/edit/#!/slack-slash-command-and-dialogs-blueprint?path=src/index.js:1:0 for posting a dialog box to type in the maven gav
func iqRecommend(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodPost:
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("[ERROR] %s\n", err)
			return
		}

		if !s.ValidateToken(config.VerificationToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch s.Command {
		case "/iq-recommend":
			//params := &slack.Msg{Text: s.Text}
			//b, err := json.Marshal(params)
			// if err != nil {
			// 	w.WriteHeader(http.StatusInternalServerError)
			// 	log.Printf("[ERROR] %s\n", err)
			// 	return
			// }
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Looking up a recommendation for: " + s.Text)) //TODO: actually look something up

			//SlashCommand Struct = https://github.com/nlopes/slack/blob/master/slash.go
			sendMessage(s.Text, s.ChannelID)

		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unsupported slash command."))
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

//https://github.com/nlopes/slack/blob/master/chat.go -> https://api.slack.com/docs/formatting
func sendMessage(pText string, pChannelID string) {
	slackAPI := slack.New(config.BotToken)

	channelID, timestamp, err := slackAPI.PostMessage(pChannelID, slack.MsgOptionText("Looking up a recommendation for: " + pText, false))
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
		return
	}
	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}