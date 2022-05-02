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

// NicLight Information about the network interface card (NIC).
type NicLight struct {
	// The account ID of the owner of the NIC.
	AccountId string `json:"AccountId,omitempty"`
	// The description of the NIC.
	Description string `json:"Description,omitempty"`
	// (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked bool                   `json:"IsSourceDestChecked,omitempty"`
	LinkNic             LinkNicLight           `json:"LinkNic,omitempty"`
	LinkPublicIp        LinkPublicIpLightForVm `json:"LinkPublicIp,omitempty"`
	// The Media Access Control (MAC) address of the NIC.
	MacAddress string `json:"MacAddress,omitempty"`
	// The ID of the Net for the NIC.
	NetId string `json:"NetId,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId,omitempty"`
	// The name of the private DNS.
	PrivateDnsName string `json:"PrivateDnsName,omitempty"`
	// The private IP or IPs of the NIC.
	PrivateIps []PrivateIpLightForVm `json:"PrivateIps,omitempty"`
	// One or more IDs of security groups for the NIC.
	SecurityGroups []SecurityGroupLight `json:"SecurityGroups,omitempty"`
	// The state of the NIC (`available` \\| `attaching` \\| `in-use` \\| `detaching`).
	State string `json:"State,omitempty"`
	// The ID of the Subnet for the NIC.
	SubnetId string `json:"SubnetId,omitempty"`
}
