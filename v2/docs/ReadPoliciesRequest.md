# ReadPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**ReadPoliciesFilters**](ReadPoliciesFilters.md) |  | [optional] 
**FirstItem** | Pointer to **int32** | The item starting the list of policies requested. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of items that can be returned in a single response (by default, 100). | [optional] 

## Methods

### NewReadPoliciesRequest

`func NewReadPoliciesRequest() *ReadPoliciesRequest`

NewReadPoliciesRequest instantiates a new ReadPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPoliciesRequestWithDefaults

`func NewReadPoliciesRequestWithDefaults() *ReadPoliciesRequest`

NewReadPoliciesRequestWithDefaults instantiates a new ReadPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadPoliciesRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadPoliciesRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadPoliciesRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadPoliciesRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadPoliciesRequest) GetFilters() ReadPoliciesFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadPoliciesRequest) GetFiltersOk() (*ReadPoliciesFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadPoliciesRequest) SetFilters(v ReadPoliciesFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadPoliciesRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFirstItem

`func (o *ReadPoliciesRequest) GetFirstItem() int32`

GetFirstItem returns the FirstItem field if non-nil, zero value otherwise.

### GetFirstItemOk

`func (o *ReadPoliciesRequest) GetFirstItemOk() (*int32, bool)`

GetFirstItemOk returns a tuple with the FirstItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItem

`func (o *ReadPoliciesRequest) SetFirstItem(v int32)`

SetFirstItem sets FirstItem field to given value.

### HasFirstItem

`func (o *ReadPoliciesRequest) HasFirstItem() bool`

HasFirstItem returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadPoliciesRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadPoliciesRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadPoliciesRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadPoliciesRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


