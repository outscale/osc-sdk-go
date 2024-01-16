# ReadPoliciesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMoreItems** | Pointer to **bool** | If true, there are more items to return using the &#x60;FirstItem&#x60; parameter in a new request. | [optional] 
**MaxResultsLimit** | Pointer to **int32** | Indicates maximum results defined for the operation. | [optional] 
**MaxResultsTruncated** | Pointer to **bool** | If true, indicates whether requested page size is more than allowed. | [optional] 
**Policies** | Pointer to [**[]Policy**](Policy.md) | Information about one or more policies. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadPoliciesResponse

`func NewReadPoliciesResponse() *ReadPoliciesResponse`

NewReadPoliciesResponse instantiates a new ReadPoliciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPoliciesResponseWithDefaults

`func NewReadPoliciesResponseWithDefaults() *ReadPoliciesResponse`

NewReadPoliciesResponseWithDefaults instantiates a new ReadPoliciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMoreItems

`func (o *ReadPoliciesResponse) GetHasMoreItems() bool`

GetHasMoreItems returns the HasMoreItems field if non-nil, zero value otherwise.

### GetHasMoreItemsOk

`func (o *ReadPoliciesResponse) GetHasMoreItemsOk() (*bool, bool)`

GetHasMoreItemsOk returns a tuple with the HasMoreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreItems

`func (o *ReadPoliciesResponse) SetHasMoreItems(v bool)`

SetHasMoreItems sets HasMoreItems field to given value.

### HasHasMoreItems

`func (o *ReadPoliciesResponse) HasHasMoreItems() bool`

HasHasMoreItems returns a boolean if a field has been set.

### GetMaxResultsLimit

`func (o *ReadPoliciesResponse) GetMaxResultsLimit() int32`

GetMaxResultsLimit returns the MaxResultsLimit field if non-nil, zero value otherwise.

### GetMaxResultsLimitOk

`func (o *ReadPoliciesResponse) GetMaxResultsLimitOk() (*int32, bool)`

GetMaxResultsLimitOk returns a tuple with the MaxResultsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsLimit

`func (o *ReadPoliciesResponse) SetMaxResultsLimit(v int32)`

SetMaxResultsLimit sets MaxResultsLimit field to given value.

### HasMaxResultsLimit

`func (o *ReadPoliciesResponse) HasMaxResultsLimit() bool`

HasMaxResultsLimit returns a boolean if a field has been set.

### GetMaxResultsTruncated

`func (o *ReadPoliciesResponse) GetMaxResultsTruncated() bool`

GetMaxResultsTruncated returns the MaxResultsTruncated field if non-nil, zero value otherwise.

### GetMaxResultsTruncatedOk

`func (o *ReadPoliciesResponse) GetMaxResultsTruncatedOk() (*bool, bool)`

GetMaxResultsTruncatedOk returns a tuple with the MaxResultsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsTruncated

`func (o *ReadPoliciesResponse) SetMaxResultsTruncated(v bool)`

SetMaxResultsTruncated sets MaxResultsTruncated field to given value.

### HasMaxResultsTruncated

`func (o *ReadPoliciesResponse) HasMaxResultsTruncated() bool`

HasMaxResultsTruncated returns a boolean if a field has been set.

### GetPolicies

`func (o *ReadPoliciesResponse) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ReadPoliciesResponse) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ReadPoliciesResponse) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ReadPoliciesResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadPoliciesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadPoliciesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadPoliciesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadPoliciesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


