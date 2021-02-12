# CreateListenerRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Listener** | [**LoadBalancerLight**](LoadBalancerLight.md) |  | 
**ListenerRule** | [**ListenerRuleForCreation**](ListenerRuleForCreation.md) |  | 
**VmIds** | **[]string** | The IDs of the backend VMs. | 

## Methods

### NewCreateListenerRuleRequest

`func NewCreateListenerRuleRequest(listener LoadBalancerLight, listenerRule ListenerRuleForCreation, vmIds []string, ) *CreateListenerRuleRequest`

NewCreateListenerRuleRequest instantiates a new CreateListenerRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListenerRuleRequestWithDefaults

`func NewCreateListenerRuleRequestWithDefaults() *CreateListenerRuleRequest`

NewCreateListenerRuleRequestWithDefaults instantiates a new CreateListenerRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateListenerRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateListenerRuleRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateListenerRuleRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateListenerRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetListener

`func (o *CreateListenerRuleRequest) GetListener() LoadBalancerLight`

GetListener returns the Listener field if non-nil, zero value otherwise.

### GetListenerOk

`func (o *CreateListenerRuleRequest) GetListenerOk() (*LoadBalancerLight, bool)`

GetListenerOk returns a tuple with the Listener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListener

`func (o *CreateListenerRuleRequest) SetListener(v LoadBalancerLight)`

SetListener sets Listener field to given value.


### GetListenerRule

`func (o *CreateListenerRuleRequest) GetListenerRule() ListenerRuleForCreation`

GetListenerRule returns the ListenerRule field if non-nil, zero value otherwise.

### GetListenerRuleOk

`func (o *CreateListenerRuleRequest) GetListenerRuleOk() (*ListenerRuleForCreation, bool)`

GetListenerRuleOk returns a tuple with the ListenerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerRule

`func (o *CreateListenerRuleRequest) SetListenerRule(v ListenerRuleForCreation)`

SetListenerRule sets ListenerRule field to given value.


### GetVmIds

`func (o *CreateListenerRuleRequest) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *CreateListenerRuleRequest) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *CreateListenerRuleRequest) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


