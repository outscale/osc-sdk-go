# CreateImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | The architecture of the OMI (by default, &#x60;i386&#x60; if you specified the &#x60;FileLocation&#x60; or &#x60;RootDeviceName&#x60; parameter). | [optional] 
**BlockDeviceMappings** | [**[]BlockDeviceMappingImage**](BlockDeviceMappingImage.md) | One or more block device mappings. | [optional] 
**Description** | **string** | A description for the new OMI. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**FileLocation** | **string** | The pre-signed URL of the OMI manifest file, or the full path to the OMI stored in a bucket. If you specify this parameter, a copy of the OMI is created in your account. You must specify only one of the following parameters: &#x60;FileLocation&#x60;, &#x60;RootDeviceName&#x60;, &#x60;SourceImageId&#x60; or &#x60;VmId&#x60;. | [optional] 
**ImageName** | **string** | A unique name for the new OMI.&lt;br /&gt; Constraints: 3-128 alphanumeric characters, underscores (_), spaces ( ), parentheses (()), slashes (/), periods (.), or dashes (-). | [optional] 
**NoReboot** | **bool** | If false, the VM shuts down before creating the OMI and then reboots. If true, the VM does not. | [optional] 
**RootDeviceName** | **string** | The name of the root device. You must specify only one of the following parameters: &#x60;FileLocation&#x60;, &#x60;RootDeviceName&#x60;, &#x60;SourceImageId&#x60; or &#x60;VmId&#x60;. | [optional] 
**SourceImageId** | **string** | The ID of the OMI you want to copy. You must specify only one of the following parameters: &#x60;FileLocation&#x60;, &#x60;RootDeviceName&#x60;, &#x60;SourceImageId&#x60; or &#x60;VmId&#x60;. | [optional] 
**SourceRegionName** | **string** | The name of the source Region, which must be the same as the Region of your account. | [optional] 
**VmId** | **string** | The ID of the VM from which you want to create the OMI. You must specify only one of the following parameters: &#x60;FileLocation&#x60;, &#x60;RootDeviceName&#x60;, &#x60;SourceImageId&#x60; or &#x60;VmId&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


