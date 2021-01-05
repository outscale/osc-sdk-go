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

// CreateCaResponse struct for CreateCaResponse
type CreateCaResponse struct {
	Ca *Ca `json:"Ca,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateCaResponse instantiates a new CreateCaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCaResponse() *CreateCaResponse {
	this := CreateCaResponse{}
	return &this
}

// NewCreateCaResponseWithDefaults instantiates a new CreateCaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCaResponseWithDefaults() *CreateCaResponse {
	this := CreateCaResponse{}
	return &this
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *CreateCaResponse) GetCa() Ca {
	if o == nil || o.Ca == nil {
		var ret Ca
		return ret
	}
	return *o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaResponse) GetCaOk() (*Ca, bool) {
	if o == nil || o.Ca == nil {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *CreateCaResponse) HasCa() bool {
	if o != nil && o.Ca != nil {
		return true
	}

	return false
}

// SetCa gets a reference to the given Ca and assigns it to the Ca field.
func (o *CreateCaResponse) SetCa(v Ca) {
	o.Ca = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateCaResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateCaResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateCaResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateCaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ca != nil {
		toSerialize["Ca"] = o.Ca
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCaResponse struct {
	value *CreateCaResponse
	isSet bool
}

func (v NullableCreateCaResponse) Get() *CreateCaResponse {
	return v.value
}

func (v *NullableCreateCaResponse) Set(val *CreateCaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCaResponse(val *CreateCaResponse) *NullableCreateCaResponse {
	return &NullableCreateCaResponse{value: val, isSet: true}
}

func (v NullableCreateCaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


