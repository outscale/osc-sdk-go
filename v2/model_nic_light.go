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

// NicLight Information about the network interface card (NIC).
type NicLight struct {
	// The account ID of the owner of the NIC.
	AccountId *string `json:"AccountId,omitempty"`
	// The description of the NIC.
	Description *string `json:"Description,omitempty"`
	// (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked *bool                   `json:"IsSourceDestChecked,omitempty"`
	LinkNic             *LinkNicLight           `json:"LinkNic,omitempty"`
	LinkPublicIp        *LinkPublicIpLightForVm `json:"LinkPublicIp,omitempty"`
	// The Media Access Control (MAC) address of the NIC.
	MacAddress *string `json:"MacAddress,omitempty"`
	// The ID of the Net for the NIC.
	NetId *string `json:"NetId,omitempty"`
	// The ID of the NIC.
	NicId *string `json:"NicId,omitempty"`
	// The name of the private DNS.
	PrivateDnsName *string `json:"PrivateDnsName,omitempty"`
	// The private IP or IPs of the NIC.
	PrivateIps *[]PrivateIpLightForVm `json:"PrivateIps,omitempty"`
	// One or more IDs of security groups for the NIC.
	SecurityGroups *[]SecurityGroupLight `json:"SecurityGroups,omitempty"`
	// The state of the NIC (`available` \\| `attaching` \\| `in-use` \\| `detaching`).
	State *string `json:"State,omitempty"`
	// The ID of the Subnet for the NIC.
	SubnetId *string `json:"SubnetId,omitempty"`
}

// NewNicLight instantiates a new NicLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNicLight() *NicLight {
	this := NicLight{}
	return &this
}

// NewNicLightWithDefaults instantiates a new NicLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNicLightWithDefaults() *NicLight {
	this := NicLight{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *NicLight) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *NicLight) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *NicLight) SetAccountId(v string) {
	o.AccountId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NicLight) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NicLight) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NicLight) SetDescription(v string) {
	o.Description = &v
}

// GetIsSourceDestChecked returns the IsSourceDestChecked field value if set, zero value otherwise.
func (o *NicLight) GetIsSourceDestChecked() bool {
	if o == nil || o.IsSourceDestChecked == nil {
		var ret bool
		return ret
	}
	return *o.IsSourceDestChecked
}

// GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetIsSourceDestCheckedOk() (*bool, bool) {
	if o == nil || o.IsSourceDestChecked == nil {
		return nil, false
	}
	return o.IsSourceDestChecked, true
}

// HasIsSourceDestChecked returns a boolean if a field has been set.
func (o *NicLight) HasIsSourceDestChecked() bool {
	if o != nil && o.IsSourceDestChecked != nil {
		return true
	}

	return false
}

// SetIsSourceDestChecked gets a reference to the given bool and assigns it to the IsSourceDestChecked field.
func (o *NicLight) SetIsSourceDestChecked(v bool) {
	o.IsSourceDestChecked = &v
}

// GetLinkNic returns the LinkNic field value if set, zero value otherwise.
func (o *NicLight) GetLinkNic() LinkNicLight {
	if o == nil || o.LinkNic == nil {
		var ret LinkNicLight
		return ret
	}
	return *o.LinkNic
}

// GetLinkNicOk returns a tuple with the LinkNic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetLinkNicOk() (*LinkNicLight, bool) {
	if o == nil || o.LinkNic == nil {
		return nil, false
	}
	return o.LinkNic, true
}

// HasLinkNic returns a boolean if a field has been set.
func (o *NicLight) HasLinkNic() bool {
	if o != nil && o.LinkNic != nil {
		return true
	}

	return false
}

// SetLinkNic gets a reference to the given LinkNicLight and assigns it to the LinkNic field.
func (o *NicLight) SetLinkNic(v LinkNicLight) {
	o.LinkNic = &v
}

// GetLinkPublicIp returns the LinkPublicIp field value if set, zero value otherwise.
func (o *NicLight) GetLinkPublicIp() LinkPublicIpLightForVm {
	if o == nil || o.LinkPublicIp == nil {
		var ret LinkPublicIpLightForVm
		return ret
	}
	return *o.LinkPublicIp
}

// GetLinkPublicIpOk returns a tuple with the LinkPublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetLinkPublicIpOk() (*LinkPublicIpLightForVm, bool) {
	if o == nil || o.LinkPublicIp == nil {
		return nil, false
	}
	return o.LinkPublicIp, true
}

// HasLinkPublicIp returns a boolean if a field has been set.
func (o *NicLight) HasLinkPublicIp() bool {
	if o != nil && o.LinkPublicIp != nil {
		return true
	}

	return false
}

// SetLinkPublicIp gets a reference to the given LinkPublicIpLightForVm and assigns it to the LinkPublicIp field.
func (o *NicLight) SetLinkPublicIp(v LinkPublicIpLightForVm) {
	o.LinkPublicIp = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *NicLight) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *NicLight) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *NicLight) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *NicLight) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *NicLight) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *NicLight) SetNetId(v string) {
	o.NetId = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *NicLight) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetNicIdOk() (*string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *NicLight) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *NicLight) SetNicId(v string) {
	o.NicId = &v
}

// GetPrivateDnsName returns the PrivateDnsName field value if set, zero value otherwise.
func (o *NicLight) GetPrivateDnsName() string {
	if o == nil || o.PrivateDnsName == nil {
		var ret string
		return ret
	}
	return *o.PrivateDnsName
}

// GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetPrivateDnsNameOk() (*string, bool) {
	if o == nil || o.PrivateDnsName == nil {
		return nil, false
	}
	return o.PrivateDnsName, true
}

// HasPrivateDnsName returns a boolean if a field has been set.
func (o *NicLight) HasPrivateDnsName() bool {
	if o != nil && o.PrivateDnsName != nil {
		return true
	}

	return false
}

// SetPrivateDnsName gets a reference to the given string and assigns it to the PrivateDnsName field.
func (o *NicLight) SetPrivateDnsName(v string) {
	o.PrivateDnsName = &v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *NicLight) GetPrivateIps() []PrivateIpLightForVm {
	if o == nil || o.PrivateIps == nil {
		var ret []PrivateIpLightForVm
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetPrivateIpsOk() (*[]PrivateIpLightForVm, bool) {
	if o == nil || o.PrivateIps == nil {
		return nil, false
	}
	return o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *NicLight) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []PrivateIpLightForVm and assigns it to the PrivateIps field.
func (o *NicLight) SetPrivateIps(v []PrivateIpLightForVm) {
	o.PrivateIps = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *NicLight) GetSecurityGroups() []SecurityGroupLight {
	if o == nil || o.SecurityGroups == nil {
		var ret []SecurityGroupLight
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetSecurityGroupsOk() (*[]SecurityGroupLight, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *NicLight) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []SecurityGroupLight and assigns it to the SecurityGroups field.
func (o *NicLight) SetSecurityGroups(v []SecurityGroupLight) {
	o.SecurityGroups = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NicLight) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NicLight) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NicLight) SetState(v string) {
	o.State = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *NicLight) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicLight) GetSubnetIdOk() (*string, bool) {
	if o == nil || o.SubnetId == nil {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *NicLight) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *NicLight) SetSubnetId(v string) {
	o.SubnetId = &v
}

func (o NicLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.IsSourceDestChecked != nil {
		toSerialize["IsSourceDestChecked"] = o.IsSourceDestChecked
	}
	if o.LinkNic != nil {
		toSerialize["LinkNic"] = o.LinkNic
	}
	if o.LinkPublicIp != nil {
		toSerialize["LinkPublicIp"] = o.LinkPublicIp
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.NicId != nil {
		toSerialize["NicId"] = o.NicId
	}
	if o.PrivateDnsName != nil {
		toSerialize["PrivateDnsName"] = o.PrivateDnsName
	}
	if o.PrivateIps != nil {
		toSerialize["PrivateIps"] = o.PrivateIps
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SubnetId != nil {
		toSerialize["SubnetId"] = o.SubnetId
	}
	return json.Marshal(toSerialize)
}

type NullableNicLight struct {
	value *NicLight
	isSet bool
}

func (v NullableNicLight) Get() *NicLight {
	return v.value
}

func (v *NullableNicLight) Set(val *NicLight) {
	v.value = val
	v.isSet = true
}

func (v NullableNicLight) IsSet() bool {
	return v.isSet
}

func (v *NullableNicLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNicLight(val *NicLight) *NullableNicLight {
	return &NullableNicLight{value: val, isSet: true}
}

func (v NullableNicLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNicLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
