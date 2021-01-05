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

// MaintenanceEvent Information about the maintenance event.
type MaintenanceEvent struct {
	// The code of the event (`system-reboot` \\| `system-maintenance`).
	Code *string `json:"Code,omitempty"`
	// The description of the event.
	Description *string `json:"Description,omitempty"`
	// The latest scheduled end time for the event.
	NotAfter *string `json:"NotAfter,omitempty"`
	// The earliest scheduled start time for the event.
	NotBefore *string `json:"NotBefore,omitempty"`
}

// NewMaintenanceEvent instantiates a new MaintenanceEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceEvent() *MaintenanceEvent {
	this := MaintenanceEvent{}
	return &this
}

// NewMaintenanceEventWithDefaults instantiates a new MaintenanceEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceEventWithDefaults() *MaintenanceEvent {
	this := MaintenanceEvent{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MaintenanceEvent) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceEvent) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MaintenanceEvent) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *MaintenanceEvent) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MaintenanceEvent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceEvent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MaintenanceEvent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MaintenanceEvent) SetDescription(v string) {
	o.Description = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *MaintenanceEvent) GetNotAfter() string {
	if o == nil || o.NotAfter == nil {
		var ret string
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceEvent) GetNotAfterOk() (*string, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *MaintenanceEvent) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given string and assigns it to the NotAfter field.
func (o *MaintenanceEvent) SetNotAfter(v string) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *MaintenanceEvent) GetNotBefore() string {
	if o == nil || o.NotBefore == nil {
		var ret string
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceEvent) GetNotBeforeOk() (*string, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *MaintenanceEvent) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given string and assigns it to the NotBefore field.
func (o *MaintenanceEvent) SetNotBefore(v string) {
	o.NotBefore = &v
}

func (o MaintenanceEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.NotAfter != nil {
		toSerialize["NotAfter"] = o.NotAfter
	}
	if o.NotBefore != nil {
		toSerialize["NotBefore"] = o.NotBefore
	}
	return json.Marshal(toSerialize)
}

type NullableMaintenanceEvent struct {
	value *MaintenanceEvent
	isSet bool
}

func (v NullableMaintenanceEvent) Get() *MaintenanceEvent {
	return v.value
}

func (v *NullableMaintenanceEvent) Set(val *MaintenanceEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceEvent(val *MaintenanceEvent) *NullableMaintenanceEvent {
	return &NullableMaintenanceEvent{value: val, isSet: true}
}

func (v NullableMaintenanceEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


