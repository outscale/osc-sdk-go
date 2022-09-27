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

// FiltersDhcpOptions One or more filters.
type FiltersDhcpOptions struct {
	// If true, lists all default DHCP options set. If false, lists all non-default DHCP options set.
	Default bool `json:"Default,omitempty"`
	// The IDs of the DHCP options sets.
	DhcpOptionsSetIds []string `json:"DhcpOptionsSetIds,omitempty"`
	// The IPs of the domain name servers used for the DHCP options sets.
	DomainNameServers []string `json:"DomainNameServers,omitempty"`
	// The domain names used for the DHCP options sets.
	DomainNames []string `json:"DomainNames,omitempty"`
	// The IPs of the log servers used for the DHCP options sets.
	LogServers []string `json:"LogServers,omitempty"`
	// The IPs of the Network Time Protocol (NTP) servers used for the DHCP options sets.
	NtpServers []string `json:"NtpServers,omitempty"`
	// The keys of the tags associated with the DHCP options sets.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the DHCP options sets.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the DHCP options sets, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
}
