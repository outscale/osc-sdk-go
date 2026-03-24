package examples_test

import (
	"encoding/base64"
	"slices"
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
// 1. Read an available subregion.
// 2. Read a public Ubuntu image.
// 3. Create a VM to use as backend and tag it.
// 4. Wait for the VM to be running.
// 5. Create a load balancer and tag it.
// 6. Link the backend VM to the load balancer.
// 7. Read the load balancer and validate the backend registration.
// 8. Read the backend VM health and validate it is reported.
// 9. Delete the load balancer.
// 10. Delete the backend VM.
func TestLoadBalancerBackend(t *testing.T) {
	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := osc.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	ctx := t.Context()
	lbName := "osc-sdk-go-test-" + RandomString(10)
	lbTagKey := "Name"
	lbTagValue := "osc-sdk-go-test-tag-" + RandomString(10)
	vmTagKey := "Name"
	vmTagValue := "osc-sdk-go-test-tag-" + RandomString(10)

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

	userData := base64.StdEncoding.EncodeToString([]byte(`#!/bin/sh
set -eu
mkdir -p /tmp/go-sdk-lb
echo ok > /tmp/go-sdk-lb/index.html
nohup python3 -m http.server 80 --directory /tmp/go-sdk-lb >/tmp/go-sdk-lb/http.log 2>&1 &
`))

	vmType := "tinav4.c1r1p2"
	minVmsCount := 1
	maxVmsCount := 1
	vmResp, err := client.CreateVms(ctx, osc.CreateVmsRequest{
		ImageId:     imageID,
		MinVmsCount: &minVmsCount,
		MaxVmsCount: &maxVmsCount,
		Placement: &osc.Placement{
			SubregionName: *subregionName,
			Tenancy:       "default",
		},
		UserData: &userData,
		VmType:   &vmType,
	}, options.WithRetryTimeout(10*time.Minute))
	require.NoError(t, err)

	vmDeleted := false
	defer func() {
		if vmDeleted {
			return
		}

		if vmResp.Vms == nil || len(*vmResp.Vms) == 0 {
			return
		}

		_, _ = client.DeleteVms(ctx, osc.DeleteVmsRequest{
			VmIds: []string{(*vmResp.Vms)[0].VmId},
		}, options.WithRetryTimeout(10*time.Minute))
	}()

	require.NotNil(t, vmResp.Vms)
	require.Len(t, *vmResp.Vms, 1)

	vmID := (*vmResp.Vms)[0].VmId
	assert.NotEmpty(t, vmID)

	_, err = client.CreateTags(ctx, osc.CreateTagsJSONRequestBody{
		ResourceIds: []string{vmID},
		Tags: []osc.ResourceTag{
			{
				Key:   vmTagKey,
				Value: vmTagValue,
			},
		},
	})
	require.NoError(t, err)

	for range 36 {
		readVmsResp, err := client.ReadVms(ctx, osc.ReadVmsRequest{
			Filters: &osc.FiltersVm{
				VmIds: &[]string{vmID},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, readVmsResp.Vms)
		require.Len(t, *readVmsResp.Vms, 1)

		vm := (*readVmsResp.Vms)[0]
		t.Logf("VM %s state: %s", vmID, vm.State)
		if vm.State == osc.VmStateRunning {
			break
		}

		time.Sleep(10 * time.Second)
	}

	lbProtocol := "TCP"
	createLBResp, err := client.CreateLoadBalancer(ctx, osc.CreateLoadBalancerRequest{
		LoadBalancerName: lbName,
		Listeners: []osc.ListenerForCreation{
			{
				BackendPort:          80,
				LoadBalancerPort:     80,
				LoadBalancerProtocol: lbProtocol,
				BackendProtocol:      &lbProtocol,
			},
		},
		SubregionNames: &[]string{*subregionName},
		Tags: &[]osc.ResourceTag{
			{
				Key:   lbTagKey,
				Value: lbTagValue,
			},
		},
	}, options.WithRetryTimeout(5*time.Minute))
	require.NoError(t, err)

	lbDeleted := false
	defer func() {
		if lbDeleted {
			return
		}

		if createLBResp.LoadBalancer == nil {
			return
		}

		_, _ = client.DeleteLoadBalancer(ctx, osc.DeleteLoadBalancerRequest{
			LoadBalancerName: lbName,
		}, options.WithRetryTimeout(5*time.Minute))
	}()

	require.NotNil(t, createLBResp.LoadBalancer)

	_, err = client.LinkLoadBalancerBackendMachines(ctx, osc.LinkLoadBalancerBackendMachinesRequest{
		LoadBalancerName: lbName,
		BackendVmIds:     &[]string{vmID},
	}, options.WithRetryTimeout(5*time.Minute))
	require.NoError(t, err)

	readLBResp, err := client.ReadLoadBalancers(ctx, osc.ReadLoadBalancersRequest{
		Filters: &osc.FiltersLoadBalancer{
			LoadBalancerNames: &[]string{lbName},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, readLBResp.LoadBalancers)
	require.Len(t, *readLBResp.LoadBalancers, 1)

	lb := (*readLBResp.LoadBalancers)[0]
	foundBackend := false
	for _, backendVMID := range lb.BackendVmIds {
		if backendVMID == vmID {
			foundBackend = true
			break
		}
	}
	assert.True(t, foundBackend, "expected backend VM %s on load balancer %s", vmID, lbName)

	foundHealthEntry := false
	for range 18 {
		healthResp, err := client.ReadVmsHealth(ctx, osc.ReadVmsHealthRequest{
			LoadBalancerName: lbName,
			BackendVmIds:     &[]string{vmID},
		})
		require.NoError(t, err)
		if healthResp.BackendVmHealth != nil {
			foundHealthEntry = slices.ContainsFunc(*healthResp.BackendVmHealth, func(entry osc.BackendVmHealth) bool {
				return entry.VmId != nil && *entry.VmId == vmID
			})
		}
		if foundHealthEntry {
			break
		}

		time.Sleep(10 * time.Second)
	}
	assert.True(t, foundHealthEntry, "expected backend health entry for VM %s", vmID)

	t.Logf("Backend VM %s registered to load balancer %s", vmID, lbName)

	deleteLBResp, err := client.DeleteLoadBalancer(ctx, osc.DeleteLoadBalancerRequest{
		LoadBalancerName: lbName,
	}, options.WithRetryTimeout(5*time.Minute))
	require.NoError(t, err)
	lbDeleted = true
	require.NotNil(t, deleteLBResp.ResponseContext)
	require.NotNil(t, deleteLBResp.ResponseContext.RequestId)
	t.Logf("Successfully deleted Load Balancer: %s", lbName)

	deleteVMResp, err := client.DeleteVms(ctx, osc.DeleteVmsRequest{
		VmIds: []string{vmID},
	}, options.WithRetryTimeout(10*time.Minute))
	require.NoError(t, err)
	vmDeleted = true
	require.NotNil(t, deleteVMResp.ResponseContext)
	require.NotNil(t, deleteVMResp.ResponseContext.RequestId)
	t.Logf("Successfully deleted VM: %s", vmID)
}
