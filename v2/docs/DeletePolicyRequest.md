# DeletePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PolicyOrn** | **string** | The OUTSCALE Resource Name (ORN) of the policy you want to delete. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | 

## Methods

### NewDeletePolicyRequest

`func NewDeletePolicyRequest(policyOrn string, ) *DeletePolicyRequest`

NewDeletePolicyRequest instantiates a new DeletePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePolicyRequestWithDefaults

`func NewDeletePolicyRequestWithDefaults() *DeletePolicyRequest`

NewDeletePolicyRequestWithDefaults instantiates a new DeletePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeletePolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeletePolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeletePolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeletePolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPolicyOrn

`func (o *DeletePolicyRequest) GetPolicyOrn() string`

GetPolicyOrn returns the PolicyOrn field if non-nil, zero value otherwise.

### GetPolicyOrnOk

`func (o *DeletePolicyRequest) GetPolicyOrnOk() (*string, bool)`

GetPolicyOrnOk returns a tuple with the PolicyOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOrn

`func (o *DeletePolicyRequest) SetPolicyOrn(v string)`

SetPolicyOrn sets PolicyOrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


