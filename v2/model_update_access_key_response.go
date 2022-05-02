/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateAccessKeyResponse struct for UpdateAccessKeyResponse
type UpdateAccessKeyResponse struct {
	AccessKey       *AccessKey       `json:"AccessKey,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateAccessKeyResponse instantiates a new UpdateAccessKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccessKeyResponse() *UpdateAccessKeyResponse {
	this := UpdateAccessKeyResponse{}
	return &this
}

// NewUpdateAccessKeyResponseWithDefaults instantiates a new UpdateAccessKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccessKeyResponseWithDefaults() *UpdateAccessKeyResponse {
	this := UpdateAccessKeyResponse{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *UpdateAccessKeyResponse) GetAccessKey() AccessKey {
	if o == nil || o.AccessKey == nil {
		var ret AccessKey
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessKeyResponse) GetAccessKeyOk() (*AccessKey, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *UpdateAccessKeyResponse) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given AccessKey and assigns it to the AccessKey field.
func (o *UpdateAccessKeyResponse) SetAccessKey(v AccessKey) {
	o.AccessKey = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateAccessKeyResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessKeyResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateAccessKeyResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateAccessKeyResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateAccessKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKey != nil {
		toSerialize["AccessKey"] = o.AccessKey
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAccessKeyResponse struct {
	value *UpdateAccessKeyResponse
	isSet bool
}

func (v NullableUpdateAccessKeyResponse) Get() *UpdateAccessKeyResponse {
	return v.value
}

func (v *NullableUpdateAccessKeyResponse) Set(val *UpdateAccessKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccessKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccessKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccessKeyResponse(val *UpdateAccessKeyResponse) *NullableUpdateAccessKeyResponse {
	return &NullableUpdateAccessKeyResponse{value: val, isSet: true}
}

func (v NullableUpdateAccessKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccessKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
