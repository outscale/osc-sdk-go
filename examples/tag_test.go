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
	"os"

	osc "github.com/outscale/osc-sdk-go/v2"
)

/* Example of adding tags on a resource (e.g. a Net) */
func ExampleTag() {
	config := osc.NewConfiguration()
	client := osc.NewAPIClient(config)
	ctx := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	println("Creating a new Net")
	netCreationOpt := osc.CreateNetRequest{IpRange: "10.0.0.0/16"}
	create, httpRes, err := client.NetApi.CreateNet(ctx).CreateNetRequest(netCreationOpt).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating Net ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	netId := create.Net.NetId
	println("Created Net", *netId)

	println("Adding tags \"project:foo\" and \"env:prod\" to Net", *netId)
	tagCreationOpts := osc.CreateTagsRequest{
		ResourceIds: []string{*netId},
		Tags: []osc.ResourceTag{
			osc.ResourceTag{Key: "project", Value: "foo"},
			osc.ResourceTag{Key: "env", Value: "prod"},
		},
	}
	_, httpRes, err = client.TagApi.CreateTags(ctx).CreateTagsRequest(tagCreationOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating Tags ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Added tags \"project:foo\" and \"env:prod\" to Net", *netId)

	println("Removing tag \"env:prod\" from Net", *netId)
	tagDeletionOpts := osc.DeleteTagsRequest{
		ResourceIds: []string{*netId},
		Tags: []osc.ResourceTag{
			osc.ResourceTag{Key: "env", Value: "prod"},
		},
	}
	_, httpRes, err = client.TagApi.DeleteTags(ctx).DeleteTagsRequest(tagDeletionOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting Tag ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Removed tag \"env:prod\" from Net")

	println("Deleting Net", *netId)
	deletionOpts := osc.DeleteNetRequest{NetId: *netId}
	_, httpRes, err = client.NetApi.DeleteNet(ctx).DeleteNetRequest(deletionOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting net ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Deleted Net", *netId)
	fmt.Println("Tag journey is over")
	// Output: Tag journey is over
}
