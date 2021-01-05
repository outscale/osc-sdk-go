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

// UpdateNicResponse struct for UpdateNicResponse
type UpdateNicResponse struct {
	Nic *Nic `json:"Nic,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateNicResponse instantiates a new UpdateNicResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNicResponse() *UpdateNicResponse {
	this := UpdateNicResponse{}
	return &this
}

// NewUpdateNicResponseWithDefaults instantiates a new UpdateNicResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNicResponseWithDefaults() *UpdateNicResponse {
	this := UpdateNicResponse{}
	return &this
}

// GetNic returns the Nic field value if set, zero value otherwise.
func (o *UpdateNicResponse) GetNic() Nic {
	if o == nil || o.Nic == nil {
		var ret Nic
		return ret
	}
	return *o.Nic
}

// GetNicOk returns a tuple with the Nic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicResponse) GetNicOk() (*Nic, bool) {
	if o == nil || o.Nic == nil {
		return nil, false
	}
	return o.Nic, true
}

// HasNic returns a boolean if a field has been set.
func (o *UpdateNicResponse) HasNic() bool {
	if o != nil && o.Nic != nil {
		return true
	}

	return false
}

// SetNic gets a reference to the given Nic and assigns it to the Nic field.
func (o *UpdateNicResponse) SetNic(v Nic) {
	o.Nic = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateNicResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateNicResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateNicResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateNicResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nic != nil {
		toSerialize["Nic"] = o.Nic
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNicResponse struct {
	value *UpdateNicResponse
	isSet bool
}

func (v NullableUpdateNicResponse) Get() *UpdateNicResponse {
	return v.value
}

func (v *NullableUpdateNicResponse) Set(val *UpdateNicResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNicResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNicResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNicResponse(val *UpdateNicResponse) *NullableUpdateNicResponse {
	return &NullableUpdateNicResponse{value: val, isSet: true}
}

func (v NullableUpdateNicResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNicResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


