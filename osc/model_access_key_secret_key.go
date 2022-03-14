/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// AccessKeySecretKey Information about the access key.
type AccessKeySecretKey struct {
	// The ID of the access key.
	AccessKeyId string `json:"AccessKeyId,omitempty"`
	// The date and time (UTC) of creation of the access key.
	CreationDate string `json:"CreationDate,omitempty"`
	// The date and time (UTC) at which the access key expires.
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	// The date and time (UTC) of the last modification of the access key.
	LastModificationDate string `json:"LastModificationDate,omitempty"`
	// The access key that enables you to send requests.
	SecretKey string `json:"SecretKey,omitempty"`
	// The state of the access key (`ACTIVE` if the key is valid for API calls, or `INACTIVE` if not).
	State string `json:"State,omitempty"`
}
