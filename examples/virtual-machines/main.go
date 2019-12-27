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

	read, httpRes, err := client.VmApi.ReadVms(auth, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading virtual machines ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("We currently have", len(read.Vms), "virtual machines:")
	for _, vm := range read.Vms {
		fmt.Printf("- id: %s, state: %s, type: %s\n", vm.VmId, vm.State, vm.VmType)
	}

	println("Creating new virtual machine")
	creationOpts := osc.CreateVmsOpts{
		CreateVmsRequest: optional.NewInterface(
			osc.CreateVmsRequest{
				ImageId:     "ami-b0d57010",
				Performance: "standard",
				VmType:      "tinav4.c1r1",
			}),
	}
	creation, httpRes, err := client.VmApi.CreateVms(auth, &creationOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating virtual machine ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
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
	read, httpRes, err = client.VmApi.ReadVms(auth, &readOpts)
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
	_, httpRes, err = client.VmApi.DeleteVms(auth, &deletionOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting virtual machine ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Virtual machine deleted")

}
