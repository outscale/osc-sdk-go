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

// ReadVmTemplatesResponse struct for ReadVmTemplatesResponse
type ReadVmTemplatesResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more VM templates.
	VmTemplates *[]VmTemplate `json:"VmTemplates,omitempty"`
}

// NewReadVmTemplatesResponse instantiates a new ReadVmTemplatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVmTemplatesResponse() *ReadVmTemplatesResponse {
	this := ReadVmTemplatesResponse{}
	return &this
}

// NewReadVmTemplatesResponseWithDefaults instantiates a new ReadVmTemplatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVmTemplatesResponseWithDefaults() *ReadVmTemplatesResponse {
	this := ReadVmTemplatesResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVmTemplatesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmTemplatesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVmTemplatesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVmTemplatesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVmTemplates returns the VmTemplates field value if set, zero value otherwise.
func (o *ReadVmTemplatesResponse) GetVmTemplates() []VmTemplate {
	if o == nil || o.VmTemplates == nil {
		var ret []VmTemplate
		return ret
	}
	return *o.VmTemplates
}

// GetVmTemplatesOk returns a tuple with the VmTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmTemplatesResponse) GetVmTemplatesOk() (*[]VmTemplate, bool) {
	if o == nil || o.VmTemplates == nil {
		return nil, false
	}
	return o.VmTemplates, true
}

// HasVmTemplates returns a boolean if a field has been set.
func (o *ReadVmTemplatesResponse) HasVmTemplates() bool {
	if o != nil && o.VmTemplates != nil {
		return true
	}

	return false
}

// SetVmTemplates gets a reference to the given []VmTemplate and assigns it to the VmTemplates field.
func (o *ReadVmTemplatesResponse) SetVmTemplates(v []VmTemplate) {
	o.VmTemplates = &v
}

func (o ReadVmTemplatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.VmTemplates != nil {
		toSerialize["VmTemplates"] = o.VmTemplates
	}
	return json.Marshal(toSerialize)
}

type NullableReadVmTemplatesResponse struct {
	value *ReadVmTemplatesResponse
	isSet bool
}

func (v NullableReadVmTemplatesResponse) Get() *ReadVmTemplatesResponse {
	return v.value
}

func (v *NullableReadVmTemplatesResponse) Set(val *ReadVmTemplatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVmTemplatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVmTemplatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVmTemplatesResponse(val *ReadVmTemplatesResponse) *NullableReadVmTemplatesResponse {
	return &NullableReadVmTemplatesResponse{value: val, isSet: true}
}

func (v NullableReadVmTemplatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVmTemplatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}