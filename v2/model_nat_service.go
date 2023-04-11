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

// NatService Information about the NAT service.
type NatService struct {
	// The ID of the NAT service.
	NatServiceId *string `json:"NatServiceId,omitempty"`
	// The ID of the Net in which the NAT service is.
	NetId *string `json:"NetId,omitempty"`
	// Information about the public IP or IPs associated with the NAT service.
	PublicIps *[]PublicIpLight `json:"PublicIps,omitempty"`
	// The state of the NAT service (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
	// The ID of the Subnet in which the NAT service is.
	SubnetId *string `json:"SubnetId,omitempty"`
	// One or more tags associated with the NAT service.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewNatService instantiates a new NatService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNatService() *NatService {
	this := NatService{}
	return &this
}

// NewNatServiceWithDefaults instantiates a new NatService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNatServiceWithDefaults() *NatService {
	this := NatService{}
	return &this
}

// GetNatServiceId returns the NatServiceId field value if set, zero value otherwise.
func (o *NatService) GetNatServiceId() string {
	if o == nil || o.NatServiceId == nil {
		var ret string
		return ret
	}
	return *o.NatServiceId
}

// GetNatServiceIdOk returns a tuple with the NatServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatService) GetNatServiceIdOk() (*string, bool) {
	if o == nil || o.NatServiceId == nil {
		return nil, false
	}
	return o.NatServiceId, true
}

// HasNatServiceId returns a boolean if a field has been set.
func (o *NatService) HasNatServiceId() bool {
	if o != nil && o.NatServiceId != nil {
		return true
	}

	return false
}

// SetNatServiceId gets a reference to the given string and assigns it to the NatServiceId field.
func (o *NatService) SetNatServiceId(v string) {
	o.NatServiceId = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *NatService) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatService) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *NatService) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *NatService) SetNetId(v string) {
	o.NetId = &v
}

// GetPublicIps returns the PublicIps field value if set, zero value otherwise.
func (o *NatService) GetPublicIps() []PublicIpLight {
	if o == nil || o.PublicIps == nil {
		var ret []PublicIpLight
		return ret
	}
	return *o.PublicIps
}

// GetPublicIpsOk returns a tuple with the PublicIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatService) GetPublicIpsOk() (*[]PublicIpLight, bool) {
	if o == nil || o.PublicIps == nil {
		return nil, false
	}
	return o.PublicIps, true
}

// HasPublicIps returns a boolean if a field has been set.
func (o *NatService) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}

// SetPublicIps gets a reference to the given []PublicIpLight and assigns it to the PublicIps field.
func (o *NatService) SetPublicIps(v []PublicIpLight) {
	o.PublicIps = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NatService) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatService) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NatService) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NatService) SetState(v string) {
	o.State = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *NatService) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatService) GetSubnetIdOk() (*string, bool) {
	if o == nil || o.SubnetId == nil {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *NatService) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *NatService) SetSubnetId(v string) {
	o.SubnetId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NatService) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatService) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NatService) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *NatService) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o NatService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NatServiceId != nil {
		toSerialize["NatServiceId"] = o.NatServiceId
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.PublicIps != nil {
		toSerialize["PublicIps"] = o.PublicIps
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SubnetId != nil {
		toSerialize["SubnetId"] = o.SubnetId
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableNatService struct {
	value *NatService
	isSet bool
}

func (v NullableNatService) Get() *NatService {
	return v.value
}

func (v *NullableNatService) Set(val *NatService) {
	v.value = val
	v.isSet = true
}

func (v NullableNatService) IsSet() bool {
	return v.isSet
}

func (v *NullableNatService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatService(val *NatService) *NullableNatService {
	return &NullableNatService{value: val, isSet: true}
}

func (v NullableNatService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
