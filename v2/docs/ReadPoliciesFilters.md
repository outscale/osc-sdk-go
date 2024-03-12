# ReadPoliciesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnlyLinked** | Pointer to **bool** | If set to true, lists only the policies attached to a user. | [optional] 
**PathPrefix** | Pointer to **string** | The path prefix you can use to filter the results, set to a slash (&#x60;/&#x60;) by default. | [optional] 
**Scope** | Pointer to **string** | The scope to filter policies (&#x60;OWS&#x60; \\| &#x60;LOCAL&#x60;). | [optional] 

## Methods

### NewReadPoliciesFilters

`func NewReadPoliciesFilters() *ReadPoliciesFilters`

NewReadPoliciesFilters instantiates a new ReadPoliciesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPoliciesFiltersWithDefaults

`func NewReadPoliciesFiltersWithDefaults() *ReadPoliciesFilters`

NewReadPoliciesFiltersWithDefaults instantiates a new ReadPoliciesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnlyLinked

`func (o *ReadPoliciesFilters) GetOnlyLinked() bool`

GetOnlyLinked returns the OnlyLinked field if non-nil, zero value otherwise.

### GetOnlyLinkedOk

`func (o *ReadPoliciesFilters) GetOnlyLinkedOk() (*bool, bool)`

GetOnlyLinkedOk returns a tuple with the OnlyLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyLinked

`func (o *ReadPoliciesFilters) SetOnlyLinked(v bool)`

SetOnlyLinked sets OnlyLinked field to given value.

### HasOnlyLinked

`func (o *ReadPoliciesFilters) HasOnlyLinked() bool`

HasOnlyLinked returns a boolean if a field has been set.

### GetPathPrefix

`func (o *ReadPoliciesFilters) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *ReadPoliciesFilters) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *ReadPoliciesFilters) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *ReadPoliciesFilters) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetScope

`func (o *ReadPoliciesFilters) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ReadPoliciesFilters) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ReadPoliciesFilters) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ReadPoliciesFilters) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


