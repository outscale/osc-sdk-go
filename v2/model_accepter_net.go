/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// AccepterNet Information about the accepter Net.
type AccepterNet struct {
	// The account ID of the owner of the accepter Net.
	AccountId *string `json:"AccountId,omitempty"`
	// The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).
	IpRange *string `json:"IpRange,omitempty"`
	// The ID of the accepter Net.
	NetId *string `json:"NetId,omitempty"`
}

// NewAccepterNet instantiates a new AccepterNet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccepterNet() *AccepterNet {
	this := AccepterNet{}
	return &this
}

// NewAccepterNetWithDefaults instantiates a new AccepterNet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccepterNetWithDefaults() *AccepterNet {
	this := AccepterNet{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AccepterNet) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccepterNet) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AccepterNet) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AccepterNet) SetAccountId(v string) {
	o.AccountId = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *AccepterNet) GetIpRange() string {
	if o == nil || o.IpRange == nil {
		var ret string
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccepterNet) GetIpRangeOk() (*string, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *AccepterNet) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *AccepterNet) SetIpRange(v string) {
	o.IpRange = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *AccepterNet) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccepterNet) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *AccepterNet) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *AccepterNet) SetNetId(v string) {
	o.NetId = &v
}

func (o AccepterNet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.IpRange != nil {
		toSerialize["IpRange"] = o.IpRange
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	return json.Marshal(toSerialize)
}

type NullableAccepterNet struct {
	value *AccepterNet
	isSet bool
}

func (v NullableAccepterNet) Get() *AccepterNet {
	return v.value
}

func (v *NullableAccepterNet) Set(val *AccepterNet) {
	v.value = val
	v.isSet = true
}

func (v NullableAccepterNet) IsSet() bool {
	return v.isSet
}

func (v *NullableAccepterNet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccepterNet(val *AccepterNet) *NullableAccepterNet {
	return &NullableAccepterNet{value: val, isSet: true}
}

func (v NullableAccepterNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccepterNet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
