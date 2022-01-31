# LinkPrivateIpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRelink** | **bool** | If true, allows an IP that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | **string** | The ID of the NIC. | 
**PrivateIps** | **[]string** | The secondary private IP or IPs you want to assign to the NIC within the IP range of the Subnet. | [optional] 
**SecondaryPrivateIpCount** | **int32** | The number of secondary private IPs to assign to the NIC. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


