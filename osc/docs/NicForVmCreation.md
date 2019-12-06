# NicForVmCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | **bool** | If &#x60;true&#x60;, the NIC is deleted when the VM is terminated. You can specify &#x60;true&#x60; only if you create a NIC when creating a VM. | [optional] 
**Description** | **string** | The description of the NIC, if you are creating a NIC when creating the VM. | [optional] 
**DeviceNumber** | **int32** | The index of the VM device for the NIC attachment (between 0 and 7, both included). This parameter is required if you create a NIC when creating the VM. | [optional] 
**NicId** | **string** | The ID of the NIC, if you are attaching an existing NIC when creating a VM. | [optional] 
**PrivateIps** | [**[]PrivateIpLight**](PrivateIpLight.md) | One or more private IP addresses to assign to the NIC, if you create a NIC when creating a VM. Only one private IP address can be the primary private IP address. | [optional] 
**SecondaryPrivateIpCount** | **int32** | The number of secondary private IP addresses, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP address in the &#x60;PrivateIps&#x60; parameter. | [optional] 
**SecurityGroupIds** | **[]string** | One or more IDs of security groups for the NIC, if you acreate a NIC when creating a VM. | [optional] 
**SubnetId** | **string** | The ID of the Subnet for the NIC, if you create a NIC when creating a VM. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


