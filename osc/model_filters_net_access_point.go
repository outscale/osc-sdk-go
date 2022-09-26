/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FiltersNetAccessPoint One or more filters.
type FiltersNetAccessPoint struct {
	// The IDs of the Net access points.
	NetAccessPointIds []string `json:"NetAccessPointIds,omitempty"`
	// The IDs of the Nets.
	NetIds []string `json:"NetIds,omitempty"`
	// The names of the services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices).
	ServiceNames []string `json:"ServiceNames,omitempty"`
	// The states of the Net access points (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	States []string `json:"States,omitempty"`
	// The keys of the tags associated with the Net access points.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Net access points.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Net access points, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
}
