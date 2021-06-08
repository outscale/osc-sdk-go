# ImageExportTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | If the OMI export task fails, an error message appears. | [optional] 
**ImageId** | Pointer to **string** | The ID of the OMI to be exported. | [optional] 
**OsuExport** | Pointer to [**OsuExportImageExportTask**](OsuExportImageExportTask.md) |  | [optional] 
**Progress** | Pointer to **int32** | The progress of the OMI export task, as a percentage. | [optional] 
**State** | Pointer to **string** | The state of the OMI export task (&#x60;pending/queued&#x60; \\| &#x60;pending&#x60; \\| &#x60;completed&#x60; \\| &#x60;failed&#x60; \\| &#x60;cancelled&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the image export task. | [optional] 
**TaskId** | Pointer to **string** | The ID of the OMI export task. | [optional] 

## Methods

### NewImageExportTask

`func NewImageExportTask() *ImageExportTask`

NewImageExportTask instantiates a new ImageExportTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageExportTaskWithDefaults

`func NewImageExportTaskWithDefaults() *ImageExportTask`

NewImageExportTaskWithDefaults instantiates a new ImageExportTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ImageExportTask) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ImageExportTask) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ImageExportTask) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ImageExportTask) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetImageId

`func (o *ImageExportTask) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageExportTask) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ImageExportTask) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ImageExportTask) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetOsuExport

`func (o *ImageExportTask) GetOsuExport() OsuExportImageExportTask`

GetOsuExport returns the OsuExport field if non-nil, zero value otherwise.

### GetOsuExportOk

`func (o *ImageExportTask) GetOsuExportOk() (*OsuExportImageExportTask, bool)`

GetOsuExportOk returns a tuple with the OsuExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuExport

`func (o *ImageExportTask) SetOsuExport(v OsuExportImageExportTask)`

SetOsuExport sets OsuExport field to given value.

### HasOsuExport

`func (o *ImageExportTask) HasOsuExport() bool`

HasOsuExport returns a boolean if a field has been set.

### GetProgress

`func (o *ImageExportTask) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ImageExportTask) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ImageExportTask) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ImageExportTask) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetState

`func (o *ImageExportTask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ImageExportTask) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ImageExportTask) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ImageExportTask) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *ImageExportTask) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImageExportTask) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImageExportTask) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImageExportTask) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskId

`func (o *ImageExportTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ImageExportTask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ImageExportTask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ImageExportTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


