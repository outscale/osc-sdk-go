# ListenerRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The type of action for the rule (always &#x60;forward&#x60;). | [optional] 
**HostNamePattern** | Pointer to **string** | A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?]. | [optional] 
**ListenerId** | Pointer to **int32** | The ID of the listener. | [optional] 
**ListenerRuleId** | Pointer to **int32** | The ID of the listener rule. | [optional] 
**ListenerRuleName** | Pointer to **string** | A human-readable name for the listener rule. | [optional] 
**PathPattern** | Pointer to **string** | A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~\&quot;&#39;@:+?]. | [optional] 
**Priority** | Pointer to **int32** | The priority level of the listener rule, between &#x60;1&#x60; and &#x60;19999&#x60; both included. Each rule must have a unique priority level. Otherwise, an error is returned. | [optional] 
**VmIds** | Pointer to **[]string** | The IDs of the backend VMs. | [optional] 

## Methods

### NewListenerRule

`func NewListenerRule() *ListenerRule`

NewListenerRule instantiates a new ListenerRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerRuleWithDefaults

`func NewListenerRuleWithDefaults() *ListenerRule`

NewListenerRuleWithDefaults instantiates a new ListenerRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ListenerRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListenerRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ListenerRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ListenerRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetHostNamePattern

`func (o *ListenerRule) GetHostNamePattern() string`

GetHostNamePattern returns the HostNamePattern field if non-nil, zero value otherwise.

### GetHostNamePatternOk

`func (o *ListenerRule) GetHostNamePatternOk() (*string, bool)`

GetHostNamePatternOk returns a tuple with the HostNamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamePattern

`func (o *ListenerRule) SetHostNamePattern(v string)`

SetHostNamePattern sets HostNamePattern field to given value.

### HasHostNamePattern

`func (o *ListenerRule) HasHostNamePattern() bool`

HasHostNamePattern returns a boolean if a field has been set.

### GetListenerId

`func (o *ListenerRule) GetListenerId() int32`

GetListenerId returns the ListenerId field if non-nil, zero value otherwise.

### GetListenerIdOk

`func (o *ListenerRule) GetListenerIdOk() (*int32, bool)`

GetListenerIdOk returns a tuple with the ListenerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerId

`func (o *ListenerRule) SetListenerId(v int32)`

SetListenerId sets ListenerId field to given value.

### HasListenerId

`func (o *ListenerRule) HasListenerId() bool`

HasListenerId returns a boolean if a field has been set.

### GetListenerRuleId

`func (o *ListenerRule) GetListenerRuleId() int32`

GetListenerRuleId returns the ListenerRuleId field if non-nil, zero value otherwise.

### GetListenerRuleIdOk

`func (o *ListenerRule) GetListenerRuleIdOk() (*int32, bool)`

GetListenerRuleIdOk returns a tuple with the ListenerRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerRuleId

`func (o *ListenerRule) SetListenerRuleId(v int32)`

SetListenerRuleId sets ListenerRuleId field to given value.

### HasListenerRuleId

`func (o *ListenerRule) HasListenerRuleId() bool`

HasListenerRuleId returns a boolean if a field has been set.

### GetListenerRuleName

`func (o *ListenerRule) GetListenerRuleName() string`

GetListenerRuleName returns the ListenerRuleName field if non-nil, zero value otherwise.

### GetListenerRuleNameOk

`func (o *ListenerRule) GetListenerRuleNameOk() (*string, bool)`

GetListenerRuleNameOk returns a tuple with the ListenerRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerRuleName

`func (o *ListenerRule) SetListenerRuleName(v string)`

SetListenerRuleName sets ListenerRuleName field to given value.

### HasListenerRuleName

`func (o *ListenerRule) HasListenerRuleName() bool`

HasListenerRuleName returns a boolean if a field has been set.

### GetPathPattern

`func (o *ListenerRule) GetPathPattern() string`

GetPathPattern returns the PathPattern field if non-nil, zero value otherwise.

### GetPathPatternOk

`func (o *ListenerRule) GetPathPatternOk() (*string, bool)`

GetPathPatternOk returns a tuple with the PathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPattern

`func (o *ListenerRule) SetPathPattern(v string)`

SetPathPattern sets PathPattern field to given value.

### HasPathPattern

`func (o *ListenerRule) HasPathPattern() bool`

HasPathPattern returns a boolean if a field has been set.

### GetPriority

`func (o *ListenerRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListenerRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ListenerRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ListenerRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetVmIds

`func (o *ListenerRule) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *ListenerRule) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *ListenerRule) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *ListenerRule) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


