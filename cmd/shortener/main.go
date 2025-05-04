package main

import (
	"github.com/VyacheslavKuzharov/url-shortener/config"
	"github.com/VyacheslavKuzharov/url-shortener/internal/app"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
