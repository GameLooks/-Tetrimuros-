
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

// Every coincap response has a top level entry called data
// and a unix timestamp in milliseconds
type coincapResp struct {
	Data      *json.RawMessage `json:"data"`
	Timestamp *Timestamp       `json:"timestamp"`
}

// fetchAndParse returns the json below the top level "data" key
// returned by the coincap api
func (c *Client) fetchAndParse(req *http.Request) (*coincapResp, error) {
	// add the gzip compression header
	req.Header.Add("Accept-Encoding", "gzip")

	// make request to the api and read the response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check if the server sent compressed data
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		// if the content encoding was gzip instantiate a new gzip reader
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
	default:
		// otherwise set the reader to the response body