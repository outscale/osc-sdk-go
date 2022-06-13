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

// ResponseContext Information about the context of the response.
type ResponseContext struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty"`
}

// NewResponseContext instantiates a new ResponseContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseContext() *ResponseContext {
	this := ResponseContext{}
	return &this
}

// NewResponseContextWithDefaults instantiates a new ResponseContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseContextWithDefaults() *ResponseContext {
	this := ResponseContext{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ResponseContext) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseContext) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ResponseContext) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ResponseContext) SetRequestId(v string) {
	o.RequestId = &v
}

func (o ResponseContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["RequestId"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableResponseContext struct {
	value *ResponseContext
	isSet bool
}

func (v NullableResponseContext) Get() *ResponseContext {
	return v.value
}

func (v *NullableResponseContext) Set(val *ResponseContext) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseContext) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseContext(val *ResponseContext) *NullableResponseContext {
	return &NullableResponseContext{value: val, isSet: true}
}

func (v NullableResponseContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
