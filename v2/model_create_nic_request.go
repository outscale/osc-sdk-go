/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateNicRequest struct for CreateNicRequest
type CreateNicRequest struct {
	// A description for the NIC.
	Description *string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The primary private IP for the NIC.<br /> This IP must be within the IP range of the Subnet that you specify with the `SubnetId` attribute.<br /> If you do not specify this attribute, a random private IP is selected within the IP range of the Subnet.
	PrivateIps *[]PrivateIpLight `json:"PrivateIps,omitempty"`
	// One or more IDs of security groups for the NIC.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The ID of the Subnet in which you want to create the NIC.
	SubnetId string `json:"SubnetId"`
}

// NewCreateNicRequest instantiates a new CreateNicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNicRequest(subnetId string) *CreateNicRequest {
	this := CreateNicRequest{}
	this.SubnetId = subnetId
	return &this
}

// NewCreateNicRequestWithDefaults instantiates a new CreateNicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNicRequestWithDefaults() *CreateNicRequest {
	this := CreateNicRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateNicRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateNicRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateNicRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateNicRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateNicRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateNicRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *CreateNicRequest) GetPrivateIps() []PrivateIpLight {
	if o == nil || o.PrivateIps == nil {
		var ret []PrivateIpLight
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetPrivateIpsOk() (*[]PrivateIpLight, bool) {
	if o == nil || o.PrivateIps == nil {
		return nil, false
	}
	return o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *CreateNicRequest) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []PrivateIpLight and assigns it to the PrivateIps field.
func (o *CreateNicRequest) SetPrivateIps(v []PrivateIpLight) {
	o.PrivateIps = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *CreateNicRequest) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *CreateNicRequest) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *CreateNicRequest) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSubnetId returns the SubnetId field value
func (o *CreateNicRequest) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *CreateNicRequest) SetSubnetId(v string) {
	o.SubnetId = v
}

func (o CreateNicRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.PrivateIps != nil {
		toSerialize["PrivateIps"] = o.PrivateIps
	}
	if o.SecurityGroupIds != nil {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	if true {
		toSerialize["SubnetId"] = o.SubnetId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNicRequest struct {
	value *CreateNicRequest
	isSet bool
}

func (v NullableCreateNicRequest) Get() *CreateNicRequest {
	return v.value
}

func (v *NullableCreateNicRequest) Set(val *CreateNicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNicRequest(val *CreateNicRequest) *NullableCreateNicRequest {
	return &NullableCreateNicRequest{value: val, isSet: true}
}

func (v NullableCreateNicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
