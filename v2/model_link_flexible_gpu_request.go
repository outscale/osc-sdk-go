/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkFlexibleGpuRequest struct for LinkFlexibleGpuRequest
type LinkFlexibleGpuRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the fGPU you want to attach.
	FlexibleGpuId string `json:"FlexibleGpuId"`
	// The ID of the VM you want to attach the fGPU to.
	VmId string `json:"VmId"`
}

// NewLinkFlexibleGpuRequest instantiates a new LinkFlexibleGpuRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkFlexibleGpuRequest(flexibleGpuId string, vmId string, ) *LinkFlexibleGpuRequest {
	this := LinkFlexibleGpuRequest{}
	this.FlexibleGpuId = flexibleGpuId
	this.VmId = vmId
	return &this
}

// NewLinkFlexibleGpuRequestWithDefaults instantiates a new LinkFlexibleGpuRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkFlexibleGpuRequestWithDefaults() *LinkFlexibleGpuRequest {
	this := LinkFlexibleGpuRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkFlexibleGpuRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkFlexibleGpuRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkFlexibleGpuRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkFlexibleGpuRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFlexibleGpuId returns the FlexibleGpuId field value
func (o *LinkFlexibleGpuRequest) GetFlexibleGpuId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FlexibleGpuId
}

// GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field value
// and a boolean to check if the value has been set.
func (o *LinkFlexibleGpuRequest) GetFlexibleGpuIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FlexibleGpuId, true
}

// SetFlexibleGpuId sets field value
func (o *LinkFlexibleGpuRequest) SetFlexibleGpuId(v string) {
	o.FlexibleGpuId = v
}

// GetVmId returns the VmId field value
func (o *LinkFlexibleGpuRequest) GetVmId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value
// and a boolean to check if the value has been set.
func (o *LinkFlexibleGpuRequest) GetVmIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VmId, true
}

// SetVmId sets field value
func (o *LinkFlexibleGpuRequest) SetVmId(v string) {
	o.VmId = v
}

func (o LinkFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["FlexibleGpuId"] = o.FlexibleGpuId
	}
	if true {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkFlexibleGpuRequest struct {
	value *LinkFlexibleGpuRequest
	isSet bool
}

func (v NullableLinkFlexibleGpuRequest) Get() *LinkFlexibleGpuRequest {
	return v.value
}

func (v *NullableLinkFlexibleGpuRequest) Set(val *LinkFlexibleGpuRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkFlexibleGpuRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkFlexibleGpuRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkFlexibleGpuRequest(val *LinkFlexibleGpuRequest) *NullableLinkFlexibleGpuRequest {
	return &NullableLinkFlexibleGpuRequest{value: val, isSet: true}
}

func (v NullableLinkFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkFlexibleGpuRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


