/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.16
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteVpnConnectionRouteRequest struct for DeleteVpnConnectionRouteRequest
type DeleteVpnConnectionRouteRequest struct {
	// The network prefix of the route to delete, in CIDR notation (for example, 10.12.0.0/16).
	DestinationIpRange string `json:"DestinationIpRange"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the target VPN connection of the static route to delete.
	VpnConnectionId string `json:"VpnConnectionId"`
}

// NewDeleteVpnConnectionRouteRequest instantiates a new DeleteVpnConnectionRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVpnConnectionRouteRequest(destinationIpRange string, vpnConnectionId string) *DeleteVpnConnectionRouteRequest {
	this := DeleteVpnConnectionRouteRequest{}
	this.DestinationIpRange = destinationIpRange
	this.VpnConnectionId = vpnConnectionId
	return &this
}

// NewDeleteVpnConnectionRouteRequestWithDefaults instantiates a new DeleteVpnConnectionRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVpnConnectionRouteRequestWithDefaults() *DeleteVpnConnectionRouteRequest {
	this := DeleteVpnConnectionRouteRequest{}
	return &this
}

// GetDestinationIpRange returns the DestinationIpRange field value
func (o *DeleteVpnConnectionRouteRequest) GetDestinationIpRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationIpRange
}

// GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field value
// and a boolean to check if the value has been set.
func (o *DeleteVpnConnectionRouteRequest) GetDestinationIpRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationIpRange, true
}

// SetDestinationIpRange sets field value
func (o *DeleteVpnConnectionRouteRequest) SetDestinationIpRange(v string) {
	o.DestinationIpRange = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteVpnConnectionRouteRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVpnConnectionRouteRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteVpnConnectionRouteRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteVpnConnectionRouteRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetVpnConnectionId returns the VpnConnectionId field value
func (o *DeleteVpnConnectionRouteRequest) GetVpnConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpnConnectionId
}

// GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field value
// and a boolean to check if the value has been set.
func (o *DeleteVpnConnectionRouteRequest) GetVpnConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpnConnectionId, true
}

// SetVpnConnectionId sets field value
func (o *DeleteVpnConnectionRouteRequest) SetVpnConnectionId(v string) {
	o.VpnConnectionId = v
}

func (o DeleteVpnConnectionRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DestinationIpRange"] = o.DestinationIpRange
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["VpnConnectionId"] = o.VpnConnectionId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteVpnConnectionRouteRequest struct {
	value *DeleteVpnConnectionRouteRequest
	isSet bool
}

func (v NullableDeleteVpnConnectionRouteRequest) Get() *DeleteVpnConnectionRouteRequest {
	return v.value
}

func (v *NullableDeleteVpnConnectionRouteRequest) Set(val *DeleteVpnConnectionRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVpnConnectionRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVpnConnectionRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVpnConnectionRouteRequest(val *DeleteVpnConnectionRouteRequest) *NullableDeleteVpnConnectionRouteRequest {
	return &NullableDeleteVpnConnectionRouteRequest{value: val, isSet: true}
}

func (v NullableDeleteVpnConnectionRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVpnConnectionRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
