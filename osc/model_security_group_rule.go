/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// SecurityGroupRule Information about the security group rule.
type SecurityGroupRule struct {
	// The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
	FromPortRange int32 `json:"FromPortRange,omitempty"`
	// The IP protocol name (`tcp`, `udp`, `icmp`, or `-1` for all protocols). By default, `-1`. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	IpProtocol string `json:"IpProtocol,omitempty"`
	// One or more IP ranges for the security group rules, in CIDR notation (for example, `10.0.0.0/16`).
	IpRanges []string `json:"IpRanges,omitempty"`
	// Information about one or more source or destination security groups.
	SecurityGroupsMembers []SecurityGroupsMember `json:"SecurityGroupsMembers,omitempty"`
	// One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices).
	ServiceIds []string `json:"ServiceIds,omitempty"`
	// The end of the port range for the TCP and UDP protocols, or an ICMP code number.
	ToPortRange int32 `json:"ToPortRange,omitempty"`
}
