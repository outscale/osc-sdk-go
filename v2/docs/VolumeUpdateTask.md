# VolumeUpdateTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | If the update volume task fails, an error message appears. | [optional] 
**CompletionDate** | Pointer to **string** | The date at which the volume update task was marked as completed. | [optional] 
**Progress** | Pointer to **int32** | The progress of the volume update task, as a percentage. | [optional] 
**StartDate** | Pointer to **string** | The creation date of the volume update task. | [optional] 
**State** | Pointer to **string** | The state of the volume (&#x60;pending&#x60; \\| &#x60;active&#x60; \\| &#x60;completed&#x60; \\| &#x60;failed&#x60; \\| &#x60;canceled&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the volume update task. | [optional] 
**TaskId** | Pointer to **string** | The ID of the volume update task in progress. Otherwise, it is not returned. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the updated volume. | [optional] 
**VolumeUpdate** | Pointer to [**VolumeUpdate**](VolumeUpdate.md) |  | [optional] 

## Methods

### NewVolumeUpdateTask

`func NewVolumeUpdateTask() *VolumeUpdateTask`

NewVolumeUpdateTask instantiates a new VolumeUpdateTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeUpdateTaskWithDefaults

`func NewVolumeUpdateTaskWithDefaults() *VolumeUpdateTask`

NewVolumeUpdateTaskWithDefaults instantiates a new VolumeUpdateTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *VolumeUpdateTask) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *VolumeUpdateTask) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *VolumeUpdateTask) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *VolumeUpdateTask) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCompletionDate

`func (o *VolumeUpdateTask) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *VolumeUpdateTask) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *VolumeUpdateTask) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *VolumeUpdateTask) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetProgress

`func (o *VolumeUpdateTask) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VolumeUpdateTask) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VolumeUpdateTask) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VolumeUpdateTask) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStartDate

`func (o *VolumeUpdateTask) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *VolumeUpdateTask) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *VolumeUpdateTask) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *VolumeUpdateTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetState

`func (o *VolumeUpdateTask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VolumeUpdateTask) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VolumeUpdateTask) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VolumeUpdateTask) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *VolumeUpdateTask) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VolumeUpdateTask) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VolumeUpdateTask) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VolumeUpdateTask) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskId

`func (o *VolumeUpdateTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *VolumeUpdateTask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *VolumeUpdateTask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *VolumeUpdateTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetVolumeId

`func (o *VolumeUpdateTask) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeUpdateTask) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeUpdateTask) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VolumeUpdateTask) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetVolumeUpdate

`func (o *VolumeUpdateTask) GetVolumeUpdate() VolumeUpdate`

GetVolumeUpdate returns the VolumeUpdate field if non-nil, zero value otherwise.

### GetVolumeUpdateOk

`func (o *VolumeUpdateTask) GetVolumeUpdateOk() (*VolumeUpdate, bool)`

GetVolumeUpdateOk returns a tuple with the VolumeUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUpdate

`func (o *VolumeUpdateTask) SetVolumeUpdate(v VolumeUpdate)`

SetVolumeUpdate sets VolumeUpdate field to given value.

### HasVolumeUpdate

`func (o *VolumeUpdateTask) HasVolumeUpdate() bool`

HasVolumeUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


