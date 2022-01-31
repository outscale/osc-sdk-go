/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteKeypairResponse struct for DeleteKeypairResponse
type DeleteKeypairResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewDeleteKeypairResponse instantiates a new DeleteKeypairResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteKeypairResponse() *DeleteKeypairResponse {
	this := DeleteKeypairResponse{}
	return &this
}

// NewDeleteKeypairResponseWithDefaults instantiates a new DeleteKeypairResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteKeypairResponseWithDefaults() *DeleteKeypairResponse {
	this := DeleteKeypairResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteKeypairResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteKeypairResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteKeypairResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteKeypairResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o DeleteKeypairResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteKeypairResponse struct {
	value *DeleteKeypairResponse
	isSet bool
}

func (v NullableDeleteKeypairResponse) Get() *DeleteKeypairResponse {
	return v.value
}

func (v *NullableDeleteKeypairResponse) Set(val *DeleteKeypairResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteKeypairResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteKeypairResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteKeypairResponse(val *DeleteKeypairResponse) *NullableDeleteKeypairResponse {
	return &NullableDeleteKeypairResponse{value: val, isSet: true}
}

func (v NullableDeleteKeypairResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteKeypairResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
