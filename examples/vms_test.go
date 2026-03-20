package examples_test

import (
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/require"
)

func TestReadVms(t *testing.T) {
	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := osc.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	read, err := client.ReadVms(t.Context(), osc.ReadVmsRequest{Filters: nil})
	require.NoError(t, err)
	require.NotNil(t, read.Vms)

	for i, vm := range *read.Vms {
		t.Logf("[%d] Id: %s; ImageId: %s", i, vm.VmId, vm.ImageId)
	}
}
