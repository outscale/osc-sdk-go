/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteVolumeRequest struct for DeleteVolumeRequest
type DeleteVolumeRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the volume you want to delete.
	VolumeId string `json:"VolumeId"`
}

// NewDeleteVolumeRequest instantiates a new DeleteVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVolumeRequest(volumeId string) *DeleteVolumeRequest {
	this := DeleteVolumeRequest{}
	this.VolumeId = volumeId
	return &this
}

// NewDeleteVolumeRequestWithDefaults instantiates a new DeleteVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVolumeRequestWithDefaults() *DeleteVolumeRequest {
	this := DeleteVolumeRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteVolumeRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVolumeRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteVolumeRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteVolumeRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetVolumeId returns the VolumeId field value
func (o *DeleteVolumeRequest) GetVolumeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value
// and a boolean to check if the value has been set.
func (o *DeleteVolumeRequest) GetVolumeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeId, true
}

// SetVolumeId sets field value
func (o *DeleteVolumeRequest) SetVolumeId(v string) {
	o.VolumeId = v
}

func (o DeleteVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["VolumeId"] = o.VolumeId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteVolumeRequest struct {
	value *DeleteVolumeRequest
	isSet bool
}

func (v NullableDeleteVolumeRequest) Get() *DeleteVolumeRequest {
	return v.value
}

func (v *NullableDeleteVolumeRequest) Set(val *DeleteVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVolumeRequest(val *DeleteVolumeRequest) *NullableDeleteVolumeRequest {
	return &NullableDeleteVolumeRequest{value: val, isSet: true}
}

func (v NullableDeleteVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
