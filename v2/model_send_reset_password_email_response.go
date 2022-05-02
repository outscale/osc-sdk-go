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

// SendResetPasswordEmailResponse struct for SendResetPasswordEmailResponse
type SendResetPasswordEmailResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewSendResetPasswordEmailResponse instantiates a new SendResetPasswordEmailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendResetPasswordEmailResponse() *SendResetPasswordEmailResponse {
	this := SendResetPasswordEmailResponse{}
	return &this
}

// NewSendResetPasswordEmailResponseWithDefaults instantiates a new SendResetPasswordEmailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendResetPasswordEmailResponseWithDefaults() *SendResetPasswordEmailResponse {
	this := SendResetPasswordEmailResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *SendResetPasswordEmailResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendResetPasswordEmailResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *SendResetPasswordEmailResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *SendResetPasswordEmailResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o SendResetPasswordEmailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableSendResetPasswordEmailResponse struct {
	value *SendResetPasswordEmailResponse
	isSet bool
}

func (v NullableSendResetPasswordEmailResponse) Get() *SendResetPasswordEmailResponse {
	return v.value
}

func (v *NullableSendResetPasswordEmailResponse) Set(val *SendResetPasswordEmailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendResetPasswordEmailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendResetPasswordEmailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendResetPasswordEmailResponse(val *SendResetPasswordEmailResponse) *NullableSendResetPasswordEmailResponse {
	return &NullableSendResetPasswordEmailResponse{value: val, isSet: true}
}

func (v NullableSendResetPasswordEmailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendResetPasswordEmailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
