/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.16
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
