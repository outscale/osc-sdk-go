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

// BsuCreated Information about the created BSU volume.
type BsuCreated struct {
	// If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The time and date of attachment of the volume to the VM.
	LinkDate *string `json:"LinkDate,omitempty"`
	// The state of the volume.
	State *string `json:"State,omitempty"`
	// The ID of the volume.
	VolumeId *string `json:"VolumeId,omitempty"`
}

// NewBsuCreated instantiates a new BsuCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBsuCreated() *BsuCreated {
	this := BsuCreated{}
	return &this
}

// NewBsuCreatedWithDefaults instantiates a new BsuCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBsuCreatedWithDefaults() *BsuCreated {
	this := BsuCreated{}
	return &this
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *BsuCreated) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuCreated) GetDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *BsuCreated) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *BsuCreated) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetLinkDate returns the LinkDate field value if set, zero value otherwise.
func (o *BsuCreated) GetLinkDate() string {
	if o == nil || o.LinkDate == nil {
		var ret string
		return ret
	}
	return *o.LinkDate
}

// GetLinkDateOk returns a tuple with the LinkDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuCreated) GetLinkDateOk() (*string, bool) {
	if o == nil || o.LinkDate == nil {
		return nil, false
	}
	return o.LinkDate, true
}

// HasLinkDate returns a boolean if a field has been set.
func (o *BsuCreated) HasLinkDate() bool {
	if o != nil && o.LinkDate != nil {
		return true
	}

	return false
}

// SetLinkDate gets a reference to the given string and assigns it to the LinkDate field.
func (o *BsuCreated) SetLinkDate(v string) {
	o.LinkDate = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BsuCreated) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuCreated) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BsuCreated) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BsuCreated) SetState(v string) {
	o.State = &v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *BsuCreated) GetVolumeId() string {
	if o == nil || o.VolumeId == nil {
		var ret string
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuCreated) GetVolumeIdOk() (*string, bool) {
	if o == nil || o.VolumeId == nil {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *BsuCreated) HasVolumeId() bool {
	if o != nil && o.VolumeId != nil {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *BsuCreated) SetVolumeId(v string) {
	o.VolumeId = &v
}

func (o BsuCreated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteOnVmDeletion != nil {
		toSerialize["DeleteOnVmDeletion"] = o.DeleteOnVmDeletion
	}
	if o.LinkDate != nil {
		toSerialize["LinkDate"] = o.LinkDate
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.VolumeId != nil {
		toSerialize["VolumeId"] = o.VolumeId
	}
	return json.Marshal(toSerialize)
}

type NullableBsuCreated struct {
	value *BsuCreated
	isSet bool
}

func (v NullableBsuCreated) Get() *BsuCreated {
	return v.value
}

func (v *NullableBsuCreated) Set(val *BsuCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableBsuCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableBsuCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsuCreated(val *BsuCreated) *NullableBsuCreated {
	return &NullableBsuCreated{value: val, isSet: true}
}

func (v NullableBsuCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsuCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
