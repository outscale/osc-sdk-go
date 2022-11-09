/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadVolumesResponse struct for ReadVolumesResponse
type ReadVolumesResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more volumes.
	Volumes *[]Volume `json:"Volumes,omitempty"`
}

// NewReadVolumesResponse instantiates a new ReadVolumesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVolumesResponse() *ReadVolumesResponse {
	this := ReadVolumesResponse{}
	return &this
}

// NewReadVolumesResponseWithDefaults instantiates a new ReadVolumesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVolumesResponseWithDefaults() *ReadVolumesResponse {
	this := ReadVolumesResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVolumesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVolumesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVolumesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVolumesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ReadVolumesResponse) GetVolumes() []Volume {
	if o == nil || o.Volumes == nil {
		var ret []Volume
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVolumesResponse) GetVolumesOk() (*[]Volume, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ReadVolumesResponse) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volume and assigns it to the Volumes field.
func (o *ReadVolumesResponse) SetVolumes(v []Volume) {
	o.Volumes = &v
}

func (o ReadVolumesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}
	return json.Marshal(toSerialize)
}

type NullableReadVolumesResponse struct {
	value *ReadVolumesResponse
	isSet bool
}

func (v NullableReadVolumesResponse) Get() *ReadVolumesResponse {
	return v.value
}

func (v *NullableReadVolumesResponse) Set(val *ReadVolumesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVolumesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVolumesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVolumesResponse(val *ReadVolumesResponse) *NullableReadVolumesResponse {
	return &NullableReadVolumesResponse{value: val, isSet: true}
}

func (v NullableReadVolumesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVolumesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
