# FiltersAccessKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyIds** | Pointer to **[]string** | The IDs of the access keys. | [optional] 
**States** | Pointer to **[]string** | The states of the access keys (&#x60;ACTIVE&#x60; \\| &#x60;INACTIVE&#x60;). | [optional] 

## Methods

### NewFiltersAccessKeys

`func NewFiltersAccessKeys() *FiltersAccessKeys`

NewFiltersAccessKeys instantiates a new FiltersAccessKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersAccessKeysWithDefaults

`func NewFiltersAccessKeysWithDefaults() *FiltersAccessKeys`

NewFiltersAccessKeysWithDefaults instantiates a new FiltersAccessKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyIds

`func (o *FiltersAccessKeys) GetAccessKeyIds() []string`

GetAccessKeyIds returns the AccessKeyIds field if non-nil, zero value otherwise.

### GetAccessKeyIdsOk

`func (o *FiltersAccessKeys) GetAccessKeyIdsOk() (*[]string, bool)`

GetAccessKeyIdsOk returns a tuple with the AccessKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyIds

`func (o *FiltersAccessKeys) SetAccessKeyIds(v []string)`

SetAccessKeyIds sets AccessKeyIds field to given value.

### HasAccessKeyIds

`func (o *FiltersAccessKeys) HasAccessKeyIds() bool`

HasAccessKeyIds returns a boolean if a field has been set.

### GetStates

`func (o *FiltersAccessKeys) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersAccessKeys) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersAccessKeys) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersAccessKeys) HasStates() bool`

HasStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


