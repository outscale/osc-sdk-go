/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// UpdateNicRequest struct for UpdateNicRequest
type UpdateNicRequest struct {
	// A new description for the NIC.
	Description string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  bool            `json:"DryRun,omitempty"`
	LinkNic LinkNicToUpdate `json:"LinkNic,omitempty"`
	// The ID of the NIC you want to modify.
	NicId string `json:"NicId"`
	// One or more IDs of security groups for the NIC.<br /> You must specify at least one group, even if you use the default security group in the Net.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
}
