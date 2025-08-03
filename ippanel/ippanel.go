package ippanel

import (
	"net/http"
	"net/url"
	"time"
)

const (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "1.0.1"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 30 * time.Second
)

type Ippanel struct {
	APIKey     string
	BaseURL    *url.URL
	HTTPClient *http.Client
}

func NewIPPanelClient(apiKey string, baseURL string) *Ippanel {
	u, _ := url.Parse(baseURL)
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
