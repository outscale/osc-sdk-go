# CreateSecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Flow** | **string** | The direction of the flow: &#x60;Inbound&#x60; or &#x60;Outbound&#x60;. You can specify &#x60;Outbound&#x60; for Nets only. | 
**FromPortRange** | **int32** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. If you specify this parameter, you cannot specify the &#x60;Rules&#x60; parameter and its subparameters. | [optional] 
**IpProtocol** | **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;, or &#x60;-1&#x60; for all protocols). By default, &#x60;-1&#x60;. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). If you specify this parameter, you cannot specify the &#x60;Rules&#x60; parameter and its subparameters. | [optional] 
**IpRange** | **string** | The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16). If you specify this parameter, you cannot specify the &#x60;Rules&#x60; parameter and its subparameters. | [optional] 
**Rules** | [**[]SecurityGroupRule**](SecurityGroupRule.md) | Information about the security group rule to create. If you specify this parent parameter and its subparameters, you cannot specify the following parent parameters: &#x60;FromPortRange&#x60;, &#x60;IpProtocol&#x60;, &#x60;IpRange&#x60;, and &#x60;ToPortRange&#x60;. | [optional] 
**SecurityGroupAccountIdToLink** | **string** | The account ID of the owner of the security group for which you want to create a rule. | [optional] 
**SecurityGroupId** | **string** | The ID of the security group for which you want to create a rule. | 
**SecurityGroupNameToLink** | **string** | The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group. | [optional] 
**ToPortRange** | **int32** | The end of the port range for the TCP and UDP protocols, or an ICMP code number. If you specify this parameter, you cannot specify the &#x60;Rules&#x60; parameter and its subparameters. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


