/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteVmsRequest struct for DeleteVmsRequest
type DeleteVmsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more IDs of VMs.
	VmIds []string `json:"VmIds"`
}

// NewDeleteVmsRequest instantiates a new DeleteVmsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVmsRequest(vmIds []string) *DeleteVmsRequest {
	this := DeleteVmsRequest{}
	this.VmIds = vmIds
	return &this
}

// NewDeleteVmsRequestWithDefaults instantiates a new DeleteVmsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVmsRequestWithDefaults() *DeleteVmsRequest {
	this := DeleteVmsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteVmsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVmsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteVmsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteVmsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetVmIds returns the VmIds field value
func (o *DeleteVmsRequest) GetVmIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value
// and a boolean to check if the value has been set.
func (o *DeleteVmsRequest) GetVmIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmIds, true
}

// SetVmIds sets field value
func (o *DeleteVmsRequest) SetVmIds(v []string) {
	o.VmIds = v
}

func (o DeleteVmsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["VmIds"] = o.VmIds
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteVmsRequest struct {
	value *DeleteVmsRequest
	isSet bool
}

func (v NullableDeleteVmsRequest) Get() *DeleteVmsRequest {
	return v.value
}

func (v *NullableDeleteVmsRequest) Set(val *DeleteVmsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVmsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVmsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVmsRequest(val *DeleteVmsRequest) *NullableDeleteVmsRequest {
	return &NullableDeleteVmsRequest{value: val, isSet: true}
}

func (v NullableDeleteVmsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVmsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
