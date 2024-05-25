package monobank

import "net/http"

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// NewClient initializes and returns a new Monobank API client
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    baseURL,
		apiKey:     apiKey,
	}
}
