# DeleteSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**SnapshotId** | **string** | The ID of the snapshot you want to delete. | 

## Methods

### NewDeleteSnapshotRequest

`func NewDeleteSnapshotRequest(snapshotId string, ) *DeleteSnapshotRequest`

NewDeleteSnapshotRequest instantiates a new DeleteSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSnapshotRequestWithDefaults

`func NewDeleteSnapshotRequestWithDefaults() *DeleteSnapshotRequest`

NewDeleteSnapshotRequestWithDefaults instantiates a new DeleteSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteSnapshotRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteSnapshotRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteSnapshotRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteSnapshotRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetSnapshotId

`func (o *DeleteSnapshotRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *DeleteSnapshotRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *DeleteSnapshotRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


