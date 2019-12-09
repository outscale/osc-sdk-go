package main

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/outscale-dev/osc-sdk-go/osc"
	"os"
)

func main() {
	client := osc.NewAPIClient(osc.NewConfiguration())
	auth := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	read, httpRes, err := client.VolumeApi.ReadVolumes(auth, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading volumes")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("We currently have", len(read.Volumes), "volumes:")
	for _, volume := range read.Volumes {
		println(volume.VolumeId)
	}

	println("Creating 10GB GP2 volume")
	// TODO find a cleaner way to pass options
	creationOpts := osc.CreateVolumeOpts{
		CreateVolumeRequest: optional.NewInterface(
			osc.CreateVolumeRequest{
				Size:          10,
				VolumeType:    "gp2",
				SubregionName: "eu-west-2a",
			}),
	}
	creation, httpRes, err := client.VolumeApi.CreateVolume(auth, &creationOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while creating volume")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, "Error while creating volume:", httpRes.Status)
		}
		os.Exit(1)
	}
	println("Created volume", creation.Volume.VolumeId)

	println("Reading created volume details")
	// TODO find a cleaner way to pass options
	readOpts := osc.ReadVolumesOpts{
		ReadVolumesRequest: optional.NewInterface(
			osc.ReadVolumesRequest{
				Filters: osc.FiltersVolume{
					VolumeIds: []string{creation.Volume.VolumeId},
				},
			}),
	}
	read, httpRes, err = client.VolumeApi.ReadVolumes(auth, &readOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading volumes")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println(creation.Volume.VolumeId, "details:")
	volume := read.Volumes[0]
	println("Id:", volume.VolumeId)
	println("Size:", volume.Size)
	println("Type:", volume.VolumeType)
	println("State:", volume.State)

	println("Deleting volume", creation.Volume.VolumeId)
	// TODO find a cleaner way to pass options
	deletionOpts := osc.DeleteVolumeOpts{
		DeleteVolumeRequest: optional.NewInterface(
			osc.DeleteVolumeRequest{
				VolumeId: creation.Volume.VolumeId,
			}),
	}
	_, httpRes, err = client.VolumeApi.DeleteVolume(auth, &deletionOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while creating volume")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, "Error while creating volume:", httpRes.Status)
		}
		os.Exit(1)
	}
	println("Deleted volume", creation.Volume.VolumeId)
}
