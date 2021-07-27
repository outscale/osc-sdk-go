# FiltersNic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptions** | **[]string** | The descriptions of the NICs. | [optional] 
**IsSourceDestCheck** | **bool** | Whether the source/destination checking is enabled (true) or disabled (false). | [optional] 
**LinkNicDeleteOnVmDeletion** | **bool** | Whether the NICs are deleted when the VMs they are attached to are terminated. | [optional] 
**LinkNicDeviceNumbers** | **[]int32** | The device numbers the NICs are attached to. | [optional] 
**LinkNicLinkNicIds** | **[]string** | The attachment IDs of the NICs. | [optional] 
**LinkNicStates** | **[]string** | The states of the attachments. | [optional] 
**LinkNicVmAccountIds** | **[]string** | The account IDs of the owners of the VMs the NICs are attached to. | [optional] 
**LinkNicVmIds** | **[]string** | The IDs of the VMs the NICs are attached to. | [optional] 
**LinkPublicIpAccountIds** | **[]string** | The account IDs of the owners of the EIPs associated with the NICs. | [optional] 
**LinkPublicIpLinkPublicIpIds** | **[]string** | The association IDs returned when the EIPs were associated with the NICs. | [optional] 
**LinkPublicIpPublicIpIds** | **[]string** | The allocation IDs returned when the EIPs were allocated to their accounts. | [optional] 
**LinkPublicIpPublicIps** | **[]string** | The EIPs associated with the NICs. | [optional] 
**MacAddresses** | **[]string** | The Media Access Control (MAC) addresses of the NICs. | [optional] 
**NetIds** | **[]string** | The IDs of the Nets where the NICs are located. | [optional] 
**NicIds** | **[]string** | The IDs of the NICs. | [optional] 
**PrivateDnsNames** | **[]string** | The private DNS names associated with the primary private IP addresses. | [optional] 
**PrivateIpsLinkPublicIpAccountIds** | **[]string** | The account IDs of the owner of the EIPs associated with the private IP addresses. | [optional] 
**PrivateIpsLinkPublicIpPublicIps** | **[]string** | The EIPs associated with the private IP addresses. | [optional] 
**PrivateIpsPrimaryIp** | **bool** | The primary private IP addresses of the NICs. | [optional] 
**PrivateIpsPrivateIps** | **[]string** | The private IP addresses of the NICs. | [optional] 
**SecurityGroupIds** | **[]string** | The IDs of the security groups associated with the NICs. | [optional] 
**SecurityGroupNames** | **[]string** | The names of the security groups associated with the NICs. | [optional] 
**States** | **[]string** | The states of the NICs. | [optional] 
**SubnetIds** | **[]string** | The IDs of the Subnets for the NICs. | [optional] 
**SubregionNames** | **[]string** | The Subregions where the NICs are located. | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the NICs. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the NICs. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the NICs, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


