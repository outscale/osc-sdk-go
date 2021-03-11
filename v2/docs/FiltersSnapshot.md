# FiltersSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAliases** | Pointer to **[]string** | The account aliases of the owners of the snapshots. | [optional] 
**AccountIds** | Pointer to **[]string** | The account IDs of the owners of the snapshots. | [optional] 
**Descriptions** | Pointer to **[]string** | The descriptions of the snapshots. | [optional] 
**PermissionsToCreateVolumeAccountIds** | Pointer to **[]string** | The account IDs of one or more users who have permissions to create volumes. | [optional] 
**PermissionsToCreateVolumeGlobalPermission** | Pointer to **bool** | If true, lists all public volumes. If false, lists all private volumes. | [optional] 
**Progresses** | Pointer to **[]int32** | The progresses of the snapshots, as a percentage. | [optional] 
**SnapshotIds** | Pointer to **[]string** | The IDs of the snapshots. | [optional] 
**States** | Pointer to **[]string** | The states of the snapshots (&#x60;in-queue&#x60; \\| &#x60;completed&#x60; \\| &#x60;error&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the snapshots. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the snapshots. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the snapshots, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**VolumeIds** | Pointer to **[]string** | The IDs of the volumes used to create the snapshots. | [optional] 
**VolumeSizes** | Pointer to **[]int32** | The sizes of the volumes used to create the snapshots, in gibibytes (GiB). | [optional] 

## Methods

### NewFiltersSnapshot

`func NewFiltersSnapshot() *FiltersSnapshot`

NewFiltersSnapshot instantiates a new FiltersSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersSnapshotWithDefaults

`func NewFiltersSnapshotWithDefaults() *FiltersSnapshot`

NewFiltersSnapshotWithDefaults instantiates a new FiltersSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAliases

`func (o *FiltersSnapshot) GetAccountAliases() []string`

GetAccountAliases returns the AccountAliases field if non-nil, zero value otherwise.

### GetAccountAliasesOk

`func (o *FiltersSnapshot) GetAccountAliasesOk() (*[]string, bool)`

GetAccountAliasesOk returns a tuple with the AccountAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAliases

`func (o *FiltersSnapshot) SetAccountAliases(v []string)`

SetAccountAliases sets AccountAliases field to given value.

### HasAccountAliases

`func (o *FiltersSnapshot) HasAccountAliases() bool`

HasAccountAliases returns a boolean if a field has been set.

### GetAccountIds

`func (o *FiltersSnapshot) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *FiltersSnapshot) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *FiltersSnapshot) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *FiltersSnapshot) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetDescriptions

`func (o *FiltersSnapshot) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersSnapshot) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *FiltersSnapshot) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *FiltersSnapshot) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetPermissionsToCreateVolumeAccountIds

`func (o *FiltersSnapshot) GetPermissionsToCreateVolumeAccountIds() []string`

GetPermissionsToCreateVolumeAccountIds returns the PermissionsToCreateVolumeAccountIds field if non-nil, zero value otherwise.

### GetPermissionsToCreateVolumeAccountIdsOk

`func (o *FiltersSnapshot) GetPermissionsToCreateVolumeAccountIdsOk() (*[]string, bool)`

GetPermissionsToCreateVolumeAccountIdsOk returns a tuple with the PermissionsToCreateVolumeAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsToCreateVolumeAccountIds

`func (o *FiltersSnapshot) SetPermissionsToCreateVolumeAccountIds(v []string)`

SetPermissionsToCreateVolumeAccountIds sets PermissionsToCreateVolumeAccountIds field to given value.

### HasPermissionsToCreateVolumeAccountIds

`func (o *FiltersSnapshot) HasPermissionsToCreateVolumeAccountIds() bool`

HasPermissionsToCreateVolumeAccountIds returns a boolean if a field has been set.

### GetPermissionsToCreateVolumeGlobalPermission

`func (o *FiltersSnapshot) GetPermissionsToCreateVolumeGlobalPermission() bool`

GetPermissionsToCreateVolumeGlobalPermission returns the PermissionsToCreateVolumeGlobalPermission field if non-nil, zero value otherwise.

### GetPermissionsToCreateVolumeGlobalPermissionOk

`func (o *FiltersSnapshot) GetPermissionsToCreateVolumeGlobalPermissionOk() (*bool, bool)`

GetPermissionsToCreateVolumeGlobalPermissionOk returns a tuple with the PermissionsToCreateVolumeGlobalPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsToCreateVolumeGlobalPermission

`func (o *FiltersSnapshot) SetPermissionsToCreateVolumeGlobalPermission(v bool)`

SetPermissionsToCreateVolumeGlobalPermission sets PermissionsToCreateVolumeGlobalPermission field to given value.

### HasPermissionsToCreateVolumeGlobalPermission

`func (o *FiltersSnapshot) HasPermissionsToCreateVolumeGlobalPermission() bool`

HasPermissionsToCreateVolumeGlobalPermission returns a boolean if a field has been set.

### GetProgresses

`func (o *FiltersSnapshot) GetProgresses() []int32`

GetProgresses returns the Progresses field if non-nil, zero value otherwise.

### GetProgressesOk

`func (o *FiltersSnapshot) GetProgressesOk() (*[]int32, bool)`

GetProgressesOk returns a tuple with the Progresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgresses

`func (o *FiltersSnapshot) SetProgresses(v []int32)`

SetProgresses sets Progresses field to given value.

### HasProgresses

`func (o *FiltersSnapshot) HasProgresses() bool`

HasProgresses returns a boolean if a field has been set.

### GetSnapshotIds

`func (o *FiltersSnapshot) GetSnapshotIds() []string`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *FiltersSnapshot) GetSnapshotIdsOk() (*[]string, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIds

`func (o *FiltersSnapshot) SetSnapshotIds(v []string)`

SetSnapshotIds sets SnapshotIds field to given value.

### HasSnapshotIds

`func (o *FiltersSnapshot) HasSnapshotIds() bool`

HasSnapshotIds returns a boolean if a field has been set.

### GetStates

`func (o *FiltersSnapshot) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersSnapshot) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersSnapshot) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersSnapshot) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersSnapshot) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersSnapshot) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersSnapshot) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersSnapshot) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersSnapshot) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersSnapshot) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersSnapshot) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersSnapshot) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersSnapshot) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersSnapshot) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersSnapshot) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersSnapshot) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumeIds

`func (o *FiltersSnapshot) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *FiltersSnapshot) GetVolumeIdsOk() (*[]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIds

`func (o *FiltersSnapshot) SetVolumeIds(v []string)`

SetVolumeIds sets VolumeIds field to given value.

### HasVolumeIds

`func (o *FiltersSnapshot) HasVolumeIds() bool`

HasVolumeIds returns a boolean if a field has been set.

### GetVolumeSizes

`func (o *FiltersSnapshot) GetVolumeSizes() []int32`

GetVolumeSizes returns the VolumeSizes field if non-nil, zero value otherwise.

### GetVolumeSizesOk

`func (o *FiltersSnapshot) GetVolumeSizesOk() (*[]int32, bool)`

GetVolumeSizesOk returns a tuple with the VolumeSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSizes

`func (o *FiltersSnapshot) SetVolumeSizes(v []int32)`

SetVolumeSizes sets VolumeSizes field to given value.

### HasVolumeSizes

`func (o *FiltersSnapshot) HasVolumeSizes() bool`

HasVolumeSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


