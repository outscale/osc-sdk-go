# CreateTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ResourceIds** | **[]string** | One or more resource IDs. | 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags to add to the specified resources. | 

## Methods

### NewCreateTagsRequest

`func NewCreateTagsRequest(resourceIds []string, tags []ResourceTag, ) *CreateTagsRequest`

NewCreateTagsRequest instantiates a new CreateTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagsRequestWithDefaults

`func NewCreateTagsRequestWithDefaults() *CreateTagsRequest`

NewCreateTagsRequestWithDefaults instantiates a new CreateTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateTagsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateTagsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateTagsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateTagsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetResourceIds

`func (o *CreateTagsRequest) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *CreateTagsRequest) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *CreateTagsRequest) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.


### GetTags

`func (o *CreateTagsRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTagsRequest) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTagsRequest) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


