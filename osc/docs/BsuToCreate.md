# BsuToCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | **bool** | By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM. | [optional] [default to true]
**Iops** | **int32** | The number of I/O operations per second (IOPS). This parameter must be specified only if you create an &#x60;io1&#x60; volume. The maximum number of IOPS allowed for &#x60;io1&#x60; volumes is &#x60;13000&#x60; with a maximum performance ratio of 300 IOPS per gibibyte. | [optional] 
**SnapshotId** | **string** | The ID of the snapshot used to create the volume. | [optional] 
**VolumeSize** | **int32** | The size of the volume, in gibibytes (GiB).&lt;br /&gt; If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.&lt;br /&gt; If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one. | [optional] 
**VolumeType** | **string** | The type of the volume (&#x60;standard&#x60; \\| &#x60;io1&#x60; \\| &#x60;gp2&#x60;). If not specified in the request, a &#x60;standard&#x60; volume is created.&lt;br /&gt; For more information about volume types, see [About Volumes &gt; Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


