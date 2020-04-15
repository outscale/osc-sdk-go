# FiltersSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAliases** | **[]string** | The account aliases of the owners of the snapshots. | [optional] 
**AccountIds** | **[]string** | The account IDs of the owners of the snapshots. | [optional] 
**Descriptions** | **[]string** | The descriptions of the snapshots. | [optional] 
**PermissionsToCreateVolumeAccountIds** | **[]string** | The account IDs of one or more users who have permissions to create volumes. | [optional] 
**PermissionsToCreateVolumeGlobalPermission** | **bool** | If &#x60;true&#x60;, lists all public volumes. If &#x60;false&#x60;, lists all private volumes. | [optional] 
**Progresses** | **[]int32** | The progresses of the snapshots, as a percentage. | [optional] 
**SnapshotIds** | **[]string** | The IDs of the snapshots. | [optional] 
**States** | **[]string** | The states of the snapshots (&#x60;in-queue&#x60; \\| &#x60;completed&#x60; \\| &#x60;error&#x60;). | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the snapshots. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the snapshots. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the snapshots, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 
**VolumeIds** | **[]string** | The IDs of the volumes used to create the snapshots. | [optional] 
**VolumeSizes** | **[]int32** | The sizes of the volumes used to create the snapshots, in gibibytes (GiB). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


