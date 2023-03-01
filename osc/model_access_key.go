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

// AccessKey Information about the access key.
type AccessKey struct {
	// The ID of the access key.
	AccessKeyId string `json:"AccessKeyId,omitempty"`
	// The date and time (UTC) of creation of the access key.
	CreationDate string `json:"CreationDate,omitempty"`
	// The date (UTC) at which the access key expires.
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	// The date and time (UTC) of the last modification of the access key.
	LastModificationDate string `json:"LastModificationDate,omitempty"`
	// The state of the access key (`ACTIVE` if the key is valid for API calls, or `INACTIVE` if not).
	State string `json:"State,omitempty"`
}
