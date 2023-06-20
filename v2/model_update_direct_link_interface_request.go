/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.27
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateDirectLinkInterfaceRequest struct for UpdateDirectLinkInterfaceRequest
type UpdateDirectLinkInterfaceRequest struct {
	// The ID of the DirectLink interface you want to update.
	DirectLinkInterfaceId string `json:"DirectLinkInterfaceId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The maximum transmission unit (MTU) of the DirectLink interface, in bytes (always `1500`).
	Mtu int32 `json:"Mtu"`
}

// NewUpdateDirectLinkInterfaceRequest instantiates a new UpdateDirectLinkInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDirectLinkInterfaceRequest(directLinkInterfaceId string, mtu int32) *UpdateDirectLinkInterfaceRequest {
	this := UpdateDirectLinkInterfaceRequest{}
	this.DirectLinkInterfaceId = directLinkInterfaceId
	this.Mtu = mtu
	return &this
}

// NewUpdateDirectLinkInterfaceRequestWithDefaults instantiates a new UpdateDirectLinkInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDirectLinkInterfaceRequestWithDefaults() *UpdateDirectLinkInterfaceRequest {
	this := UpdateDirectLinkInterfaceRequest{}
	return &this
}

// GetDirectLinkInterfaceId returns the DirectLinkInterfaceId field value
func (o *UpdateDirectLinkInterfaceRequest) GetDirectLinkInterfaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectLinkInterfaceId
}

// GetDirectLinkInterfaceIdOk returns a tuple with the DirectLinkInterfaceId field value
// and a boolean to check if the value has been set.
func (o *UpdateDirectLinkInterfaceRequest) GetDirectLinkInterfaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectLinkInterfaceId, true
}

// SetDirectLinkInterfaceId sets field value
func (o *UpdateDirectLinkInterfaceRequest) SetDirectLinkInterfaceId(v string) {
	o.DirectLinkInterfaceId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateDirectLinkInterfaceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDirectLinkInterfaceRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateDirectLinkInterfaceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateDirectLinkInterfaceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetMtu returns the Mtu field value
func (o *UpdateDirectLinkInterfaceRequest) GetMtu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value
// and a boolean to check if the value has been set.
func (o *UpdateDirectLinkInterfaceRequest) GetMtuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mtu, true
}

// SetMtu sets field value
func (o *UpdateDirectLinkInterfaceRequest) SetMtu(v int32) {
	o.Mtu = v
}

func (o UpdateDirectLinkInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DirectLinkInterfaceId"] = o.DirectLinkInterfaceId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Mtu"] = o.Mtu
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDirectLinkInterfaceRequest struct {
	value *UpdateDirectLinkInterfaceRequest
	isSet bool
}

func (v NullableUpdateDirectLinkInterfaceRequest) Get() *UpdateDirectLinkInterfaceRequest {
	return v.value
}

func (v *NullableUpdateDirectLinkInterfaceRequest) Set(val *UpdateDirectLinkInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDirectLinkInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDirectLinkInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDirectLinkInterfaceRequest(val *UpdateDirectLinkInterfaceRequest) *NullableUpdateDirectLinkInterfaceRequest {
	return &NullableUpdateDirectLinkInterfaceRequest{value: val, isSet: true}
}

func (v NullableUpdateDirectLinkInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDirectLinkInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
