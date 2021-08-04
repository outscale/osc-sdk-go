/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateVmResponse struct for UpdateVmResponse
type UpdateVmResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	Vm              *Vm              `json:"Vm,omitempty"`
}

// NewUpdateVmResponse instantiates a new UpdateVmResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVmResponse() *UpdateVmResponse {
	this := UpdateVmResponse{}
	return &this
}

// NewUpdateVmResponseWithDefaults instantiates a new UpdateVmResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVmResponseWithDefaults() *UpdateVmResponse {
	this := UpdateVmResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateVmResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateVmResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateVmResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVm returns the Vm field value if set, zero value otherwise.
func (o *UpdateVmResponse) GetVm() Vm {
	if o == nil || o.Vm == nil {
		var ret Vm
		return ret
	}
	return *o.Vm
}

// GetVmOk returns a tuple with the Vm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmResponse) GetVmOk() (*Vm, bool) {
	if o == nil || o.Vm == nil {
		return nil, false
	}
	return o.Vm, true
}

// HasVm returns a boolean if a field has been set.
func (o *UpdateVmResponse) HasVm() bool {
	if o != nil && o.Vm != nil {
		return true
	}

	return false
}

// SetVm gets a reference to the given Vm and assigns it to the Vm field.
func (o *UpdateVmResponse) SetVm(v Vm) {
	o.Vm = &v
}

func (o UpdateVmResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Vm != nil {
		toSerialize["Vm"] = o.Vm
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateVmResponse struct {
	value *UpdateVmResponse
	isSet bool
}

func (v NullableUpdateVmResponse) Get() *UpdateVmResponse {
	return v.value
}

func (v *NullableUpdateVmResponse) Set(val *UpdateVmResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVmResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVmResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVmResponse(val *UpdateVmResponse) *NullableUpdateVmResponse {
	return &NullableUpdateVmResponse{value: val, isSet: true}
}

func (v NullableUpdateVmResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVmResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}