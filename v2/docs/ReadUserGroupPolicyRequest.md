# ReadUserGroupPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PolicyName** | **string** | The name of the policy. | 
**UserGroupName** | **string** | The name of the group. | 
**UserGroupPath** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 

## Methods

### NewReadUserGroupPolicyRequest

`func NewReadUserGroupPolicyRequest(policyName string, userGroupName string, ) *ReadUserGroupPolicyRequest`

NewReadUserGroupPolicyRequest instantiates a new ReadUserGroupPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserGroupPolicyRequestWithDefaults

`func NewReadUserGroupPolicyRequestWithDefaults() *ReadUserGroupPolicyRequest`

NewReadUserGroupPolicyRequestWithDefaults instantiates a new ReadUserGroupPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadUserGroupPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadUserGroupPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadUserGroupPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadUserGroupPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicyName

`func (o *ReadUserGroupPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ReadUserGroupPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ReadUserGroupPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetUserGroupName

`func (o *ReadUserGroupPolicyRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *ReadUserGroupPolicyRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *ReadUserGroupPolicyRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.


### GetUserGroupPath

`func (o *ReadUserGroupPolicyRequest) GetUserGroupPath() string`

GetUserGroupPath returns the UserGroupPath field if non-nil, zero value otherwise.

### GetUserGroupPathOk

`func (o *ReadUserGroupPolicyRequest) GetUserGroupPathOk() (*string, bool)`

GetUserGroupPathOk returns a tuple with the UserGroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupPath

`func (o *ReadUserGroupPolicyRequest) SetUserGroupPath(v string)`

SetUserGroupPath sets UserGroupPath field to given value.

### HasUserGroupPath

`func (o *ReadUserGroupPolicyRequest) HasUserGroupPath() bool`

HasUserGroupPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


