/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UnlinkLoadBalancerBackendMachinesRequest struct for UnlinkLoadBalancerBackendMachinesRequest
type UnlinkLoadBalancerBackendMachinesRequest struct {
	//  One or more public IPs of back-end VMs.
	BackendIps *[]string `json:"BackendIps,omitempty"`
	//  One or more IDs of back-end VMs.
	BackendVmIds *[]string `json:"BackendVmIds,omitempty"`
	//  If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	//  The name of the load balancer.
	LoadBalancerName string `json:"LoadBalancerName"`
}

// NewUnlinkLoadBalancerBackendMachinesRequest instantiates a new UnlinkLoadBalancerBackendMachinesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkLoadBalancerBackendMachinesRequest(loadBalancerName string) *UnlinkLoadBalancerBackendMachinesRequest {
	this := UnlinkLoadBalancerBackendMachinesRequest{}
	this.LoadBalancerName = loadBalancerName
	return &this
}

// NewUnlinkLoadBalancerBackendMachinesRequestWithDefaults instantiates a new UnlinkLoadBalancerBackendMachinesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkLoadBalancerBackendMachinesRequestWithDefaults() *UnlinkLoadBalancerBackendMachinesRequest {
	this := UnlinkLoadBalancerBackendMachinesRequest{}
	return &this
}

// GetBackendIps returns the BackendIps field value if set, zero value otherwise.
func (o *UnlinkLoadBalancerBackendMachinesRequest) GetBackendIps() []string {
	if o == nil || o.BackendIps == nil {
		var ret []string
		return ret
	}
	return *o.BackendIps
}

// GetBackendIpsOk returns a tuple with the BackendIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkLoadBalancerBackendMachinesRequest) GetBackendIpsOk() (*[]string, bool) {
	if o == nil || o.BackendIps == nil {
		return nil, false
	}
	return o.BackendIps, true
}

// HasBackendIps returns a boolean if a field has been set.
func (o *UnlinkLoadBalancerBackendMachinesRequest) HasBackendIps() bool {
	if o != nil && o.BackendIps != nil {
		return true
	}

	return false
}

// SetBackendIps gets a reference to the given []string and assigns it to the BackendIps field.
func (o *UnlinkLoadBalancerBackendMachinesRequest) SetBackendIps(v []string) {
	o.BackendIps = &v
}

// GetBackendVmIds returns the BackendVmIds field value if set, zero value otherwise.
func (o *UnlinkLoadBalancerBackendMachinesRequest) GetBackendVmIds() []string {
	if o == nil || o.BackendVmIds == nil {
		var ret []string
		return ret
	}
	return *o.BackendVmIds
}

// GetBackendVmIdsOk returns a tuple with the BackendVmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkLoadBalancerBackendMachinesRequest) GetBackendVmIdsOk() (*[]string, bool) {
	if o == nil || o.BackendVmIds == nil {
		return nil, false
	}
	return o.BackendVmIds, true
}

// HasBackendVmIds returns a boolean if a field has been set.
func (o *UnlinkLoadBalancerBackendMachinesRequest) HasBackendVmIds() bool {
	if o != nil && o.BackendVmIds != nil {
		return true
	}

	return false
}

// SetBackendVmIds gets a reference to the given []string and assigns it to the BackendVmIds field.
func (o *UnlinkLoadBalancerBackendMachinesRequest) SetBackendVmIds(v []string) {
	o.BackendVmIds = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkLoadBalancerBackendMachinesRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkLoadBalancerBackendMachinesRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkLoadBalancerBackendMachinesRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkLoadBalancerBackendMachinesRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *UnlinkLoadBalancerBackendMachinesRequest) GetLoadBalancerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *UnlinkLoadBalancerBackendMachinesRequest) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *UnlinkLoadBalancerBackendMachinesRequest) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

func (o UnlinkLoadBalancerBackendMachinesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackendIps != nil {
		toSerialize["BackendIps"] = o.BackendIps
	}
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

type NullableUnlinkLoadBalancerBackendMachinesRequest struct {
	value *UnlinkLoadBalancerBackendMachinesRequest
	isSet bool
}

func (v NullableUnlinkLoadBalancerBackendMachinesRequest) Get() *UnlinkLoadBalancerBackendMachinesRequest {
	return v.value
}

func (v *NullableUnlinkLoadBalancerBackendMachinesRequest) Set(val *UnlinkLoadBalancerBackendMachinesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkLoadBalancerBackendMachinesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkLoadBalancerBackendMachinesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkLoadBalancerBackendMachinesRequest(val *UnlinkLoadBalancerBackendMachinesRequest) *NullableUnlinkLoadBalancerBackendMachinesRequest {
	return &NullableUnlinkLoadBalancerBackendMachinesRequest{value: val, isSet: true}
}

func (v NullableUnlinkLoadBalancerBackendMachinesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkLoadBalancerBackendMachinesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
