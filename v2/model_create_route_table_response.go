/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateRouteTableResponse struct for CreateRouteTableResponse
type CreateRouteTableResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	RouteTable      *RouteTable      `json:"RouteTable,omitempty"`
}

// NewCreateRouteTableResponse instantiates a new CreateRouteTableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRouteTableResponse() *CreateRouteTableResponse {
	this := CreateRouteTableResponse{}
	return &this
}

// NewCreateRouteTableResponseWithDefaults instantiates a new CreateRouteTableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRouteTableResponseWithDefaults() *CreateRouteTableResponse {
	this := CreateRouteTableResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateRouteTableResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteTableResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateRouteTableResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateRouteTableResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetRouteTable returns the RouteTable field value if set, zero value otherwise.
func (o *CreateRouteTableResponse) GetRouteTable() RouteTable {
	if o == nil || o.RouteTable == nil {
		var ret RouteTable
		return ret
	}
	return *o.RouteTable
}

// GetRouteTableOk returns a tuple with the RouteTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteTableResponse) GetRouteTableOk() (*RouteTable, bool) {
	if o == nil || o.RouteTable == nil {
		return nil, false
	}
	return o.RouteTable, true
}

// HasRouteTable returns a boolean if a field has been set.
func (o *CreateRouteTableResponse) HasRouteTable() bool {
	if o != nil && o.RouteTable != nil {
		return true
	}

	return false
}

// SetRouteTable gets a reference to the given RouteTable and assigns it to the RouteTable field.
func (o *CreateRouteTableResponse) SetRouteTable(v RouteTable) {
	o.RouteTable = &v
}

func (o CreateRouteTableResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.RouteTable != nil {
		toSerialize["RouteTable"] = o.RouteTable
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRouteTableResponse struct {
	value *CreateRouteTableResponse
	isSet bool
}

func (v NullableCreateRouteTableResponse) Get() *CreateRouteTableResponse {
	return v.value
}

func (v *NullableCreateRouteTableResponse) Set(val *CreateRouteTableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRouteTableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRouteTableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRouteTableResponse(val *CreateRouteTableResponse) *NullableCreateRouteTableResponse {
	return &NullableCreateRouteTableResponse{value: val, isSet: true}
}

func (v NullableCreateRouteTableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRouteTableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
