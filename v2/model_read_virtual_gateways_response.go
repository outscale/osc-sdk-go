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

// ReadVirtualGatewaysResponse struct for ReadVirtualGatewaysResponse
type ReadVirtualGatewaysResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more virtual gateways.
	VirtualGateways *[]VirtualGateway `json:"VirtualGateways,omitempty"`
}

// NewReadVirtualGatewaysResponse instantiates a new ReadVirtualGatewaysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVirtualGatewaysResponse() *ReadVirtualGatewaysResponse {
	this := ReadVirtualGatewaysResponse{}
	return &this
}

// NewReadVirtualGatewaysResponseWithDefaults instantiates a new ReadVirtualGatewaysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVirtualGatewaysResponseWithDefaults() *ReadVirtualGatewaysResponse {
	this := ReadVirtualGatewaysResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVirtualGatewaysResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVirtualGatewaysResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVirtualGatewaysResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVirtualGatewaysResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVirtualGateways returns the VirtualGateways field value if set, zero value otherwise.
func (o *ReadVirtualGatewaysResponse) GetVirtualGateways() []VirtualGateway {
	if o == nil || o.VirtualGateways == nil {
		var ret []VirtualGateway
		return ret
	}
	return *o.VirtualGateways
}

// GetVirtualGatewaysOk returns a tuple with the VirtualGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVirtualGatewaysResponse) GetVirtualGatewaysOk() (*[]VirtualGateway, bool) {
	if o == nil || o.VirtualGateways == nil {
		return nil, false
	}
	return o.VirtualGateways, true
}

// HasVirtualGateways returns a boolean if a field has been set.
func (o *ReadVirtualGatewaysResponse) HasVirtualGateways() bool {
	if o != nil && o.VirtualGateways != nil {
		return true
	}

	return false
}

// SetVirtualGateways gets a reference to the given []VirtualGateway and assigns it to the VirtualGateways field.
func (o *ReadVirtualGatewaysResponse) SetVirtualGateways(v []VirtualGateway) {
	o.VirtualGateways = &v
}

func (o ReadVirtualGatewaysResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.VirtualGateways != nil {
		toSerialize["VirtualGateways"] = o.VirtualGateways
	}
	return json.Marshal(toSerialize)
}

type NullableReadVirtualGatewaysResponse struct {
	value *ReadVirtualGatewaysResponse
	isSet bool
}

func (v NullableReadVirtualGatewaysResponse) Get() *ReadVirtualGatewaysResponse {
	return v.value
}

func (v *NullableReadVirtualGatewaysResponse) Set(val *ReadVirtualGatewaysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVirtualGatewaysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVirtualGatewaysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVirtualGatewaysResponse(val *ReadVirtualGatewaysResponse) *NullableReadVirtualGatewaysResponse {
	return &NullableReadVirtualGatewaysResponse{value: val, isSet: true}
}

func (v NullableReadVirtualGatewaysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVirtualGatewaysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
