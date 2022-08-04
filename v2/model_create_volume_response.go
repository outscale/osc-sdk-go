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

// CreateVolumeResponse struct for CreateVolumeResponse
type CreateVolumeResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	Volume          *Volume          `json:"Volume,omitempty"`
}

// NewCreateVolumeResponse instantiates a new CreateVolumeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVolumeResponse() *CreateVolumeResponse {
	this := CreateVolumeResponse{}
	return &this
}

// NewCreateVolumeResponseWithDefaults instantiates a new CreateVolumeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVolumeResponseWithDefaults() *CreateVolumeResponse {
	this := CreateVolumeResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateVolumeResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVolumeResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateVolumeResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateVolumeResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *CreateVolumeResponse) GetVolume() Volume {
	if o == nil || o.Volume == nil {
		var ret Volume
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVolumeResponse) GetVolumeOk() (*Volume, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *CreateVolumeResponse) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given Volume and assigns it to the Volume field.
func (o *CreateVolumeResponse) SetVolume(v Volume) {
	o.Volume = &v
}

func (o CreateVolumeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Volume != nil {
		toSerialize["Volume"] = o.Volume
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVolumeResponse struct {
	value *CreateVolumeResponse
	isSet bool
}

func (v NullableCreateVolumeResponse) Get() *CreateVolumeResponse {
	return v.value
}

func (v *NullableCreateVolumeResponse) Set(val *CreateVolumeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVolumeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVolumeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVolumeResponse(val *CreateVolumeResponse) *NullableCreateVolumeResponse {
	return &NullableCreateVolumeResponse{value: val, isSet: true}
}

func (v NullableCreateVolumeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVolumeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
