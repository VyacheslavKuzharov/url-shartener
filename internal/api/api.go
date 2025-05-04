package api

import (
	"github.com/VyacheslavKuzharov/url-shortener/config"
	"github.com/VyacheslavKuzharov/url-shortener/internal/api/handlers"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(mux *chi.Mux, repo repository.Interface, cfg *config.Config) {
	mux.Post(`/`, handlers.SaveLink(repo.ShortLink(), cfg))
	mux.Get(`/{shortKey}`, handlers.Redirect(repo.ShortLink()))
}
