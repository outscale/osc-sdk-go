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

// FiltersProductType One or more filters.
type FiltersProductType struct {
	// The IDs of the product types.
	ProductTypeIds *[]string `json:"ProductTypeIds,omitempty"`
}

// NewFiltersProductType instantiates a new FiltersProductType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersProductType() *FiltersProductType {
	this := FiltersProductType{}
	return &this
}

// NewFiltersProductTypeWithDefaults instantiates a new FiltersProductType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersProductTypeWithDefaults() *FiltersProductType {
	this := FiltersProductType{}
	return &this
}

// GetProductTypeIds returns the ProductTypeIds field value if set, zero value otherwise.
func (o *FiltersProductType) GetProductTypeIds() []string {
	if o == nil || o.ProductTypeIds == nil {
		var ret []string
		return ret
	}
	return *o.ProductTypeIds
}

// GetProductTypeIdsOk returns a tuple with the ProductTypeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersProductType) GetProductTypeIdsOk() (*[]string, bool) {
	if o == nil || o.ProductTypeIds == nil {
		return nil, false
	}
	return o.ProductTypeIds, true
}

// HasProductTypeIds returns a boolean if a field has been set.
func (o *FiltersProductType) HasProductTypeIds() bool {
	if o != nil && o.ProductTypeIds != nil {
		return true
	}

	return false
}

// SetProductTypeIds gets a reference to the given []string and assigns it to the ProductTypeIds field.
func (o *FiltersProductType) SetProductTypeIds(v []string) {
	o.ProductTypeIds = &v
}

func (o FiltersProductType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductTypeIds != nil {
		toSerialize["ProductTypeIds"] = o.ProductTypeIds
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersProductType struct {
	value *FiltersProductType
	isSet bool
}

func (v NullableFiltersProductType) Get() *FiltersProductType {
	return v.value
}

func (v *NullableFiltersProductType) Set(val *FiltersProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersProductType(val *FiltersProductType) *NullableFiltersProductType {
	return &NullableFiltersProductType{value: val, isSet: true}
}

func (v NullableFiltersProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
