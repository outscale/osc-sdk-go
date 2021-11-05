/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// RegisterVmsInLoadBalancerRequest struct for RegisterVmsInLoadBalancerRequest
type RegisterVmsInLoadBalancerRequest struct {
	// One or more IDs of back-end VMs.<br /> Specifying the same ID several times has no effect as each back-end VM has equal weight.
	BackendVmIds []string `json:"BackendVmIds"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The name of the load balancer.
	LoadBalancerName string `json:"LoadBalancerName"`
}

// NewRegisterVmsInLoadBalancerRequest instantiates a new RegisterVmsInLoadBalancerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterVmsInLoadBalancerRequest(backendVmIds []string, loadBalancerName string) *RegisterVmsInLoadBalancerRequest {
	this := RegisterVmsInLoadBalancerRequest{}
	this.BackendVmIds = backendVmIds
	this.LoadBalancerName = loadBalancerName
	return &this
}

// NewRegisterVmsInLoadBalancerRequestWithDefaults instantiates a new RegisterVmsInLoadBalancerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterVmsInLoadBalancerRequestWithDefaults() *RegisterVmsInLoadBalancerRequest {
	this := RegisterVmsInLoadBalancerRequest{}
	return &this
}

// GetBackendVmIds returns the BackendVmIds field value
func (o *RegisterVmsInLoadBalancerRequest) GetBackendVmIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BackendVmIds
}

// GetBackendVmIdsOk returns a tuple with the BackendVmIds field value
// and a boolean to check if the value has been set.
func (o *RegisterVmsInLoadBalancerRequest) GetBackendVmIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendVmIds, true
}

// SetBackendVmIds sets field value
func (o *RegisterVmsInLoadBalancerRequest) SetBackendVmIds(v []string) {
	o.BackendVmIds = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *RegisterVmsInLoadBalancerRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterVmsInLoadBalancerRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *RegisterVmsInLoadBalancerRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *RegisterVmsInLoadBalancerRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *RegisterVmsInLoadBalancerRequest) GetLoadBalancerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *RegisterVmsInLoadBalancerRequest) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *RegisterVmsInLoadBalancerRequest) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

func (o RegisterVmsInLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["BackendVmIds"] = o.BackendVmIds
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterVmsInLoadBalancerRequest struct {
	value *RegisterVmsInLoadBalancerRequest
	isSet bool
}

func (v NullableRegisterVmsInLoadBalancerRequest) Get() *RegisterVmsInLoadBalancerRequest {
	return v.value
}

func (v *NullableRegisterVmsInLoadBalancerRequest) Set(val *RegisterVmsInLoadBalancerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterVmsInLoadBalancerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterVmsInLoadBalancerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterVmsInLoadBalancerRequest(val *RegisterVmsInLoadBalancerRequest) *NullableRegisterVmsInLoadBalancerRequest {
	return &NullableRegisterVmsInLoadBalancerRequest{value: val, isSet: true}
}

func (v NullableRegisterVmsInLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterVmsInLoadBalancerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}