# CreateImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | The architecture of the OMI (by default, &#x60;i386&#x60;). | [optional] 
**BlockDeviceMappings** | [**[]BlockDeviceMappingImage**](BlockDeviceMappingImage.md) | One or more block device mappings. | [optional] 
**Description** | **string** | A description for the new OMI. | [optional] 
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**FileLocation** | **string** | The pre-signed URL of the OMI manifest file, or the full path to the OMI stored in an OSU bucket. If you specify this parameter, a copy of the OMI is created in your account. | [optional] 
**ImageName** | **string** | A unique name for the new OMI.&lt;br /&gt; Constraints: 3-128 alphanumeric characters, underscores (_), spaces ( ), parentheses (()), slashes (/), periods (.), or dashes (-). | [optional] 
**NoReboot** | **bool** | If &#x60;false&#x60;, the VM shuts down before creating the OMI and then reboots. If &#x60;true&#x60;, the VM does not. | [optional] 
**RootDeviceName** | **string** | The name of the root device. | [optional] 
**SourceImageId** | **string** | The ID of the OMI you want to copy. | [optional] 
**SourceRegionName** | **string** | The name of the source Region, which must be the same as the Region of your account. | [optional] 
**VmId** | **string** | The ID of the VM from which you want to create the OMI. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


