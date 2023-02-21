/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateVirtualGatewayResponse struct for CreateVirtualGatewayResponse
type CreateVirtualGatewayResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	VirtualGateway  *VirtualGateway  `json:"VirtualGateway,omitempty"`
}

// NewCreateVirtualGatewayResponse instantiates a new CreateVirtualGatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVirtualGatewayResponse() *CreateVirtualGatewayResponse {
	this := CreateVirtualGatewayResponse{}
	return &this
}

// NewCreateVirtualGatewayResponseWithDefaults instantiates a new CreateVirtualGatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVirtualGatewayResponseWithDefaults() *CreateVirtualGatewayResponse {
	this := CreateVirtualGatewayResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateVirtualGatewayResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualGatewayResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateVirtualGatewayResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateVirtualGatewayResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVirtualGateway returns the VirtualGateway field value if set, zero value otherwise.
func (o *CreateVirtualGatewayResponse) GetVirtualGateway() VirtualGateway {
	if o == nil || o.VirtualGateway == nil {
		var ret VirtualGateway
		return ret
	}
	return *o.VirtualGateway
}

// GetVirtualGatewayOk returns a tuple with the VirtualGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualGatewayResponse) GetVirtualGatewayOk() (*VirtualGateway, bool) {
	if o == nil || o.VirtualGateway == nil {
		return nil, false
	}
	return o.VirtualGateway, true
}

// HasVirtualGateway returns a boolean if a field has been set.
func (o *CreateVirtualGatewayResponse) HasVirtualGateway() bool {
	if o != nil && o.VirtualGateway != nil {
		return true
	}

	return false
}

// SetVirtualGateway gets a reference to the given VirtualGateway and assigns it to the VirtualGateway field.
func (o *CreateVirtualGatewayResponse) SetVirtualGateway(v VirtualGateway) {
	o.VirtualGateway = &v
}

func (o CreateVirtualGatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.VirtualGateway != nil {
		toSerialize["VirtualGateway"] = o.VirtualGateway
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVirtualGatewayResponse struct {
	value *CreateVirtualGatewayResponse
	isSet bool
}

func (v NullableCreateVirtualGatewayResponse) Get() *CreateVirtualGatewayResponse {
	return v.value
}

func (v *NullableCreateVirtualGatewayResponse) Set(val *CreateVirtualGatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVirtualGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVirtualGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVirtualGatewayResponse(val *CreateVirtualGatewayResponse) *NullableCreateVirtualGatewayResponse {
	return &NullableCreateVirtualGatewayResponse{value: val, isSet: true}
}

func (v NullableCreateVirtualGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVirtualGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
