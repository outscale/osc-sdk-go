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

// UnlinkPrivateIpsRequest struct for UnlinkPrivateIpsRequest
type UnlinkPrivateIpsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId"`
	// One or more secondary private IPs you want to unassign from the NIC.
	PrivateIps []string `json:"PrivateIps"`
}

// NewUnlinkPrivateIpsRequest instantiates a new UnlinkPrivateIpsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkPrivateIpsRequest(nicId string, privateIps []string) *UnlinkPrivateIpsRequest {
	this := UnlinkPrivateIpsRequest{}
	this.NicId = nicId
	this.PrivateIps = privateIps
	return &this
}

// NewUnlinkPrivateIpsRequestWithDefaults instantiates a new UnlinkPrivateIpsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkPrivateIpsRequestWithDefaults() *UnlinkPrivateIpsRequest {
	this := UnlinkPrivateIpsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkPrivateIpsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkPrivateIpsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkPrivateIpsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkPrivateIpsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNicId returns the NicId field value
func (o *UnlinkPrivateIpsRequest) GetNicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value
// and a boolean to check if the value has been set.
func (o *UnlinkPrivateIpsRequest) GetNicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NicId, true
}

// SetNicId sets field value
func (o *UnlinkPrivateIpsRequest) SetNicId(v string) {
	o.NicId = v
}

// GetPrivateIps returns the PrivateIps field value
func (o *UnlinkPrivateIpsRequest) GetPrivateIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value
// and a boolean to check if the value has been set.
func (o *UnlinkPrivateIpsRequest) GetPrivateIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateIps, true
}

// SetPrivateIps sets field value
func (o *UnlinkPrivateIpsRequest) SetPrivateIps(v []string) {
	o.PrivateIps = v
}

func (o UnlinkPrivateIpsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NicId"] = o.NicId
	}
	if true {
		toSerialize["PrivateIps"] = o.PrivateIps
	}
	return json.Marshal(toSerialize)
}

type NullableUnlinkPrivateIpsRequest struct {
	value *UnlinkPrivateIpsRequest
	isSet bool
}

func (v NullableUnlinkPrivateIpsRequest) Get() *UnlinkPrivateIpsRequest {
	return v.value
}

func (v *NullableUnlinkPrivateIpsRequest) Set(val *UnlinkPrivateIpsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkPrivateIpsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkPrivateIpsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkPrivateIpsRequest(val *UnlinkPrivateIpsRequest) *NullableUnlinkPrivateIpsRequest {
	return &NullableUnlinkPrivateIpsRequest{value: val, isSet: true}
}

func (v NullableUnlinkPrivateIpsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkPrivateIpsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
