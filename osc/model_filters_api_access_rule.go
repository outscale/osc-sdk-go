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

// FiltersApiAccessRule One or more filters.
type FiltersApiAccessRule struct {
	// One or more IDs of API access rules.
	ApiAccessRuleIds []string `json:"ApiAccessRuleIds,omitempty"`
	// One or more IDs of Client Certificate Authorities (CAs).
	CaIds []string `json:"CaIds,omitempty"`
	// One or more Client Certificate Common Names (CNs).
	Cns []string `json:"Cns,omitempty"`
	// One or more descriptions of API access rules.
	Descriptions []string `json:"Descriptions,omitempty"`
	// One or more IP ranges, in CIDR notation (for example, 192.0.2.0/16).
	IpRanges []string `json:"IpRanges,omitempty"`
}
