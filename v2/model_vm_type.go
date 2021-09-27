/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// VmType Information about the VM type.
type VmType struct {
	// Indicates whether the VM is optimized for BSU I/O.
	BsuOptimized *bool `json:"BsuOptimized,omitempty"`
	// The maximum number of private IP addresses per network interface card (NIC).
	MaxPrivateIps *int32 `json:"MaxPrivateIps,omitempty"`
	// The amount of memory, in gibibytes.
	MemorySize *float32 `json:"MemorySize,omitempty"`
	// The number of vCores.
	VcoreCount *int32 `json:"VcoreCount,omitempty"`
	// The name of the VM type.
	VmTypeName *string `json:"VmTypeName,omitempty"`
	// The maximum number of ephemeral storage disks.
	VolumeCount *int32 `json:"VolumeCount,omitempty"`
	// The size of one ephemeral storage disk, in gibibytes (GiB).
	VolumeSize *int32 `json:"VolumeSize,omitempty"`
}

// NewVmType instantiates a new VmType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmType() *VmType {
	this := VmType{}
	return &this
}

// NewVmTypeWithDefaults instantiates a new VmType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmTypeWithDefaults() *VmType {
	this := VmType{}
	return &this
}

// GetBsuOptimized returns the BsuOptimized field value if set, zero value otherwise.
func (o *VmType) GetBsuOptimized() bool {
	if o == nil || o.BsuOptimized == nil {
		var ret bool
		return ret
	}
	return *o.BsuOptimized
}

// GetBsuOptimizedOk returns a tuple with the BsuOptimized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmType) GetBsuOptimizedOk() (*bool, bool) {
	if o == nil || o.BsuOptimized == nil {
		return nil, false
	}
	return o.BsuOptimized, true
}

// HasBsuOptimized returns a boolean if a field has been set.
func (o *VmType) HasBsuOptimized() bool {
	if o != nil && o.BsuOptimized != nil {
		return true
	}

	return false
}

// SetBsuOptimized gets a reference to the given bool and assigns it to the BsuOptimized field.
func (o *VmType) SetBsuOptimized(v bool) {
	o.BsuOptimized = &v
}

// GetMaxPrivateIps returns the MaxPrivateIps field value if set, zero value otherwise.
func (o *VmType) GetMaxPrivateIps() int32 {
	if o == nil || o.MaxPrivateIps == nil {
		var ret int32
		return ret
	}
	return *o.MaxPrivateIps
}

// GetMaxPrivateIpsOk returns a tuple with the MaxPrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmType) GetMaxPrivateIpsOk() (*int32, bool) {
	if o == nil || o.MaxPrivateIps == nil {
		return nil, false
	}
	return o.MaxPrivateIps, true
}

// HasMaxPrivateIps returns a boolean if a field has been set.
func (o *VmType) HasMaxPrivateIps() bool {
	if o != nil && o.MaxPrivateIps != nil {
		return true
	}

	return false
}

// SetMaxPrivateIps gets a reference to the given int32 and assigns it to the MaxPrivateIps field.
func (o *VmType) SetMaxPrivateIps(v int32) {
	o.MaxPrivateIps = &v
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise.
func (o *VmType) GetMemorySize() float32 {
	if o == nil || o.MemorySize == nil {
		var ret float32
		return ret
	}
	return *o.MemorySize
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmType) GetMemorySizeOk() (*float32, bool) {
	if o == nil || o.MemorySize == nil {
		return nil, false
	}
	return o.MemorySize, true
}

// HasMemorySize returns a boolean if a field has been set.
func (o *VmType) HasMemorySize() bool {
	if o != nil && o.MemorySize != nil {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given float32 and assigns it to the MemorySize field.
func (o *VmType) SetMemorySize(v float32) {
	o.MemorySize = &v
}

// GetVcoreCount returns the VcoreCount field value if set, zero value otherwise.
func (o *VmType) GetVcoreCount() int32 {
	if o == nil || o.VcoreCount == nil {
		var ret int32
		return ret
	}
	return *o.VcoreCount
}

// GetVcoreCountOk returns a tuple with the VcoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmType) GetVcoreCountOk() (*int32, bool) {
	if o == nil || o.VcoreCount == nil {
		return nil, false
	}
	return o.VcoreCount, true
}

// HasVcoreCount returns a boolean if a field has been set.
func (o *VmType) HasVcoreCount() bool {
	if o != nil && o.VcoreCount != nil {
		return true
	}

	return false
}

// SetVcoreCount gets a reference to the given int32 and assigns it to the VcoreCount field.
func (o *VmType) SetVcoreCount(v int32) {
	o.VcoreCount = &v
}

// GetVmTypeName returns the VmTypeName field value if set, zero value otherwise.
func (o *VmType) GetVmTypeName() string {
	if o == nil || o.VmTypeName == nil {
		var ret string
		return ret
	}
	return *o.VmTypeName
}

// GetVmTypeNameOk returns a tuple with the VmTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmType) GetVmTypeNameOk() (*string, bool) {
	if o == nil || o.VmTypeName == nil {
		return nil, false
	}
	return o.VmTypeName, true
}

// HasVmTypeName returns a boolean if a field has been set.
func (o *VmType) HasVmTypeName() bool {
	if o != nil && o.VmTypeName != nil {
		return true
	}

	return false
}

// SetVmTypeName gets a reference to the given string and assigns it to the VmTypeName field.
func (o *VmType) SetVmTypeName(v string) {
	o.VmTypeName = &v
}

// GetVolumeCount returns the VolumeCount field value if set, zero value otherwise.
func (o *VmType) GetVolumeCount() int32 {
	if o == nil || o.VolumeCount == nil {
		var ret int32
		return ret
	}
	return *o.VolumeCount
}

// GetVolumeCountOk returns a tuple with the VolumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmType) GetVolumeCountOk() (*int32, bool) {
	if o == nil || o.VolumeCount == nil {
		return nil, false
	}
	return o.VolumeCount, true
}

// HasVolumeCount returns a boolean if a field has been set.
func (o *VmType) HasVolumeCount() bool {
	if o != nil && o.VolumeCount != nil {
		return true
	}

	return false
}

// SetVolumeCount gets a reference to the given int32 and assigns it to the VolumeCount field.
func (o *VmType) SetVolumeCount(v int32) {
	o.VolumeCount = &v
}

// GetVolumeSize returns the VolumeSize field value if set, zero value otherwise.
func (o *VmType) GetVolumeSize() int32 {
	if o == nil || o.VolumeSize == nil {
		var ret int32
		return ret
	}
	return *o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmType) GetVolumeSizeOk() (*int32, bool) {
	if o == nil || o.VolumeSize == nil {
		return nil, false
	}
	return o.VolumeSize, true
}

// HasVolumeSize returns a boolean if a field has been set.
func (o *VmType) HasVolumeSize() bool {
	if o != nil && o.VolumeSize != nil {
		return true
	}

	return false
}

// SetVolumeSize gets a reference to the given int32 and assigns it to the VolumeSize field.
func (o *VmType) SetVolumeSize(v int32) {
	o.VolumeSize = &v
}

func (o VmType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BsuOptimized != nil {
		toSerialize["BsuOptimized"] = o.BsuOptimized
	}
	if o.MaxPrivateIps != nil {
		toSerialize["MaxPrivateIps"] = o.MaxPrivateIps
	}
	if o.MemorySize != nil {
		toSerialize["MemorySize"] = o.MemorySize
	}
	if o.VcoreCount != nil {
		toSerialize["VcoreCount"] = o.VcoreCount
	}
	if o.VmTypeName != nil {
		toSerialize["VmTypeName"] = o.VmTypeName
	}
	if o.VolumeCount != nil {
		toSerialize["VolumeCount"] = o.VolumeCount
	}
	if o.VolumeSize != nil {
		toSerialize["VolumeSize"] = o.VolumeSize
	}
	return json.Marshal(toSerialize)
}

type NullableVmType struct {
	value *VmType
	isSet bool
}

func (v NullableVmType) Get() *VmType {
	return v.value
}

func (v *NullableVmType) Set(val *VmType) {
	v.value = val
	v.isSet = true
}

func (v NullableVmType) IsSet() bool {
	return v.isSet
}

func (v *NullableVmType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmType(val *VmType) *NullableVmType {
	return &NullableVmType{value: val, isSet: true}
}

func (v NullableVmType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
