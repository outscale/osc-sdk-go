package utils

import (
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware"
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware/useragent"
)

func WithUseragent(ua string) middleware.MiddlewareChainOption {
	return middleware.WithMiddleware(
		middleware.MiddlewareSlotUseragent,
		&useragent.UseragentMiddleware{Useragent: ua},
	)
}
