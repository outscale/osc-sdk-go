package examples_test

import (
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/outscale/osc-sdk-go/v3/pkg/utils"
)

func TestReadVms(t *testing.T) {
	userProfile, err := profile.New()
	if err != nil {
		panic(err)
	}

	client, err := osc.NewClient(userProfile, utils.WithLogging(&testingLogger{t}))
	if err != nil {
		panic(err)
	}

	read, err := client.ReadVms(t.Context(), osc.ReadVmsRequest{Filters: nil})
	if err != nil {
		panic(err)
	}

	for i, vm := range *read.Vms {
		t.Logf("[%d] Id: %s; ImageId: %s", i, vm.VmId, vm.ImageId)
	}
}
