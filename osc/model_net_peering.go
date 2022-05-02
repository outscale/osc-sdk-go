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

// NetPeering Information about the Net peering connection.
type NetPeering struct {
	AccepterNet AccepterNet `json:"AccepterNet,omitempty"`
	// The ID of the Net peering connection.
	NetPeeringId string          `json:"NetPeeringId,omitempty"`
	SourceNet    SourceNet       `json:"SourceNet,omitempty"`
	State        NetPeeringState `json:"State,omitempty"`
	// One or more tags associated with the Net peering connection.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
