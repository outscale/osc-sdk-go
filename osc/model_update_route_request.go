/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.21
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// UpdateRouteRequest struct for UpdateRouteRequest
type UpdateRouteRequest struct {
	// The IP range used for the destination match, in CIDR notation (for example, `10.0.0.0/24`).
	DestinationIpRange string `json:"DestinationIpRange"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The ID of an Internet service or virtual gateway attached to your Net.
	GatewayId string `json:"GatewayId,omitempty"`
	// The ID of a NAT service.
	NatServiceId string `json:"NatServiceId,omitempty"`
	// The ID of a Net peering connection.
	NetPeeringId string `json:"NetPeeringId,omitempty"`
	// The ID of a network interface card (NIC).
	NicId string `json:"NicId,omitempty"`
	// The ID of the route table.
	RouteTableId string `json:"RouteTableId"`
	// The ID of a NAT VM in your Net.
	VmId string `json:"VmId,omitempty"`
}
