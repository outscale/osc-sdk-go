/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadVmsHealthRequest struct for ReadVmsHealthRequest
type ReadVmsHealthRequest struct {
	// One or more IDs of back-end VMs.
	BackendVmIds *[]string `json:"BackendVmIds,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The name of the load balancer.
	LoadBalancerName string `json:"LoadBalancerName"`
}

// NewReadVmsHealthRequest instantiates a new ReadVmsHealthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVmsHealthRequest(loadBalancerName string) *ReadVmsHealthRequest {
	this := ReadVmsHealthRequest{}
	this.LoadBalancerName = loadBalancerName
	return &this
}

// NewReadVmsHealthRequestWithDefaults instantiates a new ReadVmsHealthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVmsHealthRequestWithDefaults() *ReadVmsHealthRequest {
	this := ReadVmsHealthRequest{}
	return &this
}

// GetBackendVmIds returns the BackendVmIds field value if set, zero value otherwise.
func (o *ReadVmsHealthRequest) GetBackendVmIds() []string {
	if o == nil || o.BackendVmIds == nil {
		var ret []string
		return ret
	}
	return *o.BackendVmIds
}

// GetBackendVmIdsOk returns a tuple with the BackendVmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsHealthRequest) GetBackendVmIdsOk() (*[]string, bool) {
	if o == nil || o.BackendVmIds == nil {
		return nil, false
	}
	return o.BackendVmIds, true
}

// HasBackendVmIds returns a boolean if a field has been set.
func (o *ReadVmsHealthRequest) HasBackendVmIds() bool {
	if o != nil && o.BackendVmIds != nil {
		return true
	}

	return false
}

// SetBackendVmIds gets a reference to the given []string and assigns it to the BackendVmIds field.
func (o *ReadVmsHealthRequest) SetBackendVmIds(v []string) {
	o.BackendVmIds = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadVmsHealthRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsHealthRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadVmsHealthRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadVmsHealthRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *ReadVmsHealthRequest) GetLoadBalancerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *ReadVmsHealthRequest) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *ReadVmsHealthRequest) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

func (o ReadVmsHealthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackendVmIds != nil {
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

type NullableReadVmsHealthRequest struct {
	value *ReadVmsHealthRequest
	isSet bool
}

func (v NullableReadVmsHealthRequest) Get() *ReadVmsHealthRequest {
	return v.value
}

func (v *NullableReadVmsHealthRequest) Set(val *ReadVmsHealthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVmsHealthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVmsHealthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVmsHealthRequest(val *ReadVmsHealthRequest) *NullableReadVmsHealthRequest {
	return &NullableReadVmsHealthRequest{value: val, isSet: true}
}

func (v NullableReadVmsHealthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVmsHealthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
