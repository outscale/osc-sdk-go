# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPortRange** | **int32** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 
**IpProtocol** | **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;) or protocol number. By default, &#x60;-1&#x60;, which means all protocols. | [optional] 
**IpRanges** | **[]string** | One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**SecurityGroupsMembers** | [**[]SecurityGroupsMember**](SecurityGroupsMember.md) | Information about one or more members of a security group. | [optional] 
**ServiceIds** | **[]string** | One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices). | [optional] 
**ToPortRange** | **int32** | The end of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


