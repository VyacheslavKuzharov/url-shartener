package http

import (
	"errors"
	"github.com/VyacheslavKuzharov/url-shortener/pkg/httpserver"
	"log"
	"os"
	"strings"
)

type HTTP struct {
	Host string
	Port string
}

func (a *HTTP) Build() *HTTP {
	if a.Host != "" && a.Port != "" {
		return a
	}

	cfg := &HTTP{}

	if os.Getenv("SERVER_ADDRESS") != "" {
		err := cfg.Set(os.Getenv("SERVER_ADDRESS"))

		if err != nil {
			log.Fatalf("SERVER_ADDRESS error: %s", err)
		}

		return cfg
	}

	cfg.Host = httpserver.DefaultHost
	cfg.Port = httpserver.DefaultPort

	return cfg
}

func (a *HTTP) String() string {
	return a.Host + ":" + a.Port
}

func (a *HTTP) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return errors.New("need address in a form host:port")
	}

	if hp[0] == "" {
		return errors.New("http Host can't be blank")
	}

	if hp[1] == "" {
		return errors.New("http Port can't be blank")
	}

	a.Host = hp[0]
	a.Port = hp[1]
	return nil
}
