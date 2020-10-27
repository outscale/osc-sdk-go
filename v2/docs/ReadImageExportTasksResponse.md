# ReadImageExportTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageExportTasks** | Pointer to [**[]ImageExportTask**](ImageExportTask.md) | Information about one or more image export tasks. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadImageExportTasksResponse

`func NewReadImageExportTasksResponse() *ReadImageExportTasksResponse`

NewReadImageExportTasksResponse instantiates a new ReadImageExportTasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadImageExportTasksResponseWithDefaults

`func NewReadImageExportTasksResponseWithDefaults() *ReadImageExportTasksResponse`

NewReadImageExportTasksResponseWithDefaults instantiates a new ReadImageExportTasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageExportTasks

`func (o *ReadImageExportTasksResponse) GetImageExportTasks() []ImageExportTask`

GetImageExportTasks returns the ImageExportTasks field if non-nil, zero value otherwise.

### GetImageExportTasksOk

`func (o *ReadImageExportTasksResponse) GetImageExportTasksOk() (*[]ImageExportTask, bool)`

GetImageExportTasksOk returns a tuple with the ImageExportTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExportTasks

`func (o *ReadImageExportTasksResponse) SetImageExportTasks(v []ImageExportTask)`

SetImageExportTasks sets ImageExportTasks field to given value.

### HasImageExportTasks

`func (o *ReadImageExportTasksResponse) HasImageExportTasks() bool`

HasImageExportTasks returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadImageExportTasksResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadImageExportTasksResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadImageExportTasksResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadImageExportTasksResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


