# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPortRange** | **int32** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 
**IpProtocol** | **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;, or &#x60;-1&#x60; for all protocols). By default, &#x60;-1&#x60;. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). | [optional] 
**IpRanges** | **[]string** | One or more IP ranges for the security group rules, in CIDR notation (for example, &#x60;10.0.0.0/16&#x60;). | [optional] 
**SecurityGroupsMembers** | [**[]SecurityGroupsMember**](SecurityGroupsMember.md) | Information about one or more members of a security group. | [optional] 
**ServiceIds** | **[]string** | One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices). | [optional] 
**ToPortRange** | **int32** | The end of the port range for the TCP and UDP protocols, or an ICMP code number. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


