/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteNetRequest struct for DeleteNetRequest
type DeleteNetRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net you want to delete.
	NetId string `json:"NetId"`
}

// NewDeleteNetRequest instantiates a new DeleteNetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteNetRequest(netId string) *DeleteNetRequest {
	this := DeleteNetRequest{}
	this.NetId = netId
	return &this
}

// NewDeleteNetRequestWithDefaults instantiates a new DeleteNetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteNetRequestWithDefaults() *DeleteNetRequest {
	this := DeleteNetRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteNetRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteNetRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteNetRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteNetRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetId returns the NetId field value
func (o *DeleteNetRequest) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *DeleteNetRequest) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *DeleteNetRequest) SetNetId(v string) {
	o.NetId = v
}

func (o DeleteNetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NetId"] = o.NetId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteNetRequest struct {
	value *DeleteNetRequest
	isSet bool
}

func (v NullableDeleteNetRequest) Get() *DeleteNetRequest {
	return v.value
}

func (v *NullableDeleteNetRequest) Set(val *DeleteNetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteNetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteNetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteNetRequest(val *DeleteNetRequest) *NullableDeleteNetRequest {
	return &NullableDeleteNetRequest{value: val, isSet: true}
}

func (v NullableDeleteNetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteNetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
