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

// RebootVmsRequest struct for RebootVmsRequest
type RebootVmsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more IDs of the VMs you want to reboot.
	VmIds []string `json:"VmIds"`
}

// NewRebootVmsRequest instantiates a new RebootVmsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRebootVmsRequest(vmIds []string) *RebootVmsRequest {
	this := RebootVmsRequest{}
	this.VmIds = vmIds
	return &this
}

// NewRebootVmsRequestWithDefaults instantiates a new RebootVmsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRebootVmsRequestWithDefaults() *RebootVmsRequest {
	this := RebootVmsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *RebootVmsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootVmsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *RebootVmsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *RebootVmsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetVmIds returns the VmIds field value
func (o *RebootVmsRequest) GetVmIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value
// and a boolean to check if the value has been set.
func (o *RebootVmsRequest) GetVmIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmIds, true
}

// SetVmIds sets field value
func (o *RebootVmsRequest) SetVmIds(v []string) {
	o.VmIds = v
}

func (o RebootVmsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["VmIds"] = o.VmIds
	}
	return json.Marshal(toSerialize)
}

type NullableRebootVmsRequest struct {
	value *RebootVmsRequest
	isSet bool
}

func (v NullableRebootVmsRequest) Get() *RebootVmsRequest {
	return v.value
}

func (v *NullableRebootVmsRequest) Set(val *RebootVmsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRebootVmsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRebootVmsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebootVmsRequest(val *RebootVmsRequest) *NullableRebootVmsRequest {
	return &NullableRebootVmsRequest{value: val, isSet: true}
}

func (v NullableRebootVmsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebootVmsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
