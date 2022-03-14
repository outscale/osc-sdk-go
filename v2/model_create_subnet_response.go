/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateSubnetResponse struct for CreateSubnetResponse
type CreateSubnetResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	Subnet          *Subnet          `json:"Subnet,omitempty"`
}

// NewCreateSubnetResponse instantiates a new CreateSubnetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubnetResponse() *CreateSubnetResponse {
	this := CreateSubnetResponse{}
	return &this
}

// NewCreateSubnetResponseWithDefaults instantiates a new CreateSubnetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubnetResponseWithDefaults() *CreateSubnetResponse {
	this := CreateSubnetResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateSubnetResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubnetResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateSubnetResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateSubnetResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *CreateSubnetResponse) GetSubnet() Subnet {
	if o == nil || o.Subnet == nil {
		var ret Subnet
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubnetResponse) GetSubnetOk() (*Subnet, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *CreateSubnetResponse) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given Subnet and assigns it to the Subnet field.
func (o *CreateSubnetResponse) SetSubnet(v Subnet) {
	o.Subnet = &v
}

func (o CreateSubnetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Subnet != nil {
		toSerialize["Subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSubnetResponse struct {
	value *CreateSubnetResponse
	isSet bool
}

func (v NullableCreateSubnetResponse) Get() *CreateSubnetResponse {
	return v.value
}

func (v *NullableCreateSubnetResponse) Set(val *CreateSubnetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubnetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubnetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubnetResponse(val *CreateSubnetResponse) *NullableCreateSubnetResponse {
	return &NullableCreateSubnetResponse{value: val, isSet: true}
}

func (v NullableCreateSubnetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubnetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
