package examples_test

import (
	"testing"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Steps done in this test:
// 1. Read an available subregion.
// 2. Create a volume.
// 3. Add a tag to the volume.
// 4. Read the volume and validate its state and tag.
// 5. Delete the volume.
func TestVolume(t *testing.T) {
	ctx := t.Context()

	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := osc.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	// Read subregions
	subregions, err := client.ReadSubregions(ctx, osc.ReadSubregionsRequest{})
	require.NoError(t, err)
	require.NotNil(t, subregions.Subregions)
	assert.NotEmpty(t, *subregions.Subregions)

	subregionName := (*subregions.Subregions)[0].SubregionName
	require.NotNil(t, subregionName)
	assert.NotEmpty(t, *subregionName)

	size := 10
	tagKey := "Name"
	tagValue := "osc-sdk-go-test-" + RandomString(10)

	// Create a volume
	createReq := osc.CreateVolumeJSONRequestBody{
		Size:          &size,
		SubregionName: *subregionName,
	}

	createResp, err := client.CreateVolume(ctx, createReq, options.WithRetryTimeout(5*time.Minute))
	require.NoError(t, err)

	deleted := false
	defer func() {
		if deleted {
			return
		}

		if createResp.Volume == nil {
			return
		}

		_, _ = client.DeleteVolume(ctx, osc.DeleteVolumeJSONRequestBody{
			VolumeId: createResp.Volume.VolumeId,
		}, options.WithRetryTimeout(5*time.Minute))
	}()

	require.NotNil(t, createResp.Volume)

	volumeID := createResp.Volume.VolumeId
	assert.NotEmpty(t, volumeID)

	t.Logf("Created volume: %s in %s", volumeID, *subregionName)

	// Add a tag to the volume
	createTagsReq := osc.CreateTagsJSONRequestBody{
		ResourceIds: []string{volumeID},
		Tags: []osc.ResourceTag{
			{
				Key:   tagKey,
				Value: tagValue,
			},
		},
	}

	_, err = client.CreateTags(ctx, createTagsReq)
	require.NoError(t, err)

	t.Logf("Created volume tag %s=%s for: %s", tagKey, tagValue, volumeID)

	// Wait until the volume becomes available
	for range 30 {
		readResp, err := client.ReadVolumes(ctx, osc.ReadVolumesJSONRequestBody{
			Filters: &osc.FiltersVolume{
				VolumeIds: &[]string{volumeID},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, readResp.Volumes)
		require.Len(t, *readResp.Volumes, 1)

		volume := (*readResp.Volumes)[0]
		t.Logf("Volume %s state: %s", volume.VolumeId, volume.State)
		if volume.State == osc.VolumeStateAvailable {
			break
		}
		if volume.State == osc.VolumeStateError {
			t.Fatalf("volume %s entered error state", volume.VolumeId)
		}

		time.Sleep(10 * time.Second)
	}

	// Read the volume
	readReq := osc.ReadVolumesJSONRequestBody{
		Filters: &osc.FiltersVolume{
			VolumeIds: &[]string{volumeID},
		},
	}

	readResp, err := client.ReadVolumes(ctx, readReq)
	require.NoError(t, err)
	require.NotNil(t, readResp.Volumes)
	require.Len(t, *readResp.Volumes, 1)

	volume := (*readResp.Volumes)[0]
	assert.Equal(t, osc.VolumeStateAvailable, volume.State)

	foundTag := false
	for _, tag := range volume.Tags {
		if tag.Key == tagKey && tag.Value == tagValue {
			foundTag = true
			break
		}
	}
	assert.True(t, foundTag, "expected tag %q=%q on volume %s", tagKey, tagValue, volumeID)

	t.Logf("Successfully read volume: %s", volumeID)

	// Delete the volume
	deleteReq := osc.DeleteVolumeJSONRequestBody{
		VolumeId: volumeID,
	}

	deleteResp, err := client.DeleteVolume(ctx, deleteReq, options.WithRetryTimeout(5*time.Minute))
	require.NoError(t, err)
	deleted = true
	require.NotNil(t, deleteResp.ResponseContext)
	assert.NotNil(t, deleteResp.ResponseContext.RequestId)

	t.Logf("Successfully deleted volume: %s", volumeID)
}
