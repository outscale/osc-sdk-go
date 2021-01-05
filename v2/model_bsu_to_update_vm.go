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

// BsuToUpdateVm Information about the BSU volume.
type BsuToUpdateVm struct {
	// If `true`, the volume is deleted when the VM is terminated.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The ID of the volume.
	VolumeId *string `json:"VolumeId,omitempty"`
}

// NewBsuToUpdateVm instantiates a new BsuToUpdateVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBsuToUpdateVm() *BsuToUpdateVm {
	this := BsuToUpdateVm{}
	return &this
}

// NewBsuToUpdateVmWithDefaults instantiates a new BsuToUpdateVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBsuToUpdateVmWithDefaults() *BsuToUpdateVm {
	this := BsuToUpdateVm{}
	return &this
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *BsuToUpdateVm) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuToUpdateVm) GetDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *BsuToUpdateVm) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *BsuToUpdateVm) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *BsuToUpdateVm) GetVolumeId() string {
	if o == nil || o.VolumeId == nil {
		var ret string
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuToUpdateVm) GetVolumeIdOk() (*string, bool) {
	if o == nil || o.VolumeId == nil {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *BsuToUpdateVm) HasVolumeId() bool {
	if o != nil && o.VolumeId != nil {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *BsuToUpdateVm) SetVolumeId(v string) {
	o.VolumeId = &v
}

func (o BsuToUpdateVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteOnVmDeletion != nil {
		toSerialize["DeleteOnVmDeletion"] = o.DeleteOnVmDeletion
	}
	if o.VolumeId != nil {
		toSerialize["VolumeId"] = o.VolumeId
	}
	return json.Marshal(toSerialize)
}

type NullableBsuToUpdateVm struct {
	value *BsuToUpdateVm
	isSet bool
}

func (v NullableBsuToUpdateVm) Get() *BsuToUpdateVm {
	return v.value
}

func (v *NullableBsuToUpdateVm) Set(val *BsuToUpdateVm) {
	v.value = val
	v.isSet = true
}

func (v NullableBsuToUpdateVm) IsSet() bool {
	return v.isSet
}

func (v *NullableBsuToUpdateVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsuToUpdateVm(val *BsuToUpdateVm) *NullableBsuToUpdateVm {
	return &NullableBsuToUpdateVm{value: val, isSet: true}
}

func (v NullableBsuToUpdateVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsuToUpdateVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


