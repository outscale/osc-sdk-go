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

// UpdateNetAccessPointResponse struct for UpdateNetAccessPointResponse
type UpdateNetAccessPointResponse struct {
	NetAccessPoint  *NetAccessPoint  `json:"NetAccessPoint,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateNetAccessPointResponse instantiates a new UpdateNetAccessPointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetAccessPointResponse() *UpdateNetAccessPointResponse {
	this := UpdateNetAccessPointResponse{}
	return &this
}

// NewUpdateNetAccessPointResponseWithDefaults instantiates a new UpdateNetAccessPointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetAccessPointResponseWithDefaults() *UpdateNetAccessPointResponse {
	this := UpdateNetAccessPointResponse{}
	return &this
}

// GetNetAccessPoint returns the NetAccessPoint field value if set, zero value otherwise.
func (o *UpdateNetAccessPointResponse) GetNetAccessPoint() NetAccessPoint {
	if o == nil || o.NetAccessPoint == nil {
		var ret NetAccessPoint
		return ret
	}
	return *o.NetAccessPoint
}

// GetNetAccessPointOk returns a tuple with the NetAccessPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetAccessPointResponse) GetNetAccessPointOk() (*NetAccessPoint, bool) {
	if o == nil || o.NetAccessPoint == nil {
		return nil, false
	}
	return o.NetAccessPoint, true
}

// HasNetAccessPoint returns a boolean if a field has been set.
func (o *UpdateNetAccessPointResponse) HasNetAccessPoint() bool {
	if o != nil && o.NetAccessPoint != nil {
		return true
	}

	return false
}

// SetNetAccessPoint gets a reference to the given NetAccessPoint and assigns it to the NetAccessPoint field.
func (o *UpdateNetAccessPointResponse) SetNetAccessPoint(v NetAccessPoint) {
	o.NetAccessPoint = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateNetAccessPointResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetAccessPointResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateNetAccessPointResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateNetAccessPointResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateNetAccessPointResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetAccessPoint != nil {
		toSerialize["NetAccessPoint"] = o.NetAccessPoint
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetAccessPointResponse struct {
	value *UpdateNetAccessPointResponse
	isSet bool
}

func (v NullableUpdateNetAccessPointResponse) Get() *UpdateNetAccessPointResponse {
	return v.value
}

func (v *NullableUpdateNetAccessPointResponse) Set(val *UpdateNetAccessPointResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetAccessPointResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetAccessPointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetAccessPointResponse(val *UpdateNetAccessPointResponse) *NullableUpdateNetAccessPointResponse {
	return &NullableUpdateNetAccessPointResponse{value: val, isSet: true}
}

func (v NullableUpdateNetAccessPointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetAccessPointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
