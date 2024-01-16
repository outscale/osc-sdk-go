# ReadPolicyVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyOrn** | **string** | The OUTSCALE Resource Name (ORN) of the policy. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | 
**VersionId** | **string** | The ID of the policy version. | 

## Methods

### NewReadPolicyVersionRequest

`func NewReadPolicyVersionRequest(policyOrn string, versionId string, ) *ReadPolicyVersionRequest`

NewReadPolicyVersionRequest instantiates a new ReadPolicyVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPolicyVersionRequestWithDefaults

`func NewReadPolicyVersionRequestWithDefaults() *ReadPolicyVersionRequest`

NewReadPolicyVersionRequestWithDefaults instantiates a new ReadPolicyVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyOrn

`func (o *ReadPolicyVersionRequest) GetPolicyOrn() string`

GetPolicyOrn returns the PolicyOrn field if non-nil, zero value otherwise.

### GetPolicyOrnOk

`func (o *ReadPolicyVersionRequest) GetPolicyOrnOk() (*string, bool)`

GetPolicyOrnOk returns a tuple with the PolicyOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOrn

`func (o *ReadPolicyVersionRequest) SetPolicyOrn(v string)`

SetPolicyOrn sets PolicyOrn field to given value.


### GetVersionId

`func (o *ReadPolicyVersionRequest) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ReadPolicyVersionRequest) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ReadPolicyVersionRequest) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


