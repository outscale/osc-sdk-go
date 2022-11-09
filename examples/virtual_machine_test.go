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

/* Virtual Machine example
- List virtual machines
- Create a new virtual machine (public cloud, no VPC)
- Show newly created virtual machine details
- Delete virtual machine

Note that to access a virtual machine, you will also need at least to provide a security group with appropriate rules (e.g. TCP port 22) and a keypair at creation.
*/

func ExampleVm() {
	config := osc.NewConfiguration()
	config.Debug = false
	client := osc.NewAPIClient(config)
	ctx := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})
	read, httpRes, err := client.VmApi.ReadVms(ctx).ReadVmsRequest(osc.ReadVmsRequest{}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading vms ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}

	Vms := *read.Vms
	if len(Vms) == 0 {
		println("No vm is currently created")
	} else {
		println("We currently have", len(Vms), "virtual machines:")
		for _, vm := range Vms {
			println("-", *vm.VmId)
		}
	}

	println("Creating a single vm")
	imageId := "ami-68ed4301"
	vmType := "tinav4.c1r1p1"
	createOpt := osc.CreateVmsRequest{
		ImageId: imageId,
		VmType:  &vmType,
	}
	creation, httpRes, err := client.VmApi.CreateVms(ctx).CreateVmsRequest(createOpt).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating vm ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	vmID := *(*creation.Vms)[0].VmId
	println("Created vm", vmID)

	println("Reading created vm details")
	readFilters := osc.FiltersVm{
		VmIds: &[]string{vmID},
	}
	readOpts := osc.ReadVmsRequest{Filters: &readFilters}
	read, httpRes, err = client.VmApi.ReadVms(ctx).ReadVmsRequest(readOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while reading vms ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	Vms = *read.Vms
	for _, vm := range Vms {
		println("- id:", *vm.VmId, "type:", *vm.VmType, "state:", *vm.State)
	}

	println("Deleting vm", vmID)
	deletionOpts := osc.DeleteVmsRequest{VmIds: []string{vmID}}
	_, httpRes, err = client.VmApi.DeleteVms(ctx).DeleteVmsRequest(deletionOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting vm ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)

			os.Exit(1)
		}
		println("Deleted vm", vmID)
	}
	fmt.Println("vm journey is over")
	// Output: vm journey is over
}
