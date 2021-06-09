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

// LinkedVolume Information about volume attachment.
type LinkedVolume struct {
	// If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.
	DeleteOnVmDeletion bool `json:"DeleteOnVmDeletion,omitempty"`
	// The name of the device.
	DeviceName string `json:"DeviceName,omitempty"`
	// The state of the attachment of the volume (`attaching` \\| `detaching` \\| `attached` \\| `detached`).
	State string `json:"State,omitempty"`
	// The ID of the VM.
	VmId string `json:"VmId,omitempty"`
	// The ID of the volume.
	VolumeId string `json:"VolumeId,omitempty"`
}
