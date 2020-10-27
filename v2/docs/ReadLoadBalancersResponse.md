# ReadLoadBalancersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancers** | Pointer to [**[]LoadBalancer**](LoadBalancer.md) | Information about one or more load balancers. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadLoadBalancersResponse

`func NewReadLoadBalancersResponse() *ReadLoadBalancersResponse`

NewReadLoadBalancersResponse instantiates a new ReadLoadBalancersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadLoadBalancersResponseWithDefaults

`func NewReadLoadBalancersResponseWithDefaults() *ReadLoadBalancersResponse`

NewReadLoadBalancersResponseWithDefaults instantiates a new ReadLoadBalancersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancers

`func (o *ReadLoadBalancersResponse) GetLoadBalancers() []LoadBalancer`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *ReadLoadBalancersResponse) GetLoadBalancersOk() (*[]LoadBalancer, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *ReadLoadBalancersResponse) SetLoadBalancers(v []LoadBalancer)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *ReadLoadBalancersResponse) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadLoadBalancersResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadLoadBalancersResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadLoadBalancersResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadLoadBalancersResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


