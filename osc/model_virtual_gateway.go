/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// VirtualGateway Information about the virtual gateway.
type VirtualGateway struct {
	// The type of VPN connection supported by the virtual gateway (only `ipsec.1` is supported).
	ConnectionType string `json:"ConnectionType,omitempty"`
	// The Net to which the virtual gateway is attached.
	NetToVirtualGatewayLinks []NetToVirtualGatewayLink `json:"NetToVirtualGatewayLinks,omitempty"`
	// The state of the virtual gateway (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State string `json:"State,omitempty"`
	// One or more tags associated with the virtual gateway.
	Tags []ResourceTag `json:"Tags,omitempty"`
	// The ID of the virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId,omitempty"`
}
