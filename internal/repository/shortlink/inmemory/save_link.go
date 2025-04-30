package inmemory

import (
	"context"
	"errors"
	"github.com/VyacheslavKuzharov/url-shortener/pkg/random"
)

func (s *MemStorage) SaveURL(ctx context.Context, originalURL string) (string, error) {
	if originalURL == "" {
		return "", errors.New("originalURL can't be blank")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	shortKey := random.GenShortKey()
	s.storage[shortKey] = originalURL

	return shortKey, nil
}
