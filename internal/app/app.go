package app

import (
	"github.com/VyacheslavKuzharov/url-shortener/config"
	"github.com/VyacheslavKuzharov/url-shortener/internal/api"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository"
	"github.com/VyacheslavKuzharov/url-shortener/pkg/httpserver"
	"github.com/VyacheslavKuzharov/url-shortener/pkg/logger"
	"github.com/go-chi/chi/v5"
	"net"
)

func Run(cfg *config.Config) {
	// Logger
	log := logger.New(cfg.Log.Level)
	log.Info().Msg("Starting URLShortener application...")

	// Repo
	repo := repository.New()

	// HTTP
	mux := chi.NewRouter()
	api.RegisterRoutes(mux, repo, cfg, log)

	httpServer := httpserver.New(mux)
	log.Info().Msgf("Start HTTP server on: %s", net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port))
	httpServer.Start(cfg.HTTP.Host, cfg.HTTP.Port)
}
