# CreateVmGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description for the VM group. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PositioningStrategy** | **string** | The positioning strategy of VMs on hypervisors. By default, or if set to &#x60;no-strategy&#x60; our orchestrator determines the most adequate position for your VMs. If set to &#x60;attract&#x60;, your VMs are deployed on the same hypervisor, which improves network performance. If set to &#x60;repulse&#x60;, your VMs are deployed on a different hypervisor, which improves fault tolerance. | [optional] [default to POSITIONING_STRATEGY_NO_STRATEGY]
**SecurityGroupIds** | **[]string** | One or more IDs of security groups for the VM group. | 
**SubnetId** | **string** | The ID of the Subnet in which you want to create the VM group. | 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags to add to the VM group. | [optional] 
**VmCount** | **int32** | The number of VMs deployed in the VM group. | 
**VmGroupName** | **string** | The name of the VM group. | 
**VmTemplateId** | **string** | The ID of the VM template used to launch VMs in the VM group. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


