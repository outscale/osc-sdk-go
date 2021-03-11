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

// UpdateSubnetResponse struct for UpdateSubnetResponse
type UpdateSubnetResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	Subnet          *Subnet          `json:"Subnet,omitempty"`
}

// NewUpdateSubnetResponse instantiates a new UpdateSubnetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubnetResponse() *UpdateSubnetResponse {
	this := UpdateSubnetResponse{}
	return &this
}

// NewUpdateSubnetResponseWithDefaults instantiates a new UpdateSubnetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubnetResponseWithDefaults() *UpdateSubnetResponse {
	this := UpdateSubnetResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateSubnetResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubnetResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateSubnetResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateSubnetResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *UpdateSubnetResponse) GetSubnet() Subnet {
	if o == nil || o.Subnet == nil {
		var ret Subnet
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubnetResponse) GetSubnetOk() (*Subnet, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *UpdateSubnetResponse) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given Subnet and assigns it to the Subnet field.
func (o *UpdateSubnetResponse) SetSubnet(v Subnet) {
	o.Subnet = &v
}

func (o UpdateSubnetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Subnet != nil {
		toSerialize["Subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSubnetResponse struct {
	value *UpdateSubnetResponse
	isSet bool
}

func (v NullableUpdateSubnetResponse) Get() *UpdateSubnetResponse {
	return v.value
}

func (v *NullableUpdateSubnetResponse) Set(val *UpdateSubnetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubnetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubnetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubnetResponse(val *UpdateSubnetResponse) *NullableUpdateSubnetResponse {
	return &NullableUpdateSubnetResponse{value: val, isSet: true}
}

func (v NullableUpdateSubnetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubnetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
