# LoadBalancerStickyCookiePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CookieExpirationPeriod** | Pointer to **int32** | The time period, in seconds, after which the cookie should be considered stale.&lt;br /&gt; If &#x60;1&#x60;, the stickiness session lasts for the duration of the browser session. | [optional] 
**PolicyName** | Pointer to **string** | The name of the stickiness policy. | [optional] 

## Methods

### NewLoadBalancerStickyCookiePolicy

`func NewLoadBalancerStickyCookiePolicy() *LoadBalancerStickyCookiePolicy`

NewLoadBalancerStickyCookiePolicy instantiates a new LoadBalancerStickyCookiePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerStickyCookiePolicyWithDefaults

`func NewLoadBalancerStickyCookiePolicyWithDefaults() *LoadBalancerStickyCookiePolicy`

NewLoadBalancerStickyCookiePolicyWithDefaults instantiates a new LoadBalancerStickyCookiePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookieExpirationPeriod

`func (o *LoadBalancerStickyCookiePolicy) GetCookieExpirationPeriod() int32`

GetCookieExpirationPeriod returns the CookieExpirationPeriod field if non-nil, zero value otherwise.

### GetCookieExpirationPeriodOk

`func (o *LoadBalancerStickyCookiePolicy) GetCookieExpirationPeriodOk() (*int32, bool)`

GetCookieExpirationPeriodOk returns a tuple with the CookieExpirationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieExpirationPeriod

`func (o *LoadBalancerStickyCookiePolicy) SetCookieExpirationPeriod(v int32)`

SetCookieExpirationPeriod sets CookieExpirationPeriod field to given value.

### HasCookieExpirationPeriod

`func (o *LoadBalancerStickyCookiePolicy) HasCookieExpirationPeriod() bool`

HasCookieExpirationPeriod returns a boolean if a field has been set.

### GetPolicyName

`func (o *LoadBalancerStickyCookiePolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *LoadBalancerStickyCookiePolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *LoadBalancerStickyCookiePolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *LoadBalancerStickyCookiePolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


