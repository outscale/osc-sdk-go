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

// CreateDhcpOptionsRequest struct for CreateDhcpOptionsRequest
type CreateDhcpOptionsRequest struct {
	// Specify a domain name (for example, MyCompany.com). You can specify only one domain name. You must specify at least one of the following parameters: `DomainName`, `DomainNameServers` or `NtpServers`.
	DomainName string `json:"DomainName,omitempty"`
	// The IPs of domain name servers. If no IPs are specified, the `OutscaleProvidedDNS` value is set by default. You must specify at least one of the following parameters: `DomainName`, `DomainNameServers` or `NtpServers`.
	DomainNameServers []string `json:"DomainNameServers,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The IPs of the Network Time Protocol (NTP) servers. You must specify at least one of the following parameters: `DomainName`, `DomainNameServers` or `NtpServers`.
	NtpServers []string `json:"NtpServers,omitempty"`
}
