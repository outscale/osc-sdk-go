# CreatePolicyVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyVersion** | Pointer to [**PolicyVersion**](PolicyVersion.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreatePolicyVersionResponse

`func NewCreatePolicyVersionResponse() *CreatePolicyVersionResponse`

NewCreatePolicyVersionResponse instantiates a new CreatePolicyVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyVersionResponseWithDefaults

`func NewCreatePolicyVersionResponseWithDefaults() *CreatePolicyVersionResponse`

NewCreatePolicyVersionResponseWithDefaults instantiates a new CreatePolicyVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyVersion

`func (o *CreatePolicyVersionResponse) GetPolicyVersion() PolicyVersion`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *CreatePolicyVersionResponse) GetPolicyVersionOk() (*PolicyVersion, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *CreatePolicyVersionResponse) SetPolicyVersion(v PolicyVersion)`

SetPolicyVersion sets PolicyVersion field to given value.

### HasPolicyVersion

`func (o *CreatePolicyVersionResponse) HasPolicyVersion() bool`

HasPolicyVersion returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreatePolicyVersionResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreatePolicyVersionResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreatePolicyVersionResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreatePolicyVersionResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


