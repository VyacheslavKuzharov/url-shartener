package baseurl

import (
	"fmt"
	"github.com/VyacheslavKuzharov/url-shortener/pkg/httpserver"
	"net"
	"os"
)

type BaseURL struct {
	Addr string
}

func (b *BaseURL) Build() *BaseURL {
	if b.Addr != "" {
		return b
	}

	if os.Getenv("BASE_URL") != "" {
		return &BaseURL{
			Addr: os.Getenv("BASE_URL"),
		}
	}

	schema := "http"
	return &BaseURL{
		Addr: fmt.Sprintf("%s://%s", schema, net.JoinHostPort(httpserver.DefaultHost, httpserver.DefaultPort)),
	}
}

func (b *BaseURL) String() string {
	return b.Addr
}

func (b *BaseURL) Set(s string) error {
	b.Addr = s
	return nil
}
