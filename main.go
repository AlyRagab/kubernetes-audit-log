package main

import (
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/nlopes/slack"
)

func webhook(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Send the audit log message to Slack
	apiToken := os.Getenv("SLACK_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")
	sclient := slack.New(apiToken, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Pretext: "Kubernetes Audit Log",
		Color:   "#3EB489",
		Fields: []slack.AttachmentField{
			{
				Value: string(reqBody),
			},
		},
	}
	sclient.PostMessage(channelID, slack.MsgOptionAttachments(attachment))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/webhook", webhook).Methods("POST")
	log.Info("Starting Server on 0.0.0.0:8080")
	http.ListenAndServe(":8080", r)
}
