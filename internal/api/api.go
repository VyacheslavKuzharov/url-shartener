package api

import (
	"github.com/VyacheslavKuzharov/url-shortener/config"
	"github.com/VyacheslavKuzharov/url-shortener/internal/api/handlers"
	"github.com/VyacheslavKuzharov/url-shortener/internal/api/middleware"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

func RegisterRoutes(mux *chi.Mux, repo repository.Interface, cfg *config.Config, log zerolog.Logger) {
	mux.Use(middleware.Logger(log))
	mux.Post(`/`, handlers.SaveLink(repo.ShortLink(), cfg, log))
	mux.Get(`/{shortKey}`, handlers.Redirect(repo.ShortLink(), log))
}
