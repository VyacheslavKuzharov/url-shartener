package app

import (
	"github.com/VyacheslavKuzharov/url-shortener/config"
	"github.com/VyacheslavKuzharov/url-shortener/internal/api"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository"
	"github.com/VyacheslavKuzharov/url-shortener/pkg/httpserver"
	"github.com/go-chi/chi/v5"
	"log"
	"net"
)

func Run(cfg *config.Config) {
	log.Println("Starting url-shortener application...")

	// Repo
	repo := repository.New()

	// HTTP
	mux := chi.NewRouter()
	api.RegisterRoutes(mux, repo, cfg)

	httpServer := httpserver.New(mux)
	log.Printf("Start HTTP server on: %s", net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port))
	httpServer.Start(cfg.HTTP.Host, cfg.HTTP.Port)
}
