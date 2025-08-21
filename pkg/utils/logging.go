package utils

import (
	"github.com/outscale/osc-sdk-go/v3/internal/middleware"
	"github.com/outscale/osc-sdk-go/v3/internal/middleware/logging"
)

func WithoutLogging() middleware.MiddlewareChainOption {
	return middleware.WithMiddleware(middleware.MiddlewareSlotLogging, nil)
}

func WithLogging(logger logging.Logger) middleware.MiddlewareChainOption {
	return middleware.WithMiddleware(
		middleware.MiddlewareSlotLogging,
		&logging.LoggingMiddleware{Logger: logger},
	)
}
