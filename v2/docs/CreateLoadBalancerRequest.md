# CreateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Listeners** | [**[]ListenerForCreation**](ListenerForCreation.md) | One or more listeners to create. | 
**LoadBalancerName** | **string** | The unique name of the load balancer (32 alphanumeric or hyphen characters maximum, but cannot start or end with a hyphen). | 
**LoadBalancerType** | Pointer to **string** | The type of load balancer: &#x60;internet-facing&#x60; or &#x60;internal&#x60;. Use this parameter only for load balancers in a Net. | [optional] 
**SecurityGroups** | Pointer to **[]string** | (Net only) One or more IDs of security groups you want to assign to the load balancer. If not specified, the default security group of the Net is assigned to the load balancer. | [optional] 
**Subnets** | Pointer to **[]string** | One or more IDs of Subnets in your Net that you want to attach to the load balancer. | [optional] 
**SubregionNames** | Pointer to **[]string** | One or more names of Subregions (currently, only one Subregion is supported). This parameter is not required if you create a load balancer in a Net. To create an internal load balancer, use the &#x60;LoadBalancerType&#x60; parameter. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags assigned to the load balancer. | [optional] 

## Methods

### NewCreateLoadBalancerRequest

`func NewCreateLoadBalancerRequest(listeners []ListenerForCreation, loadBalancerName string, ) *CreateLoadBalancerRequest`

NewCreateLoadBalancerRequest instantiates a new CreateLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerRequestWithDefaults

`func NewCreateLoadBalancerRequestWithDefaults() *CreateLoadBalancerRequest`

NewCreateLoadBalancerRequestWithDefaults instantiates a new CreateLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateLoadBalancerRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateLoadBalancerRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateLoadBalancerRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateLoadBalancerRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetListeners

`func (o *CreateLoadBalancerRequest) GetListeners() []ListenerForCreation`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *CreateLoadBalancerRequest) GetListenersOk() (*[]ListenerForCreation, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *CreateLoadBalancerRequest) SetListeners(v []ListenerForCreation)`

SetListeners sets Listeners field to given value.


### GetLoadBalancerName

`func (o *CreateLoadBalancerRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *CreateLoadBalancerRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *CreateLoadBalancerRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.


### GetLoadBalancerType

`func (o *CreateLoadBalancerRequest) GetLoadBalancerType() string`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *CreateLoadBalancerRequest) GetLoadBalancerTypeOk() (*string, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *CreateLoadBalancerRequest) SetLoadBalancerType(v string)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *CreateLoadBalancerRequest) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CreateLoadBalancerRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateLoadBalancerRequest) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateLoadBalancerRequest) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CreateLoadBalancerRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSubnets

`func (o *CreateLoadBalancerRequest) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateLoadBalancerRequest) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateLoadBalancerRequest) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *CreateLoadBalancerRequest) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSubregionNames

`func (o *CreateLoadBalancerRequest) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *CreateLoadBalancerRequest) GetSubregionNamesOk() (*[]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionNames

`func (o *CreateLoadBalancerRequest) SetSubregionNames(v []string)`

SetSubregionNames sets SubregionNames field to given value.

### HasSubregionNames

`func (o *CreateLoadBalancerRequest) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### GetTags

`func (o *CreateLoadBalancerRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateLoadBalancerRequest) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateLoadBalancerRequest) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateLoadBalancerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


