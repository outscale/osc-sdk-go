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

// RejectNetPeeringResponse struct for RejectNetPeeringResponse
type RejectNetPeeringResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewRejectNetPeeringResponse instantiates a new RejectNetPeeringResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectNetPeeringResponse() *RejectNetPeeringResponse {
	this := RejectNetPeeringResponse{}
	return &this
}

// NewRejectNetPeeringResponseWithDefaults instantiates a new RejectNetPeeringResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectNetPeeringResponseWithDefaults() *RejectNetPeeringResponse {
	this := RejectNetPeeringResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *RejectNetPeeringResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectNetPeeringResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *RejectNetPeeringResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *RejectNetPeeringResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o RejectNetPeeringResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableRejectNetPeeringResponse struct {
	value *RejectNetPeeringResponse
	isSet bool
}

func (v NullableRejectNetPeeringResponse) Get() *RejectNetPeeringResponse {
	return v.value
}

func (v *NullableRejectNetPeeringResponse) Set(val *RejectNetPeeringResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectNetPeeringResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectNetPeeringResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectNetPeeringResponse(val *RejectNetPeeringResponse) *NullableRejectNetPeeringResponse {
	return &NullableRejectNetPeeringResponse{value: val, isSet: true}
}

func (v NullableRejectNetPeeringResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectNetPeeringResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


