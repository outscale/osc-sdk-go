package osc

import (
	"fmt"

	"github.com/outscale/osc-sdk-go/v3/pkg/logger"
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/outscale/osc-sdk-go/v3/pkg/utils"
	"github.com/outscale/osc-sdk-go/v3/pkg/version"
)

func newClientRaw(
	userProfile *profile.Profile,
	opts ...middleware.MiddlewareChainOption,
) (*ClientRaw, error) {
	s, err := userProfile.GetEndpoint(profile.OscServiceApi)
	if err != nil {
		return nil, err
	}

	if s[len(s)-1] != '/' {
		s += "/"
	}

	opts = append([]middleware.MiddlewareChainOption{
		middleware.FromProfile(userProfile, profile.OscServiceApi),
		utils.WithRatelimit(5),
		utils.WithRetry(nil, nil, nil),
		utils.WithLogging(logger.Default()),
		utils.WithUseragent(fmt.Sprintf("osc-sdk-go/%s", version.SDKVersion)),
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
