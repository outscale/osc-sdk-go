# FiltersVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDates** | **[]string** | The dates and times of creation of the volumes, in ISO 8601 date-time format (for example, &#x60;2020-06-30T00:00:00.000Z&#x60;). | [optional] 
**LinkVolumeDeleteOnVmDeletion** | **bool** | Whether the volumes are deleted or not when terminating the VMs. | [optional] 
**LinkVolumeDeviceNames** | **[]string** | The VM device names. | [optional] 
**LinkVolumeLinkDates** | **[]string** | The dates and times of creation of the volumes, in ISO 8601 date-time format (for example, &#x60;2020-06-30T00:00:00.000Z&#x60;). | [optional] 
**LinkVolumeLinkStates** | **[]string** | The attachment states of the volumes (&#x60;attaching&#x60; \\| &#x60;detaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detached&#x60;). | [optional] 
**LinkVolumeVmIds** | **[]string** | One or more IDs of VMs. | [optional] 
**SnapshotIds** | **[]string** | The snapshots from which the volumes were created. | [optional] 
**SubregionNames** | **[]string** | The names of the Subregions in which the volumes were created. | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the volumes. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the volumes. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the volumes, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**VolumeIds** | **[]string** | The IDs of the volumes. | [optional] 
**VolumeSizes** | **[]int32** | The sizes of the volumes, in gibibytes (GiB). | [optional] 
**VolumeStates** | **[]string** | The states of the volumes (&#x60;creating&#x60; \\| &#x60;available&#x60; \\| &#x60;in-use&#x60; \\| &#x60;updating&#x60; \\| &#x60;deleting&#x60; \\| &#x60;error&#x60;). | [optional] 
**VolumeTypes** | **[]string** | The types of the volumes (&#x60;standard&#x60; \\| &#x60;gp2&#x60; \\| &#x60;io1&#x60;). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


