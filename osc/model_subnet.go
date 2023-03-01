/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// Subnet Information about the Subnet.
type Subnet struct {
	// The number of available IPs in the Subnets.
	AvailableIpsCount int32 `json:"AvailableIpsCount,omitempty"`
	// The IP range in the Subnet, in CIDR notation (for example, `10.0.0.0/16`).
	IpRange string `json:"IpRange,omitempty"`
	// If true, a public IP is assigned to the network interface cards (NICs) created in the specified Subnet.
	MapPublicIpOnLaunch bool `json:"MapPublicIpOnLaunch,omitempty"`
	// The ID of the Net in which the Subnet is.
	NetId string `json:"NetId,omitempty"`
	// The state of the Subnet (`pending` \\| `available` \\| `deleted`).
	State string `json:"State,omitempty"`
	// The ID of the Subnet.
	SubnetId string `json:"SubnetId,omitempty"`
	// The name of the Subregion in which the Subnet is located.
	SubregionName string `json:"SubregionName,omitempty"`
	// One or more tags associated with the Subnet.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
