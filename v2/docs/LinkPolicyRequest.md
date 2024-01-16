# LinkPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PolicyOrn** | **string** | The OUTSCALE Resource Name (ORN) of the policy. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | 
**UserName** | **string** | The name of the user you want to link the policy to (between 1 and 64 characters). | 

## Methods

### NewLinkPolicyRequest

`func NewLinkPolicyRequest(policyOrn string, userName string, ) *LinkPolicyRequest`

NewLinkPolicyRequest instantiates a new LinkPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPolicyRequestWithDefaults

`func NewLinkPolicyRequestWithDefaults() *LinkPolicyRequest`

NewLinkPolicyRequestWithDefaults instantiates a new LinkPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *LinkPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicyOrn

`func (o *LinkPolicyRequest) GetPolicyOrn() string`

GetPolicyOrn returns the PolicyOrn field if non-nil, zero value otherwise.

### GetPolicyOrnOk

`func (o *LinkPolicyRequest) GetPolicyOrnOk() (*string, bool)`

GetPolicyOrnOk returns a tuple with the PolicyOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOrn

`func (o *LinkPolicyRequest) SetPolicyOrn(v string)`

SetPolicyOrn sets PolicyOrn field to given value.


### GetUserName

`func (o *LinkPolicyRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *LinkPolicyRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *LinkPolicyRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


