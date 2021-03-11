/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateNetRequest struct for UpdateNetRequest
type UpdateNetRequest struct {
	// The ID of the DHCP options set (or `default` if you want to associate the default one).
	DhcpOptionsSetId string `json:"DhcpOptionsSetId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net.
	NetId string `json:"NetId"`
}

// NewUpdateNetRequest instantiates a new UpdateNetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetRequest(dhcpOptionsSetId string, netId string) *UpdateNetRequest {
	this := UpdateNetRequest{}
	this.DhcpOptionsSetId = dhcpOptionsSetId
	this.NetId = netId
	return &this
}

// NewUpdateNetRequestWithDefaults instantiates a new UpdateNetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetRequestWithDefaults() *UpdateNetRequest {
	this := UpdateNetRequest{}
	return &this
}

// GetDhcpOptionsSetId returns the DhcpOptionsSetId field value
func (o *UpdateNetRequest) GetDhcpOptionsSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DhcpOptionsSetId
}

// GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field value
// and a boolean to check if the value has been set.
func (o *UpdateNetRequest) GetDhcpOptionsSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DhcpOptionsSetId, true
}

// SetDhcpOptionsSetId sets field value
func (o *UpdateNetRequest) SetDhcpOptionsSetId(v string) {
	o.DhcpOptionsSetId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateNetRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateNetRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateNetRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetId returns the NetId field value
func (o *UpdateNetRequest) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *UpdateNetRequest) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *UpdateNetRequest) SetNetId(v string) {
	o.NetId = v
}

func (o UpdateNetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DhcpOptionsSetId"] = o.DhcpOptionsSetId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NetId"] = o.NetId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetRequest struct {
	value *UpdateNetRequest
	isSet bool
}

func (v NullableUpdateNetRequest) Get() *UpdateNetRequest {
	return v.value
}

func (v *NullableUpdateNetRequest) Set(val *UpdateNetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetRequest(val *UpdateNetRequest) *NullableUpdateNetRequest {
	return &NullableUpdateNetRequest{value: val, isSet: true}
}

func (v NullableUpdateNetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
