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

// UnlinkVirtualGatewayRequest struct for UnlinkVirtualGatewayRequest
type UnlinkVirtualGatewayRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net from which you want to detach the virtual gateway.
	NetId string `json:"NetId"`
	// The ID of the virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId"`
}

// NewUnlinkVirtualGatewayRequest instantiates a new UnlinkVirtualGatewayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkVirtualGatewayRequest(netId string, virtualGatewayId string) *UnlinkVirtualGatewayRequest {
	this := UnlinkVirtualGatewayRequest{}
	this.NetId = netId
	this.VirtualGatewayId = virtualGatewayId
	return &this
}

// NewUnlinkVirtualGatewayRequestWithDefaults instantiates a new UnlinkVirtualGatewayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkVirtualGatewayRequestWithDefaults() *UnlinkVirtualGatewayRequest {
	this := UnlinkVirtualGatewayRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkVirtualGatewayRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkVirtualGatewayRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkVirtualGatewayRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkVirtualGatewayRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetId returns the NetId field value
func (o *UnlinkVirtualGatewayRequest) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *UnlinkVirtualGatewayRequest) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *UnlinkVirtualGatewayRequest) SetNetId(v string) {
	o.NetId = v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value
func (o *UnlinkVirtualGatewayRequest) GetVirtualGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value
// and a boolean to check if the value has been set.
func (o *UnlinkVirtualGatewayRequest) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualGatewayId, true
}

// SetVirtualGatewayId sets field value
func (o *UnlinkVirtualGatewayRequest) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = v
}

func (o UnlinkVirtualGatewayRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NetId"] = o.NetId
	}
	if true {
		toSerialize["VirtualGatewayId"] = o.VirtualGatewayId
	}
	return json.Marshal(toSerialize)
}

type NullableUnlinkVirtualGatewayRequest struct {
	value *UnlinkVirtualGatewayRequest
	isSet bool
}

func (v NullableUnlinkVirtualGatewayRequest) Get() *UnlinkVirtualGatewayRequest {
	return v.value
}

func (v *NullableUnlinkVirtualGatewayRequest) Set(val *UnlinkVirtualGatewayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkVirtualGatewayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkVirtualGatewayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkVirtualGatewayRequest(val *UnlinkVirtualGatewayRequest) *NullableUnlinkVirtualGatewayRequest {
	return &NullableUnlinkVirtualGatewayRequest{value: val, isSet: true}
}

func (v NullableUnlinkVirtualGatewayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkVirtualGatewayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
