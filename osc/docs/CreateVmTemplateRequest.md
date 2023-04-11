# CreateVmTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCores** | **int32** | The number of vCores to use for each VM. | 
**CpuGeneration** | **string** | The processor generation to use for each VM (for example, &#x60;v4&#x60;). | 
**CpuPerformance** | **string** | The performance of the VMs (&#x60;medium&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;).  | [optional] [default to CPU_PERFORMANCE_HIGH]
**Description** | **string** | A description for the VM template. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | **string** | The ID of the OMI to use for each VM. You can find a list of OMIs by calling the [ReadImages](#readimages) method. | 
**KeypairName** | **string** | The name of the keypair to use for each VM. | [optional] 
**Ram** | **int32** | The amount of RAM to use for each VM. | 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags to add to the VM template. | [optional] 
**VmTemplateName** | **string** | The name of the VM template. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


