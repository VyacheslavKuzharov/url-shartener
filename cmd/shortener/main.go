package main

import (
	"github.com/VyacheslavKuzharov/url-shortener/internal/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("app error: %s", err)
	}
}
