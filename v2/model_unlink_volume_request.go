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

// UnlinkVolumeRequest struct for UnlinkVolumeRequest
type UnlinkVolumeRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// Forces the detachment of the volume in case of previous failure. Important: This action may damage your data or file systems.
	ForceUnlink *bool `json:"ForceUnlink,omitempty"`
	// The ID of the volume you want to detach.
	VolumeId string `json:"VolumeId"`
}

// NewUnlinkVolumeRequest instantiates a new UnlinkVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkVolumeRequest(volumeId string) *UnlinkVolumeRequest {
	this := UnlinkVolumeRequest{}
	this.VolumeId = volumeId
	return &this
}

// NewUnlinkVolumeRequestWithDefaults instantiates a new UnlinkVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkVolumeRequestWithDefaults() *UnlinkVolumeRequest {
	this := UnlinkVolumeRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkVolumeRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkVolumeRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkVolumeRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkVolumeRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetForceUnlink returns the ForceUnlink field value if set, zero value otherwise.
func (o *UnlinkVolumeRequest) GetForceUnlink() bool {
	if o == nil || o.ForceUnlink == nil {
		var ret bool
		return ret
	}
	return *o.ForceUnlink
}

// GetForceUnlinkOk returns a tuple with the ForceUnlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkVolumeRequest) GetForceUnlinkOk() (*bool, bool) {
	if o == nil || o.ForceUnlink == nil {
		return nil, false
	}
	return o.ForceUnlink, true
}

// HasForceUnlink returns a boolean if a field has been set.
func (o *UnlinkVolumeRequest) HasForceUnlink() bool {
	if o != nil && o.ForceUnlink != nil {
		return true
	}

	return false
}

// SetForceUnlink gets a reference to the given bool and assigns it to the ForceUnlink field.
func (o *UnlinkVolumeRequest) SetForceUnlink(v bool) {
	o.ForceUnlink = &v
}

// GetVolumeId returns the VolumeId field value
func (o *UnlinkVolumeRequest) GetVolumeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value
// and a boolean to check if the value has been set.
func (o *UnlinkVolumeRequest) GetVolumeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeId, true
}

// SetVolumeId sets field value
func (o *UnlinkVolumeRequest) SetVolumeId(v string) {
	o.VolumeId = v
}

func (o UnlinkVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.ForceUnlink != nil {
		toSerialize["ForceUnlink"] = o.ForceUnlink
	}
	if true {
		toSerialize["VolumeId"] = o.VolumeId
	}
	return json.Marshal(toSerialize)
}

type NullableUnlinkVolumeRequest struct {
	value *UnlinkVolumeRequest
	isSet bool
}

func (v NullableUnlinkVolumeRequest) Get() *UnlinkVolumeRequest {
	return v.value
}

func (v *NullableUnlinkVolumeRequest) Set(val *UnlinkVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkVolumeRequest(val *UnlinkVolumeRequest) *NullableUnlinkVolumeRequest {
	return &NullableUnlinkVolumeRequest{value: val, isSet: true}
}

func (v NullableUnlinkVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
