# LoadBalancerLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerName** | **string** | The name of the load balancer to which the listener is attached. | 
**LoadBalancerPort** | **int32** | The port of load balancer on which the load balancer is listening (between &#x60;1&#x60; and &#x60;65535&#x60; both included). | 

## Methods

### NewLoadBalancerLight

`func NewLoadBalancerLight(loadBalancerName string, loadBalancerPort int32, ) *LoadBalancerLight`

NewLoadBalancerLight instantiates a new LoadBalancerLight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerLightWithDefaults

`func NewLoadBalancerLightWithDefaults() *LoadBalancerLight`

NewLoadBalancerLightWithDefaults instantiates a new LoadBalancerLight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerName

`func (o *LoadBalancerLight) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *LoadBalancerLight) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *LoadBalancerLight) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.


### GetLoadBalancerPort

`func (o *LoadBalancerLight) GetLoadBalancerPort() int32`

GetLoadBalancerPort returns the LoadBalancerPort field if non-nil, zero value otherwise.

### GetLoadBalancerPortOk

`func (o *LoadBalancerLight) GetLoadBalancerPortOk() (*int32, bool)`

GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPort

`func (o *LoadBalancerLight) SetLoadBalancerPort(v int32)`

SetLoadBalancerPort sets LoadBalancerPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


