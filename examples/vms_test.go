package examples_test

import (
	"context"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/client"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
)

func TestReadVms(t *testing.T) {
	profile, err := client.NewProfileFromStrandardConfiguration("", "")
	if err != nil {
		panic(err)
	}
	profile.TlsSkipVerify = true

	client, err := client.NewOapiClient(
		client.WithProfile(profile),
		client.WithRetry(),
		client.WithRateLimit(),
	)
	if err != nil {
		panic(err)
	}

	read, err := client.ReadVms(context.TODO(), osc.ReadVmsRequest{Filters: nil})
	if err != nil {
		panic(err)
	}

	for i, vm := range *read.Vms {
		t.Logf("[%d] Id: %s; ImageId: %s", i, *vm.VmId, *vm.ImageId)
	}
}
