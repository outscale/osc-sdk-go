/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteNetAccessPointRequest struct for DeleteNetAccessPointRequest
type DeleteNetAccessPointRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net access point.
	NetAccessPointId string `json:"NetAccessPointId"`
}

// NewDeleteNetAccessPointRequest instantiates a new DeleteNetAccessPointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteNetAccessPointRequest(netAccessPointId string) *DeleteNetAccessPointRequest {
	this := DeleteNetAccessPointRequest{}
	this.NetAccessPointId = netAccessPointId
	return &this
}

// NewDeleteNetAccessPointRequestWithDefaults instantiates a new DeleteNetAccessPointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteNetAccessPointRequestWithDefaults() *DeleteNetAccessPointRequest {
	this := DeleteNetAccessPointRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteNetAccessPointRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteNetAccessPointRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteNetAccessPointRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteNetAccessPointRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetAccessPointId returns the NetAccessPointId field value
func (o *DeleteNetAccessPointRequest) GetNetAccessPointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetAccessPointId
}

// GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field value
// and a boolean to check if the value has been set.
func (o *DeleteNetAccessPointRequest) GetNetAccessPointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetAccessPointId, true
}

// SetNetAccessPointId sets field value
func (o *DeleteNetAccessPointRequest) SetNetAccessPointId(v string) {
	o.NetAccessPointId = v
}

func (o DeleteNetAccessPointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NetAccessPointId"] = o.NetAccessPointId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteNetAccessPointRequest struct {
	value *DeleteNetAccessPointRequest
	isSet bool
}

func (v NullableDeleteNetAccessPointRequest) Get() *DeleteNetAccessPointRequest {
	return v.value
}

func (v *NullableDeleteNetAccessPointRequest) Set(val *DeleteNetAccessPointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteNetAccessPointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteNetAccessPointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteNetAccessPointRequest(val *DeleteNetAccessPointRequest) *NullableDeleteNetAccessPointRequest {
	return &NullableDeleteNetAccessPointRequest{value: val, isSet: true}
}

func (v NullableDeleteNetAccessPointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteNetAccessPointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
