package examples_test

import (
	"testing"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Steps done in this test:
// 1. Create a net.
// 2. Add a tag to the net.
// 3. Create a subnet in the net.
// 4. Add a tag to the subnet.
// 5. Read the subnet and validate its tag.
// 6. Update the subnet.
// 7. Delete the subnet.
// 8. Delete the net.
func TestNetAndSubnet(t *testing.T) {
	ctx := t.Context()

	client := newOSCClient(t)

	netTagKey := "Name"
	netTagValue := "osc-sdk-go-test-tag-" + RandomString(10)
	subnetTagKey := "Name"
	subnetTagValue := "osc-sdk-go-test-tag-" + RandomString(10)

	// Create a net first (subnets need a net)
	netCreateReq := osc.CreateNetJSONRequestBody{
		IpRange: "10.0.0.0/16",
	}

	netCreateResp, err := client.CreateNet(ctx, netCreateReq)
	require.NoError(t, err)

	netDeleted := false
	defer func() {
		if netDeleted {
			return
		}

		if netCreateResp.Net == nil {
			return
		}

		_, _ = client.DeleteNet(ctx, osc.DeleteNetJSONRequestBody{
			NetId: netCreateResp.Net.NetId,
		})
	}()

	require.NotNil(t, netCreateResp.Net)

	netID := netCreateResp.Net.NetId
	assert.NotEmpty(t, netID)

	_, err = client.CreateTags(ctx, osc.CreateTagsJSONRequestBody{
		ResourceIds: []string{netID},
		Tags: []osc.ResourceTag{
			{
				Key:   netTagKey,
				Value: netTagValue,
			},
		},
	})
	require.NoError(t, err)

	t.Logf("Created net: %s", netID)

	// Small sleep to ensure net is properly propagated
	time.Sleep(2 * time.Second)

	// Create a subnet
	createReq := osc.CreateSubnetJSONRequestBody{
		NetId:   netID,
		IpRange: "10.0.1.0/24",
	}

	createResp, err := client.CreateSubnet(ctx, createReq)
	require.NoError(t, err)

	subnetDeleted := false
	defer func() {
		if subnetDeleted {
			return
		}

		if createResp.Subnet == nil {
			return
		}

		_, _ = client.DeleteSubnet(ctx, osc.DeleteSubnetJSONRequestBody{
			SubnetId: createResp.Subnet.SubnetId,
		})
	}()

	require.NotNil(t, createResp.Subnet)

	subnetID := createResp.Subnet.SubnetId
	assert.NotEmpty(t, subnetID)

	_, err = client.CreateTags(ctx, osc.CreateTagsJSONRequestBody{
		ResourceIds: []string{subnetID},
		Tags: []osc.ResourceTag{
			{
				Key:   subnetTagKey,
				Value: subnetTagValue,
			},
		},
	})
	require.NoError(t, err)

	t.Logf("Created subnet: %s", subnetID)

	// Small sleep to ensure subnet is properly propagated
	time.Sleep(2 * time.Second)

	// Read the subnet
	readReq := osc.ReadSubnetsJSONRequestBody{
		Filters: &osc.FiltersSubnet{
			SubnetIds: &[]string{subnetID},
		},
	}

	readResp, err := client.ReadSubnets(ctx, readReq)
	require.NoError(t, err)
	require.NotNil(t, readResp.Subnets)
	require.Len(t, *readResp.Subnets, 1)
	assert.Equal(t, subnetID, (*readResp.Subnets)[0].SubnetId)
	assert.Contains(t, (*readResp.Subnets)[0].Tags, osc.ResourceTag{
		Key:   subnetTagKey,
		Value: subnetTagValue,
	}, "expected tag %q=%q on subnet %s", subnetTagKey, subnetTagValue, subnetID)

	t.Logf("Successfully read subnet: %s", subnetID)

	// Update the subnet
	mapPublicIPOnLaunch := false
	updateReq := osc.UpdateSubnetJSONRequestBody{
		SubnetId:            subnetID,
		MapPublicIpOnLaunch: mapPublicIPOnLaunch,
	}

	updateResp, err := client.UpdateSubnet(ctx, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updateResp.Subnet)

	t.Logf("Successfully updated subnet: %s", subnetID)

	// Delete the subnet
	deleteReq := osc.DeleteSubnetJSONRequestBody{
		SubnetId: subnetID,
	}

	deleteResp, err := client.DeleteSubnet(ctx, deleteReq)
	require.NoError(t, err)
	subnetDeleted = true
	require.NotNil(t, deleteResp.ResponseContext)
	assert.NotNil(t, deleteResp.ResponseContext.RequestId)

	t.Logf("Successfully deleted subnet: %s", subnetID)

	// Delete the net
	netDeleteReq := osc.DeleteNetJSONRequestBody{
		NetId: netID,
	}

	netDeleteResp, err := client.DeleteNet(ctx, netDeleteReq)
	require.NoError(t, err)
	netDeleted = true
	require.NotNil(t, netDeleteResp.ResponseContext)
	assert.NotNil(t, netDeleteResp.ResponseContext.RequestId)

	t.Logf("Successfully deleted net: %s", netID)
}
