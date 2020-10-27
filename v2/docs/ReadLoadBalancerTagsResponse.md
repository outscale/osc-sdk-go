# ReadLoadBalancerTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Tags** | Pointer to [**[]LoadBalancerTag**](LoadBalancerTag.md) | Information about one or more load balancer tags. | [optional] 

## Methods

### NewReadLoadBalancerTagsResponse

`func NewReadLoadBalancerTagsResponse() *ReadLoadBalancerTagsResponse`

NewReadLoadBalancerTagsResponse instantiates a new ReadLoadBalancerTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadLoadBalancerTagsResponseWithDefaults

`func NewReadLoadBalancerTagsResponseWithDefaults() *ReadLoadBalancerTagsResponse`

NewReadLoadBalancerTagsResponseWithDefaults instantiates a new ReadLoadBalancerTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadLoadBalancerTagsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadLoadBalancerTagsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadLoadBalancerTagsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadLoadBalancerTagsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetTags

`func (o *ReadLoadBalancerTagsResponse) GetTags() []LoadBalancerTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReadLoadBalancerTagsResponse) GetTagsOk() (*[]LoadBalancerTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReadLoadBalancerTagsResponse) SetTags(v []LoadBalancerTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReadLoadBalancerTagsResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


