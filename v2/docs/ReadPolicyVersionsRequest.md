# ReadPolicyVersionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstItem** | Pointer to **int32** | The item starting the list of policies requested. | [optional] 
**PolicyOrn** | **string** | The OUTSCALE Resource Name (ORN) of the policy. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | 
**ResultsPerPage** | Pointer to **int32** | The maximum number of items that can be returned in a single response (by default, 100). | [optional] 

## Methods

### NewReadPolicyVersionsRequest

`func NewReadPolicyVersionsRequest(policyOrn string, ) *ReadPolicyVersionsRequest`

NewReadPolicyVersionsRequest instantiates a new ReadPolicyVersionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPolicyVersionsRequestWithDefaults

`func NewReadPolicyVersionsRequestWithDefaults() *ReadPolicyVersionsRequest`

NewReadPolicyVersionsRequestWithDefaults instantiates a new ReadPolicyVersionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstItem

`func (o *ReadPolicyVersionsRequest) GetFirstItem() int32`

GetFirstItem returns the FirstItem field if non-nil, zero value otherwise.

### GetFirstItemOk

`func (o *ReadPolicyVersionsRequest) GetFirstItemOk() (*int32, bool)`

GetFirstItemOk returns a tuple with the FirstItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItem

`func (o *ReadPolicyVersionsRequest) SetFirstItem(v int32)`

SetFirstItem sets FirstItem field to given value.

### HasFirstItem

`func (o *ReadPolicyVersionsRequest) HasFirstItem() bool`

HasFirstItem returns a boolean if a field has been set.

### GetPolicyOrn

`func (o *ReadPolicyVersionsRequest) GetPolicyOrn() string`

GetPolicyOrn returns the PolicyOrn field if non-nil, zero value otherwise.

### GetPolicyOrnOk

`func (o *ReadPolicyVersionsRequest) GetPolicyOrnOk() (*string, bool)`

GetPolicyOrnOk returns a tuple with the PolicyOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOrn

`func (o *ReadPolicyVersionsRequest) SetPolicyOrn(v string)`

SetPolicyOrn sets PolicyOrn field to given value.


### GetResultsPerPage

`func (o *ReadPolicyVersionsRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadPolicyVersionsRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadPolicyVersionsRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadPolicyVersionsRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


