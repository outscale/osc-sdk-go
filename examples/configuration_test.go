package examples_test

import (
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/client"
)

const TestEndpoint = "https://api.test.outscale.com/api/v1"

func TestProfile(t *testing.T) {
	profile := client.Profile{
		Region:   "not-existing-region",
		Protocol: "https",
	}

	ep, err := profile.GetEndpoint(client.OApi)
	if err != nil {
		panic(err)
	}
	if ep != "https://api.not-existing-region.outscale.com/api/v1" {
		panic(ep)
	}

	profile.Endpoints.API = TestEndpoint
	ep, err = profile.GetEndpoint(client.OApi)
	if err != nil {
		panic(err)
	}
	if ep != TestEndpoint {
		panic(ep)
	}
}
