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

// LinkedVolume Information about volume attachment.
type LinkedVolume struct {
	// If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The name of the device.
	DeviceName *string `json:"DeviceName,omitempty"`
	// The state of the attachment of the volume (`attaching` \\| `detaching` \\| `attached` \\| `detached`).
	State *string `json:"State,omitempty"`
	// The ID of the VM.
	VmId *string `json:"VmId,omitempty"`
	// The ID of the volume.
	VolumeId *string `json:"VolumeId,omitempty"`
}

// NewLinkedVolume instantiates a new LinkedVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedVolume() *LinkedVolume {
	this := LinkedVolume{}
	return &this
}

// NewLinkedVolumeWithDefaults instantiates a new LinkedVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedVolumeWithDefaults() *LinkedVolume {
	this := LinkedVolume{}
	return &this
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *LinkedVolume) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedVolume) GetDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *LinkedVolume) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *LinkedVolume) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *LinkedVolume) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedVolume) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *LinkedVolume) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *LinkedVolume) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LinkedVolume) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedVolume) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LinkedVolume) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LinkedVolume) SetState(v string) {
	o.State = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *LinkedVolume) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedVolume) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *LinkedVolume) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *LinkedVolume) SetVmId(v string) {
	o.VmId = &v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *LinkedVolume) GetVolumeId() string {
	if o == nil || o.VolumeId == nil {
		var ret string
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedVolume) GetVolumeIdOk() (*string, bool) {
	if o == nil || o.VolumeId == nil {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *LinkedVolume) HasVolumeId() bool {
	if o != nil && o.VolumeId != nil {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *LinkedVolume) SetVolumeId(v string) {
	o.VolumeId = &v
}

func (o LinkedVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteOnVmDeletion != nil {
		toSerialize["DeleteOnVmDeletion"] = o.DeleteOnVmDeletion
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	if o.VolumeId != nil {
		toSerialize["VolumeId"] = o.VolumeId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkedVolume struct {
	value *LinkedVolume
	isSet bool
}

func (v NullableLinkedVolume) Get() *LinkedVolume {
	return v.value
}

func (v *NullableLinkedVolume) Set(val *LinkedVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedVolume(val *LinkedVolume) *NullableLinkedVolume {
	return &NullableLinkedVolume{value: val, isSet: true}
}

func (v NullableLinkedVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
