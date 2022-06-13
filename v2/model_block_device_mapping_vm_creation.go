/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// BlockDeviceMappingVmCreation Information about the block device mapping.
type BlockDeviceMappingVmCreation struct {
	Bsu *BsuToCreate `json:"Bsu,omitempty"`
	// The device name for the volume. For a root device, you must use `/dev/sda1`. For other volumes, you must use `/dev/sdX` or `/dev/xvdX` (where `X` is a letter between `b` and `z`).
	DeviceName *string `json:"DeviceName,omitempty"`
	// Removes the device which is included in the block device mapping of the OMI.
	NoDevice *string `json:"NoDevice,omitempty"`
	// The name of the virtual device (`ephemeralN`).
	VirtualDeviceName *string `json:"VirtualDeviceName,omitempty"`
}

// NewBlockDeviceMappingVmCreation instantiates a new BlockDeviceMappingVmCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockDeviceMappingVmCreation() *BlockDeviceMappingVmCreation {
	this := BlockDeviceMappingVmCreation{}
	return &this
}

// NewBlockDeviceMappingVmCreationWithDefaults instantiates a new BlockDeviceMappingVmCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockDeviceMappingVmCreationWithDefaults() *BlockDeviceMappingVmCreation {
	this := BlockDeviceMappingVmCreation{}
	return &this
}

// GetBsu returns the Bsu field value if set, zero value otherwise.
func (o *BlockDeviceMappingVmCreation) GetBsu() BsuToCreate {
	if o == nil || o.Bsu == nil {
		var ret BsuToCreate
		return ret
	}
	return *o.Bsu
}

// GetBsuOk returns a tuple with the Bsu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingVmCreation) GetBsuOk() (*BsuToCreate, bool) {
	if o == nil || o.Bsu == nil {
		return nil, false
	}
	return o.Bsu, true
}

// HasBsu returns a boolean if a field has been set.
func (o *BlockDeviceMappingVmCreation) HasBsu() bool {
	if o != nil && o.Bsu != nil {
		return true
	}

	return false
}

// SetBsu gets a reference to the given BsuToCreate and assigns it to the Bsu field.
func (o *BlockDeviceMappingVmCreation) SetBsu(v BsuToCreate) {
	o.Bsu = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingVmCreation) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingVmCreation) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingVmCreation) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *BlockDeviceMappingVmCreation) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetNoDevice returns the NoDevice field value if set, zero value otherwise.
func (o *BlockDeviceMappingVmCreation) GetNoDevice() string {
	if o == nil || o.NoDevice == nil {
		var ret string
		return ret
	}
	return *o.NoDevice
}

// GetNoDeviceOk returns a tuple with the NoDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingVmCreation) GetNoDeviceOk() (*string, bool) {
	if o == nil || o.NoDevice == nil {
		return nil, false
	}
	return o.NoDevice, true
}

// HasNoDevice returns a boolean if a field has been set.
func (o *BlockDeviceMappingVmCreation) HasNoDevice() bool {
	if o != nil && o.NoDevice != nil {
		return true
	}

	return false
}

// SetNoDevice gets a reference to the given string and assigns it to the NoDevice field.
func (o *BlockDeviceMappingVmCreation) SetNoDevice(v string) {
	o.NoDevice = &v
}

// GetVirtualDeviceName returns the VirtualDeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingVmCreation) GetVirtualDeviceName() string {
	if o == nil || o.VirtualDeviceName == nil {
		var ret string
		return ret
	}
	return *o.VirtualDeviceName
}

// GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingVmCreation) GetVirtualDeviceNameOk() (*string, bool) {
	if o == nil || o.VirtualDeviceName == nil {
		return nil, false
	}
	return o.VirtualDeviceName, true
}

// HasVirtualDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingVmCreation) HasVirtualDeviceName() bool {
	if o != nil && o.VirtualDeviceName != nil {
		return true
	}

	return false
}

// SetVirtualDeviceName gets a reference to the given string and assigns it to the VirtualDeviceName field.
func (o *BlockDeviceMappingVmCreation) SetVirtualDeviceName(v string) {
	o.VirtualDeviceName = &v
}

func (o BlockDeviceMappingVmCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bsu != nil {
		toSerialize["Bsu"] = o.Bsu
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.NoDevice != nil {
		toSerialize["NoDevice"] = o.NoDevice
	}
	if o.VirtualDeviceName != nil {
		toSerialize["VirtualDeviceName"] = o.VirtualDeviceName
	}
	return json.Marshal(toSerialize)
}

type NullableBlockDeviceMappingVmCreation struct {
	value *BlockDeviceMappingVmCreation
	isSet bool
}

func (v NullableBlockDeviceMappingVmCreation) Get() *BlockDeviceMappingVmCreation {
	return v.value
}

func (v *NullableBlockDeviceMappingVmCreation) Set(val *BlockDeviceMappingVmCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockDeviceMappingVmCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockDeviceMappingVmCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockDeviceMappingVmCreation(val *BlockDeviceMappingVmCreation) *NullableBlockDeviceMappingVmCreation {
	return &NullableBlockDeviceMappingVmCreation{value: val, isSet: true}
}

func (v NullableBlockDeviceMappingVmCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockDeviceMappingVmCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
