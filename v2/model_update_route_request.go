/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.27
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateRouteRequest struct for UpdateRouteRequest
type UpdateRouteRequest struct {
	// The IP range used for the destination match, in CIDR notation (for example, `10.0.0.0/24`).
	DestinationIpRange string `json:"DestinationIpRange"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of an Internet service or virtual gateway attached to your Net.
	GatewayId *string `json:"GatewayId,omitempty"`
	// The ID of a NAT service.
	NatServiceId *string `json:"NatServiceId,omitempty"`
	// The ID of a Net peering.
	NetPeeringId *string `json:"NetPeeringId,omitempty"`
	// The ID of a network interface card (NIC).
	NicId *string `json:"NicId,omitempty"`
	// The ID of the route table.
	RouteTableId string `json:"RouteTableId"`
	// The ID of a NAT VM in your Net.
	VmId *string `json:"VmId,omitempty"`
}

// NewUpdateRouteRequest instantiates a new UpdateRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRouteRequest(destinationIpRange string, routeTableId string) *UpdateRouteRequest {
	this := UpdateRouteRequest{}
	this.DestinationIpRange = destinationIpRange
	this.RouteTableId = routeTableId
	return &this
}

// NewUpdateRouteRequestWithDefaults instantiates a new UpdateRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRouteRequestWithDefaults() *UpdateRouteRequest {
	this := UpdateRouteRequest{}
	return &this
}

// GetDestinationIpRange returns the DestinationIpRange field value
func (o *UpdateRouteRequest) GetDestinationIpRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationIpRange
}

// GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field value
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetDestinationIpRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationIpRange, true
}

// SetDestinationIpRange sets field value
func (o *UpdateRouteRequest) SetDestinationIpRange(v string) {
	o.DestinationIpRange = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateRouteRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *UpdateRouteRequest) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetNatServiceId returns the NatServiceId field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetNatServiceId() string {
	if o == nil || o.NatServiceId == nil {
		var ret string
		return ret
	}
	return *o.NatServiceId
}

// GetNatServiceIdOk returns a tuple with the NatServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetNatServiceIdOk() (*string, bool) {
	if o == nil || o.NatServiceId == nil {
		return nil, false
	}
	return o.NatServiceId, true
}

// HasNatServiceId returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasNatServiceId() bool {
	if o != nil && o.NatServiceId != nil {
		return true
	}

	return false
}

// SetNatServiceId gets a reference to the given string and assigns it to the NatServiceId field.
func (o *UpdateRouteRequest) SetNatServiceId(v string) {
	o.NatServiceId = &v
}

// GetNetPeeringId returns the NetPeeringId field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetNetPeeringId() string {
	if o == nil || o.NetPeeringId == nil {
		var ret string
		return ret
	}
	return *o.NetPeeringId
}

// GetNetPeeringIdOk returns a tuple with the NetPeeringId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetNetPeeringIdOk() (*string, bool) {
	if o == nil || o.NetPeeringId == nil {
		return nil, false
	}
	return o.NetPeeringId, true
}

// HasNetPeeringId returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasNetPeeringId() bool {
	if o != nil && o.NetPeeringId != nil {
		return true
	}

	return false
}

// SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.
func (o *UpdateRouteRequest) SetNetPeeringId(v string) {
	o.NetPeeringId = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetNicIdOk() (*string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *UpdateRouteRequest) SetNicId(v string) {
	o.NicId = &v
}

// GetRouteTableId returns the RouteTableId field value
func (o *UpdateRouteRequest) GetRouteTableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteTableId
}

// GetRouteTableIdOk returns a tuple with the RouteTableId field value
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetRouteTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteTableId, true
}

// SetRouteTableId sets field value
func (o *UpdateRouteRequest) SetRouteTableId(v string) {
	o.RouteTableId = v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *UpdateRouteRequest) SetVmId(v string) {
	o.VmId = &v
}

func (o UpdateRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DestinationIpRange"] = o.DestinationIpRange
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.GatewayId != nil {
		toSerialize["GatewayId"] = o.GatewayId
	}
	if o.NatServiceId != nil {
		toSerialize["NatServiceId"] = o.NatServiceId
	}
	if o.NetPeeringId != nil {
		toSerialize["NetPeeringId"] = o.NetPeeringId
	}
	if o.NicId != nil {
		toSerialize["NicId"] = o.NicId
	}
	if true {
		toSerialize["RouteTableId"] = o.RouteTableId
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRouteRequest struct {
	value *UpdateRouteRequest
	isSet bool
}

func (v NullableUpdateRouteRequest) Get() *UpdateRouteRequest {
	return v.value
}

func (v *NullableUpdateRouteRequest) Set(val *UpdateRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRouteRequest(val *UpdateRouteRequest) *NullableUpdateRouteRequest {
	return &NullableUpdateRouteRequest{value: val, isSet: true}
}

func (v NullableUpdateRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
