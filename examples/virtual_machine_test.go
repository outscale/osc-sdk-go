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

	"github.com/antihax/optional"
	"github.com/outscale/osc-sdk-go/osc"
)

/*
	Virtual Machine example

- List virtual machines
- Create a new virtual machine (public cloud, no VPC)
- Show newly created virtual machine details
- Delete virtual machine

Note that to access a virtual machine, you will also need at least to provide a security group with appropriate rules (e.g. TCP port 22) and a keypair at creation.
*/
func ExampleVirtualMachine() {
	configEnv := osc.NewConfigEnv()
	config, err := configEnv.Configuration()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot load configuration from environement variables")
		os.Exit(1)
	}
	ctx, err := configEnv.Context(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot set context from environement variables")
		os.Exit(1)
	}
	client := osc.NewAPIClient(config)

	read, httpRes, err := client.VmApi.ReadVms(ctx, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading virtual machines ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("We currently have", len(read.Vms), "virtual machines:")
	for _, vm := range read.Vms {
		println("- id:", vm.VmId, "(", vm.State, ")")
	}

	println("Creating new virtual machine")
	creationOpts := osc.CreateVmsOpts{
		CreateVmsRequest: optional.NewInterface(
			osc.CreateVmsRequest{
				ImageId: os.Getenv("OMI_ID"),
				VmType:  "tinav4.c1r1p1",
			}),
	}
	creation, httpRes, err := client.VmApi.CreateVms(ctx, &creationOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating virtual machine ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status, creation)
		}
		os.Exit(1)
	}
	VMID := creation.Vms[0].VmId
	println("Created virtual machine", VMID)

	println("Query virtual machine details")
	readOpts := osc.ReadVmsOpts{
		ReadVmsRequest: optional.NewInterface(
			osc.ReadVmsRequest{
				Filters: osc.FiltersVm{
					VmIds: []string{VMID},
				},
			}),
	}
	read, httpRes, err = client.VmApi.ReadVms(ctx, &readOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading virtual machines")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	vm := read.Vms[0]
	println("Virtual machine details:")
	println("- id:", vm.VmId)
	println("- type:", vm.VmType)
	println("- state:", vm.State)
	println("- image:", vm.ImageId)
	println("- performance:", vm.Performance)
	println("- private DNS:", vm.PrivateDnsName)
	println("- private IP:", vm.PrivateIp)
	println("- ...")

	println("Deleting virtual machine", VMID)
	deletionOpts := osc.DeleteVmsOpts{
		DeleteVmsRequest: optional.NewInterface(
			osc.DeleteVmsRequest{
				VmIds: []string{VMID},
			}),
	}
	_, httpRes, err = client.VmApi.DeleteVms(ctx, &deletionOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting virtual machine ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Virtual machine deleted")
	fmt.Println("Virtual machine journey is over")
	// Output: Virtual machine journey is over
}
