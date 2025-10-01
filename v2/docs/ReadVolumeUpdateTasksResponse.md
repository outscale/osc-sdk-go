# ReadVolumeUpdateTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VolumeUpdateTasks** | Pointer to [**[]VolumeUpdateTask**](VolumeUpdateTask.md) | Information about one or more volume update tasks. | [optional] 

## Methods

### NewReadVolumeUpdateTasksResponse

`func NewReadVolumeUpdateTasksResponse() *ReadVolumeUpdateTasksResponse`

NewReadVolumeUpdateTasksResponse instantiates a new ReadVolumeUpdateTasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVolumeUpdateTasksResponseWithDefaults

`func NewReadVolumeUpdateTasksResponseWithDefaults() *ReadVolumeUpdateTasksResponse`

NewReadVolumeUpdateTasksResponseWithDefaults instantiates a new ReadVolumeUpdateTasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ReadVolumeUpdateTasksResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadVolumeUpdateTasksResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadVolumeUpdateTasksResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadVolumeUpdateTasksResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadVolumeUpdateTasksResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVolumeUpdateTasksResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadVolumeUpdateTasksResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadVolumeUpdateTasksResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVolumeUpdateTasks

`func (o *ReadVolumeUpdateTasksResponse) GetVolumeUpdateTasks() []VolumeUpdateTask`

GetVolumeUpdateTasks returns the VolumeUpdateTasks field if non-nil, zero value otherwise.

### GetVolumeUpdateTasksOk

`func (o *ReadVolumeUpdateTasksResponse) GetVolumeUpdateTasksOk() (*[]VolumeUpdateTask, bool)`

GetVolumeUpdateTasksOk returns a tuple with the VolumeUpdateTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUpdateTasks

`func (o *ReadVolumeUpdateTasksResponse) SetVolumeUpdateTasks(v []VolumeUpdateTask)`

SetVolumeUpdateTasks sets VolumeUpdateTasks field to given value.

### HasVolumeUpdateTasks

`func (o *ReadVolumeUpdateTasksResponse) HasVolumeUpdateTasks() bool`

HasVolumeUpdateTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


