package xigubo

import "net/url"

func NewClient(cfg *Config) (*Client, error) {
	addressURL, err := url.Parse(cfg.Address)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.JoinPath(addressURL.String(), cfg.BathPath)
	if err != nil {
		return nil, err
	}

	client := &Client{
		baseURL:     baseURL,
		accessToken: cfg.AccessToken,
	}

	client.Events = events{client: client}
	client.Webhooks = webhooks{client: client}
	client.Targets = targets{client: client}

	return client, nil
}

type Client struct {
	baseURL     string
	accessToken string
	Webhooks    Webhooks
	Events      Events
	Targets     Targets
}
