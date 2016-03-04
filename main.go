package main

import (
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/narqo/go-dice/Godeps/_workspace/src/github.com/apex/log"
	logfmt "github.com/narqo/go-dice/Godeps/_workspace/src/github.com/apex/log/handlers/logfmt"
	"github.com/narqo/go-dice/dice"
)

var token string

func checkAuth(intoken string) bool {
	return intoken == token
}

func rollDice(w http.ResponseWriter, req *http.Request) {
	q, _ := url.ParseQuery(req.URL.RawQuery)

	intoken, ok := q["token"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !checkAuth(intoken[0]) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	dn, ok := q["text"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dc, err := dice.Parse(dn[0])
	if err != nil {
		log.WithError(err).Debug("parse")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(strconv.Itoa(dc.Roll())))
}

func main() {
	log.SetHandler(logfmt.New(os.Stderr))
	log.SetLevel(log.DebugLevel)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT was not set")
	}

	token = os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Info("$SLACK_TOKEN was not set")
	} else {
		log.WithFields(log.Fields{
			"token": token,
		}).Debug("token")
	}

	http.HandleFunc("/roll", rollDice)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.WithError(err).Fatal("Cannot listen")
	}
}
