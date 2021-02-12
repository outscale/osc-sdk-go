# CreateLoadBalancerPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CookieName** | Pointer to **string** | The name of the application cookie used for stickiness. This parameter is required if you create a stickiness policy based on an application-generated cookie. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | **string** | The name of the load balancer for which you want to create a policy. | 
**PolicyName** | **string** | The name of the policy. This name must be unique and consist of alphanumeric characters and dashes (-). | 
**PolicyType** | **string** | The type of stickiness policy you want to create: &#x60;app&#x60; or &#x60;load_balancer&#x60;. | 

## Methods

### NewCreateLoadBalancerPolicyRequest

`func NewCreateLoadBalancerPolicyRequest(loadBalancerName string, policyName string, policyType string, ) *CreateLoadBalancerPolicyRequest`

NewCreateLoadBalancerPolicyRequest instantiates a new CreateLoadBalancerPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerPolicyRequestWithDefaults

`func NewCreateLoadBalancerPolicyRequestWithDefaults() *CreateLoadBalancerPolicyRequest`

NewCreateLoadBalancerPolicyRequestWithDefaults instantiates a new CreateLoadBalancerPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookieName

`func (o *CreateLoadBalancerPolicyRequest) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *CreateLoadBalancerPolicyRequest) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *CreateLoadBalancerPolicyRequest) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *CreateLoadBalancerPolicyRequest) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateLoadBalancerPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateLoadBalancerPolicyRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateLoadBalancerPolicyRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateLoadBalancerPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *CreateLoadBalancerPolicyRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *CreateLoadBalancerPolicyRequest) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *CreateLoadBalancerPolicyRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.


### GetPolicyName

`func (o *CreateLoadBalancerPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *CreateLoadBalancerPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *CreateLoadBalancerPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetPolicyType

`func (o *CreateLoadBalancerPolicyRequest) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *CreateLoadBalancerPolicyRequest) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *CreateLoadBalancerPolicyRequest) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


