# ReadApiAccessRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccessRules** | Pointer to [**[]ApiAccessRule**](ApiAccessRule.md) | A list of API access rules. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadApiAccessRulesResponse

`func NewReadApiAccessRulesResponse() *ReadApiAccessRulesResponse`

NewReadApiAccessRulesResponse instantiates a new ReadApiAccessRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadApiAccessRulesResponseWithDefaults

`func NewReadApiAccessRulesResponseWithDefaults() *ReadApiAccessRulesResponse`

NewReadApiAccessRulesResponseWithDefaults instantiates a new ReadApiAccessRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccessRules

`func (o *ReadApiAccessRulesResponse) GetApiAccessRules() []ApiAccessRule`

GetApiAccessRules returns the ApiAccessRules field if non-nil, zero value otherwise.

### GetApiAccessRulesOk

`func (o *ReadApiAccessRulesResponse) GetApiAccessRulesOk() (*[]ApiAccessRule, bool)`

GetApiAccessRulesOk returns a tuple with the ApiAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessRules

`func (o *ReadApiAccessRulesResponse) SetApiAccessRules(v []ApiAccessRule)`

SetApiAccessRules sets ApiAccessRules field to given value.

### HasApiAccessRules

`func (o *ReadApiAccessRulesResponse) HasApiAccessRules() bool`

HasApiAccessRules returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadApiAccessRulesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadApiAccessRulesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadApiAccessRulesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadApiAccessRulesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


