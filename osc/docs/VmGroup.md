# VmGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | **string** | The date and time of creation of the VM group. | [optional] 
**Description** | **string** | The description of the VM group. | [optional] 
**PositioningStrategy** | **string** | The positioning strategy of the VMs on hypervisors. By default, or if set to &#x60;no-strategy&#x60;, TINA determines the most adequate position for the VMs. If set to &#x60;attract&#x60;, the VMs are deployed on the same hypervisor, which improves network performance. If set to &#x60;repulse&#x60;, the VMs are deployed on a different hypervisor, which improves fault tolerance. | [optional] 
**SecurityGroupIds** | **[]string** | One or more IDs of security groups for the VM group. | [optional] 
**State** | **string** | The state of the VM group (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;scaling up&#x60; \\| &#x60;scaling down&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**SubnetId** | **string** | The ID of the Subnet for the VM group. | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the VM group. | [optional] 
**VmCount** | **int32** | The number of VMs in the VM group. | [optional] 
**VmGroupId** | **string** | The ID of the VM group. | [optional] 
**VmGroupName** | **string** | The name of the VM group. | [optional] 
**VmIds** | **[]string** | The IDs of the VMs in the VM group. | [optional] 
**VmTemplateId** | **string** | The ID of the VM template used by the VM group. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


