/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// SecurityGroup Information about the security group.
type SecurityGroup struct {
	// The account ID of a user that has been granted permission.
	AccountId string `json:"AccountId,omitempty"`
	// The description of the security group.
	Description string `json:"Description,omitempty"`
	// The inbound rules associated with the security group.
	InboundRules []SecurityGroupRule `json:"InboundRules,omitempty"`
	// The ID of the Net for the security group.
	NetId string `json:"NetId,omitempty"`
	// The outbound rules associated with the security group.
	OutboundRules []SecurityGroupRule `json:"OutboundRules,omitempty"`
	// The ID of the security group.
	SecurityGroupId string `json:"SecurityGroupId,omitempty"`
	// The name of the security group.
	SecurityGroupName string `json:"SecurityGroupName,omitempty"`
	// One or more tags associated with the security group.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
