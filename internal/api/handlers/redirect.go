package handlers

import (
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository/shortlink"
	"log"
	"net/http"
)

func Redirect(repo shortlink.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Started Redirect handler. Method %s", r.Method)

		if r.Method != http.MethodGet {
			http.Error(w, "Only GET requests allowed!", http.StatusMethodNotAllowed)
			return
		}

		shortKey := r.PathValue("shortKey")

		originalURL, err := repo.GetURL(r.Context(), shortKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusGone)
			return
		}
		log.Printf("link: %s detected with shortKey: %s", originalURL, shortKey)

		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
