# CreateSecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Flow** | **string** | The direction of the flow: &#x60;Inbound&#x60; or &#x60;Outbound&#x60;. You can specify &#x60;Outbound&#x60; for Nets only. | 
**FromPortRange** | **int32** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 
**IpProtocol** | **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;) or protocol number. By default, &#x60;-1&#x60;, which means all protocols. | [optional] 
**IpRange** | **string** | The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**Rules** | [**[]SecurityGroupRule**](SecurityGroupRule.md) | Information about the security group rule to create. | [optional] 
**SecurityGroupAccountIdToLink** | **string** | The account ID of the owner of the security group for which you want to create a rule. | [optional] 
**SecurityGroupId** | **string** | The ID of the security group for which you want to create a rule. | 
**SecurityGroupNameToLink** | **string** | The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group. | [optional] 
**ToPortRange** | **int32** | The end of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


