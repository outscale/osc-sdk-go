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
	Public IP example (aka Elastic IP)

- List existing Public IP
- Create an Public IP
- Query new Public IP details
- Delete Public IP
*/
func ExamplePublicIp() {
	client := osc.NewAPIClient(osc.NewConfiguration())
	auth := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	read, httpRes, err := client.PublicIpApi.ReadPublicIps(auth, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading public ips")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("We currently have", len(read.PublicIps), "public ips:")
	for _, pip := range read.PublicIps {
		println("-", pip.PublicIp)

	}

	println("Creating new Public Ip")
	creationOpts := osc.CreatePublicIpOpts{}
	creation, httpRes, err := client.PublicIpApi.CreatePublicIp(auth, &creationOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating Public IP ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}

	println("Created public ip", creation.PublicIp.PublicIp)

	println("Reading created public ip details")
	readOpts := osc.ReadPublicIpsOpts{
		ReadPublicIpsRequest: optional.NewInterface(
			osc.ReadPublicIpsRequest{
				Filters: osc.FiltersPublicIp{
					PublicIpIds: []string{creation.PublicIp.PublicIpId},
				},
			}),
	}
	read, httpRes, err = client.PublicIpApi.ReadPublicIps(auth, &readOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while reading public ip ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println(creation.PublicIp.PublicIpId, "details:")
	pip := read.PublicIps[0]
	println("- Id:", pip.PublicIpId)
	println("- Public IP address:", pip.PublicIp)

	println("Deleting public ip", creation.PublicIp.PublicIpId)
	deletionOpts := osc.DeletePublicIpOpts{
		DeletePublicIpRequest: optional.NewInterface(
			osc.DeletePublicIpRequest{
				PublicIpId: creation.PublicIp.PublicIpId,
			}),
	}
	_, httpRes, err = client.PublicIpApi.DeletePublicIp(auth, &deletionOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting public ip ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Deleted public ip", creation.PublicIp.PublicIpId)

	fmt.Println("Public IP journey is over")
	// Output: Public IP journey is over
}
