package main

import (
	"context"
	"fmt"
	"os"

	"github.com/antihax/optional"
	"github.com/outscale-dev/osc-sdk-go/osc"
)

func main() {
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
		fmt.Printf("- (%s) %s\n", sg.SecurityGroupId, sg.SecurityGroupName)
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
		fmt.Printf("- IP protocol: %s, IP range: %s, ports: %d-%d\n", rule.IpProtocol, rule.IpRanges, rule.FromPortRange, rule.ToPortRange)
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
	println("Deleted volume", sgID)
}
