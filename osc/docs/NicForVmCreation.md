# NicForVmCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | **bool** | If true, the NIC is deleted when the VM is terminated. You can specify this parameter only for a new NIC. To modify this value for an existing NIC, see [UpdateNic](#updatenic). | [optional] 
**Description** | **string** | The description of the NIC, if you are creating a NIC when creating the VM. | [optional] 
**DeviceNumber** | **int32** | The index of the VM device for the NIC attachment (between 0 and 7, both included). This parameter is required if you create a NIC when creating the VM. | [optional] 
**NicId** | **string** | The ID of the NIC, if you are attaching an existing NIC when creating a VM. | [optional] 
**PrivateIps** | [**[]PrivateIpLight**](PrivateIpLight.md) | One or more private IPs to assign to the NIC, if you create a NIC when creating a VM. Only one private IP can be the primary private IP. | [optional] 
**SecondaryPrivateIpCount** | **int32** | The number of secondary private IPs, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP in the &#x60;PrivateIps&#x60; parameter. | [optional] 
**SecurityGroupIds** | **[]string** | One or more IDs of security groups for the NIC, if you create a NIC when creating a VM. | [optional] 
**SubnetId** | **string** | The ID of the Subnet for the NIC, if you create a NIC when creating a VM. This parameter is required if you create a NIC when creating the VM. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


