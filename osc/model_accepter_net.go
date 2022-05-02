/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// AccepterNet Information about the accepter Net.
type AccepterNet struct {
	// The account ID of the owner of the accepter Net.
	AccountId string `json:"AccountId,omitempty"`
	// The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).
	IpRange string `json:"IpRange,omitempty"`
	// The ID of the accepter Net.
	NetId string `json:"NetId,omitempty"`
}
