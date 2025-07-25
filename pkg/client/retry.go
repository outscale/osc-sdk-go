package client

import (
	"net/http"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

type ClientWithRetry struct {
	*retryablehttp.Client
}

type ClientWithRetryOption (func(*ClientWithRetry))

func WithRetryWaitMin(t time.Duration) ClientWithRetryOption {
	return func(c *ClientWithRetry) {
		c.RetryWaitMin = t
	}
}

func WithRetryWaitMax(t time.Duration) ClientWithRetryOption {
	return func(c *ClientWithRetry) {
		c.RetryWaitMax = t
	}
}

func WithRetryMax(i int) ClientWithRetryOption {
	return func(c *ClientWithRetry) {
		c.RetryMax = i
	}
}

func NewClientWithRetry(opts ...ClientWithRetryOption) *ClientWithRetry {
	client := &ClientWithRetry{
		Client: retryablehttp.NewClient(),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c ClientWithRetry) Do(req *http.Request) (*http.Response, error) {
	r, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, err
	}

	return c.Client.Do(r)
}
