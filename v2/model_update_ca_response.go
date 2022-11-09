/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateCaResponse struct for UpdateCaResponse
type UpdateCaResponse struct {
	Ca              *Ca              `json:"Ca,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateCaResponse instantiates a new UpdateCaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCaResponse() *UpdateCaResponse {
	this := UpdateCaResponse{}
	return &this
}

// NewUpdateCaResponseWithDefaults instantiates a new UpdateCaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCaResponseWithDefaults() *UpdateCaResponse {
	this := UpdateCaResponse{}
	return &this
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *UpdateCaResponse) GetCa() Ca {
	if o == nil || o.Ca == nil {
		var ret Ca
		return ret
	}
	return *o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCaResponse) GetCaOk() (*Ca, bool) {
	if o == nil || o.Ca == nil {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *UpdateCaResponse) HasCa() bool {
	if o != nil && o.Ca != nil {
		return true
	}

	return false
}

// SetCa gets a reference to the given Ca and assigns it to the Ca field.
func (o *UpdateCaResponse) SetCa(v Ca) {
	o.Ca = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateCaResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCaResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateCaResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateCaResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateCaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ca != nil {
		toSerialize["Ca"] = o.Ca
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCaResponse struct {
	value *UpdateCaResponse
	isSet bool
}

func (v NullableUpdateCaResponse) Get() *UpdateCaResponse {
	return v.value
}

func (v *NullableUpdateCaResponse) Set(val *UpdateCaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCaResponse(val *UpdateCaResponse) *NullableUpdateCaResponse {
	return &NullableUpdateCaResponse{value: val, isSet: true}
}

func (v NullableUpdateCaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
