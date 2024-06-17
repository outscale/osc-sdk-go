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
	"math/rand"
	"os"
	"time"

	osc "github.com/outscale/osc-sdk-go/v2"
)

/*
	Teardown example

- Delete VMs
- Delete Load Balancers
- Unlink and delete non main Route Tables
- Unlink Virtual Gateways
- Delete NATs
- Delete Net Access Points
- Delete/reject Net Peerings
- Delete NICs
- Delete non default Security Groups
- Unlink and delete Internet Services
- Delete linked Public IPs
- Delete Subnets
- Delete Net
*/
func ExampleTeardown() {
	configEnv := osc.NewConfigEnv()
	config, err := configEnv.Configuration()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot init configuration from env variables")
		os.Exit(1)
	}
	ctx, err := configEnv.Context(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot init context from env variables")
		os.Exit(1)
	}
	client := osc.NewAPIClient(config)

	net, httpRes, err := client.NetApi.CreateNet(ctx).CreateNetRequest(*osc.NewCreateNetRequest("10.0.0.0/16")).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating a Net")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status, httpRes.Body)
		}
		os.Exit(1)
	}
	netID := *net.Net.NetId
	_, httpRes, err = client.SubnetApi.CreateSubnet(ctx).CreateSubnetRequest(*osc.NewCreateSubnetRequest("10.0.1.0/29", netID)).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating a subnet")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status, httpRes.Body)
		}
		os.Exit(1)
	}

	Teardown(netID, client, ctx)

	filterNets := osc.FiltersNet{NetIds: &[]string{netID}}
	netsCB, httpRes, err := client.NetApi.ReadNets(ctx).ReadNetsRequest(osc.ReadNetsRequest{Filters: &filterNets}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the Net state")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status, httpRes.Body)
		}
		os.Exit(1)
	}
	if netsCB.Nets != nil {
		if *(*netsCB.Nets)[0].State != "deleting" {
			fmt.Fprintln(os.Stderr, "Teardown failed")
			os.Exit(1)
		}
	}
	// Output: Teardown completed
}

func Teardown(netID string, client *osc.APIClient, ctx context.Context) error {
	cleared := false
	var publicIps []string
	subnetsIDs, err := getSubnetsIDs(netID, client, ctx)
	if err != nil {
		return err
	}

	// Delete VMs
	vms, err := getVMs(netID, client, ctx)
	if err != nil {
		return err
	}
	var vmIds []string
	for _, vm := range *vms.Vms {
		if vm.PublicIp != nil {
			publicIps = append(publicIps, *vm.PublicIp)
		}
		vmIds = append(vmIds, *vm.VmId)
	}

	if len(vmIds) != 0 {
		_, _, err := client.VmApi.DeleteVms(ctx).DeleteVmsRequest(*osc.NewDeleteVmsRequest(vmIds)).Execute()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error deleting the VMs")
			return err
		}
		fmt.Fprintf(os.Stderr, "> VMs %v deleted\n", vmIds)
	} else {
		fmt.Fprintln(os.Stderr, "> No VMs found")
	}

	// Delete Load Balancers
	lbs, _, err := client.LoadBalancerApi.ReadLoadBalancers(ctx).ReadLoadBalancersRequest(*osc.NewReadLoadBalancersRequest()).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the Load Balancers")
		return err
	}
	for _, lb := range *lbs.LoadBalancers {
		if lb.NetId != nil && *lb.NetId == netID {
			if lb.PublicIp != nil {
				publicIps = append(publicIps, *lb.PublicIp)
			}
			_, _, err := client.LoadBalancerApi.DeleteLoadBalancer(ctx).DeleteLoadBalancerRequest(*osc.NewDeleteLoadBalancerRequest(*lb.LoadBalancerName)).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error deleting the Load Balancer \"%s\"\n", *lb.LoadBalancerName)
				return err
			}
			cleared = true
			fmt.Fprintf(os.Stderr, "> Load Balancer \"%s\" deleted\n", *lb.LoadBalancerName)
		}
	}
	if !cleared {
		fmt.Fprintln(os.Stderr, "> No Load Balancers found")
	}

	// Unlink and delete non main Route Tables
	isMainRT := false
	filterRTs := osc.FiltersRouteTable{NetIds: &[]string{netID}, LinkRouteTableMain: &isMainRT}
	routeTables, _, err := client.RouteTableApi.ReadRouteTables(ctx).ReadRouteTablesRequest(osc.ReadRouteTablesRequest{Filters: &filterRTs}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the Route Tables")
		return err
	}
	for _, routeTable := range *routeTables.RouteTables {
		for _, routeTableLink := range *routeTable.LinkRouteTables {
			_, _, err := client.RouteTableApi.UnlinkRouteTable(ctx).UnlinkRouteTableRequest(*osc.NewUnlinkRouteTableRequest(*routeTableLink.LinkRouteTableId)).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error unlinking the Route Table Link \"%s\"\n", *routeTableLink.LinkRouteTableId)
				return err
			}
			fmt.Fprintf(os.Stderr, "> Route Table Link \"%s\" unlinked\n", *routeTableLink.LinkRouteTableId)
		}
		_, _, err := client.RouteTableApi.DeleteRouteTable(ctx).DeleteRouteTableRequest(*osc.NewDeleteRouteTableRequest(*routeTable.RouteTableId)).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting the Route Table \"%s\"\n", *routeTable.RouteTableId)
			return err
		}
		fmt.Fprintf(os.Stderr, "> Route Table \"%s\" deleted\n", *routeTable.RouteTableId)
	}
	if len(*routeTables.RouteTables) == 0 {
		fmt.Fprintln(os.Stderr, "> No Route Tables found")
	}

	// Unlink Virtual Gateways
	filterVGWs := osc.FiltersVirtualGateway{LinkNetIds: &[]string{netID}}
	vgws, _, err := client.VirtualGatewayApi.ReadVirtualGateways(ctx).ReadVirtualGatewaysRequest(osc.ReadVirtualGatewaysRequest{Filters: &filterVGWs}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the Virtual Gateways")
		return err
	}
	for _, vgw := range *vgws.VirtualGateways {
		_, _, err := client.VirtualGatewayApi.UnlinkVirtualGateway(ctx).UnlinkVirtualGatewayRequest(*osc.NewUnlinkVirtualGatewayRequest(netID, *vgw.VirtualGatewayId)).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unlinking the Virtual Gateway \"%s\"\n", *vgw.VirtualGatewayId)
			return err
		}
		_, _, err = client.VirtualGatewayApi.DeleteVirtualGateway(ctx).DeleteVirtualGatewayRequest(*osc.NewDeleteVirtualGatewayRequest(*vgw.VirtualGatewayId)).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting the Virtual Gateway \"%s\"\n", *vgw.VirtualGatewayId)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "> Virtual Gateway \"%s\" deleted\n", *vgw.VirtualGatewayId)
	}
	if len(*vgws.VirtualGateways) == 0 {
		fmt.Fprintln(os.Stderr, "> No Virtual Gateways found")
	}

	// Delete NATs
	filterNats := osc.FiltersNatService{NetIds: &[]string{netID}}
	nats, _, err := client.NatServiceApi.ReadNatServices(ctx).ReadNatServicesRequest(osc.ReadNatServicesRequest{Filters: &filterNats}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the NATs")
		return err
	}
	for _, nat := range *nats.NatServices {
		if nat.PublicIps != nil {
			for _, pip := range *nat.PublicIps {
				publicIps = append(publicIps, *pip.PublicIp)
			}
		}
		_, _, err := client.NatServiceApi.DeleteNatService(ctx).DeleteNatServiceRequest(*osc.NewDeleteNatServiceRequest(*nat.NatServiceId)).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting the NAT \"%s\"\n", *nat.NatServiceId)
			return err
		}
		fmt.Fprintf(os.Stderr, "> NAT \"%s\" deleted\n", *nat.NatServiceId)
	}
	if len(*nats.NatServices) == 0 {
		fmt.Fprintln(os.Stderr, "> No NATs found")
	}

	// Delete Net Access Points
	filterNetAPs := osc.FiltersNetAccessPoint{NetIds: &[]string{netID}}
	netAPs, _, err := client.NetAccessPointApi.ReadNetAccessPoints(ctx).ReadNetAccessPointsRequest(osc.ReadNetAccessPointsRequest{Filters: &filterNetAPs}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the Net Access Points")
		return err
	}
	for _, netAP := range *netAPs.NetAccessPoints {
		_, _, err := client.NetAccessPointApi.DeleteNetAccessPoint(ctx).DeleteNetAccessPointRequest(*osc.NewDeleteNetAccessPointRequest(*netAP.NetAccessPointId)).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting the Net Access Point \"%s\"\n", *netAP.NetAccessPointId)
			return err
		}
		fmt.Fprintf(os.Stderr, "> Net Access Point \"%s\" deleted\n", *netAP.NetAccessPointId)
	}
	if len(*netAPs.NetAccessPoints) == 0 {
		fmt.Fprintln(os.Stderr, "> No Net Access Points found")
	}

	// Delete Net Peerings
	filterSourcePeers := osc.FiltersNetPeering{SourceNetNetIds: &[]string{netID}}
	filterAccepterPeers := osc.FiltersNetPeering{AccepterNetNetIds: &[]string{netID}}
	sourceNetPeers, _, err := client.NetPeeringApi.ReadNetPeerings(ctx).ReadNetPeeringsRequest(osc.ReadNetPeeringsRequest{Filters: &filterSourcePeers}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the source Net Peerings")
		return err
	}
	accepterNetPeers, _, err := client.NetPeeringApi.ReadNetPeerings(ctx).ReadNetPeeringsRequest(osc.ReadNetPeeringsRequest{Filters: &filterAccepterPeers}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the accepter Net Peerings")
		return err
	}
	clearedSource, err := teardownNetPeerings(sourceNetPeers, client, ctx)
	if err != nil {
		return err
	}
	clearedAccepters, err := teardownNetPeerings(accepterNetPeers, client, ctx)
	if err != nil {
		return err
	}
	if !(clearedSource || clearedAccepters) {
		fmt.Fprintln(os.Stderr, "> No Net Peerings found")
	}

	// Delete NICs
	filterNics := osc.FiltersNic{NetIds: &[]string{netID}}
	nics, _, _ := client.NicApi.ReadNics(ctx).ReadNicsRequest(osc.ReadNicsRequest{Filters: &filterNics}).Execute()
	err = teardownNICs(vmIds, &publicIps, nics, client, ctx)
	if err != nil {
		return err
	}
	if len(*nics.Nics) == 0 {
		fmt.Fprintln(os.Stderr, "> No NICs found")
	}

	// Unlink and delete Internet Services
	filterISs := osc.FiltersInternetService{LinkNetIds: &[]string{netID}}
	interServices, _, err := client.InternetServiceApi.ReadInternetServices(ctx).ReadInternetServicesRequest(osc.ReadInternetServicesRequest{Filters: &filterISs}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the Internet Services")
		return err
	}
	for _, interService := range *interServices.InternetServices {
		_, _, err := client.InternetServiceApi.UnlinkInternetService(ctx).UnlinkInternetServiceRequest(*osc.NewUnlinkInternetServiceRequest(*interService.InternetServiceId, netID)).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unlinking the Internet Service \"%s\"\n", *interService.InternetServiceId)
			return err
		}
		fmt.Fprintf(os.Stderr, "> Internet Service \"%s\" unlinked\n", *interService.InternetServiceId)

		_, _, err = client.InternetServiceApi.DeleteInternetService(ctx).DeleteInternetServiceRequest(*osc.NewDeleteInternetServiceRequest(*interService.InternetServiceId)).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting the Internet Service \"%s\"\n", *interService.InternetServiceId)
			return err
		}
		fmt.Fprintf(os.Stderr, "> Internet Service \"%s\" deleted\n", *interService.InternetServiceId)
	}
	if len(*interServices.InternetServices) == 0 {
		fmt.Fprintln(os.Stderr, "> No Internet Services found")
	}

	// Delete non default Security Groups
	filterSGs := osc.FiltersSecurityGroup{NetIds: &[]string{netID}}
	secGroups, _, err := client.SecurityGroupApi.ReadSecurityGroups(ctx).ReadSecurityGroupsRequest(osc.ReadSecurityGroupsRequest{Filters: &filterSGs}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the Security Groups")
		return err
	}
	cleared = false
	for _, secGroup := range *secGroups.SecurityGroups {
		if *secGroup.SecurityGroupName != "default" {
			_, _, err := client.SecurityGroupApi.DeleteSecurityGroup(ctx).DeleteSecurityGroupRequest(osc.DeleteSecurityGroupRequest{SecurityGroupId: secGroup.SecurityGroupId}).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error deleting the Security Group \"%s\"\n", *secGroup.SecurityGroupId)
				return err
			}
			fmt.Fprintf(os.Stderr, "> Security Group \"%s\" deleted\n", *secGroup.SecurityGroupId)
			cleared = true
		}
	}
	if !cleared {
		fmt.Fprintln(os.Stderr, "> No Security Groups found")
	}

	// Delete Public IPs
	if len(publicIps) != 0 {
		filterPips := osc.FiltersPublicIp{PublicIps: &publicIps}
		publicIPsCb, _, err := client.PublicIpApi.ReadPublicIps(ctx).ReadPublicIpsRequest(osc.ReadPublicIpsRequest{Filters: &filterPips}).Execute()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting the Public IPs")
			return err
		}
		for _, publicIP := range *publicIPsCb.PublicIps {
			if publicIP.LinkPublicIpId == nil {
				_, _, err := client.PublicIpApi.DeletePublicIp(ctx).DeletePublicIpRequest(osc.DeletePublicIpRequest{PublicIpId: publicIP.PublicIpId}).Execute()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error deleting the Public IP \"%s\"\n", *publicIP.PublicIp)
					return err
				}
				fmt.Fprintf(os.Stderr, "> Public IP \"%s\" deleted\n", *publicIP.PublicIp)
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "> No Public IPs found")
	}

	// Delete Subnets
	for _, subnetID := range subnetsIDs {
		_, _, err := client.SubnetApi.DeleteSubnet(ctx).DeleteSubnetRequest(*osc.NewDeleteSubnetRequest(subnetID)).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting the Subnet \"%s\"\n", subnetID)
			return err
		}
		fmt.Fprintf(os.Stderr, "> Subnet \"%s\" deleted\n", subnetID)
	}
	if len(subnetsIDs) == 0 {
		fmt.Fprintln(os.Stderr, "> No Subnets found")
	}

	// Delete Net
	_, _, err = client.NetApi.DeleteNet(ctx).DeleteNetRequest(*osc.NewDeleteNetRequest(netID)).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting the Net \"%s\"\n", netID)
		return err
	}
	fmt.Fprintf(os.Stderr, "> Net \"%s\" deleted\n", netID)
	fmt.Println("Teardown completed")

	return nil
}

func getSubnetsIDs(netID string, client *osc.APIClient, ctx context.Context) ([]string, error) {
	filterSubnets := osc.FiltersSubnet{NetIds: &[]string{netID}}

	subnets, _, err := client.SubnetApi.ReadSubnets(ctx).ReadSubnetsRequest(osc.ReadSubnetsRequest{Filters: &filterSubnets}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the Subnets")
		return nil, err
	}
	var subnetsIDs []string
	for _, subnet := range *subnets.Subnets {
		subnetsIDs = append(subnetsIDs, *subnet.SubnetId)
	}

	return subnetsIDs, nil
}

func getVMsFromIDs(vmIDs []string, client *osc.APIClient, ctx context.Context) (osc.ReadVmsResponse, error) {
	filterVms := osc.FiltersVm{VmIds: &vmIDs}
	vms, _, err := client.VmApi.ReadVms(ctx).ReadVmsRequest(osc.ReadVmsRequest{Filters: &filterVms}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the VMs")
		return vms, err
	}

	return vms, nil
}

func getVMs(netID string, client *osc.APIClient, ctx context.Context) (osc.ReadVmsResponse, error) {
	filterVms := osc.FiltersVm{NetIds: &[]string{netID}}
	vms, _, err := client.VmApi.ReadVms(ctx).ReadVmsRequest(osc.ReadVmsRequest{Filters: &filterVms}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting the VMs")
		return vms, err
	}

	return vms, nil
}

func teardownNetPeerings(netPeerings osc.ReadNetPeeringsResponse, client *osc.APIClient, ctx context.Context) (bool, error) {
	cleared := false
	for _, netPeer := range *netPeerings.NetPeerings {
		if *netPeer.State.Name == `pending-acceptance` {
			_, _, err := client.NetPeeringApi.RejectNetPeering(ctx).RejectNetPeeringRequest(*osc.NewRejectNetPeeringRequest(*netPeer.NetPeeringId)).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error rejecting the Net Peering \"%s\"\n", *netPeer.NetPeeringId)
				return cleared, err
			}
			fmt.Fprintf(os.Stderr, "> Net Peering \"%s\" rejected\n", *netPeer.NetPeeringId)
			cleared = true
		} else if *netPeer.State.Name == `active` {
			_, _, err := client.NetPeeringApi.DeleteNetPeering(ctx).DeleteNetPeeringRequest(*osc.NewDeleteNetPeeringRequest(*netPeer.NetPeeringId)).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error deleting the Net Peering \"%s\"\n", *netPeer.NetPeeringId)
				return cleared, err
			}
			fmt.Fprintf(os.Stderr, "> Net Peering \"%s\" deleted\n", *netPeer.NetPeeringId)
			cleared = true
		}
	}
	return cleared, nil
}

func deleteNIC(nicID string, client *osc.APIClient, ctx context.Context) error {
	_, _, err := client.NicApi.DeleteNic(ctx).DeleteNicRequest(*osc.NewDeleteNicRequest(nicID)).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting the NIC \"%s\"\n", nicID)
		return err
	}
	fmt.Fprintf(os.Stderr, "> NIC \"%s\" deleted\n", nicID)
	return nil
}

func teardownNICs(vmIDs []string, publicIPs *[]string, nics osc.ReadNicsResponse, client *osc.APIClient, ctx context.Context) error {
	var nicsRest []osc.Nic
	for _, nic := range *nics.Nics {
		if nic.LinkNic == nil {
			err := deleteNIC(*nic.NicId, client, ctx)
			if err != nil {
				return err
			}
		} else {
			nicsRest = append(nicsRest, nic)
		}
		if nic.LinkPublicIp != nil && nic.LinkPublicIp.PublicIp != nil {
			*publicIPs = append(*publicIPs, *nic.LinkPublicIp.PublicIp)
		}

	}
	if len(nicsRest) == 0 {
		return nil
	}

	fmt.Fprintln(os.Stderr, "> Waiting for VMs to be terminated to delete NICs...")
	for {
		allTerminated := true
		vms, err := getVMsFromIDs(vmIDs, client, ctx)
		if err != nil {
			return err
		}
		for _, vm := range *vms.Vms {
			if *vm.State != `terminated` {
				allTerminated = false
			}
		}
		if allTerminated {
			break
		}
		sleepTime := rand.Intn(10) + 1
		time.Sleep(time.Second * time.Duration(sleepTime))
	}
	for _, nic := range nicsRest {
		err := deleteNIC(*nic.NicId, client, ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
