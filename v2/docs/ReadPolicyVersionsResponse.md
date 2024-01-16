# ReadPolicyVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMoreItems** | Pointer to **bool** | If true, there are more items to return using the &#x60;FirstItem&#x60; parameter in a new request. | [optional] 
**MaxResultsLimit** | Pointer to **int32** | Indicates maximum results defined for the operation. | [optional] 
**PolicyVersions** | Pointer to [**[]PolicyVersion**](PolicyVersion.md) | A list of all the versions of the policy. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadPolicyVersionsResponse

`func NewReadPolicyVersionsResponse() *ReadPolicyVersionsResponse`

NewReadPolicyVersionsResponse instantiates a new ReadPolicyVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPolicyVersionsResponseWithDefaults

`func NewReadPolicyVersionsResponseWithDefaults() *ReadPolicyVersionsResponse`

NewReadPolicyVersionsResponseWithDefaults instantiates a new ReadPolicyVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMoreItems

`func (o *ReadPolicyVersionsResponse) GetHasMoreItems() bool`

GetHasMoreItems returns the HasMoreItems field if non-nil, zero value otherwise.

### GetHasMoreItemsOk

`func (o *ReadPolicyVersionsResponse) GetHasMoreItemsOk() (*bool, bool)`

GetHasMoreItemsOk returns a tuple with the HasMoreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreItems

`func (o *ReadPolicyVersionsResponse) SetHasMoreItems(v bool)`

SetHasMoreItems sets HasMoreItems field to given value.

### HasHasMoreItems

`func (o *ReadPolicyVersionsResponse) HasHasMoreItems() bool`

HasHasMoreItems returns a boolean if a field has been set.

### GetMaxResultsLimit

`func (o *ReadPolicyVersionsResponse) GetMaxResultsLimit() int32`

GetMaxResultsLimit returns the MaxResultsLimit field if non-nil, zero value otherwise.

### GetMaxResultsLimitOk

`func (o *ReadPolicyVersionsResponse) GetMaxResultsLimitOk() (*int32, bool)`

GetMaxResultsLimitOk returns a tuple with the MaxResultsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultsLimit

`func (o *ReadPolicyVersionsResponse) SetMaxResultsLimit(v int32)`

SetMaxResultsLimit sets MaxResultsLimit field to given value.

### HasMaxResultsLimit

`func (o *ReadPolicyVersionsResponse) HasMaxResultsLimit() bool`

HasMaxResultsLimit returns a boolean if a field has been set.

### GetPolicyVersions

`func (o *ReadPolicyVersionsResponse) GetPolicyVersions() []PolicyVersion`

GetPolicyVersions returns the PolicyVersions field if non-nil, zero value otherwise.

### GetPolicyVersionsOk

`func (o *ReadPolicyVersionsResponse) GetPolicyVersionsOk() (*[]PolicyVersion, bool)`

GetPolicyVersionsOk returns a tuple with the PolicyVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersions

`func (o *ReadPolicyVersionsResponse) SetPolicyVersions(v []PolicyVersion)`

SetPolicyVersions sets PolicyVersions field to given value.

### HasPolicyVersions

`func (o *ReadPolicyVersionsResponse) HasPolicyVersions() bool`

HasPolicyVersions returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadPolicyVersionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadPolicyVersionsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadPolicyVersionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadPolicyVersionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


