# DeleteUserPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PolicyName** | **string** | The name of the policy document you want to delete (between 1 and 128 characters). | 
**UserName** | **string** | The name of the user you want to delete the policy from. | 

## Methods

### NewDeleteUserPolicyRequest

`func NewDeleteUserPolicyRequest(policyName string, userName string, ) *DeleteUserPolicyRequest`

NewDeleteUserPolicyRequest instantiates a new DeleteUserPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserPolicyRequestWithDefaults

`func NewDeleteUserPolicyRequestWithDefaults() *DeleteUserPolicyRequest`

NewDeleteUserPolicyRequestWithDefaults instantiates a new DeleteUserPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteUserPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteUserPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteUserPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteUserPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicyName

`func (o *DeleteUserPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *DeleteUserPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *DeleteUserPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetUserName

`func (o *DeleteUserPolicyRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DeleteUserPolicyRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DeleteUserPolicyRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


