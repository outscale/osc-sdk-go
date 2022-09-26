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

// ReadQuotasRequest struct for ReadQuotasRequest
type ReadQuotasRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool         `json:"DryRun,omitempty"`
	Filters *FiltersQuota `json:"Filters,omitempty"`
}

// NewReadQuotasRequest instantiates a new ReadQuotasRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadQuotasRequest() *ReadQuotasRequest {
	this := ReadQuotasRequest{}
	return &this
}

// NewReadQuotasRequestWithDefaults instantiates a new ReadQuotasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadQuotasRequestWithDefaults() *ReadQuotasRequest {
	this := ReadQuotasRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadQuotasRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadQuotasRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadQuotasRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadQuotasRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadQuotasRequest) GetFilters() FiltersQuota {
	if o == nil || o.Filters == nil {
		var ret FiltersQuota
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadQuotasRequest) GetFiltersOk() (*FiltersQuota, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadQuotasRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersQuota and assigns it to the Filters field.
func (o *ReadQuotasRequest) SetFilters(v FiltersQuota) {
	o.Filters = &v
}

func (o ReadQuotasRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadQuotasRequest struct {
	value *ReadQuotasRequest
	isSet bool
}

func (v NullableReadQuotasRequest) Get() *ReadQuotasRequest {
	return v.value
}

func (v *NullableReadQuotasRequest) Set(val *ReadQuotasRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadQuotasRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadQuotasRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadQuotasRequest(val *ReadQuotasRequest) *NullableReadQuotasRequest {
	return &NullableReadQuotasRequest{value: val, isSet: true}
}

func (v NullableReadQuotasRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadQuotasRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
