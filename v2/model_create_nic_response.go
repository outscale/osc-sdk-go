/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.21
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateNicResponse struct for CreateNicResponse
type CreateNicResponse struct {
	Nic             *Nic             `json:"Nic,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateNicResponse instantiates a new CreateNicResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNicResponse() *CreateNicResponse {
	this := CreateNicResponse{}
	return &this
}

// NewCreateNicResponseWithDefaults instantiates a new CreateNicResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNicResponseWithDefaults() *CreateNicResponse {
	this := CreateNicResponse{}
	return &this
}

// GetNic returns the Nic field value if set, zero value otherwise.
func (o *CreateNicResponse) GetNic() Nic {
	if o == nil || o.Nic == nil {
		var ret Nic
		return ret
	}
	return *o.Nic
}

// GetNicOk returns a tuple with the Nic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicResponse) GetNicOk() (*Nic, bool) {
	if o == nil || o.Nic == nil {
		return nil, false
	}
	return o.Nic, true
}

// HasNic returns a boolean if a field has been set.
func (o *CreateNicResponse) HasNic() bool {
	if o != nil && o.Nic != nil {
		return true
	}

	return false
}

// SetNic gets a reference to the given Nic and assigns it to the Nic field.
func (o *CreateNicResponse) SetNic(v Nic) {
	o.Nic = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateNicResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateNicResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateNicResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateNicResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nic != nil {
		toSerialize["Nic"] = o.Nic
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNicResponse struct {
	value *CreateNicResponse
	isSet bool
}

func (v NullableCreateNicResponse) Get() *CreateNicResponse {
	return v.value
}

func (v *NullableCreateNicResponse) Set(val *CreateNicResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNicResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNicResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNicResponse(val *CreateNicResponse) *NullableCreateNicResponse {
	return &NullableCreateNicResponse{value: val, isSet: true}
}

func (v NullableCreateNicResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNicResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
