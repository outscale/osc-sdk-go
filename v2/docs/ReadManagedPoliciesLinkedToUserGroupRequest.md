# ReadManagedPoliciesLinkedToUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersUserGroup**](FiltersUserGroup.md) |  | [optional] 
**FirstItem** | Pointer to **int32** | The item starting the list of policies requested. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of items that can be returned in a single response (by default, &#x60;100&#x60;). | [optional] 
**UserGroupName** | **string** | The name of the group. | 

## Methods

### NewReadManagedPoliciesLinkedToUserGroupRequest

`func NewReadManagedPoliciesLinkedToUserGroupRequest(userGroupName string, ) *ReadManagedPoliciesLinkedToUserGroupRequest`

NewReadManagedPoliciesLinkedToUserGroupRequest instantiates a new ReadManagedPoliciesLinkedToUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadManagedPoliciesLinkedToUserGroupRequestWithDefaults

`func NewReadManagedPoliciesLinkedToUserGroupRequestWithDefaults() *ReadManagedPoliciesLinkedToUserGroupRequest`

NewReadManagedPoliciesLinkedToUserGroupRequestWithDefaults instantiates a new ReadManagedPoliciesLinkedToUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetFilters() FiltersUserGroup`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetFiltersOk() (*FiltersUserGroup, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) SetFilters(v FiltersUserGroup)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFirstItem

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetFirstItem() int32`

GetFirstItem returns the FirstItem field if non-nil, zero value otherwise.

### GetFirstItemOk

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetFirstItemOk() (*int32, bool)`

GetFirstItemOk returns a tuple with the FirstItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItem

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) SetFirstItem(v int32)`

SetFirstItem sets FirstItem field to given value.

### HasFirstItem

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) HasFirstItem() bool`

HasFirstItem returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.

### GetUserGroupName

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *ReadManagedPoliciesLinkedToUserGroupRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


