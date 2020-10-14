# ReadApiLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersApiLog**](FiltersApiLog.md) |  | [optional] 
**NextPageToken** | Pointer to **string** | The token to request the next page of results. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of items returned in a single page. By default, 100. | [optional] 
**With** | Pointer to [**With**](With.md) |  | [optional] 

## Methods

### NewReadApiLogsRequest

`func NewReadApiLogsRequest() *ReadApiLogsRequest`

NewReadApiLogsRequest instantiates a new ReadApiLogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadApiLogsRequestWithDefaults

`func NewReadApiLogsRequestWithDefaults() *ReadApiLogsRequest`

NewReadApiLogsRequestWithDefaults instantiates a new ReadApiLogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadApiLogsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadApiLogsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadApiLogsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadApiLogsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadApiLogsRequest) GetFilters() FiltersApiLog`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadApiLogsRequest) GetFiltersOk() (*FiltersApiLog, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadApiLogsRequest) SetFilters(v FiltersApiLog)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadApiLogsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ReadApiLogsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadApiLogsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadApiLogsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadApiLogsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadApiLogsRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadApiLogsRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadApiLogsRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadApiLogsRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.

### GetWith

`func (o *ReadApiLogsRequest) GetWith() With`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *ReadApiLogsRequest) GetWithOk() (*With, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWith

`func (o *ReadApiLogsRequest) SetWith(v With)`

SetWith sets With field to given value.

### HasWith

`func (o *ReadApiLogsRequest) HasWith() bool`

HasWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


