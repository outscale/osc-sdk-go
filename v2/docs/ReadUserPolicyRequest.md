# ReadUserPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PolicyName** | **string** | The name of the policy. | 
**UserName** | **string** | The name of the user. | 

## Methods

### NewReadUserPolicyRequest

`func NewReadUserPolicyRequest(policyName string, userName string, ) *ReadUserPolicyRequest`

NewReadUserPolicyRequest instantiates a new ReadUserPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserPolicyRequestWithDefaults

`func NewReadUserPolicyRequestWithDefaults() *ReadUserPolicyRequest`

NewReadUserPolicyRequestWithDefaults instantiates a new ReadUserPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadUserPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadUserPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadUserPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadUserPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicyName

`func (o *ReadUserPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ReadUserPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ReadUserPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetUserName

`func (o *ReadUserPolicyRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ReadUserPolicyRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ReadUserPolicyRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


