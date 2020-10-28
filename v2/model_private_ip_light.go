/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// PrivateIpLight Information about the private IP.
type PrivateIpLight struct {
	// If `true`, the IP address is the primary private IP address of the NIC.
	IsPrimary *bool `json:"IsPrimary,omitempty"`
	// The private IP address of the NIC.
	PrivateIp *string `json:"PrivateIp,omitempty"`
}

// NewPrivateIpLight instantiates a new PrivateIpLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateIpLight() *PrivateIpLight {
	this := PrivateIpLight{}
	return &this
}

// NewPrivateIpLightWithDefaults instantiates a new PrivateIpLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateIpLightWithDefaults() *PrivateIpLight {
	this := PrivateIpLight{}
	return &this
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *PrivateIpLight) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIpLight) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *PrivateIpLight) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *PrivateIpLight) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *PrivateIpLight) GetPrivateIp() string {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIpLight) GetPrivateIpOk() (*string, bool) {
	if o == nil || o.PrivateIp == nil {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *PrivateIpLight) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *PrivateIpLight) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

func (o PrivateIpLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsPrimary != nil {
		toSerialize["IsPrimary"] = o.IsPrimary
	}
	if o.PrivateIp != nil {
		toSerialize["PrivateIp"] = o.PrivateIp
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateIpLight struct {
	value *PrivateIpLight
	isSet bool
}

func (v NullablePrivateIpLight) Get() *PrivateIpLight {
	return v.value
}

func (v *NullablePrivateIpLight) Set(val *PrivateIpLight) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateIpLight) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateIpLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateIpLight(val *PrivateIpLight) *NullablePrivateIpLight {
	return &NullablePrivateIpLight{value: val, isSet: true}
}

func (v NullablePrivateIpLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateIpLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


