/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadPublicIpRangesResponse struct for ReadPublicIpRangesResponse
type ReadPublicIpRangesResponse struct {
	// The list of public IPv4 addresses used in the Region, in CIDR notation.
	PublicIps       *[]string        `json:"PublicIps,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadPublicIpRangesResponse instantiates a new ReadPublicIpRangesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadPublicIpRangesResponse() *ReadPublicIpRangesResponse {
	this := ReadPublicIpRangesResponse{}
	return &this
}

// NewReadPublicIpRangesResponseWithDefaults instantiates a new ReadPublicIpRangesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadPublicIpRangesResponseWithDefaults() *ReadPublicIpRangesResponse {
	this := ReadPublicIpRangesResponse{}
	return &this
}

// GetPublicIps returns the PublicIps field value if set, zero value otherwise.
func (o *ReadPublicIpRangesResponse) GetPublicIps() []string {
	if o == nil || o.PublicIps == nil {
		var ret []string
		return ret
	}
	return *o.PublicIps
}

// GetPublicIpsOk returns a tuple with the PublicIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadPublicIpRangesResponse) GetPublicIpsOk() (*[]string, bool) {
	if o == nil || o.PublicIps == nil {
		return nil, false
	}
	return o.PublicIps, true
}

// HasPublicIps returns a boolean if a field has been set.
func (o *ReadPublicIpRangesResponse) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}

// SetPublicIps gets a reference to the given []string and assigns it to the PublicIps field.
func (o *ReadPublicIpRangesResponse) SetPublicIps(v []string) {
	o.PublicIps = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadPublicIpRangesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadPublicIpRangesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadPublicIpRangesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadPublicIpRangesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadPublicIpRangesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicIps != nil {
		toSerialize["PublicIps"] = o.PublicIps
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadPublicIpRangesResponse struct {
	value *ReadPublicIpRangesResponse
	isSet bool
}

func (v NullableReadPublicIpRangesResponse) Get() *ReadPublicIpRangesResponse {
	return v.value
}

func (v *NullableReadPublicIpRangesResponse) Set(val *ReadPublicIpRangesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadPublicIpRangesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadPublicIpRangesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadPublicIpRangesResponse(val *ReadPublicIpRangesResponse) *NullableReadPublicIpRangesResponse {
	return &NullableReadPublicIpRangesResponse{value: val, isSet: true}
}

func (v NullableReadPublicIpRangesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadPublicIpRangesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
