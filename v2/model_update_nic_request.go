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

// UpdateNicRequest struct for UpdateNicRequest
type UpdateNicRequest struct {
	// A new description for the NIC.
	Description *string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool            `json:"DryRun,omitempty"`
	LinkNic *LinkNicToUpdate `json:"LinkNic,omitempty"`
	// The ID of the NIC you want to modify.
	NicId string `json:"NicId"`
	// One or more IDs of security groups for the NIC.<br /> You must specify at least one group, even if you use the default security group in the Net.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
}

// NewUpdateNicRequest instantiates a new UpdateNicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNicRequest(nicId string) *UpdateNicRequest {
	this := UpdateNicRequest{}
	this.NicId = nicId
	return &this
}

// NewUpdateNicRequestWithDefaults instantiates a new UpdateNicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNicRequestWithDefaults() *UpdateNicRequest {
	this := UpdateNicRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNicRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateNicRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNicRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateNicRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateNicRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateNicRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLinkNic returns the LinkNic field value if set, zero value otherwise.
func (o *UpdateNicRequest) GetLinkNic() LinkNicToUpdate {
	if o == nil || o.LinkNic == nil {
		var ret LinkNicToUpdate
		return ret
	}
	return *o.LinkNic
}

// GetLinkNicOk returns a tuple with the LinkNic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicRequest) GetLinkNicOk() (*LinkNicToUpdate, bool) {
	if o == nil || o.LinkNic == nil {
		return nil, false
	}
	return o.LinkNic, true
}

// HasLinkNic returns a boolean if a field has been set.
func (o *UpdateNicRequest) HasLinkNic() bool {
	if o != nil && o.LinkNic != nil {
		return true
	}

	return false
}

// SetLinkNic gets a reference to the given LinkNicToUpdate and assigns it to the LinkNic field.
func (o *UpdateNicRequest) SetLinkNic(v LinkNicToUpdate) {
	o.LinkNic = &v
}

// GetNicId returns the NicId field value
func (o *UpdateNicRequest) GetNicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value
// and a boolean to check if the value has been set.
func (o *UpdateNicRequest) GetNicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NicId, true
}

// SetNicId sets field value
func (o *UpdateNicRequest) SetNicId(v string) {
	o.NicId = v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *UpdateNicRequest) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicRequest) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *UpdateNicRequest) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *UpdateNicRequest) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

func (o UpdateNicRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.LinkNic != nil {
		toSerialize["LinkNic"] = o.LinkNic
	}
	if true {
		toSerialize["NicId"] = o.NicId
	}
	if o.SecurityGroupIds != nil {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNicRequest struct {
	value *UpdateNicRequest
	isSet bool
}

func (v NullableUpdateNicRequest) Get() *UpdateNicRequest {
	return v.value
}

func (v *NullableUpdateNicRequest) Set(val *UpdateNicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNicRequest(val *UpdateNicRequest) *NullableUpdateNicRequest {
	return &NullableUpdateNicRequest{value: val, isSet: true}
}

func (v NullableUpdateNicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
