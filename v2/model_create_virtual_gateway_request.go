/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateVirtualGatewayRequest struct for CreateVirtualGatewayRequest
type CreateVirtualGatewayRequest struct {
	// The type of VPN connection supported by the virtual gateway (only `ipsec.1` is supported).
	ConnectionType string `json:"ConnectionType"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewCreateVirtualGatewayRequest instantiates a new CreateVirtualGatewayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVirtualGatewayRequest(connectionType string) *CreateVirtualGatewayRequest {
	this := CreateVirtualGatewayRequest{}
	this.ConnectionType = connectionType
	return &this
}

// NewCreateVirtualGatewayRequestWithDefaults instantiates a new CreateVirtualGatewayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVirtualGatewayRequestWithDefaults() *CreateVirtualGatewayRequest {
	this := CreateVirtualGatewayRequest{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *CreateVirtualGatewayRequest) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *CreateVirtualGatewayRequest) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *CreateVirtualGatewayRequest) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateVirtualGatewayRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualGatewayRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateVirtualGatewayRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateVirtualGatewayRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o CreateVirtualGatewayRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ConnectionType"] = o.ConnectionType
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVirtualGatewayRequest struct {
	value *CreateVirtualGatewayRequest
	isSet bool
}

func (v NullableCreateVirtualGatewayRequest) Get() *CreateVirtualGatewayRequest {
	return v.value
}

func (v *NullableCreateVirtualGatewayRequest) Set(val *CreateVirtualGatewayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVirtualGatewayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVirtualGatewayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVirtualGatewayRequest(val *CreateVirtualGatewayRequest) *NullableCreateVirtualGatewayRequest {
	return &NullableCreateVirtualGatewayRequest{value: val, isSet: true}
}

func (v NullableCreateVirtualGatewayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVirtualGatewayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
