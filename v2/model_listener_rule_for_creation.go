/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ListenerRuleForCreation Information about the listener rule.
type ListenerRuleForCreation struct {
	// The type of action for the rule (always `forward`).
	Action *string `json:"Action,omitempty"`
	// A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].
	HostNamePattern *string `json:"HostNamePattern,omitempty"`
	// A human-readable name for the listener rule.
	ListenerRuleName string `json:"ListenerRuleName"`
	// A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].
	PathPattern *string `json:"PathPattern,omitempty"`
	// The priority level of the listener rule, between `1` and `19999` both included. Each rule must have a unique priority level. Otherwise, an error is returned.
	Priority int32 `json:"Priority"`
}

// NewListenerRuleForCreation instantiates a new ListenerRuleForCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListenerRuleForCreation(listenerRuleName string, priority int32) *ListenerRuleForCreation {
	this := ListenerRuleForCreation{}
	this.ListenerRuleName = listenerRuleName
	this.Priority = priority
	return &this
}

// NewListenerRuleForCreationWithDefaults instantiates a new ListenerRuleForCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenerRuleForCreationWithDefaults() *ListenerRuleForCreation {
	this := ListenerRuleForCreation{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ListenerRuleForCreation) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRuleForCreation) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ListenerRuleForCreation) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ListenerRuleForCreation) SetAction(v string) {
	o.Action = &v
}

// GetHostNamePattern returns the HostNamePattern field value if set, zero value otherwise.
func (o *ListenerRuleForCreation) GetHostNamePattern() string {
	if o == nil || o.HostNamePattern == nil {
		var ret string
		return ret
	}
	return *o.HostNamePattern
}

// GetHostNamePatternOk returns a tuple with the HostNamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRuleForCreation) GetHostNamePatternOk() (*string, bool) {
	if o == nil || o.HostNamePattern == nil {
		return nil, false
	}
	return o.HostNamePattern, true
}

// HasHostNamePattern returns a boolean if a field has been set.
func (o *ListenerRuleForCreation) HasHostNamePattern() bool {
	if o != nil && o.HostNamePattern != nil {
		return true
	}

	return false
}

// SetHostNamePattern gets a reference to the given string and assigns it to the HostNamePattern field.
func (o *ListenerRuleForCreation) SetHostNamePattern(v string) {
	o.HostNamePattern = &v
}

// GetListenerRuleName returns the ListenerRuleName field value
func (o *ListenerRuleForCreation) GetListenerRuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListenerRuleName
}

// GetListenerRuleNameOk returns a tuple with the ListenerRuleName field value
// and a boolean to check if the value has been set.
func (o *ListenerRuleForCreation) GetListenerRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListenerRuleName, true
}

// SetListenerRuleName sets field value
func (o *ListenerRuleForCreation) SetListenerRuleName(v string) {
	o.ListenerRuleName = v
}

// GetPathPattern returns the PathPattern field value if set, zero value otherwise.
func (o *ListenerRuleForCreation) GetPathPattern() string {
	if o == nil || o.PathPattern == nil {
		var ret string
		return ret
	}
	return *o.PathPattern
}

// GetPathPatternOk returns a tuple with the PathPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRuleForCreation) GetPathPatternOk() (*string, bool) {
	if o == nil || o.PathPattern == nil {
		return nil, false
	}
	return o.PathPattern, true
}

// HasPathPattern returns a boolean if a field has been set.
func (o *ListenerRuleForCreation) HasPathPattern() bool {
	if o != nil && o.PathPattern != nil {
		return true
	}

	return false
}

// SetPathPattern gets a reference to the given string and assigns it to the PathPattern field.
func (o *ListenerRuleForCreation) SetPathPattern(v string) {
	o.PathPattern = &v
}

// GetPriority returns the Priority field value
func (o *ListenerRuleForCreation) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ListenerRuleForCreation) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *ListenerRuleForCreation) SetPriority(v int32) {
	o.Priority = v
}

func (o ListenerRuleForCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.HostNamePattern != nil {
		toSerialize["HostNamePattern"] = o.HostNamePattern
	}
	if true {
		toSerialize["ListenerRuleName"] = o.ListenerRuleName
	}
	if o.PathPattern != nil {
		toSerialize["PathPattern"] = o.PathPattern
	}
	if true {
		toSerialize["Priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullableListenerRuleForCreation struct {
	value *ListenerRuleForCreation
	isSet bool
}

func (v NullableListenerRuleForCreation) Get() *ListenerRuleForCreation {
	return v.value
}

func (v *NullableListenerRuleForCreation) Set(val *ListenerRuleForCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableListenerRuleForCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableListenerRuleForCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListenerRuleForCreation(val *ListenerRuleForCreation) *NullableListenerRuleForCreation {
	return &NullableListenerRuleForCreation{value: val, isSet: true}
}

func (v NullableListenerRuleForCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListenerRuleForCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
