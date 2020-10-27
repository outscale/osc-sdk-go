# UpdateSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Snapshot** | Pointer to [**Snapshot**](Snapshot.md) |  | [optional] 

## Methods

### NewUpdateSnapshotResponse

`func NewUpdateSnapshotResponse() *UpdateSnapshotResponse`

NewUpdateSnapshotResponse instantiates a new UpdateSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotResponseWithDefaults

`func NewUpdateSnapshotResponseWithDefaults() *UpdateSnapshotResponse`

NewUpdateSnapshotResponseWithDefaults instantiates a new UpdateSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *UpdateSnapshotResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UpdateSnapshotResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UpdateSnapshotResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UpdateSnapshotResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSnapshot

`func (o *UpdateSnapshotResponse) GetSnapshot() Snapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *UpdateSnapshotResponse) GetSnapshotOk() (*Snapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *UpdateSnapshotResponse) SetSnapshot(v Snapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *UpdateSnapshotResponse) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


