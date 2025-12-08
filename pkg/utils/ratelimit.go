package utils

import (
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware"
	mrt "github.com/outscale/osc-sdk-go/v3/pkg/middleware/ratelimit"
	"go.uber.org/ratelimit"
)

func WithoutRatelimit() middleware.MiddlewareChainOption {
	return middleware.WithMiddleware(middleware.MiddlewareSlotRateLimit, nil)
}

func WithRatelimit(rate int, opts ...ratelimit.Option) middleware.MiddlewareChainOption {
	return middleware.WithMiddleware(
		middleware.MiddlewareSlotRateLimit,
		&mrt.RatelimitMiddleware{Limiter: ratelimit.New(rate, opts...)},
	)
}
