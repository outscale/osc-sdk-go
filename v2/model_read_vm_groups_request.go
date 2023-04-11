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

// ReadVmGroupsRequest struct for ReadVmGroupsRequest
type ReadVmGroupsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool           `json:"DryRun,omitempty"`
	Filters *FiltersVmGroup `json:"Filters,omitempty"`
}

// NewReadVmGroupsRequest instantiates a new ReadVmGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVmGroupsRequest() *ReadVmGroupsRequest {
	this := ReadVmGroupsRequest{}
	return &this
}

// NewReadVmGroupsRequestWithDefaults instantiates a new ReadVmGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVmGroupsRequestWithDefaults() *ReadVmGroupsRequest {
	this := ReadVmGroupsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadVmGroupsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmGroupsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadVmGroupsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadVmGroupsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadVmGroupsRequest) GetFilters() FiltersVmGroup {
	if o == nil || o.Filters == nil {
		var ret FiltersVmGroup
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmGroupsRequest) GetFiltersOk() (*FiltersVmGroup, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadVmGroupsRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersVmGroup and assigns it to the Filters field.
func (o *ReadVmGroupsRequest) SetFilters(v FiltersVmGroup) {
	o.Filters = &v
}

func (o ReadVmGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadVmGroupsRequest struct {
	value *ReadVmGroupsRequest
	isSet bool
}

func (v NullableReadVmGroupsRequest) Get() *ReadVmGroupsRequest {
	return v.value
}

func (v *NullableReadVmGroupsRequest) Set(val *ReadVmGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVmGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVmGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVmGroupsRequest(val *ReadVmGroupsRequest) *NullableReadVmGroupsRequest {
	return &NullableReadVmGroupsRequest{value: val, isSet: true}
}

func (v NullableReadVmGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVmGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
