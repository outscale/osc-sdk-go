# CreatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description for the policy. | [optional] 
**Document** | **string** | The policy document, corresponding to a JSON string that contains the policy. For more information, see [EIM Reference Information](https://docs.outscale.com/en/userguide/EIM-Reference-Information.html). | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Path** | Pointer to **string** | The path of the policy. | [optional] 
**PolicyName** | **string** | The name of the policy. | 

## Methods

### NewCreatePolicyRequest

`func NewCreatePolicyRequest(document string, policyName string, ) *CreatePolicyRequest`

NewCreatePolicyRequest instantiates a new CreatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyRequestWithDefaults

`func NewCreatePolicyRequestWithDefaults() *CreatePolicyRequest`

NewCreatePolicyRequestWithDefaults instantiates a new CreatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreatePolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocument

`func (o *CreatePolicyRequest) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *CreatePolicyRequest) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *CreatePolicyRequest) SetDocument(v string)`

SetDocument sets Document field to given value.


### GetDryRun

`func (o *CreatePolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreatePolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreatePolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreatePolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPath

`func (o *CreatePolicyRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreatePolicyRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreatePolicyRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreatePolicyRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPolicyName

`func (o *CreatePolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *CreatePolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *CreatePolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


