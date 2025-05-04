package config

import (
	"flag"
	cfgURL "github.com/VyacheslavKuzharov/url-shortener/config/base_url"
	cfgHTTP "github.com/VyacheslavKuzharov/url-shortener/config/http"
	cfgLog "github.com/VyacheslavKuzharov/url-shortener/config/log"
)

type Config struct {
	HTTP    *cfgHTTP.HTTP
	BaseURL *cfgURL.BaseURL
	Log     *cfgLog.Log
}

func New() (*Config, error) {
	addr, baseURL := parseFlags()

	return &Config{
		HTTP:    addr.Build(),
		BaseURL: baseURL.Build(),
		Log:     cfgLog.New(),
	}, nil
}

func parseFlags() (*cfgHTTP.HTTP, *cfgURL.BaseURL) {
	addr := new(cfgHTTP.HTTP)
	url := new(cfgURL.BaseURL)

	flag.Var(addr, "a", "Net address host:port")
	flag.Var(url, "b", "base address of the resulting shortened URL http://localhost:3000")
	flag.Parse()

	return addr, url
}
