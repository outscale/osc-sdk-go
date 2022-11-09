/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// DirectLinkInterfaces Information about the DirectLink interfaces.
type DirectLinkInterfaces struct {
	// The account ID of the owner of the DirectLink interface.
	AccountId string `json:"AccountId,omitempty"`
	// The BGP (Border Gateway Protocol) ASN (Autonomous System Number) on the customer's side of the DirectLink interface.
	BgpAsn int32 `json:"BgpAsn,omitempty"`
	// The BGP authentication key.
	BgpKey string `json:"BgpKey,omitempty"`
	// The IP on the customer's side of the DirectLink interface.
	ClientPrivateIp string `json:"ClientPrivateIp,omitempty"`
	// The ID of the DirectLink.
	DirectLinkId string `json:"DirectLinkId,omitempty"`
	// The ID of the DirectLink interface.
	DirectLinkInterfaceId string `json:"DirectLinkInterfaceId,omitempty"`
	// The name of the DirectLink interface.
	DirectLinkInterfaceName string `json:"DirectLinkInterfaceName,omitempty"`
	// The type of the DirectLink interface (always `private`).
	InterfaceType string `json:"InterfaceType,omitempty"`
	// The datacenter where the DirectLink interface is located.
	Location string `json:"Location,omitempty"`
	// The maximum transmission unit (MTU) of the DirectLink interface, in bytes (by default, `1500`).
	Mtu int32 `json:"Mtu,omitempty"`
	// The IP on the OUTSCALE side of the DirectLink interface.
	OutscalePrivateIp string `json:"OutscalePrivateIp,omitempty"`
	// The state of the DirectLink interface (`pending` \\| `available` \\| `deleting` \\| `deleted` \\| `confirming` \\| `rejected` \\| `expired`).
	State string `json:"State,omitempty"`
	// The ID of the target virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId,omitempty"`
	// The VLAN number associated with the DirectLink interface.
	Vlan int32 `json:"Vlan,omitempty"`
}
