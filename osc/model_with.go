/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// With The information to display in each returned log. If this parameter is not specified, all available fields are considered as true.
type With struct {
	// By default or if set to true, the account ID is displayed.
	AccountId bool `json:"AccountId,omitempty"`
	// If set to true, the duration of the call is displayed.
	CallDuration bool `json:"CallDuration,omitempty"`
	// If set to true, the access key is displayed.
	QueryAccessKey bool `json:"QueryAccessKey,omitempty"`
	// If set to true, the name of the API is displayed.
	QueryApiName bool `json:"QueryApiName,omitempty"`
	// If set to true, the version of the API is displayed.
	QueryApiVersion bool `json:"QueryApiVersion,omitempty"`
	// If set to true, the name of the call is displayed.
	QueryCallName bool `json:"QueryCallName,omitempty"`
	// If set to true, the date of the call is displayed.
	QueryDate bool `json:"QueryDate,omitempty"`
	// If set to true, the raw header of the HTTP request is displayed.
	QueryHeaderRaw bool `json:"QueryHeaderRaw,omitempty"`
	// If set to true, the size of the raw header of the HTTP request is displayed.
	QueryHeaderSize bool `json:"QueryHeaderSize,omitempty"`
	// If set to true, the IP address is displayed.
	QueryIpAddress bool `json:"QueryIpAddress,omitempty"`
	// If set to true, the raw payload of the HTTP request is displayed.
	QueryPayloadRaw bool `json:"QueryPayloadRaw,omitempty"`
	// If set to true, the size of the raw payload of the HTTP request is displayed.
	QueryPayloadSize bool `json:"QueryPayloadSize,omitempty"`
	// If set to true, the user agent of the HTTP request is displayed.
	QueryUserAgent bool `json:"QueryUserAgent,omitempty"`
	// By default ot if set to true, the request ID is displayed.
	RequestId bool `json:"RequestId,omitempty"`
	// If set to true, the size of the response is displayed.
	ResponseSize bool `json:"ResponseSize,omitempty"`
	// If set to true, the HTTP status code of the response is displayed.
	ResponseStatusCode bool `json:"ResponseStatusCode,omitempty"`
}
