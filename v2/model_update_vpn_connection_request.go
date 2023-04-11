/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateVpnConnectionRequest struct for UpdateVpnConnectionRequest
type UpdateVpnConnectionRequest struct {
	// The ID of the client gateway.
	ClientGatewayId *string `json:"ClientGatewayId,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the virtual gateway.
	VirtualGatewayId *string `json:"VirtualGatewayId,omitempty"`
	// The ID of the VPN connection you want to modify.
	VpnConnectionId string      `json:"VpnConnectionId"`
	VpnOptions      *VpnOptions `json:"VpnOptions,omitempty"`
}

// NewUpdateVpnConnectionRequest instantiates a new UpdateVpnConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVpnConnectionRequest(vpnConnectionId string) *UpdateVpnConnectionRequest {
	this := UpdateVpnConnectionRequest{}
	this.VpnConnectionId = vpnConnectionId
	return &this
}

// NewUpdateVpnConnectionRequestWithDefaults instantiates a new UpdateVpnConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVpnConnectionRequestWithDefaults() *UpdateVpnConnectionRequest {
	this := UpdateVpnConnectionRequest{}
	return &this
}

// GetClientGatewayId returns the ClientGatewayId field value if set, zero value otherwise.
func (o *UpdateVpnConnectionRequest) GetClientGatewayId() string {
	if o == nil || o.ClientGatewayId == nil {
		var ret string
		return ret
	}
	return *o.ClientGatewayId
}

// GetClientGatewayIdOk returns a tuple with the ClientGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVpnConnectionRequest) GetClientGatewayIdOk() (*string, bool) {
	if o == nil || o.ClientGatewayId == nil {
		return nil, false
	}
	return o.ClientGatewayId, true
}

// HasClientGatewayId returns a boolean if a field has been set.
func (o *UpdateVpnConnectionRequest) HasClientGatewayId() bool {
	if o != nil && o.ClientGatewayId != nil {
		return true
	}

	return false
}

// SetClientGatewayId gets a reference to the given string and assigns it to the ClientGatewayId field.
func (o *UpdateVpnConnectionRequest) SetClientGatewayId(v string) {
	o.ClientGatewayId = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateVpnConnectionRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVpnConnectionRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateVpnConnectionRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateVpnConnectionRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value if set, zero value otherwise.
func (o *UpdateVpnConnectionRequest) GetVirtualGatewayId() string {
	if o == nil || o.VirtualGatewayId == nil {
		var ret string
		return ret
	}
	return *o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVpnConnectionRequest) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil || o.VirtualGatewayId == nil {
		return nil, false
	}
	return o.VirtualGatewayId, true
}

// HasVirtualGatewayId returns a boolean if a field has been set.
func (o *UpdateVpnConnectionRequest) HasVirtualGatewayId() bool {
	if o != nil && o.VirtualGatewayId != nil {
		return true
	}

	return false
}

// SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.
func (o *UpdateVpnConnectionRequest) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = &v
}

// GetVpnConnectionId returns the VpnConnectionId field value
func (o *UpdateVpnConnectionRequest) GetVpnConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpnConnectionId
}

// GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field value
// and a boolean to check if the value has been set.
func (o *UpdateVpnConnectionRequest) GetVpnConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpnConnectionId, true
}

// SetVpnConnectionId sets field value
func (o *UpdateVpnConnectionRequest) SetVpnConnectionId(v string) {
	o.VpnConnectionId = v
}

// GetVpnOptions returns the VpnOptions field value if set, zero value otherwise.
func (o *UpdateVpnConnectionRequest) GetVpnOptions() VpnOptions {
	if o == nil || o.VpnOptions == nil {
		var ret VpnOptions
		return ret
	}
	return *o.VpnOptions
}

// GetVpnOptionsOk returns a tuple with the VpnOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVpnConnectionRequest) GetVpnOptionsOk() (*VpnOptions, bool) {
	if o == nil || o.VpnOptions == nil {
		return nil, false
	}
	return o.VpnOptions, true
}

// HasVpnOptions returns a boolean if a field has been set.
func (o *UpdateVpnConnectionRequest) HasVpnOptions() bool {
	if o != nil && o.VpnOptions != nil {
		return true
	}

	return false
}

// SetVpnOptions gets a reference to the given VpnOptions and assigns it to the VpnOptions field.
func (o *UpdateVpnConnectionRequest) SetVpnOptions(v VpnOptions) {
	o.VpnOptions = &v
}

func (o UpdateVpnConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientGatewayId != nil {
		toSerialize["ClientGatewayId"] = o.ClientGatewayId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.VirtualGatewayId != nil {
		toSerialize["VirtualGatewayId"] = o.VirtualGatewayId
	}
	if true {
		toSerialize["VpnConnectionId"] = o.VpnConnectionId
	}
	if o.VpnOptions != nil {
		toSerialize["VpnOptions"] = o.VpnOptions
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateVpnConnectionRequest struct {
	value *UpdateVpnConnectionRequest
	isSet bool
}

func (v NullableUpdateVpnConnectionRequest) Get() *UpdateVpnConnectionRequest {
	return v.value
}

func (v *NullableUpdateVpnConnectionRequest) Set(val *UpdateVpnConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVpnConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVpnConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVpnConnectionRequest(val *UpdateVpnConnectionRequest) *NullableUpdateVpnConnectionRequest {
	return &NullableUpdateVpnConnectionRequest{value: val, isSet: true}
}

func (v NullableUpdateVpnConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVpnConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
