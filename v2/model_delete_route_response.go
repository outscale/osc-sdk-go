/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteRouteResponse struct for DeleteRouteResponse
type DeleteRouteResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	RouteTable *RouteTable `json:"RouteTable,omitempty"`
}

// NewDeleteRouteResponse instantiates a new DeleteRouteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRouteResponse() *DeleteRouteResponse {
	this := DeleteRouteResponse{}
	return &this
}

// NewDeleteRouteResponseWithDefaults instantiates a new DeleteRouteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRouteResponseWithDefaults() *DeleteRouteResponse {
	this := DeleteRouteResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteRouteResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRouteResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteRouteResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteRouteResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetRouteTable returns the RouteTable field value if set, zero value otherwise.
func (o *DeleteRouteResponse) GetRouteTable() RouteTable {
	if o == nil || o.RouteTable == nil {
		var ret RouteTable
		return ret
	}
	return *o.RouteTable
}

// GetRouteTableOk returns a tuple with the RouteTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRouteResponse) GetRouteTableOk() (*RouteTable, bool) {
	if o == nil || o.RouteTable == nil {
		return nil, false
	}
	return o.RouteTable, true
}

// HasRouteTable returns a boolean if a field has been set.
func (o *DeleteRouteResponse) HasRouteTable() bool {
	if o != nil && o.RouteTable != nil {
		return true
	}

	return false
}

// SetRouteTable gets a reference to the given RouteTable and assigns it to the RouteTable field.
func (o *DeleteRouteResponse) SetRouteTable(v RouteTable) {
	o.RouteTable = &v
}

func (o DeleteRouteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.RouteTable != nil {
		toSerialize["RouteTable"] = o.RouteTable
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteRouteResponse struct {
	value *DeleteRouteResponse
	isSet bool
}

func (v NullableDeleteRouteResponse) Get() *DeleteRouteResponse {
	return v.value
}

func (v *NullableDeleteRouteResponse) Set(val *DeleteRouteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRouteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRouteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRouteResponse(val *DeleteRouteResponse) *NullableDeleteRouteResponse {
	return &NullableDeleteRouteResponse{value: val, isSet: true}
}

func (v NullableDeleteRouteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRouteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


