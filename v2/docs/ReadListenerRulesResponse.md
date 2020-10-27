# ReadListenerRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListenerRules** | Pointer to [**[]ListenerRule**](ListenerRule.md) | The list of the rules to describe. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadListenerRulesResponse

`func NewReadListenerRulesResponse() *ReadListenerRulesResponse`

NewReadListenerRulesResponse instantiates a new ReadListenerRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadListenerRulesResponseWithDefaults

`func NewReadListenerRulesResponseWithDefaults() *ReadListenerRulesResponse`

NewReadListenerRulesResponseWithDefaults instantiates a new ReadListenerRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListenerRules

`func (o *ReadListenerRulesResponse) GetListenerRules() []ListenerRule`

GetListenerRules returns the ListenerRules field if non-nil, zero value otherwise.

### GetListenerRulesOk

`func (o *ReadListenerRulesResponse) GetListenerRulesOk() (*[]ListenerRule, bool)`

GetListenerRulesOk returns a tuple with the ListenerRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerRules

`func (o *ReadListenerRulesResponse) SetListenerRules(v []ListenerRule)`

SetListenerRules sets ListenerRules field to given value.

### HasListenerRules

`func (o *ReadListenerRulesResponse) HasListenerRules() bool`

HasListenerRules returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadListenerRulesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadListenerRulesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadListenerRulesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadListenerRulesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


