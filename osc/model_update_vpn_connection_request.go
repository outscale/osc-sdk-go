/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// UpdateVpnConnectionRequest struct for UpdateVpnConnectionRequest
type UpdateVpnConnectionRequest struct {
	// The ID of the client gateway.
	ClientGatewayId string `json:"ClientGatewayId,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The ID of the virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId,omitempty"`
	// The ID of the VPN connection you want to modify.
	VpnConnectionId string     `json:"VpnConnectionId"`
	VpnOptions      VpnOptions `json:"VpnOptions,omitempty"`
}
