# ReadLinkedPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**ReadLinkedPoliciesFilters**](ReadLinkedPoliciesFilters.md) |  | [optional] 
**FirstItem** | Pointer to **int32** | The item starting the list of policies requested. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of items that can be returned in a single response (by default, 100). | [optional] 
**UserName** | Pointer to **string** | The name of the user the policies are linked to. | [optional] 

## Methods

### NewReadLinkedPoliciesRequest

`func NewReadLinkedPoliciesRequest() *ReadLinkedPoliciesRequest`

NewReadLinkedPoliciesRequest instantiates a new ReadLinkedPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadLinkedPoliciesRequestWithDefaults

`func NewReadLinkedPoliciesRequestWithDefaults() *ReadLinkedPoliciesRequest`

NewReadLinkedPoliciesRequestWithDefaults instantiates a new ReadLinkedPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadLinkedPoliciesRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadLinkedPoliciesRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadLinkedPoliciesRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadLinkedPoliciesRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadLinkedPoliciesRequest) GetFilters() ReadLinkedPoliciesFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadLinkedPoliciesRequest) GetFiltersOk() (*ReadLinkedPoliciesFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadLinkedPoliciesRequest) SetFilters(v ReadLinkedPoliciesFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadLinkedPoliciesRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFirstItem

`func (o *ReadLinkedPoliciesRequest) GetFirstItem() int32`

GetFirstItem returns the FirstItem field if non-nil, zero value otherwise.

### GetFirstItemOk

`func (o *ReadLinkedPoliciesRequest) GetFirstItemOk() (*int32, bool)`

GetFirstItemOk returns a tuple with the FirstItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItem

`func (o *ReadLinkedPoliciesRequest) SetFirstItem(v int32)`

SetFirstItem sets FirstItem field to given value.

### HasFirstItem

`func (o *ReadLinkedPoliciesRequest) HasFirstItem() bool`

HasFirstItem returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadLinkedPoliciesRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadLinkedPoliciesRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadLinkedPoliciesRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadLinkedPoliciesRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.

### GetUserName

`func (o *ReadLinkedPoliciesRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ReadLinkedPoliciesRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ReadLinkedPoliciesRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ReadLinkedPoliciesRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


