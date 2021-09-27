/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkVirtualGatewayRequest struct for LinkVirtualGatewayRequest
type LinkVirtualGatewayRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net to which you want to attach the virtual gateway.
	NetId string `json:"NetId"`
	// The ID of the virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId"`
}

// NewLinkVirtualGatewayRequest instantiates a new LinkVirtualGatewayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkVirtualGatewayRequest(netId string, virtualGatewayId string) *LinkVirtualGatewayRequest {
	this := LinkVirtualGatewayRequest{}
	this.NetId = netId
	this.VirtualGatewayId = virtualGatewayId
	return &this
}

// NewLinkVirtualGatewayRequestWithDefaults instantiates a new LinkVirtualGatewayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkVirtualGatewayRequestWithDefaults() *LinkVirtualGatewayRequest {
	this := LinkVirtualGatewayRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkVirtualGatewayRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkVirtualGatewayRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkVirtualGatewayRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkVirtualGatewayRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetId returns the NetId field value
func (o *LinkVirtualGatewayRequest) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *LinkVirtualGatewayRequest) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *LinkVirtualGatewayRequest) SetNetId(v string) {
	o.NetId = v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value
func (o *LinkVirtualGatewayRequest) GetVirtualGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value
// and a boolean to check if the value has been set.
func (o *LinkVirtualGatewayRequest) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualGatewayId, true
}

// SetVirtualGatewayId sets field value
func (o *LinkVirtualGatewayRequest) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = v
}

func (o LinkVirtualGatewayRequest) MarshalJSON() ([]byte, error) {
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

type NullableLinkVirtualGatewayRequest struct {
	value *LinkVirtualGatewayRequest
	isSet bool
}

func (v NullableLinkVirtualGatewayRequest) Get() *LinkVirtualGatewayRequest {
	return v.value
}

func (v *NullableLinkVirtualGatewayRequest) Set(val *LinkVirtualGatewayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkVirtualGatewayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkVirtualGatewayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkVirtualGatewayRequest(val *LinkVirtualGatewayRequest) *NullableLinkVirtualGatewayRequest {
	return &NullableLinkVirtualGatewayRequest{value: val, isSet: true}
}

func (v NullableLinkVirtualGatewayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkVirtualGatewayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
