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

// ReadSubnetsResponse struct for ReadSubnetsResponse
type ReadSubnetsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more Subnets.
	Subnets *[]Subnet `json:"Subnets,omitempty"`
}

// NewReadSubnetsResponse instantiates a new ReadSubnetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadSubnetsResponse() *ReadSubnetsResponse {
	this := ReadSubnetsResponse{}
	return &this
}

// NewReadSubnetsResponseWithDefaults instantiates a new ReadSubnetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadSubnetsResponseWithDefaults() *ReadSubnetsResponse {
	this := ReadSubnetsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadSubnetsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSubnetsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadSubnetsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadSubnetsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *ReadSubnetsResponse) GetSubnets() []Subnet {
	if o == nil || o.Subnets == nil {
		var ret []Subnet
		return ret
	}
	return *o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSubnetsResponse) GetSubnetsOk() (*[]Subnet, bool) {
	if o == nil || o.Subnets == nil {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *ReadSubnetsResponse) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []Subnet and assigns it to the Subnets field.
func (o *ReadSubnetsResponse) SetSubnets(v []Subnet) {
	o.Subnets = &v
}

func (o ReadSubnetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Subnets != nil {
		toSerialize["Subnets"] = o.Subnets
	}
	return json.Marshal(toSerialize)
}

type NullableReadSubnetsResponse struct {
	value *ReadSubnetsResponse
	isSet bool
}

func (v NullableReadSubnetsResponse) Get() *ReadSubnetsResponse {
	return v.value
}

func (v *NullableReadSubnetsResponse) Set(val *ReadSubnetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadSubnetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadSubnetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadSubnetsResponse(val *ReadSubnetsResponse) *NullableReadSubnetsResponse {
	return &NullableReadSubnetsResponse{value: val, isSet: true}
}

func (v NullableReadSubnetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadSubnetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
