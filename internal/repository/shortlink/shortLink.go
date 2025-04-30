package shortlink

import (
	"context"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository/shortlink/infile"
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository/shortlink/inmemory"
	"log"
)

type Repo interface {
	SaveURL(ctx context.Context, originalURL string) (string, error)
	GetURL(ctx context.Context, key string) (string, error)
}

func NewRepo(kind string) Repo {
	switch kind {
	case "memo":
		return inmemory.NewStorage()
	case "file":
		return infile.NewStorage("path/to/file")
	default:
		log.Fatal("unknown storage type")
	}

	return nil
}
