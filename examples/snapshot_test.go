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
// 4. Wait for the volume to become available.
// 5. Create a snapshot from the volume.
// 6. Add a tag to the snapshot.
// 7. Wait for the snapshot to complete.
// 8. Read the snapshot and validate its data and tag.
// 9. Delete the snapshot.
// 10. Delete the volume.
func TestSnapshot(t *testing.T) {
	ctx := t.Context()

	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := osc.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	volumeTagKey := "Name"
	volumeTagValue := "osc-sdk-go-test-tag-" + RandomString(10)
	snapshotTagKey := "Name"
	snapshotTagValue := "osc-sdk-go-test-tag-" + RandomString(10)

	subregions, err := client.ReadSubregions(ctx, osc.ReadSubregionsRequest{})
	require.NoError(t, err)
	require.NotNil(t, subregions.Subregions)
	assert.NotEmpty(t, *subregions.Subregions)

	subregionName := (*subregions.Subregions)[0].SubregionName
	require.NotNil(t, subregionName)
	assert.NotEmpty(t, *subregionName)

	size := 10
	volumeCreateResp, err := client.CreateVolume(ctx, osc.CreateVolumeJSONRequestBody{
		Size:          &size,
		SubregionName: *subregionName,
	}, options.WithRetryTimeout(5*time.Minute))
	require.NoError(t, err)

	volumeDeleted := false
	defer func() {
		if volumeDeleted {
			return
		}

		if volumeCreateResp.Volume == nil {
			return
		}

		_, _ = client.DeleteVolume(ctx, osc.DeleteVolumeJSONRequestBody{
			VolumeId: volumeCreateResp.Volume.VolumeId,
		}, options.WithRetryTimeout(5*time.Minute))
	}()

	require.NotNil(t, volumeCreateResp.Volume)

	volumeID := volumeCreateResp.Volume.VolumeId
	assert.NotEmpty(t, volumeID)

	_, err = client.CreateTags(ctx, osc.CreateTagsJSONRequestBody{
		ResourceIds: []string{volumeID},
		Tags: []osc.ResourceTag{
			{
				Key:   volumeTagKey,
				Value: volumeTagValue,
			},
		},
	})
	require.NoError(t, err)

	for range 30 {
		readVolumesResp, err := client.ReadVolumes(ctx, osc.ReadVolumesJSONRequestBody{
			Filters: &osc.FiltersVolume{
				VolumeIds: &[]string{volumeID},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, readVolumesResp.Volumes)
		require.Len(t, *readVolumesResp.Volumes, 1)

		volume := (*readVolumesResp.Volumes)[0]
		if volume.State == osc.VolumeStateAvailable {
			break
		}
		require.NotEqual(t, osc.VolumeStateError, volume.State, "volume %s entered error state", volumeID)

		time.Sleep(10 * time.Second)
	}

	t.Logf("Created volume: %s", volumeID)

	description := "osc-sdk-go-test-" + RandomString(10)
	createResp, err := client.CreateSnapshot(ctx, osc.CreateSnapshotJSONRequestBody{
		Description: &description,
		VolumeId:    &volumeID,
	}, options.WithRetryTimeout(10*time.Minute))
	require.NoError(t, err)

	snapshotDeleted := false
	defer func() {
		if snapshotDeleted {
			return
		}

		if createResp.Snapshot == nil {
			return
		}

		_, _ = client.DeleteSnapshot(ctx, osc.DeleteSnapshotJSONRequestBody{
			SnapshotId: createResp.Snapshot.SnapshotId,
		}, options.WithRetryTimeout(5*time.Minute))
	}()

	require.NotNil(t, createResp.Snapshot)

	snapshotID := createResp.Snapshot.SnapshotId
	assert.NotEmpty(t, snapshotID)

	_, err = client.CreateTags(ctx, osc.CreateTagsJSONRequestBody{
		ResourceIds: []string{snapshotID},
		Tags: []osc.ResourceTag{
			{
				Key:   snapshotTagKey,
				Value: snapshotTagValue,
			},
		},
	})
	require.NoError(t, err)

	for range 60 {
		readResp, err := client.ReadSnapshots(ctx, osc.ReadSnapshotsJSONRequestBody{
			Filters: &osc.FiltersSnapshot{
				SnapshotIds: &[]string{snapshotID},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, readResp.Snapshots)
		require.Len(t, *readResp.Snapshots, 1)

		snapshot := (*readResp.Snapshots)[0]
		t.Logf("Snapshot %s state: %s", snapshotID, snapshot.State)
		if snapshot.State == osc.SnapshotStateCompleted {
			break
		}
		require.NotEqual(t, osc.SnapshotStateError, snapshot.State, "snapshot %s entered error state", snapshotID)

		time.Sleep(10 * time.Second)
	}

	readResp, err := client.ReadSnapshots(ctx, osc.ReadSnapshotsJSONRequestBody{
		Filters: &osc.FiltersSnapshot{
			SnapshotIds: &[]string{snapshotID},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, readResp.Snapshots)
	require.Len(t, *readResp.Snapshots, 1)

	snapshot := (*readResp.Snapshots)[0]
	assert.Equal(t, snapshotID, snapshot.SnapshotId)
	assert.Equal(t, volumeID, snapshot.VolumeId)
	require.NotNil(t, snapshot.Description)
	assert.Equal(t, description, *snapshot.Description)
	require.NotNil(t, snapshot.Tags)
	assert.Contains(t, *snapshot.Tags, osc.ResourceTag{
		Key:   snapshotTagKey,
		Value: snapshotTagValue,
	}, "expected tag %q=%q on snapshot %s", snapshotTagKey, snapshotTagValue, snapshotID)

	t.Logf("Successfully read snapshot: %s", snapshotID)

	deleteResp, err := client.DeleteSnapshot(ctx, osc.DeleteSnapshotJSONRequestBody{
		SnapshotId: snapshotID,
	}, options.WithRetryTimeout(5*time.Minute))
	require.NoError(t, err)
	snapshotDeleted = true
	require.NotNil(t, deleteResp.ResponseContext)
	assert.NotNil(t, deleteResp.ResponseContext.RequestId)

	t.Logf("Successfully deleted snapshot: %s", snapshotID)

	volumeDeleteResp, err := client.DeleteVolume(ctx, osc.DeleteVolumeJSONRequestBody{
		VolumeId: volumeID,
	}, options.WithRetryTimeout(5*time.Minute))
	require.NoError(t, err)
	volumeDeleted = true
	require.NotNil(t, volumeDeleteResp.ResponseContext)
	assert.NotNil(t, volumeDeleteResp.ResponseContext.RequestId)

	t.Logf("Successfully deleted volume: %s", volumeID)
}
