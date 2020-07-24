# Nic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID of the owner of the NIC. | [optional] 
**Description** | **string** | The description of the NIC. | [optional] 
**IsSourceDestChecked** | **bool** | (Net only) If &#x60;true&#x60;, the source/destination check is enabled. If &#x60;false&#x60;, it is disabled. This value must be &#x60;false&#x60; for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**LinkNic** | [**LinkNic**](LinkNic.md) |  | [optional] 
**LinkPublicIp** | [**LinkPublicIp**](LinkPublicIp.md) |  | [optional] 
**MacAddress** | **string** | The Media Access Control (MAC) address of the NIC. | [optional] 
**NetId** | **string** | The ID of the Net for the NIC. | [optional] 
**NicId** | **string** | The ID of the NIC. | [optional] 
**PrivateDnsName** | **string** | The name of the private DNS. | [optional] 
**PrivateIps** | [**[]PrivateIp**](PrivateIp.md) | The private IP addresses of the NIC. | [optional] 
**SecurityGroups** | [**[]SecurityGroupLight**](SecurityGroupLight.md) | One or more IDs of security groups for the NIC. | [optional] 
**State** | **string** | The state of the NIC (&#x60;available&#x60; \\| &#x60;attaching&#x60; \\| &#x60;in-use&#x60; \\| &#x60;detaching&#x60;). | [optional] 
**SubnetId** | **string** | The ID of the Subnet. | [optional] 
**SubregionName** | **string** | The Subregion in which the NIC is located. | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the NIC. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


