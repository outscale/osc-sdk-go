/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// CreateImageRequest struct for CreateImageRequest
type CreateImageRequest struct {
	// The architecture of the OMI (by default, `i386`).
	Architecture string `json:"Architecture,omitempty"`
	// One or more block device mappings.
	BlockDeviceMappings []BlockDeviceMappingImage `json:"BlockDeviceMappings,omitempty"`
	// A description for the new OMI.
	Description string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The pre-signed URL of the OMI manifest file, or the full path to the OMI stored in a bucket. If you specify this parameter, a copy of the OMI is created in your account.
	FileLocation string `json:"FileLocation,omitempty"`
	// A unique name for the new OMI.<br /> Constraints: 3-128 alphanumeric characters, underscores (_), spaces ( ), parentheses (()), slashes (/), periods (.), or dashes (-).
	ImageName string `json:"ImageName,omitempty"`
	// If false, the VM shuts down before creating the OMI and then reboots. If true, the VM does not.
	NoReboot bool `json:"NoReboot,omitempty"`
	// The name of the root device.
	RootDeviceName string `json:"RootDeviceName,omitempty"`
	// The ID of the OMI you want to copy.
	SourceImageId string `json:"SourceImageId,omitempty"`
	// The name of the source Region, which must be the same as the Region of your account.
	SourceRegionName string `json:"SourceRegionName,omitempty"`
	// The ID of the VM from which you want to create the OMI.
	VmId string `json:"VmId,omitempty"`
}
