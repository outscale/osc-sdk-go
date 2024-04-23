# DeleteUserGroupPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PolicyName** | **string** | The name of the policy document you want to delete. | 
**UserGroupName** | **string** | The name of the group. | 
**UserGroupPath** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 

## Methods

### NewDeleteUserGroupPolicyRequest

`func NewDeleteUserGroupPolicyRequest(policyName string, userGroupName string, ) *DeleteUserGroupPolicyRequest`

NewDeleteUserGroupPolicyRequest instantiates a new DeleteUserGroupPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserGroupPolicyRequestWithDefaults

`func NewDeleteUserGroupPolicyRequestWithDefaults() *DeleteUserGroupPolicyRequest`

NewDeleteUserGroupPolicyRequestWithDefaults instantiates a new DeleteUserGroupPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteUserGroupPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteUserGroupPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteUserGroupPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteUserGroupPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicyName

`func (o *DeleteUserGroupPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *DeleteUserGroupPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *DeleteUserGroupPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetUserGroupName

`func (o *DeleteUserGroupPolicyRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *DeleteUserGroupPolicyRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *DeleteUserGroupPolicyRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.


### GetUserGroupPath

`func (o *DeleteUserGroupPolicyRequest) GetUserGroupPath() string`

GetUserGroupPath returns the UserGroupPath field if non-nil, zero value otherwise.

### GetUserGroupPathOk

`func (o *DeleteUserGroupPolicyRequest) GetUserGroupPathOk() (*string, bool)`

GetUserGroupPathOk returns a tuple with the UserGroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupPath

`func (o *DeleteUserGroupPolicyRequest) SetUserGroupPath(v string)`

SetUserGroupPath sets UserGroupPath field to given value.

### HasUserGroupPath

`func (o *DeleteUserGroupPolicyRequest) HasUserGroupPath() bool`

HasUserGroupPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


