package helium

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// APIURL url for v1 api
	APIURL = "api.helium.io/v1"
	// BETAURL url for v1 beta api
	BETAURL = "api.helium.wtf/v1"
	// HTTPTimeout timeout for http requests
	HTTPTimeout = 60
)

// Client
type Client struct {
	client *http.Client
	URL    string
	Key    string
}

// Option is a configuration option
type Option func(*Client)

func DefaultClient() *Client {
	client := &http.Client{
		Timeout: time.Second * HTTPTimeout,
	}
	return &Client{
		client: client,
		URL: APIURL,
	}
}

// ClientWithptions creates a new client with options
func ClientWithOptions(opts ...Option) *Client {
	c := &Client{
		client: &http.Client{
			Timeout: time.Second * HTTPTimeout,
		},
		URL: APIURL,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithKey(key string) Option {
	return func(c *Client) {
		c.Key = key
	}
}

func WithURL(url string) Option {
	return func(c *Client) {
		c.URL = url
	}
}

// Request handles http requests
func (c *Client) Request(method string, path string, params map[string]string) ([]byte, error) {
	path = fmt.Sprintf("https://%s%s", c.URL, path)
	// Create request
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if len(c.Key) > 0 {
		req.Header.Add("key", c.Key)
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	// Fetch Request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return nil, fmt.Errorf("request returned %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}