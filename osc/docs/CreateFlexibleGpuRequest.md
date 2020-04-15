# CreateFlexibleGpuRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | **bool** | If &#x60;true&#x60;, the fGPU is deleted when the VM is terminated. | [optional] [default to false]
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Generation** | **string** | The processor generation that the fGPU must be compatible with. If not specified, the oldest possible processor generation is selected (as provided by **[ReadFlexibleGpuCatalog](#readflexiblegpucatalog)** for the specified model of fGPU). | [optional] 
**ModelName** | **string** | The model of fGPU you want to allocate. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs). | 
**SubregionName** | **string** | The Subregion in which you want to create the fGPU. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


