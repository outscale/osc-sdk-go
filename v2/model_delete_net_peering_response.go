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

// DeleteNetPeeringResponse struct for DeleteNetPeeringResponse
type DeleteNetPeeringResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewDeleteNetPeeringResponse instantiates a new DeleteNetPeeringResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteNetPeeringResponse() *DeleteNetPeeringResponse {
	this := DeleteNetPeeringResponse{}
	return &this
}

// NewDeleteNetPeeringResponseWithDefaults instantiates a new DeleteNetPeeringResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteNetPeeringResponseWithDefaults() *DeleteNetPeeringResponse {
	this := DeleteNetPeeringResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteNetPeeringResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteNetPeeringResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteNetPeeringResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteNetPeeringResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o DeleteNetPeeringResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteNetPeeringResponse struct {
	value *DeleteNetPeeringResponse
	isSet bool
}

func (v NullableDeleteNetPeeringResponse) Get() *DeleteNetPeeringResponse {
	return v.value
}

func (v *NullableDeleteNetPeeringResponse) Set(val *DeleteNetPeeringResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteNetPeeringResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteNetPeeringResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteNetPeeringResponse(val *DeleteNetPeeringResponse) *NullableDeleteNetPeeringResponse {
	return &NullableDeleteNetPeeringResponse{value: val, isSet: true}
}

func (v NullableDeleteNetPeeringResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteNetPeeringResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
