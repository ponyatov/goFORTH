package main

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

var samara, utc *time.Location

func main() {

	samara, _ = time.LoadLocation("Europe/Samara")
	utc, _ = time.LoadLocation("UTC")

	http.HandleFunc("/", index)
	appengine.Main()
}

func index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fmt.Fprintf(w, "<html><h1>Превед медвед in <ul><li>%v</li><li>%v</li></ul>", now.In(samara), now.In(utc))
}
