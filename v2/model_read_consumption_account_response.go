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

// ReadConsumptionAccountResponse struct for ReadConsumptionAccountResponse
type ReadConsumptionAccountResponse struct {
	// Information about the resources consumed during the specified time period.
	ConsumptionEntries *[]ConsumptionEntry `json:"ConsumptionEntries,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadConsumptionAccountResponse instantiates a new ReadConsumptionAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadConsumptionAccountResponse() *ReadConsumptionAccountResponse {
	this := ReadConsumptionAccountResponse{}
	return &this
}

// NewReadConsumptionAccountResponseWithDefaults instantiates a new ReadConsumptionAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadConsumptionAccountResponseWithDefaults() *ReadConsumptionAccountResponse {
	this := ReadConsumptionAccountResponse{}
	return &this
}

// GetConsumptionEntries returns the ConsumptionEntries field value if set, zero value otherwise.
func (o *ReadConsumptionAccountResponse) GetConsumptionEntries() []ConsumptionEntry {
	if o == nil || o.ConsumptionEntries == nil {
		var ret []ConsumptionEntry
		return ret
	}
	return *o.ConsumptionEntries
}

// GetConsumptionEntriesOk returns a tuple with the ConsumptionEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountResponse) GetConsumptionEntriesOk() (*[]ConsumptionEntry, bool) {
	if o == nil || o.ConsumptionEntries == nil {
		return nil, false
	}
	return o.ConsumptionEntries, true
}

// HasConsumptionEntries returns a boolean if a field has been set.
func (o *ReadConsumptionAccountResponse) HasConsumptionEntries() bool {
	if o != nil && o.ConsumptionEntries != nil {
		return true
	}

	return false
}

// SetConsumptionEntries gets a reference to the given []ConsumptionEntry and assigns it to the ConsumptionEntries field.
func (o *ReadConsumptionAccountResponse) SetConsumptionEntries(v []ConsumptionEntry) {
	o.ConsumptionEntries = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadConsumptionAccountResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadConsumptionAccountResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadConsumptionAccountResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadConsumptionAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsumptionEntries != nil {
		toSerialize["ConsumptionEntries"] = o.ConsumptionEntries
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadConsumptionAccountResponse struct {
	value *ReadConsumptionAccountResponse
	isSet bool
}

func (v NullableReadConsumptionAccountResponse) Get() *ReadConsumptionAccountResponse {
	return v.value
}

func (v *NullableReadConsumptionAccountResponse) Set(val *ReadConsumptionAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadConsumptionAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadConsumptionAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadConsumptionAccountResponse(val *ReadConsumptionAccountResponse) *NullableReadConsumptionAccountResponse {
	return &NullableReadConsumptionAccountResponse{value: val, isSet: true}
}

func (v NullableReadConsumptionAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadConsumptionAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


