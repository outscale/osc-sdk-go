# ReadLoadBalancerTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerNames** | **[]string** | One or more load balancer names. | 

## Methods

### NewReadLoadBalancerTagsRequest

`func NewReadLoadBalancerTagsRequest(loadBalancerNames []string, ) *ReadLoadBalancerTagsRequest`

NewReadLoadBalancerTagsRequest instantiates a new ReadLoadBalancerTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadLoadBalancerTagsRequestWithDefaults

`func NewReadLoadBalancerTagsRequestWithDefaults() *ReadLoadBalancerTagsRequest`

NewReadLoadBalancerTagsRequestWithDefaults instantiates a new ReadLoadBalancerTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadLoadBalancerTagsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadLoadBalancerTagsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadLoadBalancerTagsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadLoadBalancerTagsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerNames

`func (o *ReadLoadBalancerTagsRequest) GetLoadBalancerNames() []string`

GetLoadBalancerNames returns the LoadBalancerNames field if non-nil, zero value otherwise.

### GetLoadBalancerNamesOk

`func (o *ReadLoadBalancerTagsRequest) GetLoadBalancerNamesOk() (*[]string, bool)`

GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerNames

`func (o *ReadLoadBalancerTagsRequest) SetLoadBalancerNames(v []string)`

SetLoadBalancerNames sets LoadBalancerNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


