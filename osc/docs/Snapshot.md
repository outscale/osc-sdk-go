# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | **string** | The account alias of the owner of the snapshot. | [optional] 
**AccountId** | **string** | The account ID of the owner of the snapshot. | [optional] 
**Description** | **string** | The description of the snapshot. | [optional] 
**PermissionsToCreateVolume** | [**PermissionsOnResource**](PermissionsOnResource.md) |  | [optional] 
**Progress** | **int32** | The progress of the snapshot, as a percentage. | [optional] 
**SnapshotId** | **string** | The ID of the snapshot. | [optional] 
**State** | **string** | The state of the snapshot (&#x60;in-queue&#x60; \\| &#x60;pending&#x60; \\| &#x60;completed&#x60;). | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the snapshot. | [optional] 
**VolumeId** | **string** | The ID of the volume used to create the snapshot. | [optional] 
**VolumeSize** | **int32** | The size of the volume used to create the snapshot, in gibibytes (GiB). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


