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

// CreateImageRequest struct for CreateImageRequest
type CreateImageRequest struct {
	// The architecture of the OMI (by default, `i386` if you specified the `FileLocation` or `RootDeviceName` parameter).
	Architecture string `json:"Architecture,omitempty"`
	// One or more block device mappings.
	BlockDeviceMappings []BlockDeviceMappingImage `json:"BlockDeviceMappings,omitempty"`
	// A description for the new OMI.
	Description string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The pre-signed URL of the OMI manifest file, or the full path to the OMI stored in a bucket. If you specify this parameter, a copy of the OMI is created in your account. You must specify only one of the following parameters: `FileLocation`, `RootDeviceName`, `SourceImageId` or `VmId`.
	FileLocation string `json:"FileLocation,omitempty"`
	// A unique name for the new OMI.<br /> Constraints: 3-128 alphanumeric characters, underscores (_), spaces ( ), parentheses (()), slashes (/), periods (.), or dashes (-).
	ImageName string `json:"ImageName,omitempty"`
	// If false, the VM shuts down before creating the OMI and then reboots. If true, the VM does not.
	NoReboot bool `json:"NoReboot,omitempty"`
	// The name of the root device. You must specify only one of the following parameters: `FileLocation`, `RootDeviceName`, `SourceImageId` or `VmId`.
	RootDeviceName string `json:"RootDeviceName,omitempty"`
	// The ID of the OMI you want to copy. You must specify only one of the following parameters: `FileLocation`, `RootDeviceName`, `SourceImageId` or `VmId`.
	SourceImageId string `json:"SourceImageId,omitempty"`
	// The name of the source Region, which must be the same as the Region of your account.
	SourceRegionName string `json:"SourceRegionName,omitempty"`
	// The ID of the VM from which you want to create the OMI. You must specify only one of the following parameters: `FileLocation`, `RootDeviceName`, `SourceImageId` or `VmId`.
	VmId string `json:"VmId,omitempty"`
}
