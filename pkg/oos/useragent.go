package oos

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/transport/http"
)

func WithUseragent(ua string) config.LoadOptionsFunc {
	return func(lo *config.LoadOptions) error {
		lo.APIOptions = append(lo.APIOptions, func(s *middleware.Stack) error {
			return s.Build.Add(userAgentSetter{ua: ua}, middleware.After)
		})
		return nil
	}
}

type userAgentSetter struct {
	ua string
}

func (s userAgentSetter) ID() string {
	return "uaSetter"
}

func (s userAgentSetter) HandleBuild(ctx context.Context, in middleware.BuildInput, next middleware.BuildHandler) (
	out middleware.BuildOutput, metadata middleware.Metadata, err error,
) {
	if req, ok := in.Request.(*http.Request); ok {
		req.Header.Set("User-Agent", s.ua)
	}
	return next.HandleBuild(ctx, in)
}

var _ middleware.BuildMiddleware = userAgentSetter{}
