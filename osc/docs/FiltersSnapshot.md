# FiltersSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAliases** | **[]string** | The account aliases of the owners of the snapshots. | [optional] 
**AccountIds** | **[]string** | The account IDs of the owners of the snapshots. | [optional] 
**Descriptions** | **[]string** | The descriptions of the snapshots. | [optional] 
**FromCreationDate** | [**time.Time**](time.Time.md) | The beginning of the time period, in ISO 8601 date-time format (for example, &#x60;2020-06-14T00:00:00.000Z&#x60;). | [optional] 
**PermissionsToCreateVolumeAccountIds** | **[]string** | The account IDs of one or more users who have permissions to create volumes. | [optional] 
**PermissionsToCreateVolumeGlobalPermission** | **bool** | If true, lists all public volumes. If false, lists all private volumes. | [optional] 
**Progresses** | **[]int32** | The progresses of the snapshots, as a percentage. | [optional] 
**SnapshotIds** | **[]string** | The IDs of the snapshots. | [optional] 
**States** | **[]string** | The states of the snapshots (&#x60;in-queue&#x60; \\| &#x60;completed&#x60; \\| &#x60;error&#x60;). | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the snapshots. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the snapshots. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the snapshots, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**ToCreationDate** | [**time.Time**](time.Time.md) | The end of the time period, in ISO 8601 date-time format (for example, &#x60;2020-06-30T00:00:00.000Z&#x60;). | [optional] 
**VolumeIds** | **[]string** | The IDs of the volumes used to create the snapshots. | [optional] 
**VolumeSizes** | **[]int32** | The sizes of the volumes used to create the snapshots, in gibibytes (GiB). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


