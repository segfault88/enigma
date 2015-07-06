package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"

	"bitbucket.org/ckvist/twilio/twiml"
	"bitbucket.org/ckvist/twilio/twirest"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func basicTwiml(c web.C, w http.ResponseWriter, r *http.Request) {
	resp := twiml.NewResponse()
	resp.Action(twiml.Say{Text: "Hello Monkey"})
	resp.Action(twiml.Play{Url: "http://demo.twilio.com/hellomonkey/monkey.mp3"})
	resp.Action(twiml.Pause{Length: 1})
	err := resp.Dial(twiml.Conference{Name: "foobar"})
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.Send(w)
}

func main() {
	log.Info("Enigma main")

	accountSid := os.Getenv("TwilioSid")
	authToken := os.Getenv("TwilioToken")

	_ = twirest.NewClient(accountSid, authToken)

	goji.Get("/hello/:name", hello)
	goji.Get("/twiml", basicTwiml)
	goji.Serve()
}
