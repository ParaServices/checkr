package checkr

import (
	"net/http"
	"net/url"
	"time"
)

const (
	defaultMaxIdleConnections = 10
	defaultRequestTimeOut     = 5
)

// Client ...
type Client struct {
	APIKey string
	client *http.Client
	// BaseURL ...
	BaseURL *url.URL
}

// createHTTPClient for connection re-use
func createHTTPClient(maxIdleConnections, requestTimeOut int) *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeOut) * time.Second,
	}

	return client
}

// ClientOpts ...
type ClientOpts struct {
	APIKey     string
	BaseURL    *url.URL
	HTTPClient struct {
		Transport struct {
			MaxIdleConnsPerHost int
		}
		TimeOut int
	}
}

const defaultBaseURL = "https://api.checkr.com/"

// NewClient ...
func NewClient(opts *ClientOpts) (*Client, error) {
	if opts.HTTPClient.Transport.MaxIdleConnsPerHost == 0 {
		opts.HTTPClient.Transport.MaxIdleConnsPerHost = defaultMaxIdleConnections
	}

	if opts.BaseURL == nil {
		u, err := url.ParseRequestURI(defaultBaseURL)
		if err != nil {
			return nil, err
		}
		opts.BaseURL = u
	}

	if opts.HTTPClient.TimeOut == 0 {
		opts.HTTPClient.TimeOut = defaultRequestTimeOut
	}

	return &Client{
		APIKey: opts.APIKey,
		client: createHTTPClient(
			opts.HTTPClient.Transport.MaxIdleConnsPerHost,
			opts.HTTPClient.TimeOut,
		),
		BaseURL: opts.BaseURL,
	}, nil
}
