# PutUserGroupPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PolicyDocument** | **string** | The policy document, corresponding to a JSON string that contains the policy. For more information, see [EIM Reference Information](https://docs.outscale.com/en/userguide/EIM-Reference-Information.html). | 
**PolicyName** | **string** | The name of the policy. | 
**UserGroupName** | **string** | The name of the group. | 
**UserGroupPath** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 

## Methods

### NewPutUserGroupPolicyRequest

`func NewPutUserGroupPolicyRequest(policyDocument string, policyName string, userGroupName string, ) *PutUserGroupPolicyRequest`

NewPutUserGroupPolicyRequest instantiates a new PutUserGroupPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutUserGroupPolicyRequestWithDefaults

`func NewPutUserGroupPolicyRequestWithDefaults() *PutUserGroupPolicyRequest`

NewPutUserGroupPolicyRequestWithDefaults instantiates a new PutUserGroupPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *PutUserGroupPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PutUserGroupPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PutUserGroupPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PutUserGroupPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicyDocument

`func (o *PutUserGroupPolicyRequest) GetPolicyDocument() string`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *PutUserGroupPolicyRequest) GetPolicyDocumentOk() (*string, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *PutUserGroupPolicyRequest) SetPolicyDocument(v string)`

SetPolicyDocument sets PolicyDocument field to given value.


### GetPolicyName

`func (o *PutUserGroupPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PutUserGroupPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PutUserGroupPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetUserGroupName

`func (o *PutUserGroupPolicyRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *PutUserGroupPolicyRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *PutUserGroupPolicyRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.


### GetUserGroupPath

`func (o *PutUserGroupPolicyRequest) GetUserGroupPath() string`

GetUserGroupPath returns the UserGroupPath field if non-nil, zero value otherwise.

### GetUserGroupPathOk

`func (o *PutUserGroupPolicyRequest) GetUserGroupPathOk() (*string, bool)`

GetUserGroupPathOk returns a tuple with the UserGroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupPath

`func (o *PutUserGroupPolicyRequest) SetUserGroupPath(v string)`

SetUserGroupPath sets UserGroupPath field to given value.

### HasUserGroupPath

`func (o *PutUserGroupPolicyRequest) HasUserGroupPath() bool`

HasUserGroupPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


