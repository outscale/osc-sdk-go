package utils

import (
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/middleware"
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware/retry"
)

func WithoutRetry() middleware.MiddlewareChainOption {
	return middleware.WithMiddleware(middleware.MiddlewareSlotRetry, nil)
}

func WithRetry(
	retryWaitMin, retryWaitMax *time.Duration,
	retryMax *int,
) middleware.MiddlewareChainOption {
	return middleware.WithMiddleware(
		middleware.MiddlewareSlotRetry,
		&retry.RetryMiddleware{
			RetryWaitMin: retryWaitMin,
			RetryWaitMax: retryWaitMax,
			RetryMax:     retryMax,
		},
	)
}
