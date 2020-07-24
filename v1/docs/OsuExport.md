# OsuExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskImageFormat** | **string** | The format of the export disk (&#x60;qcow2&#x60; \\| &#x60;vdi&#x60; \\| &#x60;vmdk&#x60;). | 
**OsuApiKey** | [**OsuApiKey**](OsuApiKey.md) |  | [optional] 
**OsuBucket** | **string** | The name of the OSU bucket you want to export the object to. | 
**OsuManifestUrl** | **string** | The URL of the manifest file. | [optional] 
**OsuPrefix** | **string** | The prefix for the key of the OSU object. This key follows this format: &#x60;prefix + object_export_task_id + &#39;.&#39; + disk_image_format&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


