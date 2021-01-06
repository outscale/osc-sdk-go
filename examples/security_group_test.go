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

/* Security Groups Example
- List all security groups
- Create new Security Group (public cloud, outside VPC)
- Add rules in newly created Security Group
- Delete created rules
- Delete Security Group
*/
func ExampleSecurityGroup() {
	client := osc.NewAPIClient(osc.NewConfiguration())
	auth := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	read, httpRes, err := client.SecurityGroupApi.ReadSecurityGroups(auth, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading security groups")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("We currently have", len(read.SecurityGroups), "security groups:")
	for _, sg := range read.SecurityGroups {
		println(sg.SecurityGroupId, sg.SecurityGroupName)
	}

	println("Creating new security group")
	creationOpts := osc.CreateSecurityGroupOpts{
		CreateSecurityGroupRequest: optional.NewInterface(
			osc.CreateSecurityGroupRequest{
				SecurityGroupName: "osc-sdk-go-example",
				Description:       "osc-sdk-go security group example",
			}),
	}
	creation, httpRes, err := client.SecurityGroupApi.CreateSecurityGroup(auth, &creationOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating security group ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	sgID := creation.SecurityGroup.SecurityGroupId
	println("Created security group", sgID)

	println("Adding new security group rule to allow an IP range to access SSH (TCP port 22)")
	ruleAddOpts := osc.CreateSecurityGroupRuleOpts{
		CreateSecurityGroupRuleRequest: optional.NewInterface(
			osc.CreateSecurityGroupRuleRequest{
				SecurityGroupId: sgID,
				Flow:            "Inbound",
				IpRange:         "10.0.0.0/24",
				IpProtocol:      "tcp",
				FromPortRange:   22,
				ToPortRange:     22,
			}),
	}
	_, httpRes, err = client.SecurityGroupRuleApi.CreateSecurityGroupRule(auth, &ruleAddOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating security group rule")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Security group rule added")

	println("Listing security group rules")
	readOpts := osc.ReadSecurityGroupsOpts{
		ReadSecurityGroupsRequest: optional.NewInterface(
			osc.ReadSecurityGroupsRequest{
				Filters: osc.FiltersSecurityGroup{
					SecurityGroupIds: []string{sgID},
				},
			}),
	}
	read, httpRes, err = client.SecurityGroupApi.ReadSecurityGroups(auth, &readOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading security groups")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Inbound rules:")
	for _, rule := range read.SecurityGroups[0].InboundRules {
		println("- IP protocol:", rule.IpProtocol, ", IPs:", len(rule.IpRanges), ", ports:", rule.FromPortRange, "-", rule.ToPortRange)
	}

	ruleDelOpts := osc.DeleteSecurityGroupRuleOpts{
		DeleteSecurityGroupRuleRequest: optional.NewInterface(
			osc.DeleteSecurityGroupRuleRequest{
				SecurityGroupId: sgID,
				Flow:            "Inbound",
				IpRange:         "10.0.0.0/24",
				IpProtocol:      "tcp",
				FromPortRange:   22,
				ToPortRange:     22,
			}),
	}
	_, httpRes, err = client.SecurityGroupRuleApi.DeleteSecurityGroupRule(auth, &ruleDelOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting security group rule")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println(("Deleted security rule"))

	println("Deleting security group", sgID)
	deletionOpts := osc.DeleteSecurityGroupOpts{
		DeleteSecurityGroupRequest: optional.NewInterface(
			osc.DeleteSecurityGroupRequest{
				SecurityGroupId: sgID,
			}),
	}
	_, httpRes, err = client.SecurityGroupApi.DeleteSecurityGroup(auth, &deletionOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting security group ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Deleted security group", sgID)
	fmt.Println("Security group journey is over")
	// Output: Security group journey is over
}
