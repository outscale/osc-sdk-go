/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteLoadBalancerPolicyRequest struct for DeleteLoadBalancerPolicyRequest
type DeleteLoadBalancerPolicyRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The name of the load balancer for which you want to delete a policy.
	LoadBalancerName string `json:"LoadBalancerName"`
	// The name of the policy you want to delete.
	PolicyName string `json:"PolicyName"`
}

// NewDeleteLoadBalancerPolicyRequest instantiates a new DeleteLoadBalancerPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteLoadBalancerPolicyRequest(loadBalancerName string, policyName string) *DeleteLoadBalancerPolicyRequest {
	this := DeleteLoadBalancerPolicyRequest{}
	this.LoadBalancerName = loadBalancerName
	this.PolicyName = policyName
	return &this
}

// NewDeleteLoadBalancerPolicyRequestWithDefaults instantiates a new DeleteLoadBalancerPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteLoadBalancerPolicyRequestWithDefaults() *DeleteLoadBalancerPolicyRequest {
	this := DeleteLoadBalancerPolicyRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteLoadBalancerPolicyRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerPolicyRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteLoadBalancerPolicyRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteLoadBalancerPolicyRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *DeleteLoadBalancerPolicyRequest) GetLoadBalancerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerPolicyRequest) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *DeleteLoadBalancerPolicyRequest) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

// GetPolicyName returns the PolicyName field value
func (o *DeleteLoadBalancerPolicyRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerPolicyRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *DeleteLoadBalancerPolicyRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

func (o DeleteLoadBalancerPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	if true {
		toSerialize["PolicyName"] = o.PolicyName
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteLoadBalancerPolicyRequest struct {
	value *DeleteLoadBalancerPolicyRequest
	isSet bool
}

func (v NullableDeleteLoadBalancerPolicyRequest) Get() *DeleteLoadBalancerPolicyRequest {
	return v.value
}

func (v *NullableDeleteLoadBalancerPolicyRequest) Set(val *DeleteLoadBalancerPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteLoadBalancerPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteLoadBalancerPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteLoadBalancerPolicyRequest(val *DeleteLoadBalancerPolicyRequest) *NullableDeleteLoadBalancerPolicyRequest {
	return &NullableDeleteLoadBalancerPolicyRequest{value: val, isSet: true}
}

func (v NullableDeleteLoadBalancerPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteLoadBalancerPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
