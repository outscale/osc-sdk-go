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

// UnlinkFlexibleGpuRequest struct for UnlinkFlexibleGpuRequest
type UnlinkFlexibleGpuRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the fGPU you want to detach from your VM.
	FlexibleGpuId string `json:"FlexibleGpuId"`
}

// NewUnlinkFlexibleGpuRequest instantiates a new UnlinkFlexibleGpuRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkFlexibleGpuRequest(flexibleGpuId string) *UnlinkFlexibleGpuRequest {
	this := UnlinkFlexibleGpuRequest{}
	this.FlexibleGpuId = flexibleGpuId
	return &this
}

// NewUnlinkFlexibleGpuRequestWithDefaults instantiates a new UnlinkFlexibleGpuRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkFlexibleGpuRequestWithDefaults() *UnlinkFlexibleGpuRequest {
	this := UnlinkFlexibleGpuRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkFlexibleGpuRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkFlexibleGpuRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkFlexibleGpuRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkFlexibleGpuRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFlexibleGpuId returns the FlexibleGpuId field value
func (o *UnlinkFlexibleGpuRequest) GetFlexibleGpuId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlexibleGpuId
}

// GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field value
// and a boolean to check if the value has been set.
func (o *UnlinkFlexibleGpuRequest) GetFlexibleGpuIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlexibleGpuId, true
}

// SetFlexibleGpuId sets field value
func (o *UnlinkFlexibleGpuRequest) SetFlexibleGpuId(v string) {
	o.FlexibleGpuId = v
}

func (o UnlinkFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["FlexibleGpuId"] = o.FlexibleGpuId
	}
	return json.Marshal(toSerialize)
}

type NullableUnlinkFlexibleGpuRequest struct {
	value *UnlinkFlexibleGpuRequest
	isSet bool
}

func (v NullableUnlinkFlexibleGpuRequest) Get() *UnlinkFlexibleGpuRequest {
	return v.value
}

func (v *NullableUnlinkFlexibleGpuRequest) Set(val *UnlinkFlexibleGpuRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkFlexibleGpuRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkFlexibleGpuRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkFlexibleGpuRequest(val *UnlinkFlexibleGpuRequest) *NullableUnlinkFlexibleGpuRequest {
	return &NullableUnlinkFlexibleGpuRequest{value: val, isSet: true}
}

func (v NullableUnlinkFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkFlexibleGpuRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
