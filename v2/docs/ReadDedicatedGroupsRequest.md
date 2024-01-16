# ReadDedicatedGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersDedicatedGroup**](FiltersDedicatedGroup.md) |  | [optional] 

## Methods

### NewReadDedicatedGroupsRequest

`func NewReadDedicatedGroupsRequest() *ReadDedicatedGroupsRequest`

NewReadDedicatedGroupsRequest instantiates a new ReadDedicatedGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadDedicatedGroupsRequestWithDefaults

`func NewReadDedicatedGroupsRequestWithDefaults() *ReadDedicatedGroupsRequest`

NewReadDedicatedGroupsRequestWithDefaults instantiates a new ReadDedicatedGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadDedicatedGroupsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadDedicatedGroupsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadDedicatedGroupsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadDedicatedGroupsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadDedicatedGroupsRequest) GetFilters() FiltersDedicatedGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadDedicatedGroupsRequest) GetFiltersOk() (*FiltersDedicatedGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadDedicatedGroupsRequest) SetFilters(v FiltersDedicatedGroup)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadDedicatedGroupsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


