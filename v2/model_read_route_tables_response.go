/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadRouteTablesResponse struct for ReadRouteTablesResponse
type ReadRouteTablesResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more route tables.
	RouteTables *[]RouteTable `json:"RouteTables,omitempty"`
}

// NewReadRouteTablesResponse instantiates a new ReadRouteTablesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadRouteTablesResponse() *ReadRouteTablesResponse {
	this := ReadRouteTablesResponse{}
	return &this
}

// NewReadRouteTablesResponseWithDefaults instantiates a new ReadRouteTablesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadRouteTablesResponseWithDefaults() *ReadRouteTablesResponse {
	this := ReadRouteTablesResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadRouteTablesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRouteTablesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadRouteTablesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadRouteTablesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetRouteTables returns the RouteTables field value if set, zero value otherwise.
func (o *ReadRouteTablesResponse) GetRouteTables() []RouteTable {
	if o == nil || o.RouteTables == nil {
		var ret []RouteTable
		return ret
	}
	return *o.RouteTables
}

// GetRouteTablesOk returns a tuple with the RouteTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRouteTablesResponse) GetRouteTablesOk() (*[]RouteTable, bool) {
	if o == nil || o.RouteTables == nil {
		return nil, false
	}
	return o.RouteTables, true
}

// HasRouteTables returns a boolean if a field has been set.
func (o *ReadRouteTablesResponse) HasRouteTables() bool {
	if o != nil && o.RouteTables != nil {
		return true
	}

	return false
}

// SetRouteTables gets a reference to the given []RouteTable and assigns it to the RouteTables field.
func (o *ReadRouteTablesResponse) SetRouteTables(v []RouteTable) {
	o.RouteTables = &v
}

func (o ReadRouteTablesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.RouteTables != nil {
		toSerialize["RouteTables"] = o.RouteTables
	}
	return json.Marshal(toSerialize)
}

type NullableReadRouteTablesResponse struct {
	value *ReadRouteTablesResponse
	isSet bool
}

func (v NullableReadRouteTablesResponse) Get() *ReadRouteTablesResponse {
	return v.value
}

func (v *NullableReadRouteTablesResponse) Set(val *ReadRouteTablesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadRouteTablesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadRouteTablesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadRouteTablesResponse(val *ReadRouteTablesResponse) *NullableReadRouteTablesResponse {
	return &NullableReadRouteTablesResponse{value: val, isSet: true}
}

func (v NullableReadRouteTablesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadRouteTablesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
