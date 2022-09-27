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

// CreateSecurityGroupRuleRequest struct for CreateSecurityGroupRuleRequest
type CreateSecurityGroupRuleRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The direction of the flow: `Inbound` or `Outbound`. You can specify `Outbound` for Nets only.
	Flow string `json:"Flow"`
	// The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. If you specify this parameter, you cannot specify the `Rules` parameter and its subparameters.
	FromPortRange int32 `json:"FromPortRange,omitempty"`
	// The IP protocol name (`tcp`, `udp`, `icmp`, or `-1` for all protocols). By default, `-1`. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). If you specify this parameter, you cannot specify the `Rules` parameter and its subparameters.
	IpProtocol string `json:"IpProtocol,omitempty"`
	// The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16). If you specify this parameter, you cannot specify the `Rules` parameter and its subparameters.
	IpRange string `json:"IpRange,omitempty"`
	// Information about the security group rule to create. If you specify this parent parameter and its subparameters, you cannot specify the following parent parameters: `FromPortRange`, `IpProtocol`, `IpRange`, and `ToPortRange`.
	Rules []SecurityGroupRule `json:"Rules,omitempty"`
	// The account ID of the owner of the security group for which you want to create a rule.
	SecurityGroupAccountIdToLink string `json:"SecurityGroupAccountIdToLink,omitempty"`
	// The ID of the security group for which you want to create a rule.
	SecurityGroupId string `json:"SecurityGroupId"`
	// The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group.
	SecurityGroupNameToLink string `json:"SecurityGroupNameToLink,omitempty"`
	// The end of the port range for the TCP and UDP protocols, or an ICMP code number. If you specify this parameter, you cannot specify the `Rules` parameter and its subparameters.
	ToPortRange int32 `json:"ToPortRange,omitempty"`
}
