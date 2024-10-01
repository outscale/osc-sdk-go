# ReadUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMoreItems** | Pointer to **bool** | If true, there are more items to return using the &#x60;FirstItem&#x60; parameter in a new request. | [optional] 
**MaxResultsLimit** | Pointer to **int32** | Indicates maximum results defined for the operation. | [optional] 
**MaxResultsTruncated** | Pointer to **bool** | If true, indicates whether requested page size is more than allowed. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) | A list of EIM users. | [optional] 

## Methods

### NewReadUsersResponse

`func NewReadUsersResponse() *ReadUsersResponse`

NewReadUsersResponse instantiates a new ReadUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUsersResponseWithDefaults

`func NewReadUsersResponseWithDefaults() *ReadUsersResponse`

NewReadUsersResponseWithDefaults instantiates a new ReadUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMoreItems

`func (o *ReadUsersResponse) GetHasMoreItems() bool`

GetHasMoreItems returns the HasMoreItems field if non-nil, zero value otherwise.

### GetHasMoreItemsOk

`func (o *ReadUsersResponse) GetHasMoreItemsOk() (*bool, bool)`

GetHasMoreItemsOk returns a tuple with the HasMoreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreItems

`func (o *ReadUsersResponse) SetHasMoreItems(v bool)`

SetHasMoreItems sets HasMoreItems field to given value.

### HasHasMoreItems

`func (o *ReadUsersResponse) HasHasMoreItems() bool`

HasHasMoreItems returns a boolean if a field has been set.

### GetMaxResultsLimit

`func (o *ReadUsersResponse) GetMaxResultsLimit() int32`

GetMaxResultsLimit returns the MaxResultsLimit field if non-nil, zero value otherwise.

### GetMaxResultsLimitOk

`func (o *ReadUsersResponse) GetMaxResultsLimitOk() (*int32, bool)`

GetMaxResultsLimitOk returns a tuple with the MaxResultsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsLimit

`func (o *ReadUsersResponse) SetMaxResultsLimit(v int32)`

SetMaxResultsLimit sets MaxResultsLimit field to given value.

### HasMaxResultsLimit

`func (o *ReadUsersResponse) HasMaxResultsLimit() bool`

HasMaxResultsLimit returns a boolean if a field has been set.

### GetMaxResultsTruncated

`func (o *ReadUsersResponse) GetMaxResultsTruncated() bool`

GetMaxResultsTruncated returns the MaxResultsTruncated field if non-nil, zero value otherwise.

### GetMaxResultsTruncatedOk

`func (o *ReadUsersResponse) GetMaxResultsTruncatedOk() (*bool, bool)`

GetMaxResultsTruncatedOk returns a tuple with the MaxResultsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsTruncated

`func (o *ReadUsersResponse) SetMaxResultsTruncated(v bool)`

SetMaxResultsTruncated sets MaxResultsTruncated field to given value.

### HasMaxResultsTruncated

`func (o *ReadUsersResponse) HasMaxResultsTruncated() bool`

HasMaxResultsTruncated returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadUsersResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadUsersResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadUsersResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadUsersResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetUsers

`func (o *ReadUsersResponse) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ReadUsersResponse) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ReadUsersResponse) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ReadUsersResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


