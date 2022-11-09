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

// PublicIp Information about the public IP.
type PublicIp struct {
	// (Required in a Net) The ID representing the association of the public IP with the VM or the NIC.
	LinkPublicIpId *string `json:"LinkPublicIpId,omitempty"`
	// The account ID of the owner of the NIC.
	NicAccountId *string `json:"NicAccountId,omitempty"`
	// The ID of the NIC the public IP is associated with (if any).
	NicId *string `json:"NicId,omitempty"`
	// The private IP associated with the public IP.
	PrivateIp *string `json:"PrivateIp,omitempty"`
	// The public IP.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The allocation ID of the public IP.
	PublicIpId *string `json:"PublicIpId,omitempty"`
	// One or more tags associated with the public IP.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the VM the public IP is associated with (if any).
	VmId *string `json:"VmId,omitempty"`
}

// NewPublicIp instantiates a new PublicIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIp() *PublicIp {
	this := PublicIp{}
	return &this
}

// NewPublicIpWithDefaults instantiates a new PublicIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIpWithDefaults() *PublicIp {
	this := PublicIp{}
	return &this
}

// GetLinkPublicIpId returns the LinkPublicIpId field value if set, zero value otherwise.
func (o *PublicIp) GetLinkPublicIpId() string {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret
	}
	return *o.LinkPublicIpId
}

// GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetLinkPublicIpIdOk() (*string, bool) {
	if o == nil || o.LinkPublicIpId == nil {
		return nil, false
	}
	return o.LinkPublicIpId, true
}

// HasLinkPublicIpId returns a boolean if a field has been set.
func (o *PublicIp) HasLinkPublicIpId() bool {
	if o != nil && o.LinkPublicIpId != nil {
		return true
	}

	return false
}

// SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.
func (o *PublicIp) SetLinkPublicIpId(v string) {
	o.LinkPublicIpId = &v
}

// GetNicAccountId returns the NicAccountId field value if set, zero value otherwise.
func (o *PublicIp) GetNicAccountId() string {
	if o == nil || o.NicAccountId == nil {
		var ret string
		return ret
	}
	return *o.NicAccountId
}

// GetNicAccountIdOk returns a tuple with the NicAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetNicAccountIdOk() (*string, bool) {
	if o == nil || o.NicAccountId == nil {
		return nil, false
	}
	return o.NicAccountId, true
}

// HasNicAccountId returns a boolean if a field has been set.
func (o *PublicIp) HasNicAccountId() bool {
	if o != nil && o.NicAccountId != nil {
		return true
	}

	return false
}

// SetNicAccountId gets a reference to the given string and assigns it to the NicAccountId field.
func (o *PublicIp) SetNicAccountId(v string) {
	o.NicAccountId = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *PublicIp) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetNicIdOk() (*string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *PublicIp) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *PublicIp) SetNicId(v string) {
	o.NicId = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *PublicIp) GetPrivateIp() string {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetPrivateIpOk() (*string, bool) {
	if o == nil || o.PrivateIp == nil {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *PublicIp) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *PublicIp) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *PublicIp) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *PublicIp) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *PublicIp) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicIpId returns the PublicIpId field value if set, zero value otherwise.
func (o *PublicIp) GetPublicIpId() string {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpId
}

// GetPublicIpIdOk returns a tuple with the PublicIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetPublicIpIdOk() (*string, bool) {
	if o == nil || o.PublicIpId == nil {
		return nil, false
	}
	return o.PublicIpId, true
}

// HasPublicIpId returns a boolean if a field has been set.
func (o *PublicIp) HasPublicIpId() bool {
	if o != nil && o.PublicIpId != nil {
		return true
	}

	return false
}

// SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.
func (o *PublicIp) SetPublicIpId(v string) {
	o.PublicIpId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PublicIp) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PublicIp) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *PublicIp) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *PublicIp) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *PublicIp) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *PublicIp) SetVmId(v string) {
	o.VmId = &v
}

func (o PublicIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkPublicIpId != nil {
		toSerialize["LinkPublicIpId"] = o.LinkPublicIpId
	}
	if o.NicAccountId != nil {
		toSerialize["NicAccountId"] = o.NicAccountId
	}
	if o.NicId != nil {
		toSerialize["NicId"] = o.NicId
	}
	if o.PrivateIp != nil {
		toSerialize["PrivateIp"] = o.PrivateIp
	}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.PublicIpId != nil {
		toSerialize["PublicIpId"] = o.PublicIpId
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullablePublicIp struct {
	value *PublicIp
	isSet bool
}

func (v NullablePublicIp) Get() *PublicIp {
	return v.value
}

func (v *NullablePublicIp) Set(val *PublicIp) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIp) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIp(val *PublicIp) *NullablePublicIp {
	return &NullablePublicIp{value: val, isSet: true}
}

func (v NullablePublicIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
