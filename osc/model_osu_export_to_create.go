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

// OsuExportToCreate Information about the OOS export task to create.
type OsuExportToCreate struct {
	// The format of the export disk (`qcow2` \\| `raw`).
	DiskImageFormat string    `json:"DiskImageFormat"`
	OsuApiKey       OsuApiKey `json:"OsuApiKey,omitempty"`
	// The name of the OOS bucket where you want to export the object.
	OsuBucket string `json:"OsuBucket"`
	// The URL of the manifest file.
	OsuManifestUrl string `json:"OsuManifestUrl,omitempty"`
	// The prefix for the key of the OOS object.
	OsuPrefix string `json:"OsuPrefix,omitempty"`
}
