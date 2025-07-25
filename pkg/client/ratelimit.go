package client

import (
	"net/http"

	"go.uber.org/ratelimit"
)

type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type ClientWithRateLimit struct {
	client HttpRequestDoer
	rl     ratelimit.Limiter
}

type ClientWithRateLimitOption (func(c *ClientWithRateLimit))

func WithRatelimit(limit int) ClientWithRateLimitOption {
	return func(c *ClientWithRateLimit) {
		c.rl = ratelimit.New(limit)
	}
}

func WithClient(client HttpRequestDoer) ClientWithRateLimitOption {
	return func(c *ClientWithRateLimit) {
		c.client = client
	}
}

func NewClientWithRateLimit(
	opts ...ClientWithRateLimitOption,
) ClientWithRateLimit {
	var c ClientWithRateLimit

	for _, opt := range opts {
		opt(&c)
	}

	if c.client == nil {
		c.client = http.DefaultClient
	}

	if c.rl == nil {
		c.rl = ratelimit.New(100)
	}

	return c
}

func (c ClientWithRateLimit) Do(req *http.Request) (*http.Response, error) {
	_ = c.rl.Take()
	return c.client.Do(req)
}
