// Package retry provides a middleware for http.RoundTripper that retries requests.
package retry

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type RetryMiddleware struct {
	RetryWaitMin *time.Duration // Minimum time to wait
	RetryWaitMax *time.Duration // Maximum time to wait
	RetryMax     *int           // Maximum number of retries
	RetryTimeout *time.Duration // Maximum time to retry for
}

type timeoutInjector struct {
	inner   http.RoundTripper
	timeout time.Duration
}

func (t *timeoutInjector) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(req.Context(), t.timeout)
	defer cancel()
	return t.inner.RoundTrip(req.WithContext(ctx))
}

func (r *RetryMiddleware) Decorate(next http.RoundTripper) http.RoundTripper {
	rc := retryablehttp.NewClient()
	rc.HTTPClient = &http.Client{
		Transport: next,
	}
	rc.Logger = nil

	if r.RetryWaitMin != nil {
		rc.RetryWaitMin = *r.RetryWaitMin
	}

	if r.RetryWaitMax != nil {
		rc.RetryWaitMax = *r.RetryWaitMax
	}

	if r.RetryMax != nil {
		rc.RetryMax = *r.RetryMax
	}

	var roundTripper http.RoundTripper
	roundTripper = &retryablehttp.RoundTripper{Client: rc}

	if r.RetryTimeout != nil {
		roundTripper = &timeoutInjector{
			inner:   roundTripper,
			timeout: *r.RetryTimeout,
		}
	}

	return roundTripper
}
