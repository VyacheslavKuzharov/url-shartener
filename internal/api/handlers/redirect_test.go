package handlers

import (
	"context"
	"errors"
	"github.com/VyacheslavKuzharov/url-shortener/config/log"
	"github.com/VyacheslavKuzharov/url-shortener/pkg/logger"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	l := logger.New(log.InfoLevel)

	originalURL := "https://practicum.yandex.ru/"
	success := func(ctx context.Context, key string) (string, error) { return originalURL, nil }
	fail := func(ctx context.Context, key string) (string, error) { return "", errors.New("short key not found") }

	testCases := []struct {
		name           string
		method         string
		request        string
		expectedCode   int
		expectedHeader string
		repo           *MockShortLinkRepo
	}{
		{
			name:           "when happy path: correct response",
			method:         http.MethodGet,
			request:        "/qwerty",
			expectedCode:   http.StatusTemporaryRedirect,
			expectedHeader: originalURL,
			repo:           &MockShortLinkRepo{getURL: success},
		},
		{
			name:           "when unhappy path: incorrect request method",
			method:         http.MethodPost,
			request:        "/qwerty",
			expectedCode:   http.StatusMethodNotAllowed,
			expectedHeader: "",
			repo:           &MockShortLinkRepo{getURL: success},
		},
		{
			name:           "when unhappy path: short key not found",
			method:         http.MethodGet,
			request:        "/qwerty",
			expectedCode:   http.StatusGone,
			expectedHeader: "",
			repo:           &MockShortLinkRepo{getURL: fail},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.request, nil)
			w := httptest.NewRecorder()

			h := Redirect(tc.repo, l)
			h(w, r)

			res := w.Result()
			defer res.Body.Close()

			// check response code
			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			// check response header
			assert.Equal(t, tc.expectedHeader, res.Header.Get("Location"))
		})
	}
}
