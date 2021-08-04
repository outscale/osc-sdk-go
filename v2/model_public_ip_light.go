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

// PublicIpLight Information about the public IP.
type PublicIpLight struct {
	// The External IP address (EIP) associated with the NAT service.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The allocation ID of the EIP associated with the NAT service.
	PublicIpId *string `json:"PublicIpId,omitempty"`
}

// NewPublicIpLight instantiates a new PublicIpLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIpLight() *PublicIpLight {
	this := PublicIpLight{}
	return &this
}

// NewPublicIpLightWithDefaults instantiates a new PublicIpLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIpLightWithDefaults() *PublicIpLight {
	this := PublicIpLight{}
	return &this
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *PublicIpLight) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIpLight) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *PublicIpLight) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *PublicIpLight) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicIpId returns the PublicIpId field value if set, zero value otherwise.
func (o *PublicIpLight) GetPublicIpId() string {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpId
}

// GetPublicIpIdOk returns a tuple with the PublicIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIpLight) GetPublicIpIdOk() (*string, bool) {
	if o == nil || o.PublicIpId == nil {
		return nil, false
	}
	return o.PublicIpId, true
}

// HasPublicIpId returns a boolean if a field has been set.
func (o *PublicIpLight) HasPublicIpId() bool {
	if o != nil && o.PublicIpId != nil {
		return true
	}

	return false
}

// SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.
func (o *PublicIpLight) SetPublicIpId(v string) {
	o.PublicIpId = &v
}

func (o PublicIpLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.PublicIpId != nil {
		toSerialize["PublicIpId"] = o.PublicIpId
	}
	return json.Marshal(toSerialize)
}

type NullablePublicIpLight struct {
	value *PublicIpLight
	isSet bool
}

func (v NullablePublicIpLight) Get() *PublicIpLight {
	return v.value
}

func (v *NullablePublicIpLight) Set(val *PublicIpLight) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIpLight) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIpLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIpLight(val *PublicIpLight) *NullablePublicIpLight {
	return &NullablePublicIpLight{value: val, isSet: true}
}

func (v NullablePublicIpLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIpLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}