# PutUserPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PolicyDocument** | **string** | The policy document, corresponding to a JSON string that contains the policy. For more information, see [EIM Reference Information](https://docs.outscale.com/en/userguide/EIM-Reference-Information.html) and [EIM Policy Generator](https://docs.outscale.com/en/userguide/EIM-Policy-Generator.html). | 
**PolicyName** | **string** | The name of the policy (between 1 and 128 characters). | 
**UserName** | **string** | The name of the user. | 

## Methods

### NewPutUserPolicyRequest

`func NewPutUserPolicyRequest(policyDocument string, policyName string, userName string, ) *PutUserPolicyRequest`

NewPutUserPolicyRequest instantiates a new PutUserPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutUserPolicyRequestWithDefaults

`func NewPutUserPolicyRequestWithDefaults() *PutUserPolicyRequest`

NewPutUserPolicyRequestWithDefaults instantiates a new PutUserPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *PutUserPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PutUserPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PutUserPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PutUserPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicyDocument

`func (o *PutUserPolicyRequest) GetPolicyDocument() string`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *PutUserPolicyRequest) GetPolicyDocumentOk() (*string, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *PutUserPolicyRequest) SetPolicyDocument(v string)`

SetPolicyDocument sets PolicyDocument field to given value.


### GetPolicyName

`func (o *PutUserPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PutUserPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PutUserPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetUserName

`func (o *PutUserPolicyRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PutUserPolicyRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PutUserPolicyRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


