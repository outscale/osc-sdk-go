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

// DeleteVmTemplateRequest struct for DeleteVmTemplateRequest
type DeleteVmTemplateRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the VM template you want to delete.
	VmTemplateId string `json:"VmTemplateId"`
}

// NewDeleteVmTemplateRequest instantiates a new DeleteVmTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVmTemplateRequest(vmTemplateId string) *DeleteVmTemplateRequest {
	this := DeleteVmTemplateRequest{}
	this.VmTemplateId = vmTemplateId
	return &this
}

// NewDeleteVmTemplateRequestWithDefaults instantiates a new DeleteVmTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVmTemplateRequestWithDefaults() *DeleteVmTemplateRequest {
	this := DeleteVmTemplateRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteVmTemplateRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVmTemplateRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteVmTemplateRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteVmTemplateRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetVmTemplateId returns the VmTemplateId field value
func (o *DeleteVmTemplateRequest) GetVmTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmTemplateId
}

// GetVmTemplateIdOk returns a tuple with the VmTemplateId field value
// and a boolean to check if the value has been set.
func (o *DeleteVmTemplateRequest) GetVmTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmTemplateId, true
}

// SetVmTemplateId sets field value
func (o *DeleteVmTemplateRequest) SetVmTemplateId(v string) {
	o.VmTemplateId = v
}

func (o DeleteVmTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["VmTemplateId"] = o.VmTemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteVmTemplateRequest struct {
	value *DeleteVmTemplateRequest
	isSet bool
}

func (v NullableDeleteVmTemplateRequest) Get() *DeleteVmTemplateRequest {
	return v.value
}

func (v *NullableDeleteVmTemplateRequest) Set(val *DeleteVmTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVmTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVmTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVmTemplateRequest(val *DeleteVmTemplateRequest) *NullableDeleteVmTemplateRequest {
	return &NullableDeleteVmTemplateRequest{value: val, isSet: true}
}

func (v NullableDeleteVmTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVmTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}