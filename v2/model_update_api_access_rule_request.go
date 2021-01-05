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

// UpdateApiAccessRuleRequest struct for UpdateApiAccessRuleRequest
type UpdateApiAccessRuleRequest struct {
	// The ID of the API access rule you want to update.
	ApiAccessRuleId string `json:"ApiAccessRuleId"`
	// One or more IDs of Client Certificate Authorities (CAs).
	CaIds *[]string `json:"CaIds,omitempty"`
	// One or more Client Certificate Common Names (CNs).
	Cns *[]string `json:"Cns,omitempty"`
	// A new description for the API access rule.
	Description *string `json:"Description,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more IP ranges, in CIDR notation (for example, 192.0.2.0/16).
	IpRanges *[]string `json:"IpRanges,omitempty"`
}

// NewUpdateApiAccessRuleRequest instantiates a new UpdateApiAccessRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiAccessRuleRequest(apiAccessRuleId string, ) *UpdateApiAccessRuleRequest {
	this := UpdateApiAccessRuleRequest{}
	this.ApiAccessRuleId = apiAccessRuleId
	return &this
}

// NewUpdateApiAccessRuleRequestWithDefaults instantiates a new UpdateApiAccessRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiAccessRuleRequestWithDefaults() *UpdateApiAccessRuleRequest {
	this := UpdateApiAccessRuleRequest{}
	return &this
}

// GetApiAccessRuleId returns the ApiAccessRuleId field value
func (o *UpdateApiAccessRuleRequest) GetApiAccessRuleId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApiAccessRuleId
}

// GetApiAccessRuleIdOk returns a tuple with the ApiAccessRuleId field value
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessRuleRequest) GetApiAccessRuleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiAccessRuleId, true
}

// SetApiAccessRuleId sets field value
func (o *UpdateApiAccessRuleRequest) SetApiAccessRuleId(v string) {
	o.ApiAccessRuleId = v
}

// GetCaIds returns the CaIds field value if set, zero value otherwise.
func (o *UpdateApiAccessRuleRequest) GetCaIds() []string {
	if o == nil || o.CaIds == nil {
		var ret []string
		return ret
	}
	return *o.CaIds
}

// GetCaIdsOk returns a tuple with the CaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessRuleRequest) GetCaIdsOk() (*[]string, bool) {
	if o == nil || o.CaIds == nil {
		return nil, false
	}
	return o.CaIds, true
}

// HasCaIds returns a boolean if a field has been set.
func (o *UpdateApiAccessRuleRequest) HasCaIds() bool {
	if o != nil && o.CaIds != nil {
		return true
	}

	return false
}

// SetCaIds gets a reference to the given []string and assigns it to the CaIds field.
func (o *UpdateApiAccessRuleRequest) SetCaIds(v []string) {
	o.CaIds = &v
}

// GetCns returns the Cns field value if set, zero value otherwise.
func (o *UpdateApiAccessRuleRequest) GetCns() []string {
	if o == nil || o.Cns == nil {
		var ret []string
		return ret
	}
	return *o.Cns
}

// GetCnsOk returns a tuple with the Cns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessRuleRequest) GetCnsOk() (*[]string, bool) {
	if o == nil || o.Cns == nil {
		return nil, false
	}
	return o.Cns, true
}

// HasCns returns a boolean if a field has been set.
func (o *UpdateApiAccessRuleRequest) HasCns() bool {
	if o != nil && o.Cns != nil {
		return true
	}

	return false
}

// SetCns gets a reference to the given []string and assigns it to the Cns field.
func (o *UpdateApiAccessRuleRequest) SetCns(v []string) {
	o.Cns = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateApiAccessRuleRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateApiAccessRuleRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateApiAccessRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateApiAccessRuleRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessRuleRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateApiAccessRuleRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateApiAccessRuleRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *UpdateApiAccessRuleRequest) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessRuleRequest) GetIpRangesOk() (*[]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *UpdateApiAccessRuleRequest) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *UpdateApiAccessRuleRequest) SetIpRanges(v []string) {
	o.IpRanges = &v
}

func (o UpdateApiAccessRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ApiAccessRuleId"] = o.ApiAccessRuleId
	}
	if o.CaIds != nil {
		toSerialize["CaIds"] = o.CaIds
	}
	if o.Cns != nil {
		toSerialize["Cns"] = o.Cns
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.IpRanges != nil {
		toSerialize["IpRanges"] = o.IpRanges
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateApiAccessRuleRequest struct {
	value *UpdateApiAccessRuleRequest
	isSet bool
}

func (v NullableUpdateApiAccessRuleRequest) Get() *UpdateApiAccessRuleRequest {
	return v.value
}

func (v *NullableUpdateApiAccessRuleRequest) Set(val *UpdateApiAccessRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiAccessRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiAccessRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiAccessRuleRequest(val *UpdateApiAccessRuleRequest) *NullableUpdateApiAccessRuleRequest {
	return &NullableUpdateApiAccessRuleRequest{value: val, isSet: true}
}

func (v NullableUpdateApiAccessRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiAccessRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


