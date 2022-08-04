/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.21
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FiltersInternetService One or more filters.
type FiltersInternetService struct {
	// The IDs of the Internet services.
	InternetServiceIds []string `json:"InternetServiceIds,omitempty"`
	// The IDs of the Nets the Internet services are attached to.
	LinkNetIds []string `json:"LinkNetIds,omitempty"`
	// The current states of the attachments between the Internet services and the Nets (only `available`, if the Internet gateway is attached to a VPC).
	LinkStates []string `json:"LinkStates,omitempty"`
	// The keys of the tags associated with the Internet services.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Internet services.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Internet services, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
}
