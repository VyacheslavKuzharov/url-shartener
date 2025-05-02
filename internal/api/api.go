package api

import (
	"github.com/VyacheslavKuzharov/url-shortener/internal/api/handlers"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(mux *chi.Mux, repo repository.Interface) {
	mux.Post(`/`, handlers.SaveLink(repo.ShortLink()))
	mux.Get(`/{shortKey}`, handlers.Redirect(repo.ShortLink()))
}
