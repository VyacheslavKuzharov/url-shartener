package handlers

import (
	"fmt"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository/shortlink"
	"io"
	"log"
	"net/http"
)

func SaveLink(repo shortlink.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Started ShortenUrl handler. Method %s", r.Method)

		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed!", http.StatusMethodNotAllowed)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		originalURL := string(b)
		if originalURL == "" {
			http.Error(w, "URL parameter is missing", http.StatusBadRequest)
			return
		}

		shortKey, err := repo.SaveURL(r.Context(), originalURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("link: %s successfuly savad with shortKey: %s", originalURL, shortKey)

		// Construct the full shortened URL
		shortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		if _, err = w.Write([]byte(shortenedURL)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
