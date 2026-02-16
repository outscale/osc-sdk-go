package oks

import (
	"fmt"

	"github.com/outscale/osc-sdk-go/v3/pkg/logger"
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware"
	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/outscale/osc-sdk-go/v3/pkg/version"
)

func newClientRaw(
	userProfile *profile.Profile,
	opts ...middleware.MiddlewareChainOption,
) (*ClientRaw, error) {
	s, err := userProfile.GetEndpoint(profile.OscServiceOKS)
	if err != nil {
		return nil, err
	}

	if s[len(s)-1] != '/' {
		s += "/"
	}

	opts = append([]middleware.MiddlewareChainOption{
		middleware.FromProfile(userProfile, profile.OscServiceOKS),
		options.WithRatelimit(5),
		options.WithRetry(nil, nil, nil),
		options.WithLogging(logger.Default()),
		options.WithUseragent(fmt.Sprintf("osc-sdk-go/%s", version.Version)),
	}, opts...)

	m, err := middleware.NewMiddlewareChain(opts...)
	if err != nil {
		return nil, err
	}

	return &ClientRaw{
		Server: s,
		Client: m,
	}, nil
}
