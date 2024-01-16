# ReadRouteTablesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersRouteTable**](FiltersRouteTable.md) |  | [optional] 
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of logs returned in a single response (between &#x60;1&#x60;and &#x60;1000&#x60;, both included). By default, &#x60;100&#x60;. | [optional] 

## Methods

### NewReadRouteTablesRequest

`func NewReadRouteTablesRequest() *ReadRouteTablesRequest`

NewReadRouteTablesRequest instantiates a new ReadRouteTablesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadRouteTablesRequestWithDefaults

`func NewReadRouteTablesRequestWithDefaults() *ReadRouteTablesRequest`

NewReadRouteTablesRequestWithDefaults instantiates a new ReadRouteTablesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadRouteTablesRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadRouteTablesRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadRouteTablesRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadRouteTablesRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadRouteTablesRequest) GetFilters() FiltersRouteTable`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadRouteTablesRequest) GetFiltersOk() (*FiltersRouteTable, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadRouteTablesRequest) SetFilters(v FiltersRouteTable)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadRouteTablesRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ReadRouteTablesRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadRouteTablesRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadRouteTablesRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadRouteTablesRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadRouteTablesRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadRouteTablesRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadRouteTablesRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadRouteTablesRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


