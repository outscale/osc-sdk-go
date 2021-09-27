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

// DirectLinkInterface Information about the DirectLink interface.
type DirectLinkInterface struct {
	// The BGP (Border Gateway Protocol) ASN (Autonomous System Number) on the customer's side of the DirectLink interface.
	BgpAsn int32 `json:"BgpAsn"`
	// The BGP authentication key.
	BgpKey *string `json:"BgpKey,omitempty"`
	// The IP address on the customer's side of the DirectLink interface.
	ClientPrivateIp *string `json:"ClientPrivateIp,omitempty"`
	// The name of the DirectLink interface.
	DirectLinkInterfaceName string `json:"DirectLinkInterfaceName"`
	// The IP address on the OUTSCALE side of the DirectLink interface.
	OutscalePrivateIp *string `json:"OutscalePrivateIp,omitempty"`
	// The ID of the target virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId"`
	// The VLAN number associated with the DirectLink interface.
	Vlan int32 `json:"Vlan"`
}

// NewDirectLinkInterface instantiates a new DirectLinkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectLinkInterface(bgpAsn int32, directLinkInterfaceName string, virtualGatewayId string, vlan int32) *DirectLinkInterface {
	this := DirectLinkInterface{}
	this.BgpAsn = bgpAsn
	this.DirectLinkInterfaceName = directLinkInterfaceName
	this.VirtualGatewayId = virtualGatewayId
	this.Vlan = vlan
	return &this
}

// NewDirectLinkInterfaceWithDefaults instantiates a new DirectLinkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectLinkInterfaceWithDefaults() *DirectLinkInterface {
	this := DirectLinkInterface{}
	return &this
}

// GetBgpAsn returns the BgpAsn field value
func (o *DirectLinkInterface) GetBgpAsn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BgpAsn
}

// GetBgpAsnOk returns a tuple with the BgpAsn field value
// and a boolean to check if the value has been set.
func (o *DirectLinkInterface) GetBgpAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BgpAsn, true
}

// SetBgpAsn sets field value
func (o *DirectLinkInterface) SetBgpAsn(v int32) {
	o.BgpAsn = v
}

// GetBgpKey returns the BgpKey field value if set, zero value otherwise.
func (o *DirectLinkInterface) GetBgpKey() string {
	if o == nil || o.BgpKey == nil {
		var ret string
		return ret
	}
	return *o.BgpKey
}

// GetBgpKeyOk returns a tuple with the BgpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterface) GetBgpKeyOk() (*string, bool) {
	if o == nil || o.BgpKey == nil {
		return nil, false
	}
	return o.BgpKey, true
}

// HasBgpKey returns a boolean if a field has been set.
func (o *DirectLinkInterface) HasBgpKey() bool {
	if o != nil && o.BgpKey != nil {
		return true
	}

	return false
}

// SetBgpKey gets a reference to the given string and assigns it to the BgpKey field.
func (o *DirectLinkInterface) SetBgpKey(v string) {
	o.BgpKey = &v
}

// GetClientPrivateIp returns the ClientPrivateIp field value if set, zero value otherwise.
func (o *DirectLinkInterface) GetClientPrivateIp() string {
	if o == nil || o.ClientPrivateIp == nil {
		var ret string
		return ret
	}
	return *o.ClientPrivateIp
}

// GetClientPrivateIpOk returns a tuple with the ClientPrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterface) GetClientPrivateIpOk() (*string, bool) {
	if o == nil || o.ClientPrivateIp == nil {
		return nil, false
	}
	return o.ClientPrivateIp, true
}

// HasClientPrivateIp returns a boolean if a field has been set.
func (o *DirectLinkInterface) HasClientPrivateIp() bool {
	if o != nil && o.ClientPrivateIp != nil {
		return true
	}

	return false
}

// SetClientPrivateIp gets a reference to the given string and assigns it to the ClientPrivateIp field.
func (o *DirectLinkInterface) SetClientPrivateIp(v string) {
	o.ClientPrivateIp = &v
}

// GetDirectLinkInterfaceName returns the DirectLinkInterfaceName field value
func (o *DirectLinkInterface) GetDirectLinkInterfaceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectLinkInterfaceName
}

// GetDirectLinkInterfaceNameOk returns a tuple with the DirectLinkInterfaceName field value
// and a boolean to check if the value has been set.
func (o *DirectLinkInterface) GetDirectLinkInterfaceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectLinkInterfaceName, true
}

// SetDirectLinkInterfaceName sets field value
func (o *DirectLinkInterface) SetDirectLinkInterfaceName(v string) {
	o.DirectLinkInterfaceName = v
}

// GetOutscalePrivateIp returns the OutscalePrivateIp field value if set, zero value otherwise.
func (o *DirectLinkInterface) GetOutscalePrivateIp() string {
	if o == nil || o.OutscalePrivateIp == nil {
		var ret string
		return ret
	}
	return *o.OutscalePrivateIp
}

// GetOutscalePrivateIpOk returns a tuple with the OutscalePrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterface) GetOutscalePrivateIpOk() (*string, bool) {
	if o == nil || o.OutscalePrivateIp == nil {
		return nil, false
	}
	return o.OutscalePrivateIp, true
}

// HasOutscalePrivateIp returns a boolean if a field has been set.
func (o *DirectLinkInterface) HasOutscalePrivateIp() bool {
	if o != nil && o.OutscalePrivateIp != nil {
		return true
	}

	return false
}

// SetOutscalePrivateIp gets a reference to the given string and assigns it to the OutscalePrivateIp field.
func (o *DirectLinkInterface) SetOutscalePrivateIp(v string) {
	o.OutscalePrivateIp = &v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value
func (o *DirectLinkInterface) GetVirtualGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value
// and a boolean to check if the value has been set.
func (o *DirectLinkInterface) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualGatewayId, true
}

// SetVirtualGatewayId sets field value
func (o *DirectLinkInterface) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = v
}

// GetVlan returns the Vlan field value
func (o *DirectLinkInterface) GetVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *DirectLinkInterface) GetVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vlan, true
}

// SetVlan sets field value
func (o *DirectLinkInterface) SetVlan(v int32) {
	o.Vlan = v
}

func (o DirectLinkInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["BgpAsn"] = o.BgpAsn
	}
	if o.BgpKey != nil {
		toSerialize["BgpKey"] = o.BgpKey
	}
	if o.ClientPrivateIp != nil {
		toSerialize["ClientPrivateIp"] = o.ClientPrivateIp
	}
	if true {
		toSerialize["DirectLinkInterfaceName"] = o.DirectLinkInterfaceName
	}
	if o.OutscalePrivateIp != nil {
		toSerialize["OutscalePrivateIp"] = o.OutscalePrivateIp
	}
	if true {
		toSerialize["VirtualGatewayId"] = o.VirtualGatewayId
	}
	if true {
		toSerialize["Vlan"] = o.Vlan
	}
	return json.Marshal(toSerialize)
}

type NullableDirectLinkInterface struct {
	value *DirectLinkInterface
	isSet bool
}

func (v NullableDirectLinkInterface) Get() *DirectLinkInterface {
	return v.value
}

func (v *NullableDirectLinkInterface) Set(val *DirectLinkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectLinkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectLinkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectLinkInterface(val *DirectLinkInterface) *NullableDirectLinkInterface {
	return &NullableDirectLinkInterface{value: val, isSet: true}
}

func (v NullableDirectLinkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectLinkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
