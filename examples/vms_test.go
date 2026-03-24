package examples_test

import (
	"sort"
	"testing"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Steps done in this test:
// 1. Read the list of VMs.
// 2. Validate the returned VM collection.
// 3. Log each VM ID and image ID.
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

// Steps done in this test:
// 1. Read an available subregion.
// 2. Read a public Ubuntu image.
// 3. Create a VM.
// 4. Add a tag to the VM.
// 5. Wait for the VM to be running.
// 6. Read the VM and validate its state and tag.
// 7. Delete the VM.
func TestVm(t *testing.T) {
	ctx := t.Context()

	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := osc.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	subregions, err := client.ReadSubregions(ctx, osc.ReadSubregionsRequest{})
	require.NoError(t, err)
	require.NotNil(t, subregions.Subregions)
	assert.NotEmpty(t, *subregions.Subregions)

	subregionName := (*subregions.Subregions)[0].SubregionName
	require.NotNil(t, subregionName)
	assert.NotEmpty(t, *subregionName)

	resultsPerPage := 10
	publicImages := true
	availableImages := []osc.ImageState{osc.ImageStateAvailable}
	imageResp, err := client.ReadImages(ctx, osc.ReadImagesRequest{
		Filters: &osc.FiltersImage{
			AccountAliases:                      &[]string{"Outscale"},
			ImageNames:                          &[]string{"Ubuntu*"},
			PermissionsToLaunchGlobalPermission: &publicImages,
			States:                              &availableImages,
		},
		ResultsPerPage: &resultsPerPage,
	})
	require.NoError(t, err)
	require.NotNil(t, imageResp.Images)
	assert.NotEmpty(t, *imageResp.Images)

	images := append([]osc.Image(nil), (*imageResp.Images)...)
	sort.Slice(images, func(i, j int) bool {
		return images[i].CreationDate.Time.After(images[j].CreationDate.Time)
	})
	imageID := images[0].ImageId
	assert.NotEmpty(t, imageID)

	vmType := "tinav4.c1r1p2"
	minVmsCount := 1
	maxVmsCount := 1
	tagKey := "Name"
	tagValue := "osc-sdk-go-test-tag-" + RandomString(10)

	createResp, err := client.CreateVms(ctx, osc.CreateVmsRequest{
		ImageId:     imageID,
		MinVmsCount: &minVmsCount,
		MaxVmsCount: &maxVmsCount,
		Placement: &osc.Placement{
			SubregionName: *subregionName,
			Tenancy:       "default",
		},
		VmType: &vmType,
	}, options.WithRetryTimeout(10*time.Minute))
	require.NoError(t, err)

	deleted := false
	defer func() {
		if deleted {
			return
		}

		if createResp.Vms == nil || len(*createResp.Vms) == 0 {
			return
		}

		_, _ = client.DeleteVms(ctx, osc.DeleteVmsRequest{
			VmIds: []string{(*createResp.Vms)[0].VmId},
		}, options.WithRetryTimeout(10*time.Minute))
	}()

	require.NotNil(t, createResp.Vms)
	require.Len(t, *createResp.Vms, 1)

	vmID := (*createResp.Vms)[0].VmId
	assert.NotEmpty(t, vmID)

	_, err = client.CreateTags(ctx, osc.CreateTagsJSONRequestBody{
		ResourceIds: []string{vmID},
		Tags: []osc.ResourceTag{
			{
				Key:   tagKey,
				Value: tagValue,
			},
		},
	})
	require.NoError(t, err)

	for range 36 {
		readResp, err := client.ReadVms(ctx, osc.ReadVmsRequest{
			Filters: &osc.FiltersVm{
				VmIds: &[]string{vmID},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, readResp.Vms)
		require.Len(t, *readResp.Vms, 1)

		vm := (*readResp.Vms)[0]
		t.Logf("VM %s state: %s", vmID, vm.State)
		if vm.State == osc.VmStateRunning {
			break
		}
		require.NotEqual(t, osc.VmStateTerminated, vm.State, "vm %s entered unexpected state %s", vmID, vm.State)
		require.NotEqual(t, osc.VmStateShuttingDown, vm.State, "vm %s entered unexpected state %s", vmID, vm.State)

		time.Sleep(10 * time.Second)
	}

	readResp, err := client.ReadVms(ctx, osc.ReadVmsRequest{
		Filters: &osc.FiltersVm{
			VmIds: &[]string{vmID},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, readResp.Vms)
	require.Len(t, *readResp.Vms, 1)

	vm := (*readResp.Vms)[0]
	assert.Equal(t, osc.VmStateRunning, vm.State)
	assert.Equal(t, imageID, vm.ImageId)
	assert.Contains(t, vm.Tags, osc.ResourceTag{
		Key:   tagKey,
		Value: tagValue,
	}, "expected tag %q=%q on VM %s", tagKey, tagValue, vmID)

	t.Logf("Successfully read VM: %s", vmID)

	deleteResp, err := client.DeleteVms(ctx, osc.DeleteVmsRequest{
		VmIds: []string{vmID},
	}, options.WithRetryTimeout(10*time.Minute))
	require.NoError(t, err)
	deleted = true
	require.NotNil(t, deleteResp.ResponseContext)
	assert.NotNil(t, deleteResp.ResponseContext.RequestId)

	t.Logf("Successfully deleted VM: %s", vmID)
}
