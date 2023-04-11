/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// Nic Information about the NIC.
type Nic struct {
	// The account ID of the owner of the NIC.
	AccountId string `json:"AccountId,omitempty"`
	// The description of the NIC.
	Description string `json:"Description,omitempty"`
	// (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked bool         `json:"IsSourceDestChecked,omitempty"`
	LinkNic             LinkNic      `json:"LinkNic,omitempty"`
	LinkPublicIp        LinkPublicIp `json:"LinkPublicIp,omitempty"`
	// The Media Access Control (MAC) address of the NIC.
	MacAddress string `json:"MacAddress,omitempty"`
	// The ID of the Net for the NIC.
	NetId string `json:"NetId,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId,omitempty"`
	// The name of the private DNS.
	PrivateDnsName string `json:"PrivateDnsName,omitempty"`
	// The private IPs of the NIC.
	PrivateIps []PrivateIp `json:"PrivateIps,omitempty"`
	// One or more IDs of security groups for the NIC.
	SecurityGroups []SecurityGroupLight `json:"SecurityGroups,omitempty"`
	// The state of the NIC (`available` \\| `attaching` \\| `in-use` \\| `detaching`).
	State string `json:"State,omitempty"`
	// The ID of the Subnet.
	SubnetId string `json:"SubnetId,omitempty"`
	// The Subregion in which the NIC is located.
	SubregionName string `json:"SubregionName,omitempty"`
	// One or more tags associated with the NIC.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
