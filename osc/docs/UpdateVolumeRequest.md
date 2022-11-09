# UpdateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Iops** | **int32** | The new number of I/O operations per second (IOPS). This parameter can be specified only if you update an &#x60;io1&#x60; volume. The maximum number of IOPS allowed for &#x60;io1&#x60; volumes is &#x60;13000&#x60; with a maximum performance ratio of 300 IOPS per gibibyte. This modification is instantaneous on a cold volume, not on a hot one. | [optional] 
**Size** | **int32** | (cold volume only) The new size of the volume, in gibibytes (GiB). This value must be equal to or greater than the current size of the volume. This modification is not instantaneous. | [optional] 
**VolumeId** | **string** | The ID of the volume you want to update. | 
**VolumeType** | **string** | (cold volume only) The new type of the volume (&#x60;standard&#x60; \\| &#x60;io1&#x60; \\| &#x60;gp2&#x60;). This modification is instantaneous. If you update to an &#x60;io1&#x60; volume, you must also specify the &#x60;Iops&#x60; parameter. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


