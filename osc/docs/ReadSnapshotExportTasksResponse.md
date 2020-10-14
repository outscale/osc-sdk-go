# ReadSnapshotExportTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**SnapshotExportTasks** | Pointer to [**[]SnapshotExportTask**](SnapshotExportTask.md) | Information about one or more snapshot export tasks. | [optional] 

## Methods

### NewReadSnapshotExportTasksResponse

`func NewReadSnapshotExportTasksResponse() *ReadSnapshotExportTasksResponse`

NewReadSnapshotExportTasksResponse instantiates a new ReadSnapshotExportTasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSnapshotExportTasksResponseWithDefaults

`func NewReadSnapshotExportTasksResponseWithDefaults() *ReadSnapshotExportTasksResponse`

NewReadSnapshotExportTasksResponseWithDefaults instantiates a new ReadSnapshotExportTasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadSnapshotExportTasksResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSnapshotExportTasksResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadSnapshotExportTasksResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadSnapshotExportTasksResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSnapshotExportTasks

`func (o *ReadSnapshotExportTasksResponse) GetSnapshotExportTasks() []SnapshotExportTask`

GetSnapshotExportTasks returns the SnapshotExportTasks field if non-nil, zero value otherwise.

### GetSnapshotExportTasksOk

`func (o *ReadSnapshotExportTasksResponse) GetSnapshotExportTasksOk() (*[]SnapshotExportTask, bool)`

GetSnapshotExportTasksOk returns a tuple with the SnapshotExportTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExportTasks

`func (o *ReadSnapshotExportTasksResponse) SetSnapshotExportTasks(v []SnapshotExportTask)`

SetSnapshotExportTasks sets SnapshotExportTasks field to given value.

### HasSnapshotExportTasks

`func (o *ReadSnapshotExportTasksResponse) HasSnapshotExportTasks() bool`

HasSnapshotExportTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


