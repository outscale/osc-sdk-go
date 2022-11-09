/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateDirectLinkInterfaceResponse struct for CreateDirectLinkInterfaceResponse
type CreateDirectLinkInterfaceResponse struct {
	DirectLinkInterface *DirectLinkInterfaces `json:"DirectLinkInterface,omitempty"`
	ResponseContext     *ResponseContext      `json:"ResponseContext,omitempty"`
}

// NewCreateDirectLinkInterfaceResponse instantiates a new CreateDirectLinkInterfaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDirectLinkInterfaceResponse() *CreateDirectLinkInterfaceResponse {
	this := CreateDirectLinkInterfaceResponse{}
	return &this
}

// NewCreateDirectLinkInterfaceResponseWithDefaults instantiates a new CreateDirectLinkInterfaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDirectLinkInterfaceResponseWithDefaults() *CreateDirectLinkInterfaceResponse {
	this := CreateDirectLinkInterfaceResponse{}
	return &this
}

// GetDirectLinkInterface returns the DirectLinkInterface field value if set, zero value otherwise.
func (o *CreateDirectLinkInterfaceResponse) GetDirectLinkInterface() DirectLinkInterfaces {
	if o == nil || o.DirectLinkInterface == nil {
		var ret DirectLinkInterfaces
		return ret
	}
	return *o.DirectLinkInterface
}

// GetDirectLinkInterfaceOk returns a tuple with the DirectLinkInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkInterfaceResponse) GetDirectLinkInterfaceOk() (*DirectLinkInterfaces, bool) {
	if o == nil || o.DirectLinkInterface == nil {
		return nil, false
	}
	return o.DirectLinkInterface, true
}

// HasDirectLinkInterface returns a boolean if a field has been set.
func (o *CreateDirectLinkInterfaceResponse) HasDirectLinkInterface() bool {
	if o != nil && o.DirectLinkInterface != nil {
		return true
	}

	return false
}

// SetDirectLinkInterface gets a reference to the given DirectLinkInterfaces and assigns it to the DirectLinkInterface field.
func (o *CreateDirectLinkInterfaceResponse) SetDirectLinkInterface(v DirectLinkInterfaces) {
	o.DirectLinkInterface = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateDirectLinkInterfaceResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkInterfaceResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateDirectLinkInterfaceResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateDirectLinkInterfaceResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateDirectLinkInterfaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DirectLinkInterface != nil {
		toSerialize["DirectLinkInterface"] = o.DirectLinkInterface
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDirectLinkInterfaceResponse struct {
	value *CreateDirectLinkInterfaceResponse
	isSet bool
}

func (v NullableCreateDirectLinkInterfaceResponse) Get() *CreateDirectLinkInterfaceResponse {
	return v.value
}

func (v *NullableCreateDirectLinkInterfaceResponse) Set(val *CreateDirectLinkInterfaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDirectLinkInterfaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDirectLinkInterfaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDirectLinkInterfaceResponse(val *CreateDirectLinkInterfaceResponse) *NullableCreateDirectLinkInterfaceResponse {
	return &NullableCreateDirectLinkInterfaceResponse{value: val, isSet: true}
}

func (v NullableCreateDirectLinkInterfaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDirectLinkInterfaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
