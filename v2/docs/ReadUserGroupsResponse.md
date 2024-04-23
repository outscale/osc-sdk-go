# ReadUserGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMoreItems** | Pointer to **bool** | If true, there are more items to return using the &#x60;FirstItem&#x60; parameter in a new request. | [optional] 
**MaxResultsLimit** | Pointer to **int32** | Indicates maximum results defined for the operation. | [optional] 
**MaxResultsTruncated** | Pointer to **bool** | If true, indicates whether requested page size is more than allowed. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**UserGroups** | Pointer to [**[]UserGroup**](UserGroup.md) | A list of user groups. | [optional] 

## Methods

### NewReadUserGroupsResponse

`func NewReadUserGroupsResponse() *ReadUserGroupsResponse`

NewReadUserGroupsResponse instantiates a new ReadUserGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserGroupsResponseWithDefaults

`func NewReadUserGroupsResponseWithDefaults() *ReadUserGroupsResponse`

NewReadUserGroupsResponseWithDefaults instantiates a new ReadUserGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMoreItems

`func (o *ReadUserGroupsResponse) GetHasMoreItems() bool`

GetHasMoreItems returns the HasMoreItems field if non-nil, zero value otherwise.

### GetHasMoreItemsOk

`func (o *ReadUserGroupsResponse) GetHasMoreItemsOk() (*bool, bool)`

GetHasMoreItemsOk returns a tuple with the HasMoreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreItems

`func (o *ReadUserGroupsResponse) SetHasMoreItems(v bool)`

SetHasMoreItems sets HasMoreItems field to given value.

### HasHasMoreItems

`func (o *ReadUserGroupsResponse) HasHasMoreItems() bool`

HasHasMoreItems returns a boolean if a field has been set.

### GetMaxResultsLimit

`func (o *ReadUserGroupsResponse) GetMaxResultsLimit() int32`

GetMaxResultsLimit returns the MaxResultsLimit field if non-nil, zero value otherwise.

### GetMaxResultsLimitOk

`func (o *ReadUserGroupsResponse) GetMaxResultsLimitOk() (*int32, bool)`

GetMaxResultsLimitOk returns a tuple with the MaxResultsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsLimit

`func (o *ReadUserGroupsResponse) SetMaxResultsLimit(v int32)`

SetMaxResultsLimit sets MaxResultsLimit field to given value.

### HasMaxResultsLimit

`func (o *ReadUserGroupsResponse) HasMaxResultsLimit() bool`

HasMaxResultsLimit returns a boolean if a field has been set.

### GetMaxResultsTruncated

`func (o *ReadUserGroupsResponse) GetMaxResultsTruncated() bool`

GetMaxResultsTruncated returns the MaxResultsTruncated field if non-nil, zero value otherwise.

### GetMaxResultsTruncatedOk

`func (o *ReadUserGroupsResponse) GetMaxResultsTruncatedOk() (*bool, bool)`

GetMaxResultsTruncatedOk returns a tuple with the MaxResultsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsTruncated

`func (o *ReadUserGroupsResponse) SetMaxResultsTruncated(v bool)`

SetMaxResultsTruncated sets MaxResultsTruncated field to given value.

### HasMaxResultsTruncated

`func (o *ReadUserGroupsResponse) HasMaxResultsTruncated() bool`

HasMaxResultsTruncated returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadUserGroupsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadUserGroupsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadUserGroupsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadUserGroupsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetUserGroups

`func (o *ReadUserGroupsResponse) GetUserGroups() []UserGroup`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *ReadUserGroupsResponse) GetUserGroupsOk() (*[]UserGroup, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *ReadUserGroupsResponse) SetUserGroups(v []UserGroup)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *ReadUserGroupsResponse) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


