package repository

import (
	"github.com/VyacheslavKuzharov/url-shortener/internal/repository/shortlink"
)

type Interface interface {
	ShortLink() shortlink.Repo
}

type Repo struct {
	shortLink shortlink.Repo
}

func New() Interface {
	return &Repo{
		shortLink: shortlink.NewRepo("memo"),
	}
}

func (r *Repo) ShortLink() shortlink.Repo {
	return r.shortLink
}
