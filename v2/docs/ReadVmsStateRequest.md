# ReadVmsStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllVms** | Pointer to **bool** | If true, includes the status of all VMs. If false, only includes the status of running VMs. | [optional] [default to false]
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersVmsState**](FiltersVmsState.md) |  | [optional] 
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of logs returned in a single response (between &#x60;1&#x60; and &#x60;1000&#x60;, both included). | [optional] 

## Methods

### NewReadVmsStateRequest

`func NewReadVmsStateRequest() *ReadVmsStateRequest`

NewReadVmsStateRequest instantiates a new ReadVmsStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadVmsStateRequestWithDefaults

`func NewReadVmsStateRequestWithDefaults() *ReadVmsStateRequest`

NewReadVmsStateRequestWithDefaults instantiates a new ReadVmsStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllVms

`func (o *ReadVmsStateRequest) GetAllVms() bool`

GetAllVms returns the AllVms field if non-nil, zero value otherwise.

### GetAllVmsOk

`func (o *ReadVmsStateRequest) GetAllVmsOk() (*bool, bool)`

GetAllVmsOk returns a tuple with the AllVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllVms

`func (o *ReadVmsStateRequest) SetAllVms(v bool)`

SetAllVms sets AllVms field to given value.

### HasAllVms

`func (o *ReadVmsStateRequest) HasAllVms() bool`

HasAllVms returns a boolean if a field has been set.

### GetDryRun

`func (o *ReadVmsStateRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadVmsStateRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadVmsStateRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadVmsStateRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadVmsStateRequest) GetFilters() FiltersVmsState`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadVmsStateRequest) GetFiltersOk() (*FiltersVmsState, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadVmsStateRequest) SetFilters(v FiltersVmsState)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadVmsStateRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ReadVmsStateRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadVmsStateRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadVmsStateRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadVmsStateRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadVmsStateRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadVmsStateRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadVmsStateRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadVmsStateRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


