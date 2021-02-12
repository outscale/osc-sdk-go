# DeleteLoadBalancerListenersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer for which you want to delete listeners. | 
**LoadBalancerPorts** | **[]int32** | One or more port numbers of the listeners you want to delete. | 

## Methods

### NewDeleteLoadBalancerListenersRequest

`func NewDeleteLoadBalancerListenersRequest(loadBalancerName string, loadBalancerPorts []int32, ) *DeleteLoadBalancerListenersRequest`

NewDeleteLoadBalancerListenersRequest instantiates a new DeleteLoadBalancerListenersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteLoadBalancerListenersRequestWithDefaults

`func NewDeleteLoadBalancerListenersRequestWithDefaults() *DeleteLoadBalancerListenersRequest`

NewDeleteLoadBalancerListenersRequestWithDefaults instantiates a new DeleteLoadBalancerListenersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteLoadBalancerListenersRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteLoadBalancerListenersRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteLoadBalancerListenersRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteLoadBalancerListenersRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *DeleteLoadBalancerListenersRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *DeleteLoadBalancerListenersRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *DeleteLoadBalancerListenersRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.


### GetLoadBalancerPorts

`func (o *DeleteLoadBalancerListenersRequest) GetLoadBalancerPorts() []int32`

GetLoadBalancerPorts returns the LoadBalancerPorts field if non-nil, zero value otherwise.

### GetLoadBalancerPortsOk

`func (o *DeleteLoadBalancerListenersRequest) GetLoadBalancerPortsOk() (*[]int32, bool)`

GetLoadBalancerPortsOk returns a tuple with the LoadBalancerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPorts

`func (o *DeleteLoadBalancerListenersRequest) SetLoadBalancerPorts(v []int32)`

SetLoadBalancerPorts sets LoadBalancerPorts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


