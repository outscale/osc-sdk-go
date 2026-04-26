package osc

import (
	"fmt"

	"github.com/google/uuid"
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
	s, err := userProfile.GetEndpoint(profile.OscServiceApi)
	if err != nil {
		return nil, err
	}

	if s[len(s)-1] != '/' {
		s += "/"
	}

	opts = append([]middleware.MiddlewareChainOption{
		middleware.FromProfile(userProfile, profile.OscServiceApi),
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

type clientTokenSetter interface {
	SetClientToken() error
}

func (r *CreateNatServiceRequest) SetClientToken() error {
	return setClientToken(&r.ClientToken)
}

func (r *CreateSnapshotRequest) SetClientToken() error {
	return setClientToken(&r.ClientToken)
}

func (r *CreateVmsRequest) SetClientToken() error {
	return setClientToken(&r.ClientToken)
}

func (r *CreateVolumeRequest) SetClientToken() error {
	return setClientToken(&r.ClientToken)
}

func setClientToken(clientToken **string) error {
	if clientToken == nil || *clientToken != nil {
		return nil
	}

	u, err := uuid.NewV7()
	if err != nil {
		return err
	}

	us := u.String()
	*clientToken = &us

	return nil
}
