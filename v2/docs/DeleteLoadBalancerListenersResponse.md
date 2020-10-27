# DeleteLoadBalancerListenersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | Pointer to [**LoadBalancer**](LoadBalancer.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewDeleteLoadBalancerListenersResponse

`func NewDeleteLoadBalancerListenersResponse() *DeleteLoadBalancerListenersResponse`

NewDeleteLoadBalancerListenersResponse instantiates a new DeleteLoadBalancerListenersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteLoadBalancerListenersResponseWithDefaults

`func NewDeleteLoadBalancerListenersResponseWithDefaults() *DeleteLoadBalancerListenersResponse`

NewDeleteLoadBalancerListenersResponseWithDefaults instantiates a new DeleteLoadBalancerListenersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *DeleteLoadBalancerListenersResponse) GetLoadBalancer() LoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *DeleteLoadBalancerListenersResponse) GetLoadBalancerOk() (*LoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *DeleteLoadBalancerListenersResponse) SetLoadBalancer(v LoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *DeleteLoadBalancerListenersResponse) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetResponseContext

`func (o *DeleteLoadBalancerListenersResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *DeleteLoadBalancerListenersResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *DeleteLoadBalancerListenersResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *DeleteLoadBalancerListenersResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


