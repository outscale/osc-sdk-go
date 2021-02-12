# UpdateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PermissionsToCreateVolume** | [**PermissionsOnResourceCreation**](PermissionsOnResourceCreation.md) |  | 
**SnapshotId** | **string** | The ID of the snapshot. | 

## Methods

### NewUpdateSnapshotRequest

`func NewUpdateSnapshotRequest(permissionsToCreateVolume PermissionsOnResourceCreation, snapshotId string, ) *UpdateSnapshotRequest`

NewUpdateSnapshotRequest instantiates a new UpdateSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotRequestWithDefaults

`func NewUpdateSnapshotRequestWithDefaults() *UpdateSnapshotRequest`

NewUpdateSnapshotRequestWithDefaults instantiates a new UpdateSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateSnapshotRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateSnapshotRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateSnapshotRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateSnapshotRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPermissionsToCreateVolume

`func (o *UpdateSnapshotRequest) GetPermissionsToCreateVolume() PermissionsOnResourceCreation`

GetPermissionsToCreateVolume returns the PermissionsToCreateVolume field if non-nil, zero value otherwise.

### GetPermissionsToCreateVolumeOk

`func (o *UpdateSnapshotRequest) GetPermissionsToCreateVolumeOk() (*PermissionsOnResourceCreation, bool)`

GetPermissionsToCreateVolumeOk returns a tuple with the PermissionsToCreateVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsToCreateVolume

`func (o *UpdateSnapshotRequest) SetPermissionsToCreateVolume(v PermissionsOnResourceCreation)`

SetPermissionsToCreateVolume sets PermissionsToCreateVolume field to given value.


### GetSnapshotId

`func (o *UpdateSnapshotRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *UpdateSnapshotRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *UpdateSnapshotRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


