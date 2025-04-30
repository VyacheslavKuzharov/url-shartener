package api

import (
	"github.com/VyacheslavKuzharov/url-shortener/internal/api/handlers"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, repo repository.Interface) {
	mux.HandleFunc(`/`, handlers.SaveLink(repo.ShortLink()))
	mux.HandleFunc(`/{shortKey}`, handlers.Redirect(repo.ShortLink()))
}
