package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

var (
	WebhookEP string = "https://hooks.slack.com/services/XXXXXXXXX/YYYYYYYYY/ZZZZZZZZZZZZZZZZZZZZZZZZ"
)

// write your code
func routine() {
	type Payload struct {
		Channel   string `json:"channel"`
		Username  string `json:"username"`
		IconEmoji string `json:"icon_emoji`
		Text      string `json:"text"`
	}

	p := Payload{
		Channel:   "#test",
		Username:  "DashBot",
		IconEmoji: ":runner:",
		Text:      "DashButton is pushed",
	}

	b, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	values := url.Values{}
	values.Add("payload", string(b))

	_, err = http.PostForm(WebhookEP, values)
	if err != nil {
		log.Fatal(err)
	}
}
