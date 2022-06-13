# FiltersSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIds** | **[]string** | The account IDs of the owners of the security groups. | [optional] 
**Descriptions** | **[]string** | The descriptions of the security groups. | [optional] 
**InboundRuleAccountIds** | **[]string** | The account IDs that have been granted permissions. | [optional] 
**InboundRuleFromPortRanges** | **[]int32** | The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers. | [optional] 
**InboundRuleIpRanges** | **[]string** | The IP ranges that have been granted permissions, in CIDR notation (for example, &#x60;10.0.0.0/24&#x60;). | [optional] 
**InboundRuleProtocols** | **[]string** | The IP protocols for the permissions (&#x60;tcp&#x60; \\| &#x60;udp&#x60; \\| &#x60;icmp&#x60;, or a protocol number, or &#x60;-1&#x60; for all protocols). | [optional] 
**InboundRuleSecurityGroupIds** | **[]string** | The IDs of the security groups that have been granted permissions. | [optional] 
**InboundRuleSecurityGroupNames** | **[]string** | The names of the security groups that have been granted permissions. | [optional] 
**InboundRuleToPortRanges** | **[]int32** | The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers. | [optional] 
**NetIds** | **[]string** | The IDs of the Nets specified when the security groups were created. | [optional] 
**OutboundRuleAccountIds** | **[]string** | The account IDs that have been granted permissions. | [optional] 
**OutboundRuleFromPortRanges** | **[]int32** | The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers. | [optional] 
**OutboundRuleIpRanges** | **[]string** | The IP ranges that have been granted permissions, in CIDR notation (for example, &#x60;10.0.0.0/24&#x60;). | [optional] 
**OutboundRuleProtocols** | **[]string** | The IP protocols for the permissions (&#x60;tcp&#x60; \\| &#x60;udp&#x60; \\| &#x60;icmp&#x60;, or a protocol number, or &#x60;-1&#x60; for all protocols). | [optional] 
**OutboundRuleSecurityGroupIds** | **[]string** | The IDs of the security groups that have been granted permissions. | [optional] 
**OutboundRuleSecurityGroupNames** | **[]string** | The names of the security groups that have been granted permissions. | [optional] 
**OutboundRuleToPortRanges** | **[]int32** | The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers. | [optional] 
**SecurityGroupIds** | **[]string** | The IDs of the security groups. | [optional] 
**SecurityGroupNames** | **[]string** | The names of the security groups. | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the security groups. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the security groups. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the security groups, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


