package main

import (
	"io/ioutil"
	"github.com/nlopes/slack"
	"net/http"
	"log"
	"bytes"
	"strings"
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
			w.WriteHeader(http.StatusOK)

			rr := lookupRecommendation(s.Text)

			if (len(rr.Remediation.VersionChanges) == 0) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte("No recommendation found for: " + s.Text))
			} else {
				//SlashCommand Struct = https://github.com/nlopes/slack/blob/master/slash.go
				sendMessage(s.Text, 
					s.ChannelID, 
					rr.Remediation.VersionChanges[0].Type,
					rr.Remediation.VersionChanges[0].Data.Component.ComponentIdentifier.Coordinates.Version,
				    lookupAllVersions(s.Text))
			}

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
func sendMessage(pText string, pChannelID string, pRemediationType string, pRecommendedVersion string, pAllVersions []string) {
	slackAPI := slack.New(config.BotToken)

	attachment := slack.Attachment{
		Text: pRecommendedVersion,
		Title: pRemediationType + " Recommended Version:",
	}

	allVersAttachment := slack.Attachment{
		Text: strings.Join(pAllVersions, " | "),
		Title:    "All Versions",
	}

	//todo: send to the bot
	channelID, timestamp, err := slackAPI.PostMessage(pChannelID, 
													slack.MsgOptionText("Looking up a recommendation for: " + pText, false),
													slack.MsgOptionAttachments(allVersAttachment, attachment))
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
		return
	}
	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func lookupRecommendation(text string) RemediationResponse {
	httpClient := http.Client{}
	//url := "http://localhost:8060/iq/api/v2/organizations"
	
	url := "http://localhost:60359/iq/api/v2/components/remediation/application/e06a119c75d04d97b8d8c11b62719752"
	//url := "http://localhost:8060/iq/api/v2/components/remediation/application/e06a119c75d04d97b8d8c11b62719752"

	book := CI{ComponentIdentifier: 
		ComponentIdentifier{ Format: "maven",
							Coordinates: MavenCoordinate{ GroupID: "org.apache.logging.log4j",
															ArtifactID: text,
															Classifier: "",
															Extension : "jar",
															Version: "2.8",
														},
							},
	}

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(book.ToComponentIdentifierJSON()))
	r.SetBasicAuth("admin", "admin123")
	r.Header.Add("Content-Type", "application/json")

	orgs, err := httpClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(orgs.Body)
	if err != nil {
		log.Println("[ERROR] could not read body")
	}

	remediationResponse := FromRemediationResponseJSON(body)

	//log.Printf("[DEBUG] Orgs: %s\n", body)
	//log.Printf("[DEBUG] RR: %s\n", remediationResponse)

	return remediationResponse
}

func lookupAllVersions(text string) []string {
	httpClient := http.Client{}
	
	url := "http://localhost:60359/iq/api/v2/components/versions"
	//url := "http://localhost:8060/iq/api/v2/components/versions"

	book :=  
		ComponentIdentifier{ Format: "maven",
							Coordinates: MavenCoordinate{ GroupID: "org.apache.logging.log4j",
															ArtifactID: text,
															Classifier: "",
															Extension : "jar",
															Version: "2.8",
														},
	}

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(book.ToComponentIdentifierJSON()))
	r.SetBasicAuth("admin", "admin123")
	r.Header.Add("Content-Type", "application/json")

	orgs, err := httpClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(orgs.Body)
	if err != nil {
		log.Println("[ERROR] could not read body")
	}

	allVersions := FromAllVersionsJSON(body)

	//log.Printf("[DEBUG] Orgs: %s\n", body)
	//log.Printf("[DEBUG] RR: %s\n", allVersions)

	return allVersions
}
