# LinkPrivateIpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRelink** | **bool** | If &#x60;true&#x60;, allows an IP address that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified. | [optional] 
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | **string** | The ID of the NIC. | 
**PrivateIps** | **[]string** | The secondary private IP address or addresses you want to assign to the NIC within the IP address range of the Subnet. | [optional] 
**SecondaryPrivateIpCount** | **int32** | The number of secondary private IP addresses to assign to the NIC. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


