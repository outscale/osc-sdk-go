# ReadPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**Policy**](Policy.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadPolicyResponse

`func NewReadPolicyResponse() *ReadPolicyResponse`

NewReadPolicyResponse instantiates a new ReadPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPolicyResponseWithDefaults

`func NewReadPolicyResponseWithDefaults() *ReadPolicyResponse`

NewReadPolicyResponseWithDefaults instantiates a new ReadPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *ReadPolicyResponse) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ReadPolicyResponse) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ReadPolicyResponse) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ReadPolicyResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadPolicyResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadPolicyResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadPolicyResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadPolicyResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


