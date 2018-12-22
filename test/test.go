package main

import (
	"log"
	"time"
)

var samara, utc *time.Location

func main() {
	samara, _ = time.LoadLocation("Europe/Samara")
	utc, _ = time.LoadLocation("UTC")

	log.Println(time.Now().In(utc))

}
