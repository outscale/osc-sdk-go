/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// BlockDeviceMappingCreated Information about the created block device mapping.
type BlockDeviceMappingCreated struct {
	Bsu *BsuCreated `json:"Bsu,omitempty"`
	// The name of the device.
	DeviceName *string `json:"DeviceName,omitempty"`
}

// NewBlockDeviceMappingCreated instantiates a new BlockDeviceMappingCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockDeviceMappingCreated() *BlockDeviceMappingCreated {
	this := BlockDeviceMappingCreated{}
	return &this
}

// NewBlockDeviceMappingCreatedWithDefaults instantiates a new BlockDeviceMappingCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockDeviceMappingCreatedWithDefaults() *BlockDeviceMappingCreated {
	this := BlockDeviceMappingCreated{}
	return &this
}

// GetBsu returns the Bsu field value if set, zero value otherwise.
func (o *BlockDeviceMappingCreated) GetBsu() BsuCreated {
	if o == nil || o.Bsu == nil {
		var ret BsuCreated
		return ret
	}
	return *o.Bsu
}

// GetBsuOk returns a tuple with the Bsu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingCreated) GetBsuOk() (*BsuCreated, bool) {
	if o == nil || o.Bsu == nil {
		return nil, false
	}
	return o.Bsu, true
}

// HasBsu returns a boolean if a field has been set.
func (o *BlockDeviceMappingCreated) HasBsu() bool {
	if o != nil && o.Bsu != nil {
		return true
	}

	return false
}

// SetBsu gets a reference to the given BsuCreated and assigns it to the Bsu field.
func (o *BlockDeviceMappingCreated) SetBsu(v BsuCreated) {
	o.Bsu = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingCreated) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingCreated) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingCreated) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *BlockDeviceMappingCreated) SetDeviceName(v string) {
	o.DeviceName = &v
}

func (o BlockDeviceMappingCreated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bsu != nil {
		toSerialize["Bsu"] = o.Bsu
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	return json.Marshal(toSerialize)
}

type NullableBlockDeviceMappingCreated struct {
	value *BlockDeviceMappingCreated
	isSet bool
}

func (v NullableBlockDeviceMappingCreated) Get() *BlockDeviceMappingCreated {
	return v.value
}

func (v *NullableBlockDeviceMappingCreated) Set(val *BlockDeviceMappingCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockDeviceMappingCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockDeviceMappingCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockDeviceMappingCreated(val *BlockDeviceMappingCreated) *NullableBlockDeviceMappingCreated {
	return &NullableBlockDeviceMappingCreated{value: val, isSet: true}
}

func (v NullableBlockDeviceMappingCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockDeviceMappingCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
