package logger

import (
	"github.com/VyacheslavKuzharov/url-shortener/config/log"
	"github.com/rs/zerolog"
	"os"
)

func New(level log.Level) zerolog.Logger {
	var l zerolog.Level

	switch level {
	case log.ErrorLevel:
		l = zerolog.ErrorLevel
	case log.WarnLevel:
		l = zerolog.WarnLevel
	case log.InfoLevel:
		l = zerolog.InfoLevel
	default:
		l = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(l)

	loggerOutput := zerolog.ConsoleWriter{Out: os.Stderr}
	logger := zerolog.New(loggerOutput).With().Timestamp().Logger()

	return logger
}
