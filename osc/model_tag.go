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

// Tag Information about the tag.
type Tag struct {
	// The key of the tag, with a minimum of 1 character.
	Key string `json:"Key,omitempty"`
	// The ID of the resource.
	ResourceId string `json:"ResourceId,omitempty"`
	// The type of the resource.
	ResourceType string `json:"ResourceType,omitempty"`
	// The value of the tag, between 0 and 255 characters.
	Value string `json:"Value,omitempty"`
}
