# CreateLoadBalancerTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerNames** | **[]string** | One or more load balancer names. | 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags to add to the specified load balancers. | 

## Methods

### NewCreateLoadBalancerTagsRequest

`func NewCreateLoadBalancerTagsRequest(loadBalancerNames []string, tags []ResourceTag, ) *CreateLoadBalancerTagsRequest`

NewCreateLoadBalancerTagsRequest instantiates a new CreateLoadBalancerTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerTagsRequestWithDefaults

`func NewCreateLoadBalancerTagsRequestWithDefaults() *CreateLoadBalancerTagsRequest`

NewCreateLoadBalancerTagsRequestWithDefaults instantiates a new CreateLoadBalancerTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateLoadBalancerTagsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateLoadBalancerTagsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateLoadBalancerTagsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateLoadBalancerTagsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerNames

`func (o *CreateLoadBalancerTagsRequest) GetLoadBalancerNames() []string`

GetLoadBalancerNames returns the LoadBalancerNames field if non-nil, zero value otherwise.

### GetLoadBalancerNamesOk

`func (o *CreateLoadBalancerTagsRequest) GetLoadBalancerNamesOk() (*[]string, bool)`

GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerNames

`func (o *CreateLoadBalancerTagsRequest) SetLoadBalancerNames(v []string)`

SetLoadBalancerNames sets LoadBalancerNames field to given value.


### GetTags

`func (o *CreateLoadBalancerTagsRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateLoadBalancerTagsRequest) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateLoadBalancerTagsRequest) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


