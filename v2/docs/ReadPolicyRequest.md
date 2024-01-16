# ReadPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyOrn** | **string** | The OUTSCALE Resource Name (ORN) of the policy. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | 

## Methods

### NewReadPolicyRequest

`func NewReadPolicyRequest(policyOrn string, ) *ReadPolicyRequest`

NewReadPolicyRequest instantiates a new ReadPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPolicyRequestWithDefaults

`func NewReadPolicyRequestWithDefaults() *ReadPolicyRequest`

NewReadPolicyRequestWithDefaults instantiates a new ReadPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyOrn

`func (o *ReadPolicyRequest) GetPolicyOrn() string`

GetPolicyOrn returns the PolicyOrn field if non-nil, zero value otherwise.

### GetPolicyOrnOk

`func (o *ReadPolicyRequest) GetPolicyOrnOk() (*string, bool)`

GetPolicyOrnOk returns a tuple with the PolicyOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOrn

`func (o *ReadPolicyRequest) SetPolicyOrn(v string)`

SetPolicyOrn sets PolicyOrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


