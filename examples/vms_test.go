package examples_test

import (
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/stretchr/testify/require"
)

// Steps done in this test:
// 1. Read the list of VMs.
// 2. Validate the returned VM collection.
// 3. Log each VM ID and image ID.
func TestReadVms(t *testing.T) {
	client := newOSCClient(t)

	read, err := client.ReadVms(t.Context(), osc.ReadVmsRequest{Filters: nil})
	require.NoError(t, err)
	require.NotNil(t, read.Vms)

	for i, vm := range *read.Vms {
		t.Logf("[%d] Id: %s; ImageId: %s", i, vm.VmId, vm.ImageId)
	}
}
