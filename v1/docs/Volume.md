# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iops** | **int32** | The number of I/O operations per second (IOPS):&lt;br /&gt; - For &#x60;io1&#x60; volumes, the number of provisioned IOPS&lt;br /&gt; - For &#x60;gp2&#x60; volumes, the baseline performance of the volume | [optional] 
**LinkedVolumes** | [**[]LinkedVolume**](LinkedVolume.md) | Information about your volume attachment. | [optional] 
**Size** | **int32** | The size of the volume, in gibibytes (GiB). | [optional] 
**SnapshotId** | **string** | The snapshot from which the volume was created. | [optional] 
**State** | **string** | The state of the volume (&#x60;creating&#x60; \\| &#x60;available&#x60; \\| &#x60;in-use&#x60; \\| &#x60;deleting&#x60; \\| &#x60;error&#x60;). | [optional] 
**SubregionName** | **string** | The Subregion in which the volume was created. | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the volume. | [optional] 
**VolumeId** | **string** | The ID of the volume. | [optional] 
**VolumeType** | **string** | The type of the volume (&#x60;standard&#x60; \\| &#x60;gp2&#x60; \\| &#x60;io1&#x60;). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


