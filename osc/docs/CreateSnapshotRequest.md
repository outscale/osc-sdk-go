# CreateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description for the snapshot. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**FileLocation** | **string** | (When importing) The pre-signed URL of the snapshot you want to import, or the normal URL of the snapshot if you have permission on the OOS bucket. For more information, see [Configuring a Pre-signed URL](https://docs.outscale.com/en/userguide/Configuring-a-Pre-signed-URL.html) or [Managing Access to Your Buckets and Objects](https://docs.outscale.com/en/userguide/Managing-Access-to-Your-Buckets-and-Objects.html). | [optional] 
**SnapshotSize** | **int64** | (When importing) The size of the snapshot you want to create in your account, in bytes. This size must be greater than or equal to the size of the original, uncompressed snapshot. | [optional] 
**SourceRegionName** | **string** | (When copying) The name of the source Region, which must be the same as the Region of your account. | [optional] 
**SourceSnapshotId** | **string** | (When copying) The ID of the snapshot you want to copy. | [optional] 
**VolumeId** | **string** | (When creating) The ID of the volume you want to create a snapshot of. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


