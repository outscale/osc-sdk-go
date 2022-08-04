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

// LinkPrivateIpsRequest struct for LinkPrivateIpsRequest
type LinkPrivateIpsRequest struct {
	// If true, allows an IP that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified.
	AllowRelink bool `json:"AllowRelink,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId"`
	// The secondary private IP or IPs you want to assign to the NIC within the IP range of the Subnet.
	PrivateIps []string `json:"PrivateIps,omitempty"`
	// The number of secondary private IPs to assign to the NIC.
	SecondaryPrivateIpCount int32 `json:"SecondaryPrivateIpCount,omitempty"`
}
