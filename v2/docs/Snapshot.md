# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | Pointer to **string** | The account alias of the owner of the snapshot. | [optional] 
**AccountId** | Pointer to **string** | The account ID of the owner of the snapshot. | [optional] 
**CreationDate** | Pointer to **string** | The date and time of creation of the snapshot. | [optional] 
**Description** | Pointer to **string** | The description of the snapshot. | [optional] 
**PermissionsToCreateVolume** | Pointer to [**PermissionsOnResource**](PermissionsOnResource.md) |  | [optional] 
**Progress** | Pointer to **int32** | The progress of the snapshot, as a percentage. | [optional] 
**SnapshotId** | Pointer to **string** | The ID of the snapshot. | [optional] 
**State** | Pointer to **string** | The state of the snapshot (&#x60;in-queue&#x60; \\| &#x60;completed&#x60; \\| &#x60;error&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the snapshot. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume used to create the snapshot. | [optional] 
**VolumeSize** | Pointer to **int32** | The size of the volume used to create the snapshot, in gibibytes (GiB). | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot() *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *Snapshot) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *Snapshot) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *Snapshot) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *Snapshot) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetAccountId

`func (o *Snapshot) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Snapshot) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Snapshot) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Snapshot) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreationDate

`func (o *Snapshot) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Snapshot) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Snapshot) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Snapshot) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *Snapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Snapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Snapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Snapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissionsToCreateVolume

`func (o *Snapshot) GetPermissionsToCreateVolume() PermissionsOnResource`

GetPermissionsToCreateVolume returns the PermissionsToCreateVolume field if non-nil, zero value otherwise.

### GetPermissionsToCreateVolumeOk

`func (o *Snapshot) GetPermissionsToCreateVolumeOk() (*PermissionsOnResource, bool)`

GetPermissionsToCreateVolumeOk returns a tuple with the PermissionsToCreateVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsToCreateVolume

`func (o *Snapshot) SetPermissionsToCreateVolume(v PermissionsOnResource)`

SetPermissionsToCreateVolume sets PermissionsToCreateVolume field to given value.

### HasPermissionsToCreateVolume

`func (o *Snapshot) HasPermissionsToCreateVolume() bool`

HasPermissionsToCreateVolume returns a boolean if a field has been set.

### GetProgress

`func (o *Snapshot) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Snapshot) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Snapshot) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Snapshot) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSnapshotId

`func (o *Snapshot) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *Snapshot) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *Snapshot) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *Snapshot) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetState

`func (o *Snapshot) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Snapshot) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Snapshot) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Snapshot) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *Snapshot) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Snapshot) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Snapshot) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Snapshot) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumeId

`func (o *Snapshot) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *Snapshot) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *Snapshot) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *Snapshot) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetVolumeSize

`func (o *Snapshot) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *Snapshot) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *Snapshot) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *Snapshot) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


