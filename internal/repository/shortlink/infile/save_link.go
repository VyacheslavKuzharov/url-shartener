package infile

import (
	"context"
	"errors"
)

func (s *FileStorage) SaveURL(ctx context.Context, originalURL string) (string, error) {
	if originalURL == "" {
		return "", errors.New("originalURL can't be blank")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	return "", nil
}
