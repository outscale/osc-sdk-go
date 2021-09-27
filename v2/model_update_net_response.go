/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateNetResponse struct for UpdateNetResponse
type UpdateNetResponse struct {
	Net             *Net             `json:"Net,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateNetResponse instantiates a new UpdateNetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetResponse() *UpdateNetResponse {
	this := UpdateNetResponse{}
	return &this
}

// NewUpdateNetResponseWithDefaults instantiates a new UpdateNetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetResponseWithDefaults() *UpdateNetResponse {
	this := UpdateNetResponse{}
	return &this
}

// GetNet returns the Net field value if set, zero value otherwise.
func (o *UpdateNetResponse) GetNet() Net {
	if o == nil || o.Net == nil {
		var ret Net
		return ret
	}
	return *o.Net
}

// GetNetOk returns a tuple with the Net field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetResponse) GetNetOk() (*Net, bool) {
	if o == nil || o.Net == nil {
		return nil, false
	}
	return o.Net, true
}

// HasNet returns a boolean if a field has been set.
func (o *UpdateNetResponse) HasNet() bool {
	if o != nil && o.Net != nil {
		return true
	}

	return false
}

// SetNet gets a reference to the given Net and assigns it to the Net field.
func (o *UpdateNetResponse) SetNet(v Net) {
	o.Net = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateNetResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateNetResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateNetResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateNetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Net != nil {
		toSerialize["Net"] = o.Net
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetResponse struct {
	value *UpdateNetResponse
	isSet bool
}

func (v NullableUpdateNetResponse) Get() *UpdateNetResponse {
	return v.value
}

func (v *NullableUpdateNetResponse) Set(val *UpdateNetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetResponse(val *UpdateNetResponse) *NullableUpdateNetResponse {
	return &NullableUpdateNetResponse{value: val, isSet: true}
}

func (v NullableUpdateNetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
