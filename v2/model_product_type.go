/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ProductType Information about the product type.
type ProductType struct {
	// The description of the product type.
	Description *string `json:"Description,omitempty"`
	// The ID of the product type.
	ProductTypeId *string `json:"ProductTypeId,omitempty"`
	// The vendor of the product type.
	Vendor *string `json:"Vendor,omitempty"`
}

// NewProductType instantiates a new ProductType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductType() *ProductType {
	this := ProductType{}
	return &this
}

// NewProductTypeWithDefaults instantiates a new ProductType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTypeWithDefaults() *ProductType {
	this := ProductType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductType) SetDescription(v string) {
	o.Description = &v
}

// GetProductTypeId returns the ProductTypeId field value if set, zero value otherwise.
func (o *ProductType) GetProductTypeId() string {
	if o == nil || o.ProductTypeId == nil {
		var ret string
		return ret
	}
	return *o.ProductTypeId
}

// GetProductTypeIdOk returns a tuple with the ProductTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductType) GetProductTypeIdOk() (*string, bool) {
	if o == nil || o.ProductTypeId == nil {
		return nil, false
	}
	return o.ProductTypeId, true
}

// HasProductTypeId returns a boolean if a field has been set.
func (o *ProductType) HasProductTypeId() bool {
	if o != nil && o.ProductTypeId != nil {
		return true
	}

	return false
}

// SetProductTypeId gets a reference to the given string and assigns it to the ProductTypeId field.
func (o *ProductType) SetProductTypeId(v string) {
	o.ProductTypeId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ProductType) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductType) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ProductType) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *ProductType) SetVendor(v string) {
	o.Vendor = &v
}

func (o ProductType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ProductTypeId != nil {
		toSerialize["ProductTypeId"] = o.ProductTypeId
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableProductType struct {
	value *ProductType
	isSet bool
}

func (v NullableProductType) Get() *ProductType {
	return v.value
}

func (v *NullableProductType) Set(val *ProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductType(val *ProductType) *NullableProductType {
	return &NullableProductType{value: val, isSet: true}
}

func (v NullableProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
