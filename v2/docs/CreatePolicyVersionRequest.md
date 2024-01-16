# CreatePolicyVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | **string** | The policy document, corresponding to a JSON string that contains the policy. For more information, see [EIM Reference Information](https://docs.outscale.com/en/userguide/EIM-Reference-Information.html). | 
**PolicyOrn** | **string** | The OUTSCALE Resource Name (ORN) of the policy. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | 
**SetAsDefault** | Pointer to **bool** | If set to true, the new policy version is set as the default version and becomes the operative one. | [optional] 

## Methods

### NewCreatePolicyVersionRequest

`func NewCreatePolicyVersionRequest(document string, policyOrn string, ) *CreatePolicyVersionRequest`

NewCreatePolicyVersionRequest instantiates a new CreatePolicyVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyVersionRequestWithDefaults

`func NewCreatePolicyVersionRequestWithDefaults() *CreatePolicyVersionRequest`

NewCreatePolicyVersionRequestWithDefaults instantiates a new CreatePolicyVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *CreatePolicyVersionRequest) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *CreatePolicyVersionRequest) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *CreatePolicyVersionRequest) SetDocument(v string)`

SetDocument sets Document field to given value.


### GetPolicyOrn

`func (o *CreatePolicyVersionRequest) GetPolicyOrn() string`

GetPolicyOrn returns the PolicyOrn field if non-nil, zero value otherwise.

### GetPolicyOrnOk

`func (o *CreatePolicyVersionRequest) GetPolicyOrnOk() (*string, bool)`

GetPolicyOrnOk returns a tuple with the PolicyOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOrn

`func (o *CreatePolicyVersionRequest) SetPolicyOrn(v string)`

SetPolicyOrn sets PolicyOrn field to given value.


### GetSetAsDefault

`func (o *CreatePolicyVersionRequest) GetSetAsDefault() bool`

GetSetAsDefault returns the SetAsDefault field if non-nil, zero value otherwise.

### GetSetAsDefaultOk

`func (o *CreatePolicyVersionRequest) GetSetAsDefaultOk() (*bool, bool)`

GetSetAsDefaultOk returns a tuple with the SetAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAsDefault

`func (o *CreatePolicyVersionRequest) SetSetAsDefault(v bool)`

SetSetAsDefault sets SetAsDefault field to given value.

### HasSetAsDefault

`func (o *CreatePolicyVersionRequest) HasSetAsDefault() bool`

HasSetAsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


