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

// Image Information about the OMI.
type Image struct {
	// The account alias of the owner of the OMI.
	AccountAlias string `json:"AccountAlias,omitempty"`
	// The account ID of the owner of the OMI.
	AccountId string `json:"AccountId,omitempty"`
	// The architecture of the OMI (by default, `i386`).
	Architecture string `json:"Architecture,omitempty"`
	// One or more block device mappings.
	BlockDeviceMappings []BlockDeviceMappingImage `json:"BlockDeviceMappings,omitempty"`
	// The date and time of creation of the OMI.
	CreationDate string `json:"CreationDate,omitempty"`
	// The description of the OMI.
	Description string `json:"Description,omitempty"`
	// The location of the bucket where the OMI files are stored.
	FileLocation string `json:"FileLocation,omitempty"`
	// The ID of the OMI.
	ImageId string `json:"ImageId,omitempty"`
	// The name of the OMI.
	ImageName string `json:"ImageName,omitempty"`
	// The type of the OMI.
	ImageType           string                `json:"ImageType,omitempty"`
	PermissionsToLaunch PermissionsOnResource `json:"PermissionsToLaunch,omitempty"`
	// The product code associated with the OMI (`0001` Linux/Unix \\| `0002` Windows \\| `0004` Linux/Oracle \\| `0005` Windows 10).
	ProductCodes []string `json:"ProductCodes,omitempty"`
	// The name of the root device.
	RootDeviceName string `json:"RootDeviceName,omitempty"`
	// The type of root device used by the OMI (always `bsu`).
	RootDeviceType string `json:"RootDeviceType,omitempty"`
	// The state of the OMI (`pending` \\| `available` \\| `failed`).
	State        string       `json:"State,omitempty"`
	StateComment StateComment `json:"StateComment,omitempty"`
	// One or more tags associated with the OMI.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
