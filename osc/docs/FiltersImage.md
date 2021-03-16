# FiltersImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAliases** | **[]string** | The account aliases of the owners of the OMIs. | [optional] 
**AccountIds** | **[]string** | The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described. | [optional] 
**Architectures** | **[]string** | The architectures of the OMIs (&#x60;i386&#x60; \\| &#x60;x86_64&#x60;). | [optional] 
**BlockDeviceMappingDeleteOnVmDeletion** | **bool** | Indicates whether the block device mapping is deleted when terminating the VM. | [optional] 
**BlockDeviceMappingDeviceNames** | **[]string** | The device names for the volumes. | [optional] 
**BlockDeviceMappingSnapshotIds** | **[]string** | The IDs of the snapshots used to create the volumes. | [optional] 
**BlockDeviceMappingVolumeSizes** | **[]int32** | The sizes of the volumes, in gibibytes (GiB). | [optional] 
**BlockDeviceMappingVolumeTypes** | **[]string** | The types of volumes (&#x60;standard&#x60; \\| &#x60;gp2&#x60; \\| &#x60;io1&#x60;). | [optional] 
**Descriptions** | **[]string** | The descriptions of the OMIs, provided when they were created. | [optional] 
**FileLocations** | **[]string** | The locations where the OMI files are stored on Object Storage Unit (OSU). | [optional] 
**Hypervisors** | **[]string** | The hypervisor type of the OMI (always &#x60;xen&#x60;). | [optional] 
**ImageIds** | **[]string** | The IDs of the OMIs. | [optional] 
**ImageNames** | **[]string** | The names of the OMIs, provided when they were created. | [optional] 
**PermissionsToLaunchAccountIds** | **[]string** | The account IDs of the users who have launch permissions for the OMIs. | [optional] 
**PermissionsToLaunchGlobalPermission** | **bool** | If true, lists all public OMIs. If false, lists all private OMIs. | [optional] 
**ProductCodes** | **[]string** | The product code associated with the OMI (&#x60;0001&#x60; Linux/Unix \\| &#x60;0002&#x60; Windows \\| &#x60;0004&#x60; Linux/Oracle \\| &#x60;0005&#x60; Windows 10). | [optional] 
**RootDeviceNames** | **[]string** | The device names of the root devices (for example, &#x60;/dev/sda1&#x60;). | [optional] 
**RootDeviceTypes** | **[]string** | The types of root device used by the OMIs (always &#x60;bsu&#x60;). | [optional] 
**States** | **[]string** | The states of the OMIs (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;failed&#x60;). | [optional] 
**TagKeys** | **[]string** | The keys of the tags associated with the OMIs. | [optional] 
**TagValues** | **[]string** | The values of the tags associated with the OMIs. | [optional] 
**Tags** | **[]string** | The key/value combination of the tags associated with the OMIs, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**VirtualizationTypes** | **[]string** | The virtualization types (always &#x60;hvm&#x60;). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


