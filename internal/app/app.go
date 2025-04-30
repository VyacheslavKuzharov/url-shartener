package app

import (
	"github.com/VyacheslavKuzharov/url-shortener/internal/api"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository"
	"github.com/VyacheslavKuzharov/url-shortener/pkg/httpserver"
	"log"
	"net"
	"net/http"
)

func Run() error {
	log.Println("Starting url-shortener application...")

	// Repo
	repo := repository.New()

	// HTTP
	mux := http.NewServeMux()
	api.RegisterRoutes(mux, repo)

	httpServer := httpserver.New(mux)
	log.Printf("Start HTTP server on: %s", net.JoinHostPort(httpserver.DefaultHost, httpserver.DefaultPort))
	httpServer.Start(httpserver.DefaultHost, httpserver.DefaultPort)

	return nil
}
