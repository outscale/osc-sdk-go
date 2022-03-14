/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateFlexibleGpuResponse struct for UpdateFlexibleGpuResponse
type UpdateFlexibleGpuResponse struct {
	FlexibleGpu     *FlexibleGpu     `json:"FlexibleGpu,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateFlexibleGpuResponse instantiates a new UpdateFlexibleGpuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFlexibleGpuResponse() *UpdateFlexibleGpuResponse {
	this := UpdateFlexibleGpuResponse{}
	return &this
}

// NewUpdateFlexibleGpuResponseWithDefaults instantiates a new UpdateFlexibleGpuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFlexibleGpuResponseWithDefaults() *UpdateFlexibleGpuResponse {
	this := UpdateFlexibleGpuResponse{}
	return &this
}

// GetFlexibleGpu returns the FlexibleGpu field value if set, zero value otherwise.
func (o *UpdateFlexibleGpuResponse) GetFlexibleGpu() FlexibleGpu {
	if o == nil || o.FlexibleGpu == nil {
		var ret FlexibleGpu
		return ret
	}
	return *o.FlexibleGpu
}

// GetFlexibleGpuOk returns a tuple with the FlexibleGpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFlexibleGpuResponse) GetFlexibleGpuOk() (*FlexibleGpu, bool) {
	if o == nil || o.FlexibleGpu == nil {
		return nil, false
	}
	return o.FlexibleGpu, true
}

// HasFlexibleGpu returns a boolean if a field has been set.
func (o *UpdateFlexibleGpuResponse) HasFlexibleGpu() bool {
	if o != nil && o.FlexibleGpu != nil {
		return true
	}

	return false
}

// SetFlexibleGpu gets a reference to the given FlexibleGpu and assigns it to the FlexibleGpu field.
func (o *UpdateFlexibleGpuResponse) SetFlexibleGpu(v FlexibleGpu) {
	o.FlexibleGpu = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateFlexibleGpuResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFlexibleGpuResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateFlexibleGpuResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateFlexibleGpuResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlexibleGpu != nil {
		toSerialize["FlexibleGpu"] = o.FlexibleGpu
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFlexibleGpuResponse struct {
	value *UpdateFlexibleGpuResponse
	isSet bool
}

func (v NullableUpdateFlexibleGpuResponse) Get() *UpdateFlexibleGpuResponse {
	return v.value
}

func (v *NullableUpdateFlexibleGpuResponse) Set(val *UpdateFlexibleGpuResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFlexibleGpuResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFlexibleGpuResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFlexibleGpuResponse(val *UpdateFlexibleGpuResponse) *NullableUpdateFlexibleGpuResponse {
	return &NullableUpdateFlexibleGpuResponse{value: val, isSet: true}
}

func (v NullableUpdateFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFlexibleGpuResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
