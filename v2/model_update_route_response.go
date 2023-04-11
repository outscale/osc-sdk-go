/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateRouteResponse struct for UpdateRouteResponse
type UpdateRouteResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	RouteTable      *RouteTable      `json:"RouteTable,omitempty"`
}

// NewUpdateRouteResponse instantiates a new UpdateRouteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRouteResponse() *UpdateRouteResponse {
	this := UpdateRouteResponse{}
	return &this
}

// NewUpdateRouteResponseWithDefaults instantiates a new UpdateRouteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRouteResponseWithDefaults() *UpdateRouteResponse {
	this := UpdateRouteResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateRouteResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateRouteResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateRouteResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetRouteTable returns the RouteTable field value if set, zero value otherwise.
func (o *UpdateRouteResponse) GetRouteTable() RouteTable {
	if o == nil || o.RouteTable == nil {
		var ret RouteTable
		return ret
	}
	return *o.RouteTable
}

// GetRouteTableOk returns a tuple with the RouteTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteResponse) GetRouteTableOk() (*RouteTable, bool) {
	if o == nil || o.RouteTable == nil {
		return nil, false
	}
	return o.RouteTable, true
}

// HasRouteTable returns a boolean if a field has been set.
func (o *UpdateRouteResponse) HasRouteTable() bool {
	if o != nil && o.RouteTable != nil {
		return true
	}

	return false
}

// SetRouteTable gets a reference to the given RouteTable and assigns it to the RouteTable field.
func (o *UpdateRouteResponse) SetRouteTable(v RouteTable) {
	o.RouteTable = &v
}

func (o UpdateRouteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.RouteTable != nil {
		toSerialize["RouteTable"] = o.RouteTable
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRouteResponse struct {
	value *UpdateRouteResponse
	isSet bool
}

func (v NullableUpdateRouteResponse) Get() *UpdateRouteResponse {
	return v.value
}

func (v *NullableUpdateRouteResponse) Set(val *UpdateRouteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRouteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRouteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRouteResponse(val *UpdateRouteResponse) *NullableUpdateRouteResponse {
	return &NullableUpdateRouteResponse{value: val, isSet: true}
}

func (v NullableUpdateRouteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRouteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
