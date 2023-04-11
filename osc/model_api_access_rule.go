/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// ApiAccessRule Information about the API access rule.
type ApiAccessRule struct {
	//  The ID of the API access rule.
	ApiAccessRuleId string `json:"ApiAccessRuleId,omitempty"`
	// One or more IDs of Client Certificate Authorities (CAs) used for the API access rule.
	CaIds []string `json:"CaIds,omitempty"`
	// One or more Client Certificate Common Names (CNs).
	Cns []string `json:"Cns,omitempty"`
	// The description of the API access rule.
	Description string `json:"Description,omitempty"`
	// One or more IP ranges used for the API access rule, in CIDR notation (for example, `192.0.2.0/16`).
	IpRanges []string `json:"IpRanges,omitempty"`
}
