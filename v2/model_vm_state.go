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

// VmState Information about the state of the VM.
type VmState struct {
	// The current state of the VM (`InService` \\| `OutOfService` \\| `Unknown`).
	CurrentState *string `json:"CurrentState,omitempty"`
	// The previous state of the VM (`InService` \\| `OutOfService` \\| `Unknown`).
	PreviousState *string `json:"PreviousState,omitempty"`
	// The ID of the VM.
	VmId *string `json:"VmId,omitempty"`
}

// NewVmState instantiates a new VmState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmState() *VmState {
	this := VmState{}
	return &this
}

// NewVmStateWithDefaults instantiates a new VmState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmStateWithDefaults() *VmState {
	this := VmState{}
	return &this
}

// GetCurrentState returns the CurrentState field value if set, zero value otherwise.
func (o *VmState) GetCurrentState() string {
	if o == nil || o.CurrentState == nil {
		var ret string
		return ret
	}
	return *o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmState) GetCurrentStateOk() (*string, bool) {
	if o == nil || o.CurrentState == nil {
		return nil, false
	}
	return o.CurrentState, true
}

// HasCurrentState returns a boolean if a field has been set.
func (o *VmState) HasCurrentState() bool {
	if o != nil && o.CurrentState != nil {
		return true
	}

	return false
}

// SetCurrentState gets a reference to the given string and assigns it to the CurrentState field.
func (o *VmState) SetCurrentState(v string) {
	o.CurrentState = &v
}

// GetPreviousState returns the PreviousState field value if set, zero value otherwise.
func (o *VmState) GetPreviousState() string {
	if o == nil || o.PreviousState == nil {
		var ret string
		return ret
	}
	return *o.PreviousState
}

// GetPreviousStateOk returns a tuple with the PreviousState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmState) GetPreviousStateOk() (*string, bool) {
	if o == nil || o.PreviousState == nil {
		return nil, false
	}
	return o.PreviousState, true
}

// HasPreviousState returns a boolean if a field has been set.
func (o *VmState) HasPreviousState() bool {
	if o != nil && o.PreviousState != nil {
		return true
	}

	return false
}

// SetPreviousState gets a reference to the given string and assigns it to the PreviousState field.
func (o *VmState) SetPreviousState(v string) {
	o.PreviousState = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *VmState) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmState) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *VmState) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *VmState) SetVmId(v string) {
	o.VmId = &v
}

func (o VmState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentState != nil {
		toSerialize["CurrentState"] = o.CurrentState
	}
	if o.PreviousState != nil {
		toSerialize["PreviousState"] = o.PreviousState
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableVmState struct {
	value *VmState
	isSet bool
}

func (v NullableVmState) Get() *VmState {
	return v.value
}

func (v *NullableVmState) Set(val *VmState) {
	v.value = val
	v.isSet = true
}

func (v NullableVmState) IsSet() bool {
	return v.isSet
}

func (v *NullableVmState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmState(val *VmState) *NullableVmState {
	return &NullableVmState{value: val, isSet: true}
}

func (v NullableVmState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
