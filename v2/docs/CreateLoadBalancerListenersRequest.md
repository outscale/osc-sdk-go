# CreateLoadBalancerListenersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Listeners** | [**[]ListenerForCreation**](ListenerForCreation.md) | One or more listeners for the load balancer. | 
**LoadBalancerName** | **string** | The name of the load balancer for which you want to create listeners. | 

## Methods

### NewCreateLoadBalancerListenersRequest

`func NewCreateLoadBalancerListenersRequest(listeners []ListenerForCreation, loadBalancerName string, ) *CreateLoadBalancerListenersRequest`

NewCreateLoadBalancerListenersRequest instantiates a new CreateLoadBalancerListenersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerListenersRequestWithDefaults

`func NewCreateLoadBalancerListenersRequestWithDefaults() *CreateLoadBalancerListenersRequest`

NewCreateLoadBalancerListenersRequestWithDefaults instantiates a new CreateLoadBalancerListenersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateLoadBalancerListenersRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateLoadBalancerListenersRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateLoadBalancerListenersRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateLoadBalancerListenersRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetListeners

`func (o *CreateLoadBalancerListenersRequest) GetListeners() []ListenerForCreation`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *CreateLoadBalancerListenersRequest) GetListenersOk() (*[]ListenerForCreation, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *CreateLoadBalancerListenersRequest) SetListeners(v []ListenerForCreation)`

SetListeners sets Listeners field to given value.


### GetLoadBalancerName

`func (o *CreateLoadBalancerListenersRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *CreateLoadBalancerListenersRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *CreateLoadBalancerListenersRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


