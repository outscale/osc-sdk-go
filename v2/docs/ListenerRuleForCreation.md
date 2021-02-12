# ListenerRuleForCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The type of action for the rule (always &#x60;forward&#x60;). | [optional] 
**HostNamePattern** | Pointer to **string** | A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].  | [optional] 
**ListenerRuleName** | **string** | A human-readable name for the listener rule. | 
**PathPattern** | Pointer to **string** | A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~\&quot;&#39;@:+?]. | [optional] 
**Priority** | **int32** | The priority level of the listener rule, between &#x60;1&#x60; and &#x60;19999&#x60; both included. Each rule must have a unique priority level. Otherwise, an error is returned. | 

## Methods

### NewListenerRuleForCreation

`func NewListenerRuleForCreation(listenerRuleName string, priority int32, ) *ListenerRuleForCreation`

NewListenerRuleForCreation instantiates a new ListenerRuleForCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerRuleForCreationWithDefaults

`func NewListenerRuleForCreationWithDefaults() *ListenerRuleForCreation`

NewListenerRuleForCreationWithDefaults instantiates a new ListenerRuleForCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ListenerRuleForCreation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListenerRuleForCreation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ListenerRuleForCreation) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ListenerRuleForCreation) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetHostNamePattern

`func (o *ListenerRuleForCreation) GetHostNamePattern() string`

GetHostNamePattern returns the HostNamePattern field if non-nil, zero value otherwise.

### GetHostNamePatternOk

`func (o *ListenerRuleForCreation) GetHostNamePatternOk() (*string, bool)`

GetHostNamePatternOk returns a tuple with the HostNamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamePattern

`func (o *ListenerRuleForCreation) SetHostNamePattern(v string)`

SetHostNamePattern sets HostNamePattern field to given value.

### HasHostNamePattern

`func (o *ListenerRuleForCreation) HasHostNamePattern() bool`

HasHostNamePattern returns a boolean if a field has been set.

### GetListenerRuleName

`func (o *ListenerRuleForCreation) GetListenerRuleName() string`

GetListenerRuleName returns the ListenerRuleName field if non-nil, zero value otherwise.

### GetListenerRuleNameOk

`func (o *ListenerRuleForCreation) GetListenerRuleNameOk() (*string, bool)`

GetListenerRuleNameOk returns a tuple with the ListenerRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerRuleName

`func (o *ListenerRuleForCreation) SetListenerRuleName(v string)`

SetListenerRuleName sets ListenerRuleName field to given value.


### GetPathPattern

`func (o *ListenerRuleForCreation) GetPathPattern() string`

GetPathPattern returns the PathPattern field if non-nil, zero value otherwise.

### GetPathPatternOk

`func (o *ListenerRuleForCreation) GetPathPatternOk() (*string, bool)`

GetPathPatternOk returns a tuple with the PathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPattern

`func (o *ListenerRuleForCreation) SetPathPattern(v string)`

SetPathPattern sets PathPattern field to given value.

### HasPathPattern

`func (o *ListenerRuleForCreation) HasPathPattern() bool`

HasPathPattern returns a boolean if a field has been set.

### GetPriority

`func (o *ListenerRuleForCreation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListenerRuleForCreation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ListenerRuleForCreation) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


