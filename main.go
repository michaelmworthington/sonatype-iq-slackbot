package main

import (
	"os"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/caarlos0/env"
)

//global vars: https://stackoverflow.com/questions/9539633/global-variables-get-command-line-argument-and-print-it
var config slackVars

type slackVars struct {
	VerificationToken string   `env:"VERIFICATION_TOKEN"`
	BotToken          string   `env:"BOT_TOKEN"`
}

func main() {
	//https://www.lynda.com/Go-tutorials/Basic-Docker-workflow-Docker-commands/672415/682819-4.html
	//flag.StringVar(&verificationToken, "token", "YOUR_VERIFICATION_TOKEN_HERE", "Your Slash Verification Token")
	//flag.Parse()
	
	//Load Environment Variables
	//https://godoc.org/github.com/joho/godotenv
	//https://www.mycodesmells.com/post/loading-config-from-env-in-go
	if err := godotenv.Load(); err != nil {
        log.Println("File .env not found, reading configuration from ENV")
	}

    if err := env.Parse(&config); err != nil {
        log.Fatalln("Failed to parse ENV")
	}

	//the router, call the index function when the request comes to /
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/iq-recommend", iqRecommend)
	http.HandleFunc("/api/slack-event-subscriptions", slackEventSubscriptions)

	//Logging
	log.Println("[INFO] Server listening on port: " + port())
	//Open a webserver on this port
	log.Fatal(http.ListenAndServe(port(), nil))

}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}