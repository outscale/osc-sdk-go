# ReadUserPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyDocument** | Pointer to **string** | The policy document, providing a description of the policy. | [optional] 
**PolicyName** | Pointer to **string** | The name of the inline policy. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**UserName** | Pointer to **string** | The name of the user in which the inline policy is included. | [optional] 

## Methods

### NewReadUserPolicyResponse

`func NewReadUserPolicyResponse() *ReadUserPolicyResponse`

NewReadUserPolicyResponse instantiates a new ReadUserPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserPolicyResponseWithDefaults

`func NewReadUserPolicyResponseWithDefaults() *ReadUserPolicyResponse`

NewReadUserPolicyResponseWithDefaults instantiates a new ReadUserPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyDocument

`func (o *ReadUserPolicyResponse) GetPolicyDocument() string`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *ReadUserPolicyResponse) GetPolicyDocumentOk() (*string, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *ReadUserPolicyResponse) SetPolicyDocument(v string)`

SetPolicyDocument sets PolicyDocument field to given value.

### HasPolicyDocument

`func (o *ReadUserPolicyResponse) HasPolicyDocument() bool`

HasPolicyDocument returns a boolean if a field has been set.

### GetPolicyName

`func (o *ReadUserPolicyResponse) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ReadUserPolicyResponse) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ReadUserPolicyResponse) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ReadUserPolicyResponse) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadUserPolicyResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadUserPolicyResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadUserPolicyResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadUserPolicyResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetUserName

`func (o *ReadUserPolicyResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ReadUserPolicyResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ReadUserPolicyResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ReadUserPolicyResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


