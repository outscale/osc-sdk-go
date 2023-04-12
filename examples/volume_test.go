/*
BSD 3-Clause License

Copyright (c) 2021, Outscale SAS
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

  - Redistributions of source code must retain the above copyright notice, this
    list of conditions and the following disclaimer.

  - Redistributions in binary form must reproduce the above copyright notice,
    this list of conditions and the following disclaimer in the documentation
    and/or other materials provided with the distribution.

  - Neither the name of the copyright holder nor the names of its
    contributors may be used to endorse or promote products derived from
    this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package osc_test

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/outscale/osc-sdk-go/osc"
	"os"
)

/*
	Volumes Example

- List existing volumes
- Create a 10GB volume
- Query new volume details
- Delete new volume
*/
func ExampleVolume() {
	configEnv := osc.NewConfigEnv()
	config, err := configEnv.Configuration()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create configuration: %s\n", err.Error())
		os.Exit(1)
	}
	client := osc.NewAPIClient(config)
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
		println("-", volume.VolumeId)
	}

	println("Creating 10GB GP2 volume")
	creationOpts := osc.CreateVolumeOpts{
		CreateVolumeRequest: optional.NewInterface(
			osc.CreateVolumeRequest{
				Size:          10,
				VolumeType:    "gp2",
				SubregionName: fmt.Sprintf("%sa", *configEnv.Region),
			}),
	}
	creation, httpRes, err := client.VolumeApi.CreateVolume(auth, &creationOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating volume ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Created volume", creation.Volume.VolumeId)

	println("Reading created volume details")
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
		fmt.Fprint(os.Stderr, "Error while reading volumes ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println(creation.Volume.VolumeId, "details:")
	volume := read.Volumes[0]
	println("- Id:", volume.VolumeId)
	println("- Size:", volume.Size)
	println("- Type:", volume.VolumeType)
	println("- State:", volume.State)

	println("Deleting volume", creation.Volume.VolumeId)
	deletionOpts := osc.DeleteVolumeOpts{
		DeleteVolumeRequest: optional.NewInterface(
			osc.DeleteVolumeRequest{
				VolumeId: creation.Volume.VolumeId,
			}),
	}
	_, httpRes, err = client.VolumeApi.DeleteVolume(auth, &deletionOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting volume ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Deleted volume", creation.Volume.VolumeId)
	fmt.Println("Volume journey is over")
	// Output: Volume journey is over
}
