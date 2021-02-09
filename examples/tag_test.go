/*
  BSD 3-Clause License

  Copyright (c) 2021, Outscale SAS
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:

  * Redistributions of source code must retain the above copyright notice, this
    list of conditions and the following disclaimer.

  * Redistributions in binary form must reproduce the above copyright notice,
    this list of conditions and the following disclaimer in the documentation
    and/or other materials provided with the distribution.

  * Neither the name of the copyright holder nor the names of its
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
	"os"

	"github.com/antihax/optional"
	"github.com/outscale/osc-sdk-go/osc"
)

/* Example of adding tags on a resource (e.g. a virtual volume)
   1. create a 10GB volume
   2. add some tags on newly created volume
   3. read all tags from volume
   4. delete some tags on volume
   5. delete volume
*/
func ExampleTag() {
	client := osc.NewAPIClient(osc.NewConfiguration())
	auth := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	println("Creating 10GB GP2 volume")
	volumeCreationOpts := osc.CreateVolumeOpts{
		CreateVolumeRequest: optional.NewInterface(
			osc.CreateVolumeRequest{
				Size:          10,
				VolumeType:    "gp2",
				SubregionName: "eu-west-2a",
			}),
	}
	volumeCreation, httpRes, err := client.VolumeApi.CreateVolume(auth, &volumeCreationOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating volume ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Created volume", volumeCreation.Volume.VolumeId)

	println("Adding tag \"project:foo\" and \"env:prod\" to volume")
	tagCreationOpts := osc.CreateTagsOpts{
		CreateTagsRequest: optional.NewInterface(
			osc.CreateTagsRequest{
				Tags: []osc.ResourceTag{
					osc.ResourceTag{Key: "project", Value: "foo"},
					osc.ResourceTag{Key: "env", Value: "prod"},
				},
				ResourceIds: []string{volumeCreation.Volume.VolumeId},
			},
		),
	}
	_, httpRes, err = client.TagApi.CreateTags(auth, &tagCreationOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating tags ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Tags added")

	println("Reading all tags from volume")
	readOpts := osc.ReadTagsOpts{
		ReadTagsRequest: optional.NewInterface(
			osc.ReadTagsRequest{
				Filters: osc.FiltersTag{
					ResourceIds: []string{volumeCreation.Volume.VolumeId},
				},
			}),
	}
	read, httpRes, err := client.TagApi.ReadTags(auth, &readOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while reading tags ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("all tags on volume", volumeCreation.Volume.VolumeId, ":")
	for _, tag := range read.Tags {
		println("-", tag.Key, ":", tag.Value)
	}

	println("Removing tag \"env:prod\" from volume")
	tagDeletionOpts := osc.DeleteTagsOpts{
		DeleteTagsRequest: optional.NewInterface(
			osc.DeleteTagsRequest{
				ResourceIds: []string{volumeCreation.Volume.VolumeId},
				Tags: []osc.ResourceTag{
					osc.ResourceTag{Key: "env", Value: "prod"},
				},
			},
		),
	}
	_, httpRes, err = client.TagApi.DeleteTags(auth, &tagDeletionOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while removing tag ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Tag removed")

	println("Deleting volume", volumeCreation.Volume.VolumeId)
	deletionOpts := osc.DeleteVolumeOpts{
		DeleteVolumeRequest: optional.NewInterface(
			osc.DeleteVolumeRequest{
				VolumeId: volumeCreation.Volume.VolumeId,
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
	println("Deleted volume", volumeCreation.Volume.VolumeId)
	fmt.Println("Volume journey is over")
	// Output: Volume journey is over
}
