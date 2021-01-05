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

// BackendVmHealth Information about the health of a back-end VM.
type BackendVmHealth struct {
	// The description of the state of the back-end VM.
	Description *string `json:"Description,omitempty"`
	// The state of the back-end VM (`InService` \\| `OutOfService` \\| `Unknown`).
	State *string `json:"State,omitempty"`
	// Information about the cause of `OutOfService` VMs.<br /> Specifically, whether the cause is Elastic Load Balancing or the VM (`ELB` \\| `Instance` \\| `N/A`).
	StateReason *string `json:"StateReason,omitempty"`
	// The ID of the back-end VM.
	VmId *string `json:"VmId,omitempty"`
}

// NewBackendVmHealth instantiates a new BackendVmHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackendVmHealth() *BackendVmHealth {
	this := BackendVmHealth{}
	return &this
}

// NewBackendVmHealthWithDefaults instantiates a new BackendVmHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackendVmHealthWithDefaults() *BackendVmHealth {
	this := BackendVmHealth{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BackendVmHealth) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendVmHealth) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BackendVmHealth) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BackendVmHealth) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BackendVmHealth) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendVmHealth) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BackendVmHealth) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BackendVmHealth) SetState(v string) {
	o.State = &v
}

// GetStateReason returns the StateReason field value if set, zero value otherwise.
func (o *BackendVmHealth) GetStateReason() string {
	if o == nil || o.StateReason == nil {
		var ret string
		return ret
	}
	return *o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendVmHealth) GetStateReasonOk() (*string, bool) {
	if o == nil || o.StateReason == nil {
		return nil, false
	}
	return o.StateReason, true
}

// HasStateReason returns a boolean if a field has been set.
func (o *BackendVmHealth) HasStateReason() bool {
	if o != nil && o.StateReason != nil {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given string and assigns it to the StateReason field.
func (o *BackendVmHealth) SetStateReason(v string) {
	o.StateReason = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *BackendVmHealth) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendVmHealth) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *BackendVmHealth) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *BackendVmHealth) SetVmId(v string) {
	o.VmId = &v
}

func (o BackendVmHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.StateReason != nil {
		toSerialize["StateReason"] = o.StateReason
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableBackendVmHealth struct {
	value *BackendVmHealth
	isSet bool
}

func (v NullableBackendVmHealth) Get() *BackendVmHealth {
	return v.value
}

func (v *NullableBackendVmHealth) Set(val *BackendVmHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableBackendVmHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableBackendVmHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackendVmHealth(val *BackendVmHealth) *NullableBackendVmHealth {
	return &NullableBackendVmHealth{value: val, isSet: true}
}

func (v NullableBackendVmHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackendVmHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


