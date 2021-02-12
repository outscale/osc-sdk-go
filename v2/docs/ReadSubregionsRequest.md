# ReadSubregionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersSubregion**](FiltersSubregion.md) |  | [optional] 

## Methods

### NewReadSubregionsRequest

`func NewReadSubregionsRequest() *ReadSubregionsRequest`

NewReadSubregionsRequest instantiates a new ReadSubregionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSubregionsRequestWithDefaults

`func NewReadSubregionsRequestWithDefaults() *ReadSubregionsRequest`

NewReadSubregionsRequestWithDefaults instantiates a new ReadSubregionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadSubregionsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadSubregionsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadSubregionsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadSubregionsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadSubregionsRequest) GetFilters() FiltersSubregion`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadSubregionsRequest) GetFiltersOk() (*FiltersSubregion, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadSubregionsRequest) SetFilters(v FiltersSubregion)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadSubregionsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


