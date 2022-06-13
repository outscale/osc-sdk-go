/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// ReadApiLogsRequest struct for ReadApiLogsRequest
type ReadApiLogsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  bool          `json:"DryRun,omitempty"`
	Filters FiltersApiLog `json:"Filters,omitempty"`
	// The token to request the next page of results.
	NextPageToken string `json:"NextPageToken,omitempty"`
	// The maximum number of logs returned in a single response (between `1`and `1000`, both included). By default, `100`.
	ResultsPerPage int32 `json:"ResultsPerPage,omitempty"`
	With           With  `json:"With,omitempty"`
}
