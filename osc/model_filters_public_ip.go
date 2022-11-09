/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FiltersPublicIp One or more filters.
type FiltersPublicIp struct {
	// The IDs representing the associations of public IPs with VMs or NICs.
	LinkPublicIpIds []string `json:"LinkPublicIpIds,omitempty"`
	// The account IDs of the owners of the NICs.
	NicAccountIds []string `json:"NicAccountIds,omitempty"`
	// The IDs of the NICs.
	NicIds []string `json:"NicIds,omitempty"`
	// Whether the public IPs are for use in the public Cloud or in a Net.
	Placements []string `json:"Placements,omitempty"`
	// The private IPs associated with the public IPs.
	PrivateIps []string `json:"PrivateIps,omitempty"`
	// The IDs of the public IPs.
	PublicIpIds []string `json:"PublicIpIds,omitempty"`
	// The public IPs.
	PublicIps []string `json:"PublicIps,omitempty"`
	// The keys of the tags associated with the public IPs.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the public IPs.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the public IPs, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
	// The IDs of the VMs.
	VmIds []string `json:"VmIds,omitempty"`
}
