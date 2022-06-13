/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FiltersNet One or more filters.
type FiltersNet struct {
	// The IDs of the DHCP options sets.
	DhcpOptionsSetIds []string `json:"DhcpOptionsSetIds,omitempty"`
	// The IP ranges for the Nets, in CIDR notation (for example, `10.0.0.0/16`).
	IpRanges []string `json:"IpRanges,omitempty"`
	// If true, the Net used is the default one.
	IsDefault bool `json:"IsDefault,omitempty"`
	// The IDs of the Nets.
	NetIds []string `json:"NetIds,omitempty"`
	// The states of the Nets (`pending` \\| `available`).
	States []string `json:"States,omitempty"`
	// The keys of the tags associated with the Nets.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Nets.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Nets, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
}
