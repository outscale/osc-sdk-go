/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateVmsResponse struct for CreateVmsResponse
type CreateVmsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more created VMs.
	Vms *[]Vm `json:"Vms,omitempty"`
}

// NewCreateVmsResponse instantiates a new CreateVmsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVmsResponse() *CreateVmsResponse {
	this := CreateVmsResponse{}
	return &this
}

// NewCreateVmsResponseWithDefaults instantiates a new CreateVmsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVmsResponseWithDefaults() *CreateVmsResponse {
	this := CreateVmsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateVmsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateVmsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateVmsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVms returns the Vms field value if set, zero value otherwise.
func (o *CreateVmsResponse) GetVms() []Vm {
	if o == nil || o.Vms == nil {
		var ret []Vm
		return ret
	}
	return *o.Vms
}

// GetVmsOk returns a tuple with the Vms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsResponse) GetVmsOk() (*[]Vm, bool) {
	if o == nil || o.Vms == nil {
		return nil, false
	}
	return o.Vms, true
}

// HasVms returns a boolean if a field has been set.
func (o *CreateVmsResponse) HasVms() bool {
	if o != nil && o.Vms != nil {
		return true
	}

	return false
}

// SetVms gets a reference to the given []Vm and assigns it to the Vms field.
func (o *CreateVmsResponse) SetVms(v []Vm) {
	o.Vms = &v
}

func (o CreateVmsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Vms != nil {
		toSerialize["Vms"] = o.Vms
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVmsResponse struct {
	value *CreateVmsResponse
	isSet bool
}

func (v NullableCreateVmsResponse) Get() *CreateVmsResponse {
	return v.value
}

func (v *NullableCreateVmsResponse) Set(val *CreateVmsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVmsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVmsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVmsResponse(val *CreateVmsResponse) *NullableCreateVmsResponse {
	return &NullableCreateVmsResponse{value: val, isSet: true}
}

func (v NullableCreateVmsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVmsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
