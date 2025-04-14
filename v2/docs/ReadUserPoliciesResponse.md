# ReadUserPoliciesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyNames** | Pointer to **[]string** | A list of policy names. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadUserPoliciesResponse

`func NewReadUserPoliciesResponse() *ReadUserPoliciesResponse`

NewReadUserPoliciesResponse instantiates a new ReadUserPoliciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserPoliciesResponseWithDefaults

`func NewReadUserPoliciesResponseWithDefaults() *ReadUserPoliciesResponse`

NewReadUserPoliciesResponseWithDefaults instantiates a new ReadUserPoliciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyNames

`func (o *ReadUserPoliciesResponse) GetPolicyNames() []string`

GetPolicyNames returns the PolicyNames field if non-nil, zero value otherwise.

### GetPolicyNamesOk

`func (o *ReadUserPoliciesResponse) GetPolicyNamesOk() (*[]string, bool)`

GetPolicyNamesOk returns a tuple with the PolicyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNames

`func (o *ReadUserPoliciesResponse) SetPolicyNames(v []string)`

SetPolicyNames sets PolicyNames field to given value.

### HasPolicyNames

`func (o *ReadUserPoliciesResponse) HasPolicyNames() bool`

HasPolicyNames returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadUserPoliciesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadUserPoliciesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadUserPoliciesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadUserPoliciesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


