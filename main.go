package main

import (
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/narqo/go-dice/dice"
	"github.com/narqo/go-dice/Godeps/_workspace/src/github.com/apex/log"
	logHandler "github.com/narqo/go-dice/Godeps/_workspace/src/github.com/apex/log/handlers/text"
)

var port string

func rollDice(w http.ResponseWriter, req *http.Request) {
	q, _ := url.ParseQuery(req.URL.RawQuery)
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
	log.SetHandler(logHandler.New(os.Stderr))
	log.SetLevel(log.DebugLevel)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT was not set")
	}

	http.HandleFunc("/roll", rollDice)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.WithError(err).Fatal("Cannot listen")
	}
}
