/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UnlinkInternetServiceRequest struct for UnlinkInternetServiceRequest
type UnlinkInternetServiceRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Internet service you want to detach.
	InternetServiceId string `json:"InternetServiceId"`
	// The ID of the Net from which you want to detach the Internet service.
	NetId string `json:"NetId"`
}

// NewUnlinkInternetServiceRequest instantiates a new UnlinkInternetServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkInternetServiceRequest(internetServiceId string, netId string) *UnlinkInternetServiceRequest {
	this := UnlinkInternetServiceRequest{}
	this.InternetServiceId = internetServiceId
	this.NetId = netId
	return &this
}

// NewUnlinkInternetServiceRequestWithDefaults instantiates a new UnlinkInternetServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkInternetServiceRequestWithDefaults() *UnlinkInternetServiceRequest {
	this := UnlinkInternetServiceRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkInternetServiceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkInternetServiceRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkInternetServiceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkInternetServiceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetInternetServiceId returns the InternetServiceId field value
func (o *UnlinkInternetServiceRequest) GetInternetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InternetServiceId
}

// GetInternetServiceIdOk returns a tuple with the InternetServiceId field value
// and a boolean to check if the value has been set.
func (o *UnlinkInternetServiceRequest) GetInternetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternetServiceId, true
}

// SetInternetServiceId sets field value
func (o *UnlinkInternetServiceRequest) SetInternetServiceId(v string) {
	o.InternetServiceId = v
}

// GetNetId returns the NetId field value
func (o *UnlinkInternetServiceRequest) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *UnlinkInternetServiceRequest) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *UnlinkInternetServiceRequest) SetNetId(v string) {
	o.NetId = v
}

func (o UnlinkInternetServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["InternetServiceId"] = o.InternetServiceId
	}
	if true {
		toSerialize["NetId"] = o.NetId
	}
	return json.Marshal(toSerialize)
}

type NullableUnlinkInternetServiceRequest struct {
	value *UnlinkInternetServiceRequest
	isSet bool
}

func (v NullableUnlinkInternetServiceRequest) Get() *UnlinkInternetServiceRequest {
	return v.value
}

func (v *NullableUnlinkInternetServiceRequest) Set(val *UnlinkInternetServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkInternetServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkInternetServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkInternetServiceRequest(val *UnlinkInternetServiceRequest) *NullableUnlinkInternetServiceRequest {
	return &NullableUnlinkInternetServiceRequest{value: val, isSet: true}
}

func (v NullableUnlinkInternetServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkInternetServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
