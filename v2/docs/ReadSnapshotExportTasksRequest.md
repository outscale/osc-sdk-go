# ReadSnapshotExportTasksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersExportTask**](FiltersExportTask.md) |  | [optional] 
**NextPageToken** | Pointer to **string** | The token to request the next page of results. Each token refers to a specific page. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of logs returned in a single response (between &#x60;1&#x60; and &#x60;1000&#x60;, both included). | [optional] 

## Methods

### NewReadSnapshotExportTasksRequest

`func NewReadSnapshotExportTasksRequest() *ReadSnapshotExportTasksRequest`

NewReadSnapshotExportTasksRequest instantiates a new ReadSnapshotExportTasksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSnapshotExportTasksRequestWithDefaults

`func NewReadSnapshotExportTasksRequestWithDefaults() *ReadSnapshotExportTasksRequest`

NewReadSnapshotExportTasksRequestWithDefaults instantiates a new ReadSnapshotExportTasksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadSnapshotExportTasksRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadSnapshotExportTasksRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadSnapshotExportTasksRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadSnapshotExportTasksRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadSnapshotExportTasksRequest) GetFilters() FiltersExportTask`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadSnapshotExportTasksRequest) GetFiltersOk() (*FiltersExportTask, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadSnapshotExportTasksRequest) SetFilters(v FiltersExportTask)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadSnapshotExportTasksRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ReadSnapshotExportTasksRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadSnapshotExportTasksRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReadSnapshotExportTasksRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReadSnapshotExportTasksRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadSnapshotExportTasksRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadSnapshotExportTasksRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadSnapshotExportTasksRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadSnapshotExportTasksRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


