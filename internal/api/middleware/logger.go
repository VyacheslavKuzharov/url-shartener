package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func Logger(log zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log.Info().Msg("logger middleware started")

		logFn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				log.Info().
					Int("status", ww.Status()).
					Str("method", r.Method).
					Str("path", r.URL.Path).
					Str("uri", r.RequestURI).
					Str("query", r.URL.RawQuery).
					Dur("duration", time.Since(start)).
					Int("bytes", ww.BytesWritten()).
					Msg("request completed")
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(logFn)
	}
}
