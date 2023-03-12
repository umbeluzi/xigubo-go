package xigubo

import "net/http"

func DefaultConfig() *Config {
	return &Config{}
}

type Config struct {
	Address     string
	BathPath    string
	AccessToken string
	HTTPClient  *http.Client
}
