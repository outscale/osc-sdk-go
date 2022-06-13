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

// DeleteApiAccessRuleRequest struct for DeleteApiAccessRuleRequest
type DeleteApiAccessRuleRequest struct {
	// The ID of the API access rule you want to delete.
	ApiAccessRuleId string `json:"ApiAccessRuleId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewDeleteApiAccessRuleRequest instantiates a new DeleteApiAccessRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteApiAccessRuleRequest(apiAccessRuleId string) *DeleteApiAccessRuleRequest {
	this := DeleteApiAccessRuleRequest{}
	this.ApiAccessRuleId = apiAccessRuleId
	return &this
}

// NewDeleteApiAccessRuleRequestWithDefaults instantiates a new DeleteApiAccessRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteApiAccessRuleRequestWithDefaults() *DeleteApiAccessRuleRequest {
	this := DeleteApiAccessRuleRequest{}
	return &this
}

// GetApiAccessRuleId returns the ApiAccessRuleId field value
func (o *DeleteApiAccessRuleRequest) GetApiAccessRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiAccessRuleId
}

// GetApiAccessRuleIdOk returns a tuple with the ApiAccessRuleId field value
// and a boolean to check if the value has been set.
func (o *DeleteApiAccessRuleRequest) GetApiAccessRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiAccessRuleId, true
}

// SetApiAccessRuleId sets field value
func (o *DeleteApiAccessRuleRequest) SetApiAccessRuleId(v string) {
	o.ApiAccessRuleId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteApiAccessRuleRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteApiAccessRuleRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteApiAccessRuleRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteApiAccessRuleRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o DeleteApiAccessRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ApiAccessRuleId"] = o.ApiAccessRuleId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteApiAccessRuleRequest struct {
	value *DeleteApiAccessRuleRequest
	isSet bool
}

func (v NullableDeleteApiAccessRuleRequest) Get() *DeleteApiAccessRuleRequest {
	return v.value
}

func (v *NullableDeleteApiAccessRuleRequest) Set(val *DeleteApiAccessRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteApiAccessRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteApiAccessRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteApiAccessRuleRequest(val *DeleteApiAccessRuleRequest) *NullableDeleteApiAccessRuleRequest {
	return &NullableDeleteApiAccessRuleRequest{value: val, isSet: true}
}

func (v NullableDeleteApiAccessRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteApiAccessRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
