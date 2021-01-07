/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// AccessKey Information about the access key.
type AccessKey struct {
	// The ID of the access key.
	AccessKeyId string `json:"AccessKeyId,omitempty"`
	// The date and time of creation of the access key.
	CreationDate string `json:"CreationDate,omitempty"`
	// The date at which the access key expires.
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	// The date and time of the last modification of the access key.
	LastModificationDate string `json:"LastModificationDate,omitempty"`
	// The state of the access key (`ACTIVE` if the key is valid for API calls, or `INACTIVE` if not).
	State string `json:"State,omitempty"`
}
