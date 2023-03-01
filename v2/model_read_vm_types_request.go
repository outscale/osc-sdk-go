/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadVmTypesRequest struct for ReadVmTypesRequest
type ReadVmTypesRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool          `json:"DryRun,omitempty"`
	Filters *FiltersVmType `json:"Filters,omitempty"`
}

// NewReadVmTypesRequest instantiates a new ReadVmTypesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVmTypesRequest() *ReadVmTypesRequest {
	this := ReadVmTypesRequest{}
	return &this
}

// NewReadVmTypesRequestWithDefaults instantiates a new ReadVmTypesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVmTypesRequestWithDefaults() *ReadVmTypesRequest {
	this := ReadVmTypesRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadVmTypesRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmTypesRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadVmTypesRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadVmTypesRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadVmTypesRequest) GetFilters() FiltersVmType {
	if o == nil || o.Filters == nil {
		var ret FiltersVmType
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmTypesRequest) GetFiltersOk() (*FiltersVmType, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadVmTypesRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersVmType and assigns it to the Filters field.
func (o *ReadVmTypesRequest) SetFilters(v FiltersVmType) {
	o.Filters = &v
}

func (o ReadVmTypesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadVmTypesRequest struct {
	value *ReadVmTypesRequest
	isSet bool
}

func (v NullableReadVmTypesRequest) Get() *ReadVmTypesRequest {
	return v.value
}

func (v *NullableReadVmTypesRequest) Set(val *ReadVmTypesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVmTypesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVmTypesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVmTypesRequest(val *ReadVmTypesRequest) *NullableReadVmTypesRequest {
	return &NullableReadVmTypesRequest{value: val, isSet: true}
}

func (v NullableReadVmTypesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVmTypesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
