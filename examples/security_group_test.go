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

/*
	Security Groups Example

- List all security groups
- Create new Security Group (public cloud, outside VPC)
- Add rules in newly created Security Group
- Delete created rules
- Delete Security Group
*/
func ExampleSecurityGroup() {
	config := osc.NewConfiguration()
	config.Debug = false
	client := osc.NewAPIClient(config)
	ctx := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})
	read, httpRes, err := client.SecurityGroupApi.ReadSecurityGroups(ctx).ReadSecurityGroupsRequest(osc.ReadSecurityGroupsRequest{}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading security groups")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status, httpRes.Body)
		}
		os.Exit(1)
	}

	securityGroups := *read.SecurityGroups
	if len(securityGroups) == 0 {
		println("No security group is currently created")
	} else {
		println("We currently have", len(securityGroups), "security groups:")
		for _, sg := range securityGroups {
			println("-", *sg.SecurityGroupId, *sg.SecurityGroupName)
		}
	}

	println("Creating a single security group")
	securityGroupName := "OscSdkExample-" + RandomString(10)
	createOpt := osc.CreateSecurityGroupRequest{
		SecurityGroupName: securityGroupName,
		Description:       "Security Group test with osc-sdk-go",
	}
	creation, httpRes, err := client.SecurityGroupApi.CreateSecurityGroup(ctx).CreateSecurityGroupRequest(createOpt).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating security group ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Created security group", creation.SecurityGroup.SecurityGroupName)

	println("Reading created security group detail")
	readFilters := osc.FiltersSecurityGroup{
		SecurityGroupNames: &[]string{securityGroupName},
	}
	readOpts := osc.ReadSecurityGroupsRequest{Filters: &readFilters}
	read, httpRes, err = client.SecurityGroupApi.ReadSecurityGroups(ctx).ReadSecurityGroupsRequest(readOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while reading security groups ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	securityGroups = *read.SecurityGroups
	for _, sg := range securityGroups {
		println("- id:", *sg.SecurityGroupId, "name:", *sg.SecurityGroupName, "description:", *sg.Description)
	}

	println("Creating a security group rule allowing port 22 from any source")
	flow := "Inbound"
	ipProtocol := "tcp"
	ipRange := "0.0.0.0/0" // BAD practice, avoid this in real applications
	var port int32 = 22
	ruleOpt := osc.CreateSecurityGroupRuleRequest{
		Flow:            flow,
		SecurityGroupId: *creation.SecurityGroup.SecurityGroupId,
		IpProtocol:      &ipProtocol,
		IpRange:         &ipRange,
		FromPortRange:   &port,
		ToPortRange:     &port,
	}
	_, httpRes, err = client.SecurityGroupRuleApi.CreateSecurityGroupRule(ctx).CreateSecurityGroupRuleRequest(ruleOpt).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating security group rule ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status, httpRes.Body)
		}
		os.Exit(1)
	}
	println("Created security group rule")

	println("Deleting security group rule")
	ruleDeletionOpts := osc.DeleteSecurityGroupRuleRequest{
		Flow:            flow,
		SecurityGroupId: *creation.SecurityGroup.SecurityGroupId,
		IpProtocol:      &ipProtocol,
		IpRange:         &ipRange,
		FromPortRange:   &port,
		ToPortRange:     &port,
	}
	_, httpRes, err = client.SecurityGroupRuleApi.DeleteSecurityGroupRule(ctx).DeleteSecurityGroupRuleRequest(ruleDeletionOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting security group rule ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)

			os.Exit(1)
		}
	}
	println("Deleted security group rule")

	println("Deleting security group", securityGroupName)
	deletionOpts := osc.DeleteSecurityGroupRequest{SecurityGroupName: &securityGroupName}
	_, httpRes, err = client.SecurityGroupApi.DeleteSecurityGroup(ctx).DeleteSecurityGroupRequest(deletionOpts).Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting security group ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)

			os.Exit(1)
		}
	}
	println("Deleted security group", securityGroupName)
	fmt.Println("security group journey is over")
	// Output: security group journey is over
}
