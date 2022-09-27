/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// ImageExportTask Information about the OMI export task.
type ImageExportTask struct {
	// If the OMI export task fails, an error message appears.
	Comment string `json:"Comment,omitempty"`
	// The ID of the OMI to be exported.
	ImageId   string                   `json:"ImageId,omitempty"`
	OsuExport OsuExportImageExportTask `json:"OsuExport,omitempty"`
	// The progress of the OMI export task, as a percentage.
	Progress int32 `json:"Progress,omitempty"`
	// The state of the OMI export task (`pending/queued` \\| `pending` \\| `completed` \\| `failed` \\| `cancelled`).
	State string `json:"State,omitempty"`
	// One or more tags associated with the image export task.
	Tags []ResourceTag `json:"Tags,omitempty"`
	// The ID of the OMI export task.
	TaskId string `json:"TaskId,omitempty"`
}
