package ufc

import (
	"net/http"
)

type Config struct {
	HTTPClient *http.Client
	Version    int
}

func NewConfig(version int) *Config {
	return &Config{
		HTTPClient: http.DefaultClient,
		Version:    version,
	}
}
