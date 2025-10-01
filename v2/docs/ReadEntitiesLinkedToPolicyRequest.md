# ReadEntitiesLinkedToPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitiesType** | Pointer to **[]string** | The type of entity linked to the policy you want to get information about. | [optional] 
**FirstItem** | Pointer to **int32** | The item starting the list of entities requested. | [optional] 
**PolicyOrn** | **string** | The OUTSCALE Resource Name (ORN) of the policy. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of items that can be returned in a single response (by default, 100). | [optional] 

## Methods

### NewReadEntitiesLinkedToPolicyRequest

`func NewReadEntitiesLinkedToPolicyRequest(policyOrn string, ) *ReadEntitiesLinkedToPolicyRequest`

NewReadEntitiesLinkedToPolicyRequest instantiates a new ReadEntitiesLinkedToPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadEntitiesLinkedToPolicyRequestWithDefaults

`func NewReadEntitiesLinkedToPolicyRequestWithDefaults() *ReadEntitiesLinkedToPolicyRequest`

NewReadEntitiesLinkedToPolicyRequestWithDefaults instantiates a new ReadEntitiesLinkedToPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitiesType

`func (o *ReadEntitiesLinkedToPolicyRequest) GetEntitiesType() []string`

GetEntitiesType returns the EntitiesType field if non-nil, zero value otherwise.

### GetEntitiesTypeOk

`func (o *ReadEntitiesLinkedToPolicyRequest) GetEntitiesTypeOk() (*[]string, bool)`

GetEntitiesTypeOk returns a tuple with the EntitiesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitiesType

`func (o *ReadEntitiesLinkedToPolicyRequest) SetEntitiesType(v []string)`

SetEntitiesType sets EntitiesType field to given value.

### HasEntitiesType

`func (o *ReadEntitiesLinkedToPolicyRequest) HasEntitiesType() bool`

HasEntitiesType returns a boolean if a field has been set.

### GetFirstItem

`func (o *ReadEntitiesLinkedToPolicyRequest) GetFirstItem() int32`

GetFirstItem returns the FirstItem field if non-nil, zero value otherwise.

### GetFirstItemOk

`func (o *ReadEntitiesLinkedToPolicyRequest) GetFirstItemOk() (*int32, bool)`

GetFirstItemOk returns a tuple with the FirstItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItem

`func (o *ReadEntitiesLinkedToPolicyRequest) SetFirstItem(v int32)`

SetFirstItem sets FirstItem field to given value.

### HasFirstItem

`func (o *ReadEntitiesLinkedToPolicyRequest) HasFirstItem() bool`

HasFirstItem returns a boolean if a field has been set.

### GetPolicyOrn

`func (o *ReadEntitiesLinkedToPolicyRequest) GetPolicyOrn() string`

GetPolicyOrn returns the PolicyOrn field if non-nil, zero value otherwise.

### GetPolicyOrnOk

`func (o *ReadEntitiesLinkedToPolicyRequest) GetPolicyOrnOk() (*string, bool)`

GetPolicyOrnOk returns a tuple with the PolicyOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOrn

`func (o *ReadEntitiesLinkedToPolicyRequest) SetPolicyOrn(v string)`

SetPolicyOrn sets PolicyOrn field to given value.


### GetResponseContext

`func (o *ReadEntitiesLinkedToPolicyRequest) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadEntitiesLinkedToPolicyRequest) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadEntitiesLinkedToPolicyRequest) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadEntitiesLinkedToPolicyRequest) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadEntitiesLinkedToPolicyRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadEntitiesLinkedToPolicyRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadEntitiesLinkedToPolicyRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadEntitiesLinkedToPolicyRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


