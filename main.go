package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/narqo/go-dice/dice"
)

func rollDice(w http.ResponseWriter, req *http.Request) {
	q, _ := url.ParseQuery(req.URL.RawQuery)
	dn, ok := q["dice"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dc, err := dice.Parse(dn[0])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(strconv.Itoa(dc.Roll())))
}

func main() {
	http.HandleFunc("/roll", rollDice)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Cannot lister:", err)
	}
}
