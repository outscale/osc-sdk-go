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

// UnlinkPrivateIpsResponse struct for UnlinkPrivateIpsResponse
type UnlinkPrivateIpsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUnlinkPrivateIpsResponse instantiates a new UnlinkPrivateIpsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkPrivateIpsResponse() *UnlinkPrivateIpsResponse {
	this := UnlinkPrivateIpsResponse{}
	return &this
}

// NewUnlinkPrivateIpsResponseWithDefaults instantiates a new UnlinkPrivateIpsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkPrivateIpsResponseWithDefaults() *UnlinkPrivateIpsResponse {
	this := UnlinkPrivateIpsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UnlinkPrivateIpsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkPrivateIpsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UnlinkPrivateIpsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UnlinkPrivateIpsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UnlinkPrivateIpsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUnlinkPrivateIpsResponse struct {
	value *UnlinkPrivateIpsResponse
	isSet bool
}

func (v NullableUnlinkPrivateIpsResponse) Get() *UnlinkPrivateIpsResponse {
	return v.value
}

func (v *NullableUnlinkPrivateIpsResponse) Set(val *UnlinkPrivateIpsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkPrivateIpsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkPrivateIpsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkPrivateIpsResponse(val *UnlinkPrivateIpsResponse) *NullableUnlinkPrivateIpsResponse {
	return &NullableUnlinkPrivateIpsResponse{value: val, isSet: true}
}

func (v NullableUnlinkPrivateIpsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkPrivateIpsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
