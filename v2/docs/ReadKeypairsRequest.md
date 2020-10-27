# ReadKeypairsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersKeypair**](FiltersKeypair.md) |  | [optional] 

## Methods

### NewReadKeypairsRequest

`func NewReadKeypairsRequest() *ReadKeypairsRequest`

NewReadKeypairsRequest instantiates a new ReadKeypairsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadKeypairsRequestWithDefaults

`func NewReadKeypairsRequestWithDefaults() *ReadKeypairsRequest`

NewReadKeypairsRequestWithDefaults instantiates a new ReadKeypairsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadKeypairsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadKeypairsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadKeypairsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadKeypairsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadKeypairsRequest) GetFilters() FiltersKeypair`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadKeypairsRequest) GetFiltersOk() (*FiltersKeypair, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadKeypairsRequest) SetFilters(v FiltersKeypair)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadKeypairsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


