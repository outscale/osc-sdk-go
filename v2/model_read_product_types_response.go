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

// ReadProductTypesResponse struct for ReadProductTypesResponse
type ReadProductTypesResponse struct {
	// Information about one or more product types.
	ProductTypes    *[]ProductType   `json:"ProductTypes,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadProductTypesResponse instantiates a new ReadProductTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadProductTypesResponse() *ReadProductTypesResponse {
	this := ReadProductTypesResponse{}
	return &this
}

// NewReadProductTypesResponseWithDefaults instantiates a new ReadProductTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadProductTypesResponseWithDefaults() *ReadProductTypesResponse {
	this := ReadProductTypesResponse{}
	return &this
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *ReadProductTypesResponse) GetProductTypes() []ProductType {
	if o == nil || o.ProductTypes == nil {
		var ret []ProductType
		return ret
	}
	return *o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadProductTypesResponse) GetProductTypesOk() (*[]ProductType, bool) {
	if o == nil || o.ProductTypes == nil {
		return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *ReadProductTypesResponse) HasProductTypes() bool {
	if o != nil && o.ProductTypes != nil {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []ProductType and assigns it to the ProductTypes field.
func (o *ReadProductTypesResponse) SetProductTypes(v []ProductType) {
	o.ProductTypes = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadProductTypesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadProductTypesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadProductTypesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadProductTypesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadProductTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductTypes != nil {
		toSerialize["ProductTypes"] = o.ProductTypes
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadProductTypesResponse struct {
	value *ReadProductTypesResponse
	isSet bool
}

func (v NullableReadProductTypesResponse) Get() *ReadProductTypesResponse {
	return v.value
}

func (v *NullableReadProductTypesResponse) Set(val *ReadProductTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadProductTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadProductTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadProductTypesResponse(val *ReadProductTypesResponse) *NullableReadProductTypesResponse {
	return &NullableReadProductTypesResponse{value: val, isSet: true}
}

func (v NullableReadProductTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadProductTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
