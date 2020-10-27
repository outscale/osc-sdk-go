# ReadSnapshotsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Snapshots** | Pointer to [**[]Snapshot**](Snapshot.md) | Information about one or more snapshots and their permissions. | [optional] 

## Methods

### NewReadSnapshotsResponse

`func NewReadSnapshotsResponse() *ReadSnapshotsResponse`

NewReadSnapshotsResponse instantiates a new ReadSnapshotsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSnapshotsResponseWithDefaults

`func NewReadSnapshotsResponseWithDefaults() *ReadSnapshotsResponse`

NewReadSnapshotsResponseWithDefaults instantiates a new ReadSnapshotsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadSnapshotsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSnapshotsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadSnapshotsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadSnapshotsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSnapshots

`func (o *ReadSnapshotsResponse) GetSnapshots() []Snapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ReadSnapshotsResponse) GetSnapshotsOk() (*[]Snapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ReadSnapshotsResponse) SetSnapshots(v []Snapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ReadSnapshotsResponse) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


