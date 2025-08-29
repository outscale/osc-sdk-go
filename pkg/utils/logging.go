package utils

import (
	"github.com/outscale/osc-sdk-go/v3/internal/middleware"
	"github.com/outscale/osc-sdk-go/v3/pkg/logger"
)

func WithoutLogging() middleware.MiddlewareChainOption {
	return middleware.WithLogger(nil)
}

func WithLogging(logger logger.Logger) middleware.MiddlewareChainOption {
	return middleware.WithLogger(logger)
}
