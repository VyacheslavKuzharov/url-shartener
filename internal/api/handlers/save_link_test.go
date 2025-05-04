package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/VyacheslavKuzharov/url-shortener/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveLink(t *testing.T) {
	shortKey := `qwerty`
	expectedBody := fmt.Sprintf("http://localhost:8080/%s", shortKey)
	originalURL := "https://practicum.yandex.ru/"
	cfg, _ := config.New()
	success := func(ctx context.Context, originalURL string) (string, error) { return shortKey, nil }

	testCases := []struct {
		name         string
		method       string
		reqBody      io.Reader
		expectedCode int
		expectedBody string
		repo         *MockShortLinkRepo
	}{
		{
			name:         "when happy path: correct response",
			method:       http.MethodPost,
			reqBody:      bytes.NewReader([]byte(originalURL)),
			expectedCode: http.StatusCreated,
			expectedBody: expectedBody,
			repo:         &MockShortLinkRepo{saveURL: success},
		},
		{
			name:         "when unhappy path: incorrect request method",
			method:       http.MethodGet,
			reqBody:      bytes.NewReader([]byte(originalURL)),
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Only POST requests allowed!\n",
			repo:         &MockShortLinkRepo{saveURL: success},
		},
		{
			name:         "when unhappy path: empty reqBody",
			method:       http.MethodPost,
			reqBody:      bytes.NewReader([]byte("")),
			expectedCode: http.StatusBadRequest,
			expectedBody: "URL parameter is missing\n",
			repo:         &MockShortLinkRepo{saveURL: success},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", tc.reqBody)
			w := httptest.NewRecorder()

			h := SaveLink(tc.repo, cfg)
			h(w, r)

			res := w.Result()
			// check response code
			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			// check response body
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedBody, string(resBody))
		})
	}
}
