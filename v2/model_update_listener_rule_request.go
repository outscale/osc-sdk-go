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

// UpdateListenerRuleRequest struct for UpdateListenerRuleRequest
type UpdateListenerRuleRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].
	HostPattern *string `json:"HostPattern,omitempty"`
	// The name of the listener rule.
	ListenerRuleName string `json:"ListenerRuleName"`
	// A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~\"'@:+?].
	PathPattern *string `json:"PathPattern,omitempty"`
}

// NewUpdateListenerRuleRequest instantiates a new UpdateListenerRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateListenerRuleRequest(listenerRuleName string, ) *UpdateListenerRuleRequest {
	this := UpdateListenerRuleRequest{}
	this.ListenerRuleName = listenerRuleName
	return &this
}

// NewUpdateListenerRuleRequestWithDefaults instantiates a new UpdateListenerRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateListenerRuleRequestWithDefaults() *UpdateListenerRuleRequest {
	this := UpdateListenerRuleRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateListenerRuleRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateListenerRuleRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateListenerRuleRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateListenerRuleRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetHostPattern returns the HostPattern field value if set, zero value otherwise.
func (o *UpdateListenerRuleRequest) GetHostPattern() string {
	if o == nil || o.HostPattern == nil {
		var ret string
		return ret
	}
	return *o.HostPattern
}

// GetHostPatternOk returns a tuple with the HostPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateListenerRuleRequest) GetHostPatternOk() (*string, bool) {
	if o == nil || o.HostPattern == nil {
		return nil, false
	}
	return o.HostPattern, true
}

// HasHostPattern returns a boolean if a field has been set.
func (o *UpdateListenerRuleRequest) HasHostPattern() bool {
	if o != nil && o.HostPattern != nil {
		return true
	}

	return false
}

// SetHostPattern gets a reference to the given string and assigns it to the HostPattern field.
func (o *UpdateListenerRuleRequest) SetHostPattern(v string) {
	o.HostPattern = &v
}

// GetListenerRuleName returns the ListenerRuleName field value
func (o *UpdateListenerRuleRequest) GetListenerRuleName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ListenerRuleName
}

// GetListenerRuleNameOk returns a tuple with the ListenerRuleName field value
// and a boolean to check if the value has been set.
func (o *UpdateListenerRuleRequest) GetListenerRuleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListenerRuleName, true
}

// SetListenerRuleName sets field value
func (o *UpdateListenerRuleRequest) SetListenerRuleName(v string) {
	o.ListenerRuleName = v
}

// GetPathPattern returns the PathPattern field value if set, zero value otherwise.
func (o *UpdateListenerRuleRequest) GetPathPattern() string {
	if o == nil || o.PathPattern == nil {
		var ret string
		return ret
	}
	return *o.PathPattern
}

// GetPathPatternOk returns a tuple with the PathPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateListenerRuleRequest) GetPathPatternOk() (*string, bool) {
	if o == nil || o.PathPattern == nil {
		return nil, false
	}
	return o.PathPattern, true
}

// HasPathPattern returns a boolean if a field has been set.
func (o *UpdateListenerRuleRequest) HasPathPattern() bool {
	if o != nil && o.PathPattern != nil {
		return true
	}

	return false
}

// SetPathPattern gets a reference to the given string and assigns it to the PathPattern field.
func (o *UpdateListenerRuleRequest) SetPathPattern(v string) {
	o.PathPattern = &v
}

func (o UpdateListenerRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.HostPattern != nil {
		toSerialize["HostPattern"] = o.HostPattern
	}
	if true {
		toSerialize["ListenerRuleName"] = o.ListenerRuleName
	}
	if o.PathPattern != nil {
		toSerialize["PathPattern"] = o.PathPattern
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateListenerRuleRequest struct {
	value *UpdateListenerRuleRequest
	isSet bool
}

func (v NullableUpdateListenerRuleRequest) Get() *UpdateListenerRuleRequest {
	return v.value
}

func (v *NullableUpdateListenerRuleRequest) Set(val *UpdateListenerRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateListenerRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateListenerRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateListenerRuleRequest(val *UpdateListenerRuleRequest) *NullableUpdateListenerRuleRequest {
	return &NullableUpdateListenerRuleRequest{value: val, isSet: true}
}

func (v NullableUpdateListenerRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateListenerRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


