# ReadUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersUsers**](FiltersUsers.md) |  | [optional] 
**FirstItem** | Pointer to **int32** | The item starting the list of users requested. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of items that can be returned in a single response (by default, &#x60;100&#x60;). | [optional] 

## Methods

### NewReadUsersRequest

`func NewReadUsersRequest() *ReadUsersRequest`

NewReadUsersRequest instantiates a new ReadUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUsersRequestWithDefaults

`func NewReadUsersRequestWithDefaults() *ReadUsersRequest`

NewReadUsersRequestWithDefaults instantiates a new ReadUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadUsersRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadUsersRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadUsersRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadUsersRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadUsersRequest) GetFilters() FiltersUsers`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadUsersRequest) GetFiltersOk() (*FiltersUsers, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadUsersRequest) SetFilters(v FiltersUsers)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadUsersRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFirstItem

`func (o *ReadUsersRequest) GetFirstItem() int32`

GetFirstItem returns the FirstItem field if non-nil, zero value otherwise.

### GetFirstItemOk

`func (o *ReadUsersRequest) GetFirstItemOk() (*int32, bool)`

GetFirstItemOk returns a tuple with the FirstItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItem

`func (o *ReadUsersRequest) SetFirstItem(v int32)`

SetFirstItem sets FirstItem field to given value.

### HasFirstItem

`func (o *ReadUsersRequest) HasFirstItem() bool`

HasFirstItem returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadUsersRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadUsersRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadUsersRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadUsersRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


