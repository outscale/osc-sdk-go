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

// VpnConnection Information about a VPN connection.
type VpnConnection struct {
	// Example configuration for the client gateway.
	ClientGatewayConfiguration *string `json:"ClientGatewayConfiguration,omitempty"`
	// The ID of the client gateway used on the client end of the connection.
	ClientGatewayId *string `json:"ClientGatewayId,omitempty"`
	// The type of VPN connection (always `ipsec.1`).
	ConnectionType *string `json:"ConnectionType,omitempty"`
	// Information about one or more static routes associated with the VPN connection, if any.
	Routes *[]RouteLight `json:"Routes,omitempty"`
	// The state of the VPN connection (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
	// If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute).
	StaticRoutesOnly *bool `json:"StaticRoutesOnly,omitempty"`
	// One or more tags associated with the VPN connection.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// Information about the current state of one or more of the VPN tunnels.
	VgwTelemetries *[]VgwTelemetry `json:"VgwTelemetries,omitempty"`
	// The ID of the virtual gateway used on the OUTSCALE end of the connection.
	VirtualGatewayId *string `json:"VirtualGatewayId,omitempty"`
	// The ID of the VPN connection.
	VpnConnectionId *string     `json:"VpnConnectionId,omitempty"`
	VpnOptions      *VpnOptions `json:"VpnOptions,omitempty"`
}

// NewVpnConnection instantiates a new VpnConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpnConnection() *VpnConnection {
	this := VpnConnection{}
	return &this
}

// NewVpnConnectionWithDefaults instantiates a new VpnConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpnConnectionWithDefaults() *VpnConnection {
	this := VpnConnection{}
	return &this
}

// GetClientGatewayConfiguration returns the ClientGatewayConfiguration field value if set, zero value otherwise.
func (o *VpnConnection) GetClientGatewayConfiguration() string {
	if o == nil || o.ClientGatewayConfiguration == nil {
		var ret string
		return ret
	}
	return *o.ClientGatewayConfiguration
}

// GetClientGatewayConfigurationOk returns a tuple with the ClientGatewayConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetClientGatewayConfigurationOk() (*string, bool) {
	if o == nil || o.ClientGatewayConfiguration == nil {
		return nil, false
	}
	return o.ClientGatewayConfiguration, true
}

// HasClientGatewayConfiguration returns a boolean if a field has been set.
func (o *VpnConnection) HasClientGatewayConfiguration() bool {
	if o != nil && o.ClientGatewayConfiguration != nil {
		return true
	}

	return false
}

// SetClientGatewayConfiguration gets a reference to the given string and assigns it to the ClientGatewayConfiguration field.
func (o *VpnConnection) SetClientGatewayConfiguration(v string) {
	o.ClientGatewayConfiguration = &v
}

// GetClientGatewayId returns the ClientGatewayId field value if set, zero value otherwise.
func (o *VpnConnection) GetClientGatewayId() string {
	if o == nil || o.ClientGatewayId == nil {
		var ret string
		return ret
	}
	return *o.ClientGatewayId
}

// GetClientGatewayIdOk returns a tuple with the ClientGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetClientGatewayIdOk() (*string, bool) {
	if o == nil || o.ClientGatewayId == nil {
		return nil, false
	}
	return o.ClientGatewayId, true
}

// HasClientGatewayId returns a boolean if a field has been set.
func (o *VpnConnection) HasClientGatewayId() bool {
	if o != nil && o.ClientGatewayId != nil {
		return true
	}

	return false
}

// SetClientGatewayId gets a reference to the given string and assigns it to the ClientGatewayId field.
func (o *VpnConnection) SetClientGatewayId(v string) {
	o.ClientGatewayId = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *VpnConnection) GetConnectionType() string {
	if o == nil || o.ConnectionType == nil {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil || o.ConnectionType == nil {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *VpnConnection) HasConnectionType() bool {
	if o != nil && o.ConnectionType != nil {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *VpnConnection) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *VpnConnection) GetRoutes() []RouteLight {
	if o == nil || o.Routes == nil {
		var ret []RouteLight
		return ret
	}
	return *o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetRoutesOk() (*[]RouteLight, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *VpnConnection) HasRoutes() bool {
	if o != nil && o.Routes != nil {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []RouteLight and assigns it to the Routes field.
func (o *VpnConnection) SetRoutes(v []RouteLight) {
	o.Routes = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VpnConnection) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VpnConnection) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VpnConnection) SetState(v string) {
	o.State = &v
}

// GetStaticRoutesOnly returns the StaticRoutesOnly field value if set, zero value otherwise.
func (o *VpnConnection) GetStaticRoutesOnly() bool {
	if o == nil || o.StaticRoutesOnly == nil {
		var ret bool
		return ret
	}
	return *o.StaticRoutesOnly
}

// GetStaticRoutesOnlyOk returns a tuple with the StaticRoutesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetStaticRoutesOnlyOk() (*bool, bool) {
	if o == nil || o.StaticRoutesOnly == nil {
		return nil, false
	}
	return o.StaticRoutesOnly, true
}

// HasStaticRoutesOnly returns a boolean if a field has been set.
func (o *VpnConnection) HasStaticRoutesOnly() bool {
	if o != nil && o.StaticRoutesOnly != nil {
		return true
	}

	return false
}

// SetStaticRoutesOnly gets a reference to the given bool and assigns it to the StaticRoutesOnly field.
func (o *VpnConnection) SetStaticRoutesOnly(v bool) {
	o.StaticRoutesOnly = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VpnConnection) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VpnConnection) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *VpnConnection) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVgwTelemetries returns the VgwTelemetries field value if set, zero value otherwise.
func (o *VpnConnection) GetVgwTelemetries() []VgwTelemetry {
	if o == nil || o.VgwTelemetries == nil {
		var ret []VgwTelemetry
		return ret
	}
	return *o.VgwTelemetries
}

// GetVgwTelemetriesOk returns a tuple with the VgwTelemetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetVgwTelemetriesOk() (*[]VgwTelemetry, bool) {
	if o == nil || o.VgwTelemetries == nil {
		return nil, false
	}
	return o.VgwTelemetries, true
}

// HasVgwTelemetries returns a boolean if a field has been set.
func (o *VpnConnection) HasVgwTelemetries() bool {
	if o != nil && o.VgwTelemetries != nil {
		return true
	}

	return false
}

// SetVgwTelemetries gets a reference to the given []VgwTelemetry and assigns it to the VgwTelemetries field.
func (o *VpnConnection) SetVgwTelemetries(v []VgwTelemetry) {
	o.VgwTelemetries = &v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value if set, zero value otherwise.
func (o *VpnConnection) GetVirtualGatewayId() string {
	if o == nil || o.VirtualGatewayId == nil {
		var ret string
		return ret
	}
	return *o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil || o.VirtualGatewayId == nil {
		return nil, false
	}
	return o.VirtualGatewayId, true
}

// HasVirtualGatewayId returns a boolean if a field has been set.
func (o *VpnConnection) HasVirtualGatewayId() bool {
	if o != nil && o.VirtualGatewayId != nil {
		return true
	}

	return false
}

// SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.
func (o *VpnConnection) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = &v
}

// GetVpnConnectionId returns the VpnConnectionId field value if set, zero value otherwise.
func (o *VpnConnection) GetVpnConnectionId() string {
	if o == nil || o.VpnConnectionId == nil {
		var ret string
		return ret
	}
	return *o.VpnConnectionId
}

// GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetVpnConnectionIdOk() (*string, bool) {
	if o == nil || o.VpnConnectionId == nil {
		return nil, false
	}
	return o.VpnConnectionId, true
}

// HasVpnConnectionId returns a boolean if a field has been set.
func (o *VpnConnection) HasVpnConnectionId() bool {
	if o != nil && o.VpnConnectionId != nil {
		return true
	}

	return false
}

// SetVpnConnectionId gets a reference to the given string and assigns it to the VpnConnectionId field.
func (o *VpnConnection) SetVpnConnectionId(v string) {
	o.VpnConnectionId = &v
}

// GetVpnOptions returns the VpnOptions field value if set, zero value otherwise.
func (o *VpnConnection) GetVpnOptions() VpnOptions {
	if o == nil || o.VpnOptions == nil {
		var ret VpnOptions
		return ret
	}
	return *o.VpnOptions
}

// GetVpnOptionsOk returns a tuple with the VpnOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnConnection) GetVpnOptionsOk() (*VpnOptions, bool) {
	if o == nil || o.VpnOptions == nil {
		return nil, false
	}
	return o.VpnOptions, true
}

// HasVpnOptions returns a boolean if a field has been set.
func (o *VpnConnection) HasVpnOptions() bool {
	if o != nil && o.VpnOptions != nil {
		return true
	}

	return false
}

// SetVpnOptions gets a reference to the given VpnOptions and assigns it to the VpnOptions field.
func (o *VpnConnection) SetVpnOptions(v VpnOptions) {
	o.VpnOptions = &v
}

func (o VpnConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientGatewayConfiguration != nil {
		toSerialize["ClientGatewayConfiguration"] = o.ClientGatewayConfiguration
	}
	if o.ClientGatewayId != nil {
		toSerialize["ClientGatewayId"] = o.ClientGatewayId
	}
	if o.ConnectionType != nil {
		toSerialize["ConnectionType"] = o.ConnectionType
	}
	if o.Routes != nil {
		toSerialize["Routes"] = o.Routes
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.StaticRoutesOnly != nil {
		toSerialize["StaticRoutesOnly"] = o.StaticRoutesOnly
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.VgwTelemetries != nil {
		toSerialize["VgwTelemetries"] = o.VgwTelemetries
	}
	if o.VirtualGatewayId != nil {
		toSerialize["VirtualGatewayId"] = o.VirtualGatewayId
	}
	if o.VpnConnectionId != nil {
		toSerialize["VpnConnectionId"] = o.VpnConnectionId
	}
	if o.VpnOptions != nil {
		toSerialize["VpnOptions"] = o.VpnOptions
	}
	return json.Marshal(toSerialize)
}

type NullableVpnConnection struct {
	value *VpnConnection
	isSet bool
}

func (v NullableVpnConnection) Get() *VpnConnection {
	return v.value
}

func (v *NullableVpnConnection) Set(val *VpnConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnConnection(val *VpnConnection) *NullableVpnConnection {
	return &NullableVpnConnection{value: val, isSet: true}
}

func (v NullableVpnConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
