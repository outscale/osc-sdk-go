/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkPrivateIpsRequest struct for LinkPrivateIpsRequest
type LinkPrivateIpsRequest struct {
	// If true, allows an IP that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified.
	AllowRelink *bool `json:"AllowRelink,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId"`
	// The secondary private IP or IPs you want to assign to the NIC within the IP range of the Subnet.
	PrivateIps *[]string `json:"PrivateIps,omitempty"`
	// The number of secondary private IPs to assign to the NIC.
	SecondaryPrivateIpCount *int32 `json:"SecondaryPrivateIpCount,omitempty"`
}

// NewLinkPrivateIpsRequest instantiates a new LinkPrivateIpsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkPrivateIpsRequest(nicId string) *LinkPrivateIpsRequest {
	this := LinkPrivateIpsRequest{}
	this.NicId = nicId
	return &this
}

// NewLinkPrivateIpsRequestWithDefaults instantiates a new LinkPrivateIpsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkPrivateIpsRequestWithDefaults() *LinkPrivateIpsRequest {
	this := LinkPrivateIpsRequest{}
	return &this
}

// GetAllowRelink returns the AllowRelink field value if set, zero value otherwise.
func (o *LinkPrivateIpsRequest) GetAllowRelink() bool {
	if o == nil || o.AllowRelink == nil {
		var ret bool
		return ret
	}
	return *o.AllowRelink
}

// GetAllowRelinkOk returns a tuple with the AllowRelink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetAllowRelinkOk() (*bool, bool) {
	if o == nil || o.AllowRelink == nil {
		return nil, false
	}
	return o.AllowRelink, true
}

// HasAllowRelink returns a boolean if a field has been set.
func (o *LinkPrivateIpsRequest) HasAllowRelink() bool {
	if o != nil && o.AllowRelink != nil {
		return true
	}

	return false
}

// SetAllowRelink gets a reference to the given bool and assigns it to the AllowRelink field.
func (o *LinkPrivateIpsRequest) SetAllowRelink(v bool) {
	o.AllowRelink = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkPrivateIpsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkPrivateIpsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkPrivateIpsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNicId returns the NicId field value
func (o *LinkPrivateIpsRequest) GetNicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetNicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NicId, true
}

// SetNicId sets field value
func (o *LinkPrivateIpsRequest) SetNicId(v string) {
	o.NicId = v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *LinkPrivateIpsRequest) GetPrivateIps() []string {
	if o == nil || o.PrivateIps == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetPrivateIpsOk() (*[]string, bool) {
	if o == nil || o.PrivateIps == nil {
		return nil, false
	}
	return o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *LinkPrivateIpsRequest) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.
func (o *LinkPrivateIpsRequest) SetPrivateIps(v []string) {
	o.PrivateIps = &v
}

// GetSecondaryPrivateIpCount returns the SecondaryPrivateIpCount field value if set, zero value otherwise.
func (o *LinkPrivateIpsRequest) GetSecondaryPrivateIpCount() int32 {
	if o == nil || o.SecondaryPrivateIpCount == nil {
		var ret int32
		return ret
	}
	return *o.SecondaryPrivateIpCount
}

// GetSecondaryPrivateIpCountOk returns a tuple with the SecondaryPrivateIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetSecondaryPrivateIpCountOk() (*int32, bool) {
	if o == nil || o.SecondaryPrivateIpCount == nil {
		return nil, false
	}
	return o.SecondaryPrivateIpCount, true
}

// HasSecondaryPrivateIpCount returns a boolean if a field has been set.
func (o *LinkPrivateIpsRequest) HasSecondaryPrivateIpCount() bool {
	if o != nil && o.SecondaryPrivateIpCount != nil {
		return true
	}

	return false
}

// SetSecondaryPrivateIpCount gets a reference to the given int32 and assigns it to the SecondaryPrivateIpCount field.
func (o *LinkPrivateIpsRequest) SetSecondaryPrivateIpCount(v int32) {
	o.SecondaryPrivateIpCount = &v
}

func (o LinkPrivateIpsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowRelink != nil {
		toSerialize["AllowRelink"] = o.AllowRelink
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NicId"] = o.NicId
	}
	if o.PrivateIps != nil {
		toSerialize["PrivateIps"] = o.PrivateIps
	}
	if o.SecondaryPrivateIpCount != nil {
		toSerialize["SecondaryPrivateIpCount"] = o.SecondaryPrivateIpCount
	}
	return json.Marshal(toSerialize)
}

type NullableLinkPrivateIpsRequest struct {
	value *LinkPrivateIpsRequest
	isSet bool
}

func (v NullableLinkPrivateIpsRequest) Get() *LinkPrivateIpsRequest {
	return v.value
}

func (v *NullableLinkPrivateIpsRequest) Set(val *LinkPrivateIpsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkPrivateIpsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkPrivateIpsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkPrivateIpsRequest(val *LinkPrivateIpsRequest) *NullableLinkPrivateIpsRequest {
	return &NullableLinkPrivateIpsRequest{value: val, isSet: true}
}

func (v NullableLinkPrivateIpsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkPrivateIpsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
