/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// SecurityGroupLight Information about the security group.
type SecurityGroupLight struct {
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty"`
	// The name of the security group.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty"`
}

// NewSecurityGroupLight instantiates a new SecurityGroupLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupLight() *SecurityGroupLight {
	this := SecurityGroupLight{}
	return &this
}

// NewSecurityGroupLightWithDefaults instantiates a new SecurityGroupLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupLightWithDefaults() *SecurityGroupLight {
	this := SecurityGroupLight{}
	return &this
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *SecurityGroupLight) GetSecurityGroupId() string {
	if o == nil || o.SecurityGroupId == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLight) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || o.SecurityGroupId == nil {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *SecurityGroupLight) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId != nil {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *SecurityGroupLight) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetSecurityGroupName returns the SecurityGroupName field value if set, zero value otherwise.
func (o *SecurityGroupLight) GetSecurityGroupName() string {
	if o == nil || o.SecurityGroupName == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupName
}

// GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLight) GetSecurityGroupNameOk() (*string, bool) {
	if o == nil || o.SecurityGroupName == nil {
		return nil, false
	}
	return o.SecurityGroupName, true
}

// HasSecurityGroupName returns a boolean if a field has been set.
func (o *SecurityGroupLight) HasSecurityGroupName() bool {
	if o != nil && o.SecurityGroupName != nil {
		return true
	}

	return false
}

// SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.
func (o *SecurityGroupLight) SetSecurityGroupName(v string) {
	o.SecurityGroupName = &v
}

func (o SecurityGroupLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecurityGroupId != nil {
		toSerialize["SecurityGroupId"] = o.SecurityGroupId
	}
	if o.SecurityGroupName != nil {
		toSerialize["SecurityGroupName"] = o.SecurityGroupName
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityGroupLight struct {
	value *SecurityGroupLight
	isSet bool
}

func (v NullableSecurityGroupLight) Get() *SecurityGroupLight {
	return v.value
}

func (v *NullableSecurityGroupLight) Set(val *SecurityGroupLight) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupLight) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupLight(val *SecurityGroupLight) *NullableSecurityGroupLight {
	return &NullableSecurityGroupLight{value: val, isSet: true}
}

func (v NullableSecurityGroupLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


