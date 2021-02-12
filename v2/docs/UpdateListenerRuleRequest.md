# UpdateListenerRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**HostPattern** | Pointer to **string** | A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?]. | [optional] 
**ListenerRuleName** | **string** | The name of the listener rule. | 
**PathPattern** | Pointer to **string** | A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~\&quot;&#39;@:+?]. | [optional] 

## Methods

### NewUpdateListenerRuleRequest

`func NewUpdateListenerRuleRequest(listenerRuleName string, ) *UpdateListenerRuleRequest`

NewUpdateListenerRuleRequest instantiates a new UpdateListenerRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListenerRuleRequestWithDefaults

`func NewUpdateListenerRuleRequestWithDefaults() *UpdateListenerRuleRequest`

NewUpdateListenerRuleRequestWithDefaults instantiates a new UpdateListenerRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateListenerRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateListenerRuleRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateListenerRuleRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateListenerRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetHostPattern

`func (o *UpdateListenerRuleRequest) GetHostPattern() string`

GetHostPattern returns the HostPattern field if non-nil, zero value otherwise.

### GetHostPatternOk

`func (o *UpdateListenerRuleRequest) GetHostPatternOk() (*string, bool)`

GetHostPatternOk returns a tuple with the HostPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPattern

`func (o *UpdateListenerRuleRequest) SetHostPattern(v string)`

SetHostPattern sets HostPattern field to given value.

### HasHostPattern

`func (o *UpdateListenerRuleRequest) HasHostPattern() bool`

HasHostPattern returns a boolean if a field has been set.

### GetListenerRuleName

`func (o *UpdateListenerRuleRequest) GetListenerRuleName() string`

GetListenerRuleName returns the ListenerRuleName field if non-nil, zero value otherwise.

### GetListenerRuleNameOk

`func (o *UpdateListenerRuleRequest) GetListenerRuleNameOk() (*string, bool)`

GetListenerRuleNameOk returns a tuple with the ListenerRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerRuleName

`func (o *UpdateListenerRuleRequest) SetListenerRuleName(v string)`

SetListenerRuleName sets ListenerRuleName field to given value.


### GetPathPattern

`func (o *UpdateListenerRuleRequest) GetPathPattern() string`

GetPathPattern returns the PathPattern field if non-nil, zero value otherwise.

### GetPathPatternOk

`func (o *UpdateListenerRuleRequest) GetPathPatternOk() (*string, bool)`

GetPathPatternOk returns a tuple with the PathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPattern

`func (o *UpdateListenerRuleRequest) SetPathPattern(v string)`

SetPathPattern sets PathPattern field to given value.

### HasPathPattern

`func (o *UpdateListenerRuleRequest) HasPathPattern() bool`

HasPathPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


