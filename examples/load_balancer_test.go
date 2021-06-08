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
	"math/rand"
	"os"
	"time"

	osc "github.com/outscale/osc-sdk-go/v2"
)

/* Load balancer example
- List existing load balancers
- Create a load balancer
- Query load balancer details
- Delete load balancer
*/
func ExampleLoadBalancer() {
	config := osc.NewConfiguration()
	config.Debug = false
	client := osc.NewAPIClient(config)
	ctx := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})
	loadBalancerName := "OscSdkExample-" + RandomString(10)
	loadBalancerSubRegion := "eu-west-2a"
	read, httpRes, err := client.LoadBalancerApi.ReadLoadBalancers(ctx).ReadLoadBalancersRequest(osc.ReadLoadBalancersRequest{}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading load balancers")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}

	loadBalancers := *read.LoadBalancers
	if len(loadBalancers) == 0 {
		println("No load balancer is currently created")
	} else {
		println("We currently have", len(loadBalancers), "load balancers:")
		for _, lb := range loadBalancers {
			println("-", lb.LoadBalancerName)
		}
	}

	println("Creating a single load balancer")
	listener1 := osc.ListenerForCreation{
		BackendPort:          80,
		LoadBalancerPort:     80,
		LoadBalancerProtocol: "TCP",
	}

	createOpt := osc.CreateLoadBalancerRequest{
		LoadBalancerName: loadBalancerName,
		Listeners:        []osc.ListenerForCreation{listener1},
		SubregionNames:   &[]string{loadBalancerSubRegion},
	}
	creation, httpRes, err := client.LoadBalancerApi.CreateLoadBalancer(ctx).CreateLoadBalancerRequest(createOpt).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating load balancer ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Created load balancer", creation.LoadBalancer.LoadBalancerName)

	println("Reading created load balancer detail")
	readFilters := osc.FiltersLoadBalancer{
		LoadBalancerNames: &[]string{loadBalancerName},
	}
	readOpts := osc.ReadLoadBalancersRequest{Filters: &readFilters}
	read, httpRes, err = client.LoadBalancerApi.ReadLoadBalancers(ctx).ReadLoadBalancersRequest(readOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while reading load balancers ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	loadBalancers = *read.LoadBalancers
	for _, lb := range loadBalancers {
		println("- name:", *lb.LoadBalancerName, "DNS:", *lb.DnsName)
	}

	println("Deleting load balancer", loadBalancerName)
	deletionOpts := osc.DeleteLoadBalancerRequest{LoadBalancerName: loadBalancerName}
	_, httpRes, err = client.LoadBalancerApi.DeleteLoadBalancer(ctx).DeleteLoadBalancerRequest(deletionOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting load balancer ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)

			os.Exit(1)
		}
		println("Deleted load balancer", creation.LoadBalancer.LoadBalancerName)
	}
	fmt.Println("Load balancer journey is over")
	// Output: Load balancer journey is over
}

func RandomString(n int) string {
	rand.Seed(int64(time.Now().Nanosecond()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	var letters_len = len(letters)
	r := make([]rune, n)
	for i := range r {
		r[i] = letters[rand.Intn(letters_len)]
	}
	return string(r)
}
