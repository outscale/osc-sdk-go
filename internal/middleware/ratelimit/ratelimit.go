// Package ratelimit // Package ratelimit provides a middleware for rate limiting HTTP requests.
// It uses the go.uber.org/ratelimit library to implement the rate limiting logic.
package ratelimit

import (
	"net/http"

	"go.uber.org/ratelimit"
)

type RatelimitMiddleware struct {
	ratelimit.Limiter
}

type innerRatelimit struct {
	inner http.RoundTripper
	rt    ratelimit.Limiter
}

func (rt *innerRatelimit) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = rt.rt.Take()
	return rt.inner.RoundTrip(req)
}

func (r *RatelimitMiddleware) Decorate(next http.RoundTripper) http.RoundTripper {
	return &innerRatelimit{
		inner: next,
		rt:    r.Limiter,
	}
}
