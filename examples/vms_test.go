package examples_test

import (
	"context"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
)

func TestReadVms(t *testing.T) {
	userProfile, err := profile.NewProfileFromStrandardConfiguration("", "")
	if err != nil {
		panic(err)
	}

	client, err := osc.NewClient(userProfile)
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
