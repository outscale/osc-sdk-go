/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LoadBalancerStickyCookiePolicy Information about the stickiness policy.
type LoadBalancerStickyCookiePolicy struct {
	// The time period, in seconds, after which the cookie should be considered stale.<br /> If `1`, the stickiness session lasts for the duration of the browser session.
	CookieExpirationPeriod *int32 `json:"CookieExpirationPeriod,omitempty"`
	// The name of the stickiness policy.
	PolicyName *string `json:"PolicyName,omitempty"`
}

// NewLoadBalancerStickyCookiePolicy instantiates a new LoadBalancerStickyCookiePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerStickyCookiePolicy() *LoadBalancerStickyCookiePolicy {
	this := LoadBalancerStickyCookiePolicy{}
	return &this
}

// NewLoadBalancerStickyCookiePolicyWithDefaults instantiates a new LoadBalancerStickyCookiePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerStickyCookiePolicyWithDefaults() *LoadBalancerStickyCookiePolicy {
	this := LoadBalancerStickyCookiePolicy{}
	return &this
}

// GetCookieExpirationPeriod returns the CookieExpirationPeriod field value if set, zero value otherwise.
func (o *LoadBalancerStickyCookiePolicy) GetCookieExpirationPeriod() int32 {
	if o == nil || o.CookieExpirationPeriod == nil {
		var ret int32
		return ret
	}
	return *o.CookieExpirationPeriod
}

// GetCookieExpirationPeriodOk returns a tuple with the CookieExpirationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerStickyCookiePolicy) GetCookieExpirationPeriodOk() (*int32, bool) {
	if o == nil || o.CookieExpirationPeriod == nil {
		return nil, false
	}
	return o.CookieExpirationPeriod, true
}

// HasCookieExpirationPeriod returns a boolean if a field has been set.
func (o *LoadBalancerStickyCookiePolicy) HasCookieExpirationPeriod() bool {
	if o != nil && o.CookieExpirationPeriod != nil {
		return true
	}

	return false
}

// SetCookieExpirationPeriod gets a reference to the given int32 and assigns it to the CookieExpirationPeriod field.
func (o *LoadBalancerStickyCookiePolicy) SetCookieExpirationPeriod(v int32) {
	o.CookieExpirationPeriod = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *LoadBalancerStickyCookiePolicy) GetPolicyName() string {
	if o == nil || o.PolicyName == nil {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerStickyCookiePolicy) GetPolicyNameOk() (*string, bool) {
	if o == nil || o.PolicyName == nil {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *LoadBalancerStickyCookiePolicy) HasPolicyName() bool {
	if o != nil && o.PolicyName != nil {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *LoadBalancerStickyCookiePolicy) SetPolicyName(v string) {
	o.PolicyName = &v
}

func (o LoadBalancerStickyCookiePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CookieExpirationPeriod != nil {
		toSerialize["CookieExpirationPeriod"] = o.CookieExpirationPeriod
	}
	if o.PolicyName != nil {
		toSerialize["PolicyName"] = o.PolicyName
	}
	return json.Marshal(toSerialize)
}

type NullableLoadBalancerStickyCookiePolicy struct {
	value *LoadBalancerStickyCookiePolicy
	isSet bool
}

func (v NullableLoadBalancerStickyCookiePolicy) Get() *LoadBalancerStickyCookiePolicy {
	return v.value
}

func (v *NullableLoadBalancerStickyCookiePolicy) Set(val *LoadBalancerStickyCookiePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerStickyCookiePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerStickyCookiePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerStickyCookiePolicy(val *LoadBalancerStickyCookiePolicy) *NullableLoadBalancerStickyCookiePolicy {
	return &NullableLoadBalancerStickyCookiePolicy{value: val, isSet: true}
}

func (v NullableLoadBalancerStickyCookiePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerStickyCookiePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
