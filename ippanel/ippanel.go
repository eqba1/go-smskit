package ippanel

import (
	"net/http"
	"net/url"
	"time"
)

var (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "1.0.1"

	Endpoint = "https://edge.ippanel.com/v1/api"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 30 * time.Second
)

type Ippanel struct {
	APIKey     string
	BaseURL    *url.URL
	HTTPClient *http.Client
}

func New(apiKey string, baseURL ...string) *Ippanel {
	if len(baseURL) > 0 && baseURL[0] != "" {
		Endpoint = baseURL[0]
	}

	u, _ := url.Parse(Endpoint)

	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   httpClientTimeout,
	}

	return &Ippanel{
		APIKey:     apiKey,
		BaseURL:    u,
		HTTPClient: client,
	}
}
