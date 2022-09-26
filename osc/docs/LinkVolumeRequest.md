# LinkVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | **string** | The name of the device. For a root device, you must use &#x60;/dev/sda1&#x60;. For other volumes, you must use &#x60;/dev/sdX&#x60;, &#x60;/dev/sdXY&#x60;, &#x60;/dev/xvdX&#x60;, or &#x60;/dev/xvdXY&#x60; (where &#x60;X&#x60; is a letter between &#x60;b&#x60; and &#x60;z&#x60; and where &#x60;Y&#x60; is a letter between &#x60;a&#x60; and &#x60;z&#x60;). | 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VmId** | **string** | The ID of the VM you want to attach the volume to. | 
**VolumeId** | **string** | The ID of the volume you want to attach. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


