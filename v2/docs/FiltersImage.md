# FiltersImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAliases** | Pointer to **[]string** | The account aliases of the owners of the OMIs. | [optional] 
**AccountIds** | Pointer to **[]string** | The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described. | [optional] 
**Architectures** | Pointer to **[]string** | The architectures of the OMIs (&#x60;i386&#x60; \\| &#x60;x86_64&#x60;). | [optional] 
**BlockDeviceMappingDeleteOnVmDeletion** | Pointer to **bool** | Indicates whether the block device mapping is deleted when terminating the VM. | [optional] 
**BlockDeviceMappingDeviceNames** | Pointer to **[]string** | The device names for the volumes. | [optional] 
**BlockDeviceMappingSnapshotIds** | Pointer to **[]string** | The IDs of the snapshots used to create the volumes. | [optional] 
**BlockDeviceMappingVolumeSizes** | Pointer to **[]int32** | The sizes of the volumes, in gibibytes (GiB). | [optional] 
**BlockDeviceMappingVolumeTypes** | Pointer to **[]string** | The types of volumes (&#x60;standard&#x60; \\| &#x60;gp2&#x60; \\| &#x60;io1&#x60;). | [optional] 
**Descriptions** | Pointer to **[]string** | The descriptions of the OMIs, provided when they were created. | [optional] 
**FileLocations** | Pointer to **[]string** | The locations of the buckets where the OMI files are stored. | [optional] 
**ImageIds** | Pointer to **[]string** | The IDs of the OMIs. | [optional] 
**ImageNames** | Pointer to **[]string** | The names of the OMIs, provided when they were created. | [optional] 
**PermissionsToLaunchAccountIds** | Pointer to **[]string** | The account IDs of the users who have launch permissions for the OMIs. | [optional] 
**PermissionsToLaunchGlobalPermission** | Pointer to **bool** | If true, lists all public OMIs. If false, lists all private OMIs. | [optional] 
**RootDeviceNames** | Pointer to **[]string** | The device names of the root devices (for example, &#x60;/dev/sda1&#x60;). | [optional] 
**RootDeviceTypes** | Pointer to **[]string** | The types of root device used by the OMIs (always &#x60;bsu&#x60;). | [optional] 
**States** | Pointer to **[]string** | The states of the OMIs (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;failed&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the OMIs. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the OMIs. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the OMIs, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**VirtualizationTypes** | Pointer to **[]string** | The virtualization types (always &#x60;hvm&#x60;). | [optional] 

## Methods

### NewFiltersImage

`func NewFiltersImage() *FiltersImage`

NewFiltersImage instantiates a new FiltersImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersImageWithDefaults

`func NewFiltersImageWithDefaults() *FiltersImage`

NewFiltersImageWithDefaults instantiates a new FiltersImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAliases

`func (o *FiltersImage) GetAccountAliases() []string`

GetAccountAliases returns the AccountAliases field if non-nil, zero value otherwise.

### GetAccountAliasesOk

`func (o *FiltersImage) GetAccountAliasesOk() (*[]string, bool)`

GetAccountAliasesOk returns a tuple with the AccountAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAliases

`func (o *FiltersImage) SetAccountAliases(v []string)`

SetAccountAliases sets AccountAliases field to given value.

### HasAccountAliases

`func (o *FiltersImage) HasAccountAliases() bool`

HasAccountAliases returns a boolean if a field has been set.

### GetAccountIds

`func (o *FiltersImage) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *FiltersImage) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *FiltersImage) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *FiltersImage) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetArchitectures

`func (o *FiltersImage) GetArchitectures() []string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *FiltersImage) GetArchitecturesOk() (*[]string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *FiltersImage) SetArchitectures(v []string)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *FiltersImage) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersImage) GetBlockDeviceMappingDeleteOnVmDeletion() bool`

GetBlockDeviceMappingDeleteOnVmDeletion returns the BlockDeviceMappingDeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetBlockDeviceMappingDeleteOnVmDeletionOk

`func (o *FiltersImage) GetBlockDeviceMappingDeleteOnVmDeletionOk() (*bool, bool)`

GetBlockDeviceMappingDeleteOnVmDeletionOk returns a tuple with the BlockDeviceMappingDeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersImage) SetBlockDeviceMappingDeleteOnVmDeletion(v bool)`

SetBlockDeviceMappingDeleteOnVmDeletion sets BlockDeviceMappingDeleteOnVmDeletion field to given value.

### HasBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersImage) HasBlockDeviceMappingDeleteOnVmDeletion() bool`

HasBlockDeviceMappingDeleteOnVmDeletion returns a boolean if a field has been set.

### GetBlockDeviceMappingDeviceNames

`func (o *FiltersImage) GetBlockDeviceMappingDeviceNames() []string`

GetBlockDeviceMappingDeviceNames returns the BlockDeviceMappingDeviceNames field if non-nil, zero value otherwise.

### GetBlockDeviceMappingDeviceNamesOk

`func (o *FiltersImage) GetBlockDeviceMappingDeviceNamesOk() (*[]string, bool)`

GetBlockDeviceMappingDeviceNamesOk returns a tuple with the BlockDeviceMappingDeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingDeviceNames

`func (o *FiltersImage) SetBlockDeviceMappingDeviceNames(v []string)`

SetBlockDeviceMappingDeviceNames sets BlockDeviceMappingDeviceNames field to given value.

### HasBlockDeviceMappingDeviceNames

`func (o *FiltersImage) HasBlockDeviceMappingDeviceNames() bool`

HasBlockDeviceMappingDeviceNames returns a boolean if a field has been set.

### GetBlockDeviceMappingSnapshotIds

`func (o *FiltersImage) GetBlockDeviceMappingSnapshotIds() []string`

GetBlockDeviceMappingSnapshotIds returns the BlockDeviceMappingSnapshotIds field if non-nil, zero value otherwise.

### GetBlockDeviceMappingSnapshotIdsOk

`func (o *FiltersImage) GetBlockDeviceMappingSnapshotIdsOk() (*[]string, bool)`

GetBlockDeviceMappingSnapshotIdsOk returns a tuple with the BlockDeviceMappingSnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingSnapshotIds

`func (o *FiltersImage) SetBlockDeviceMappingSnapshotIds(v []string)`

SetBlockDeviceMappingSnapshotIds sets BlockDeviceMappingSnapshotIds field to given value.

### HasBlockDeviceMappingSnapshotIds

`func (o *FiltersImage) HasBlockDeviceMappingSnapshotIds() bool`

HasBlockDeviceMappingSnapshotIds returns a boolean if a field has been set.

### GetBlockDeviceMappingVolumeSizes

`func (o *FiltersImage) GetBlockDeviceMappingVolumeSizes() []int32`

GetBlockDeviceMappingVolumeSizes returns the BlockDeviceMappingVolumeSizes field if non-nil, zero value otherwise.

### GetBlockDeviceMappingVolumeSizesOk

`func (o *FiltersImage) GetBlockDeviceMappingVolumeSizesOk() (*[]int32, bool)`

GetBlockDeviceMappingVolumeSizesOk returns a tuple with the BlockDeviceMappingVolumeSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingVolumeSizes

`func (o *FiltersImage) SetBlockDeviceMappingVolumeSizes(v []int32)`

SetBlockDeviceMappingVolumeSizes sets BlockDeviceMappingVolumeSizes field to given value.

### HasBlockDeviceMappingVolumeSizes

`func (o *FiltersImage) HasBlockDeviceMappingVolumeSizes() bool`

HasBlockDeviceMappingVolumeSizes returns a boolean if a field has been set.

### GetBlockDeviceMappingVolumeTypes

`func (o *FiltersImage) GetBlockDeviceMappingVolumeTypes() []string`

GetBlockDeviceMappingVolumeTypes returns the BlockDeviceMappingVolumeTypes field if non-nil, zero value otherwise.

### GetBlockDeviceMappingVolumeTypesOk

`func (o *FiltersImage) GetBlockDeviceMappingVolumeTypesOk() (*[]string, bool)`

GetBlockDeviceMappingVolumeTypesOk returns a tuple with the BlockDeviceMappingVolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMappingVolumeTypes

`func (o *FiltersImage) SetBlockDeviceMappingVolumeTypes(v []string)`

SetBlockDeviceMappingVolumeTypes sets BlockDeviceMappingVolumeTypes field to given value.

### HasBlockDeviceMappingVolumeTypes

`func (o *FiltersImage) HasBlockDeviceMappingVolumeTypes() bool`

HasBlockDeviceMappingVolumeTypes returns a boolean if a field has been set.

### GetDescriptions

`func (o *FiltersImage) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersImage) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *FiltersImage) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *FiltersImage) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetFileLocations

`func (o *FiltersImage) GetFileLocations() []string`

GetFileLocations returns the FileLocations field if non-nil, zero value otherwise.

### GetFileLocationsOk

`func (o *FiltersImage) GetFileLocationsOk() (*[]string, bool)`

GetFileLocationsOk returns a tuple with the FileLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocations

`func (o *FiltersImage) SetFileLocations(v []string)`

SetFileLocations sets FileLocations field to given value.

### HasFileLocations

`func (o *FiltersImage) HasFileLocations() bool`

HasFileLocations returns a boolean if a field has been set.

### GetImageIds

`func (o *FiltersImage) GetImageIds() []string`

GetImageIds returns the ImageIds field if non-nil, zero value otherwise.

### GetImageIdsOk

`func (o *FiltersImage) GetImageIdsOk() (*[]string, bool)`

GetImageIdsOk returns a tuple with the ImageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIds

`func (o *FiltersImage) SetImageIds(v []string)`

SetImageIds sets ImageIds field to given value.

### HasImageIds

`func (o *FiltersImage) HasImageIds() bool`

HasImageIds returns a boolean if a field has been set.

### GetImageNames

`func (o *FiltersImage) GetImageNames() []string`

GetImageNames returns the ImageNames field if non-nil, zero value otherwise.

### GetImageNamesOk

`func (o *FiltersImage) GetImageNamesOk() (*[]string, bool)`

GetImageNamesOk returns a tuple with the ImageNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageNames

`func (o *FiltersImage) SetImageNames(v []string)`

SetImageNames sets ImageNames field to given value.

### HasImageNames

`func (o *FiltersImage) HasImageNames() bool`

HasImageNames returns a boolean if a field has been set.

### GetPermissionsToLaunchAccountIds

`func (o *FiltersImage) GetPermissionsToLaunchAccountIds() []string`

GetPermissionsToLaunchAccountIds returns the PermissionsToLaunchAccountIds field if non-nil, zero value otherwise.

### GetPermissionsToLaunchAccountIdsOk

`func (o *FiltersImage) GetPermissionsToLaunchAccountIdsOk() (*[]string, bool)`

GetPermissionsToLaunchAccountIdsOk returns a tuple with the PermissionsToLaunchAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsToLaunchAccountIds

`func (o *FiltersImage) SetPermissionsToLaunchAccountIds(v []string)`

SetPermissionsToLaunchAccountIds sets PermissionsToLaunchAccountIds field to given value.

### HasPermissionsToLaunchAccountIds

`func (o *FiltersImage) HasPermissionsToLaunchAccountIds() bool`

HasPermissionsToLaunchAccountIds returns a boolean if a field has been set.

### GetPermissionsToLaunchGlobalPermission

`func (o *FiltersImage) GetPermissionsToLaunchGlobalPermission() bool`

GetPermissionsToLaunchGlobalPermission returns the PermissionsToLaunchGlobalPermission field if non-nil, zero value otherwise.

### GetPermissionsToLaunchGlobalPermissionOk

`func (o *FiltersImage) GetPermissionsToLaunchGlobalPermissionOk() (*bool, bool)`

GetPermissionsToLaunchGlobalPermissionOk returns a tuple with the PermissionsToLaunchGlobalPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsToLaunchGlobalPermission

`func (o *FiltersImage) SetPermissionsToLaunchGlobalPermission(v bool)`

SetPermissionsToLaunchGlobalPermission sets PermissionsToLaunchGlobalPermission field to given value.

### HasPermissionsToLaunchGlobalPermission

`func (o *FiltersImage) HasPermissionsToLaunchGlobalPermission() bool`

HasPermissionsToLaunchGlobalPermission returns a boolean if a field has been set.

### GetRootDeviceNames

`func (o *FiltersImage) GetRootDeviceNames() []string`

GetRootDeviceNames returns the RootDeviceNames field if non-nil, zero value otherwise.

### GetRootDeviceNamesOk

`func (o *FiltersImage) GetRootDeviceNamesOk() (*[]string, bool)`

GetRootDeviceNamesOk returns a tuple with the RootDeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceNames

`func (o *FiltersImage) SetRootDeviceNames(v []string)`

SetRootDeviceNames sets RootDeviceNames field to given value.

### HasRootDeviceNames

`func (o *FiltersImage) HasRootDeviceNames() bool`

HasRootDeviceNames returns a boolean if a field has been set.

### GetRootDeviceTypes

`func (o *FiltersImage) GetRootDeviceTypes() []string`

GetRootDeviceTypes returns the RootDeviceTypes field if non-nil, zero value otherwise.

### GetRootDeviceTypesOk

`func (o *FiltersImage) GetRootDeviceTypesOk() (*[]string, bool)`

GetRootDeviceTypesOk returns a tuple with the RootDeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceTypes

`func (o *FiltersImage) SetRootDeviceTypes(v []string)`

SetRootDeviceTypes sets RootDeviceTypes field to given value.

### HasRootDeviceTypes

`func (o *FiltersImage) HasRootDeviceTypes() bool`

HasRootDeviceTypes returns a boolean if a field has been set.

### GetStates

`func (o *FiltersImage) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersImage) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersImage) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersImage) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersImage) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersImage) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersImage) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersImage) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersImage) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersImage) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersImage) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersImage) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersImage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersImage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersImage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVirtualizationTypes

`func (o *FiltersImage) GetVirtualizationTypes() []string`

GetVirtualizationTypes returns the VirtualizationTypes field if non-nil, zero value otherwise.

### GetVirtualizationTypesOk

`func (o *FiltersImage) GetVirtualizationTypesOk() (*[]string, bool)`

GetVirtualizationTypesOk returns a tuple with the VirtualizationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationTypes

`func (o *FiltersImage) SetVirtualizationTypes(v []string)`

SetVirtualizationTypes sets VirtualizationTypes field to given value.

### HasVirtualizationTypes

`func (o *FiltersImage) HasVirtualizationTypes() bool`

HasVirtualizationTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


