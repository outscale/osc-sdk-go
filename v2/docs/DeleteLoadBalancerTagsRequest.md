# DeleteLoadBalancerTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerNames** | **[]string** | One or more load balancer names. | 
**Tags** | [**[]ResourceLoadBalancerTag**](ResourceLoadBalancerTag.md) | One or more tags to delete from the load balancers. | 

## Methods

### NewDeleteLoadBalancerTagsRequest

`func NewDeleteLoadBalancerTagsRequest(loadBalancerNames []string, tags []ResourceLoadBalancerTag, ) *DeleteLoadBalancerTagsRequest`

NewDeleteLoadBalancerTagsRequest instantiates a new DeleteLoadBalancerTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteLoadBalancerTagsRequestWithDefaults

`func NewDeleteLoadBalancerTagsRequestWithDefaults() *DeleteLoadBalancerTagsRequest`

NewDeleteLoadBalancerTagsRequestWithDefaults instantiates a new DeleteLoadBalancerTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteLoadBalancerTagsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteLoadBalancerTagsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteLoadBalancerTagsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteLoadBalancerTagsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerNames

`func (o *DeleteLoadBalancerTagsRequest) GetLoadBalancerNames() []string`

GetLoadBalancerNames returns the LoadBalancerNames field if non-nil, zero value otherwise.

### GetLoadBalancerNamesOk

`func (o *DeleteLoadBalancerTagsRequest) GetLoadBalancerNamesOk() (*[]string, bool)`

GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerNames

`func (o *DeleteLoadBalancerTagsRequest) SetLoadBalancerNames(v []string)`

SetLoadBalancerNames sets LoadBalancerNames field to given value.


### GetTags

`func (o *DeleteLoadBalancerTagsRequest) GetTags() []ResourceLoadBalancerTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeleteLoadBalancerTagsRequest) GetTagsOk() (*[]ResourceLoadBalancerTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeleteLoadBalancerTagsRequest) SetTags(v []ResourceLoadBalancerTag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


