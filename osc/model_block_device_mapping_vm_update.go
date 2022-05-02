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

// BlockDeviceMappingVmUpdate Information about the block device mapping.
type BlockDeviceMappingVmUpdate struct {
	Bsu BsuToUpdateVm `json:"Bsu,omitempty"`
	// The device name for the volume. For a root device, you must use `/dev/sda1`. For other volumes, you must use `/dev/sdX` or `/dev/xvdX` (where `X` is a letter between `b` and `z`).
	DeviceName string `json:"DeviceName,omitempty"`
	// Removes the device which is included in the block device mapping of the OMI.
	NoDevice string `json:"NoDevice,omitempty"`
	// The name of the virtual device (ephemeralN).
	VirtualDeviceName string `json:"VirtualDeviceName,omitempty"`
}
