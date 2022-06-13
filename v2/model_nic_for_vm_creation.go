/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// NicForVmCreation Information about the network interface card (NIC) when creating a virtual machine (VM).
type NicForVmCreation struct {
	// If true, the NIC is deleted when the VM is terminated. You can specify this parameter only for a new NIC. To modify this value for an existing NIC, see [UpdateNic](#updatenic).
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The description of the NIC, if you are creating a NIC when creating the VM.
	Description *string `json:"Description,omitempty"`
	// The index of the VM device for the NIC attachment (between `0` and `7`, both included). This parameter is required if you create a NIC when creating the VM.
	DeviceNumber *int32 `json:"DeviceNumber,omitempty"`
	// The ID of the NIC, if you are attaching an existing NIC when creating a VM.
	NicId *string `json:"NicId,omitempty"`
	// One or more private IPs to assign to the NIC, if you create a NIC when creating a VM. Only one private IP can be the primary private IP.
	PrivateIps *[]PrivateIpLight `json:"PrivateIps,omitempty"`
	// The number of secondary private IPs, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP in the `PrivateIps` parameter.
	SecondaryPrivateIpCount *int32 `json:"SecondaryPrivateIpCount,omitempty"`
	// One or more IDs of security groups for the NIC, if you create a NIC when creating a VM.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The ID of the Subnet for the NIC, if you create a NIC when creating a VM. This parameter is required if you create a NIC when creating the VM.
	SubnetId *string `json:"SubnetId,omitempty"`
}

// NewNicForVmCreation instantiates a new NicForVmCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNicForVmCreation() *NicForVmCreation {
	this := NicForVmCreation{}
	return &this
}

// NewNicForVmCreationWithDefaults instantiates a new NicForVmCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNicForVmCreationWithDefaults() *NicForVmCreation {
	this := NicForVmCreation{}
	return &this
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *NicForVmCreation) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *NicForVmCreation) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *NicForVmCreation) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NicForVmCreation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NicForVmCreation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NicForVmCreation) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceNumber returns the DeviceNumber field value if set, zero value otherwise.
func (o *NicForVmCreation) GetDeviceNumber() int32 {
	if o == nil || o.DeviceNumber == nil {
		var ret int32
		return ret
	}
	return *o.DeviceNumber
}

// GetDeviceNumberOk returns a tuple with the DeviceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetDeviceNumberOk() (*int32, bool) {
	if o == nil || o.DeviceNumber == nil {
		return nil, false
	}
	return o.DeviceNumber, true
}

// HasDeviceNumber returns a boolean if a field has been set.
func (o *NicForVmCreation) HasDeviceNumber() bool {
	if o != nil && o.DeviceNumber != nil {
		return true
	}

	return false
}

// SetDeviceNumber gets a reference to the given int32 and assigns it to the DeviceNumber field.
func (o *NicForVmCreation) SetDeviceNumber(v int32) {
	o.DeviceNumber = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *NicForVmCreation) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetNicIdOk() (*string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *NicForVmCreation) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *NicForVmCreation) SetNicId(v string) {
	o.NicId = &v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *NicForVmCreation) GetPrivateIps() []PrivateIpLight {
	if o == nil || o.PrivateIps == nil {
		var ret []PrivateIpLight
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetPrivateIpsOk() (*[]PrivateIpLight, bool) {
	if o == nil || o.PrivateIps == nil {
		return nil, false
	}
	return o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *NicForVmCreation) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []PrivateIpLight and assigns it to the PrivateIps field.
func (o *NicForVmCreation) SetPrivateIps(v []PrivateIpLight) {
	o.PrivateIps = &v
}

// GetSecondaryPrivateIpCount returns the SecondaryPrivateIpCount field value if set, zero value otherwise.
func (o *NicForVmCreation) GetSecondaryPrivateIpCount() int32 {
	if o == nil || o.SecondaryPrivateIpCount == nil {
		var ret int32
		return ret
	}
	return *o.SecondaryPrivateIpCount
}

// GetSecondaryPrivateIpCountOk returns a tuple with the SecondaryPrivateIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetSecondaryPrivateIpCountOk() (*int32, bool) {
	if o == nil || o.SecondaryPrivateIpCount == nil {
		return nil, false
	}
	return o.SecondaryPrivateIpCount, true
}

// HasSecondaryPrivateIpCount returns a boolean if a field has been set.
func (o *NicForVmCreation) HasSecondaryPrivateIpCount() bool {
	if o != nil && o.SecondaryPrivateIpCount != nil {
		return true
	}

	return false
}

// SetSecondaryPrivateIpCount gets a reference to the given int32 and assigns it to the SecondaryPrivateIpCount field.
func (o *NicForVmCreation) SetSecondaryPrivateIpCount(v int32) {
	o.SecondaryPrivateIpCount = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *NicForVmCreation) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *NicForVmCreation) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *NicForVmCreation) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *NicForVmCreation) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetSubnetIdOk() (*string, bool) {
	if o == nil || o.SubnetId == nil {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *NicForVmCreation) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *NicForVmCreation) SetSubnetId(v string) {
	o.SubnetId = &v
}

func (o NicForVmCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteOnVmDeletion != nil {
		toSerialize["DeleteOnVmDeletion"] = o.DeleteOnVmDeletion
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DeviceNumber != nil {
		toSerialize["DeviceNumber"] = o.DeviceNumber
	}
	if o.NicId != nil {
		toSerialize["NicId"] = o.NicId
	}
	if o.PrivateIps != nil {
		toSerialize["PrivateIps"] = o.PrivateIps
	}
	if o.SecondaryPrivateIpCount != nil {
		toSerialize["SecondaryPrivateIpCount"] = o.SecondaryPrivateIpCount
	}
	if o.SecurityGroupIds != nil {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	if o.SubnetId != nil {
		toSerialize["SubnetId"] = o.SubnetId
	}
	return json.Marshal(toSerialize)
}

type NullableNicForVmCreation struct {
	value *NicForVmCreation
	isSet bool
}

func (v NullableNicForVmCreation) Get() *NicForVmCreation {
	return v.value
}

func (v *NullableNicForVmCreation) Set(val *NicForVmCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableNicForVmCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableNicForVmCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNicForVmCreation(val *NicForVmCreation) *NullableNicForVmCreation {
	return &NullableNicForVmCreation{value: val, isSet: true}
}

func (v NullableNicForVmCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNicForVmCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
