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

// LinkNic Information about the NIC attachment.
type LinkNic struct {
	// If true, the NIC is deleted when the VM is terminated.
	DeleteOnVmDeletion bool `json:"DeleteOnVmDeletion,omitempty"`
	// The device index for the NIC attachment (between 1 and 7, both included).
	DeviceNumber int32 `json:"DeviceNumber,omitempty"`
	// The ID of the NIC to attach.
	LinkNicId string `json:"LinkNicId,omitempty"`
	// The state of the attachment (`attaching` \\| `attached` \\| `detaching` \\| `detached`).
	State string `json:"State,omitempty"`
	// The account ID of the owner of the VM.
	VmAccountId string `json:"VmAccountId,omitempty"`
	// The ID of the VM.
	VmId string `json:"VmId,omitempty"`
}
