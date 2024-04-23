# ReadUserGroupPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**FirstItem** | Pointer to **int32** | The item starting the list of policies requested. | [optional] 
**ResultsPerPage** | Pointer to **int32** | The maximum number of items that can be returned in a single response (by default, &#x60;100&#x60;). | [optional] 
**UserGroupName** | **string** | The name of the group. | 
**UserGroupPath** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 

## Methods

### NewReadUserGroupPoliciesRequest

`func NewReadUserGroupPoliciesRequest(userGroupName string, ) *ReadUserGroupPoliciesRequest`

NewReadUserGroupPoliciesRequest instantiates a new ReadUserGroupPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserGroupPoliciesRequestWithDefaults

`func NewReadUserGroupPoliciesRequestWithDefaults() *ReadUserGroupPoliciesRequest`

NewReadUserGroupPoliciesRequestWithDefaults instantiates a new ReadUserGroupPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadUserGroupPoliciesRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadUserGroupPoliciesRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadUserGroupPoliciesRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadUserGroupPoliciesRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFirstItem

`func (o *ReadUserGroupPoliciesRequest) GetFirstItem() int32`

GetFirstItem returns the FirstItem field if non-nil, zero value otherwise.

### GetFirstItemOk

`func (o *ReadUserGroupPoliciesRequest) GetFirstItemOk() (*int32, bool)`

GetFirstItemOk returns a tuple with the FirstItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItem

`func (o *ReadUserGroupPoliciesRequest) SetFirstItem(v int32)`

SetFirstItem sets FirstItem field to given value.

### HasFirstItem

`func (o *ReadUserGroupPoliciesRequest) HasFirstItem() bool`

HasFirstItem returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *ReadUserGroupPoliciesRequest) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadUserGroupPoliciesRequest) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *ReadUserGroupPoliciesRequest) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.

### HasResultsPerPage

`func (o *ReadUserGroupPoliciesRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.

### GetUserGroupName

`func (o *ReadUserGroupPoliciesRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *ReadUserGroupPoliciesRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *ReadUserGroupPoliciesRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.


### GetUserGroupPath

`func (o *ReadUserGroupPoliciesRequest) GetUserGroupPath() string`

GetUserGroupPath returns the UserGroupPath field if non-nil, zero value otherwise.

### GetUserGroupPathOk

`func (o *ReadUserGroupPoliciesRequest) GetUserGroupPathOk() (*string, bool)`

GetUserGroupPathOk returns a tuple with the UserGroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupPath

`func (o *ReadUserGroupPoliciesRequest) SetUserGroupPath(v string)`

SetUserGroupPath sets UserGroupPath field to given value.

### HasUserGroupPath

`func (o *ReadUserGroupPoliciesRequest) HasUserGroupPath() bool`

HasUserGroupPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


