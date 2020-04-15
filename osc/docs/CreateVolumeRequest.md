# CreateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Iops** | **int32** | The number of I/O operations per second (IOPS). This parameter must be specified only if you create an &#x60;io1&#x60; volume. The maximum number of IOPS allowed for &#x60;io1&#x60; volumes is &#x60;13000&#x60;. | [optional] 
**Size** | **int32** | The size of the volume, in gibibytes (GiB). The maximum allowed size for a volume is 14,901 GiB. This parameter is required if the volume is not created from a snapshot (&#x60;SnapshotId&#x60; unspecified).  | [optional] 
**SnapshotId** | **string** | The ID of the snapshot from which you want to create the volume. | [optional] 
**SubregionName** | **string** | The Subregion in which you want to create the volume. | 
**VolumeType** | **string** | The type of volume you want to create (&#x60;io1&#x60; \\| &#x60;gp2&#x60; \\| &#x60;standard&#x60;). If not specified, a &#x60;standard&#x60; volume is created.&lt;br /&gt; For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


