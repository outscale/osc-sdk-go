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

// ReadDhcpOptionsResponse struct for ReadDhcpOptionsResponse
type ReadDhcpOptionsResponse struct {
	// Information about one or more DHCP options sets.
	DhcpOptionsSets *[]DhcpOptionsSet `json:"DhcpOptionsSets,omitempty"`
	ResponseContext *ResponseContext  `json:"ResponseContext,omitempty"`
}

// NewReadDhcpOptionsResponse instantiates a new ReadDhcpOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadDhcpOptionsResponse() *ReadDhcpOptionsResponse {
	this := ReadDhcpOptionsResponse{}
	return &this
}

// NewReadDhcpOptionsResponseWithDefaults instantiates a new ReadDhcpOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadDhcpOptionsResponseWithDefaults() *ReadDhcpOptionsResponse {
	this := ReadDhcpOptionsResponse{}
	return &this
}

// GetDhcpOptionsSets returns the DhcpOptionsSets field value if set, zero value otherwise.
func (o *ReadDhcpOptionsResponse) GetDhcpOptionsSets() []DhcpOptionsSet {
	if o == nil || o.DhcpOptionsSets == nil {
		var ret []DhcpOptionsSet
		return ret
	}
	return *o.DhcpOptionsSets
}

// GetDhcpOptionsSetsOk returns a tuple with the DhcpOptionsSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadDhcpOptionsResponse) GetDhcpOptionsSetsOk() (*[]DhcpOptionsSet, bool) {
	if o == nil || o.DhcpOptionsSets == nil {
		return nil, false
	}
	return o.DhcpOptionsSets, true
}

// HasDhcpOptionsSets returns a boolean if a field has been set.
func (o *ReadDhcpOptionsResponse) HasDhcpOptionsSets() bool {
	if o != nil && o.DhcpOptionsSets != nil {
		return true
	}

	return false
}

// SetDhcpOptionsSets gets a reference to the given []DhcpOptionsSet and assigns it to the DhcpOptionsSets field.
func (o *ReadDhcpOptionsResponse) SetDhcpOptionsSets(v []DhcpOptionsSet) {
	o.DhcpOptionsSets = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadDhcpOptionsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadDhcpOptionsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadDhcpOptionsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadDhcpOptionsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadDhcpOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DhcpOptionsSets != nil {
		toSerialize["DhcpOptionsSets"] = o.DhcpOptionsSets
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadDhcpOptionsResponse struct {
	value *ReadDhcpOptionsResponse
	isSet bool
}

func (v NullableReadDhcpOptionsResponse) Get() *ReadDhcpOptionsResponse {
	return v.value
}

func (v *NullableReadDhcpOptionsResponse) Set(val *ReadDhcpOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadDhcpOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadDhcpOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadDhcpOptionsResponse(val *ReadDhcpOptionsResponse) *NullableReadDhcpOptionsResponse {
	return &NullableReadDhcpOptionsResponse{value: val, isSet: true}
}

func (v NullableReadDhcpOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadDhcpOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
