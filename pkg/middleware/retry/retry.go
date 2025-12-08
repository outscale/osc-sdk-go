// Package retry provides a middleware for http.RoundTripper that retries requests.
package retry

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type RetryMiddleware struct {
	RetryWaitMin *time.Duration // Minimum time to wait
	RetryWaitMax *time.Duration // Maximum time to wait
	RetryMax     *int           // Maximum number of retries
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

	return &retryablehttp.RoundTripper{Client: rc}
}
