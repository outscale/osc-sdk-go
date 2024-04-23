# ReadUserGroupPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**InlinePolicy**](InlinePolicy.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadUserGroupPolicyResponse

`func NewReadUserGroupPolicyResponse() *ReadUserGroupPolicyResponse`

NewReadUserGroupPolicyResponse instantiates a new ReadUserGroupPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserGroupPolicyResponseWithDefaults

`func NewReadUserGroupPolicyResponseWithDefaults() *ReadUserGroupPolicyResponse`

NewReadUserGroupPolicyResponseWithDefaults instantiates a new ReadUserGroupPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *ReadUserGroupPolicyResponse) GetPolicy() InlinePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ReadUserGroupPolicyResponse) GetPolicyOk() (*InlinePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ReadUserGroupPolicyResponse) SetPolicy(v InlinePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ReadUserGroupPolicyResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadUserGroupPolicyResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadUserGroupPolicyResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadUserGroupPolicyResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadUserGroupPolicyResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


