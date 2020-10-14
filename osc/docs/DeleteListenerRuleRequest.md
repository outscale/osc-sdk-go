# DeleteListenerRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ListenerRuleName** | **string** | The name of the rule you want to delete. | 

## Methods

### NewDeleteListenerRuleRequest

`func NewDeleteListenerRuleRequest(listenerRuleName string, ) *DeleteListenerRuleRequest`

NewDeleteListenerRuleRequest instantiates a new DeleteListenerRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteListenerRuleRequestWithDefaults

`func NewDeleteListenerRuleRequestWithDefaults() *DeleteListenerRuleRequest`

NewDeleteListenerRuleRequestWithDefaults instantiates a new DeleteListenerRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteListenerRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteListenerRuleRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteListenerRuleRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteListenerRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetListenerRuleName

`func (o *DeleteListenerRuleRequest) GetListenerRuleName() string`

GetListenerRuleName returns the ListenerRuleName field if non-nil, zero value otherwise.

### GetListenerRuleNameOk

`func (o *DeleteListenerRuleRequest) GetListenerRuleNameOk() (*string, bool)`

GetListenerRuleNameOk returns a tuple with the ListenerRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerRuleName

`func (o *DeleteListenerRuleRequest) SetListenerRuleName(v string)`

SetListenerRuleName sets ListenerRuleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


