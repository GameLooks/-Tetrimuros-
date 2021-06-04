
// Package coincap provides a client for interacting with the CoinCap V2 API
package coincap

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var baseURL = "https://api.coincap.io/v2"

// Client is a rest client for the CoinCap V2 API
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient returns a new client for interacting with the CoinCap API
// If no httpClient is passed it will use http.DefaultClient
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

// SetBaseURL allows the setting of custom api base paths
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}