# DeleteTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ResourceIds** | **[]string** | One or more resource IDs. | 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags to delete (if you set a tag value, only the tags matching exactly this value are deleted). | 

## Methods

### NewDeleteTagsRequest

`func NewDeleteTagsRequest(resourceIds []string, tags []ResourceTag, ) *DeleteTagsRequest`

NewDeleteTagsRequest instantiates a new DeleteTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTagsRequestWithDefaults

`func NewDeleteTagsRequestWithDefaults() *DeleteTagsRequest`

NewDeleteTagsRequestWithDefaults instantiates a new DeleteTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteTagsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteTagsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteTagsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteTagsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetResourceIds

`func (o *DeleteTagsRequest) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *DeleteTagsRequest) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *DeleteTagsRequest) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.


### GetTags

`func (o *DeleteTagsRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeleteTagsRequest) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeleteTagsRequest) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


