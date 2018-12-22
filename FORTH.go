package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	//	"google.golang.org/appengine"
)

var samara, utc *time.Location

const html = `<html><head>
	<title>FORTH/go</title>
	<link rel="stylesheet" type="text/css" href="/static/css.css">
	<link rel="manifest" href="/static/app.manifest">
	<link rel="icon" href="/static/lego.png" type="image/png">
	<meta name="theme-color" content="black">
	<meta name="viewport" content="width=device-width, initial-scale=1">
</head>`

const htlog = `
</div>
`

func main() {

	samara, _ = time.LoadLocation("Europe/Samara")
	utc, _ = time.LoadLocation("UTC")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", index)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func manifest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
{
	"short_name": "FORTH/go",
	"name": "FORTH/go",
	"icons": [ { "src": "/static/lego.png", "type": "image/png" } ],	
	"orientation": "portrait",
	"display": "standalone",
	"background_color": "black",
	"theme_color": "black"
}
	`)
}

const pad = `
<div>
<form id=cmd action="" method=post novalidate>
<textarea autocorrect="off" autofocus id=pad name="pad" rows="3"></textarea>
<input id=go type=submit value="GO">
</form>
</div>
`

func index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fmt.Fprintf(w, `
	%s
	<div id=log><h1>Превед медвед in <ul><li>%v</li><li>%v</li></ul></div>
	%s
	`, html, now.In(samara), now.In(utc), pad)
}
