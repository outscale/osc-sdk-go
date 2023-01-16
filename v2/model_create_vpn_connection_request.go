/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateVpnConnectionRequest struct for CreateVpnConnectionRequest
type CreateVpnConnectionRequest struct {
	// The ID of the client gateway.
	ClientGatewayId string `json:"ClientGatewayId"`
	// The type of VPN connection (only `ipsec.1` is supported).
	ConnectionType string `json:"ConnectionType"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute).
	StaticRoutesOnly *bool `json:"StaticRoutesOnly,omitempty"`
	// The ID of the virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId"`
}

// NewCreateVpnConnectionRequest instantiates a new CreateVpnConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVpnConnectionRequest(clientGatewayId string, connectionType string, virtualGatewayId string) *CreateVpnConnectionRequest {
	this := CreateVpnConnectionRequest{}
	this.ClientGatewayId = clientGatewayId
	this.ConnectionType = connectionType
	this.VirtualGatewayId = virtualGatewayId
	return &this
}

// NewCreateVpnConnectionRequestWithDefaults instantiates a new CreateVpnConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVpnConnectionRequestWithDefaults() *CreateVpnConnectionRequest {
	this := CreateVpnConnectionRequest{}
	return &this
}

// GetClientGatewayId returns the ClientGatewayId field value
func (o *CreateVpnConnectionRequest) GetClientGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientGatewayId
}

// GetClientGatewayIdOk returns a tuple with the ClientGatewayId field value
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionRequest) GetClientGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientGatewayId, true
}

// SetClientGatewayId sets field value
func (o *CreateVpnConnectionRequest) SetClientGatewayId(v string) {
	o.ClientGatewayId = v
}

// GetConnectionType returns the ConnectionType field value
func (o *CreateVpnConnectionRequest) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionRequest) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *CreateVpnConnectionRequest) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateVpnConnectionRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateVpnConnectionRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateVpnConnectionRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetStaticRoutesOnly returns the StaticRoutesOnly field value if set, zero value otherwise.
func (o *CreateVpnConnectionRequest) GetStaticRoutesOnly() bool {
	if o == nil || o.StaticRoutesOnly == nil {
		var ret bool
		return ret
	}
	return *o.StaticRoutesOnly
}

// GetStaticRoutesOnlyOk returns a tuple with the StaticRoutesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionRequest) GetStaticRoutesOnlyOk() (*bool, bool) {
	if o == nil || o.StaticRoutesOnly == nil {
		return nil, false
	}
	return o.StaticRoutesOnly, true
}

// HasStaticRoutesOnly returns a boolean if a field has been set.
func (o *CreateVpnConnectionRequest) HasStaticRoutesOnly() bool {
	if o != nil && o.StaticRoutesOnly != nil {
		return true
	}

	return false
}

// SetStaticRoutesOnly gets a reference to the given bool and assigns it to the StaticRoutesOnly field.
func (o *CreateVpnConnectionRequest) SetStaticRoutesOnly(v bool) {
	o.StaticRoutesOnly = &v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value
func (o *CreateVpnConnectionRequest) GetVirtualGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionRequest) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualGatewayId, true
}

// SetVirtualGatewayId sets field value
func (o *CreateVpnConnectionRequest) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = v
}

func (o CreateVpnConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClientGatewayId"] = o.ClientGatewayId
	}
	if true {
		toSerialize["ConnectionType"] = o.ConnectionType
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.StaticRoutesOnly != nil {
		toSerialize["StaticRoutesOnly"] = o.StaticRoutesOnly
	}
	if true {
		toSerialize["VirtualGatewayId"] = o.VirtualGatewayId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVpnConnectionRequest struct {
	value *CreateVpnConnectionRequest
	isSet bool
}

func (v NullableCreateVpnConnectionRequest) Get() *CreateVpnConnectionRequest {
	return v.value
}

func (v *NullableCreateVpnConnectionRequest) Set(val *CreateVpnConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVpnConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVpnConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVpnConnectionRequest(val *CreateVpnConnectionRequest) *NullableCreateVpnConnectionRequest {
	return &NullableCreateVpnConnectionRequest{value: val, isSet: true}
}

func (v NullableCreateVpnConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVpnConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
