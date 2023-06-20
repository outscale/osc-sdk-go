/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.27
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateNetAccessPointResponse struct for CreateNetAccessPointResponse
type CreateNetAccessPointResponse struct {
	NetAccessPoint  *NetAccessPoint  `json:"NetAccessPoint,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateNetAccessPointResponse instantiates a new CreateNetAccessPointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetAccessPointResponse() *CreateNetAccessPointResponse {
	this := CreateNetAccessPointResponse{}
	return &this
}

// NewCreateNetAccessPointResponseWithDefaults instantiates a new CreateNetAccessPointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetAccessPointResponseWithDefaults() *CreateNetAccessPointResponse {
	this := CreateNetAccessPointResponse{}
	return &this
}

// GetNetAccessPoint returns the NetAccessPoint field value if set, zero value otherwise.
func (o *CreateNetAccessPointResponse) GetNetAccessPoint() NetAccessPoint {
	if o == nil || o.NetAccessPoint == nil {
		var ret NetAccessPoint
		return ret
	}
	return *o.NetAccessPoint
}

// GetNetAccessPointOk returns a tuple with the NetAccessPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetAccessPointResponse) GetNetAccessPointOk() (*NetAccessPoint, bool) {
	if o == nil || o.NetAccessPoint == nil {
		return nil, false
	}
	return o.NetAccessPoint, true
}

// HasNetAccessPoint returns a boolean if a field has been set.
func (o *CreateNetAccessPointResponse) HasNetAccessPoint() bool {
	if o != nil && o.NetAccessPoint != nil {
		return true
	}

	return false
}

// SetNetAccessPoint gets a reference to the given NetAccessPoint and assigns it to the NetAccessPoint field.
func (o *CreateNetAccessPointResponse) SetNetAccessPoint(v NetAccessPoint) {
	o.NetAccessPoint = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateNetAccessPointResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetAccessPointResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateNetAccessPointResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateNetAccessPointResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateNetAccessPointResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetAccessPoint != nil {
		toSerialize["NetAccessPoint"] = o.NetAccessPoint
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetAccessPointResponse struct {
	value *CreateNetAccessPointResponse
	isSet bool
}

func (v NullableCreateNetAccessPointResponse) Get() *CreateNetAccessPointResponse {
	return v.value
}

func (v *NullableCreateNetAccessPointResponse) Set(val *CreateNetAccessPointResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetAccessPointResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetAccessPointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetAccessPointResponse(val *CreateNetAccessPointResponse) *NullableCreateNetAccessPointResponse {
	return &NullableCreateNetAccessPointResponse{value: val, isSet: true}
}

func (v NullableCreateNetAccessPointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetAccessPointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
