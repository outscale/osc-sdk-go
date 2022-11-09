/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadVmsStateRequest struct for ReadVmsStateRequest
type ReadVmsStateRequest struct {
	// If true, includes the status of all VMs. By default or if set to false, only includes the status of running VMs.
	AllVms *bool `json:"AllVms,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool            `json:"DryRun,omitempty"`
	Filters *FiltersVmsState `json:"Filters,omitempty"`
}

// NewReadVmsStateRequest instantiates a new ReadVmsStateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVmsStateRequest() *ReadVmsStateRequest {
	this := ReadVmsStateRequest{}
	var allVms bool = false
	this.AllVms = &allVms
	return &this
}

// NewReadVmsStateRequestWithDefaults instantiates a new ReadVmsStateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVmsStateRequestWithDefaults() *ReadVmsStateRequest {
	this := ReadVmsStateRequest{}
	var allVms bool = false
	this.AllVms = &allVms
	return &this
}

// GetAllVms returns the AllVms field value if set, zero value otherwise.
func (o *ReadVmsStateRequest) GetAllVms() bool {
	if o == nil || o.AllVms == nil {
		var ret bool
		return ret
	}
	return *o.AllVms
}

// GetAllVmsOk returns a tuple with the AllVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsStateRequest) GetAllVmsOk() (*bool, bool) {
	if o == nil || o.AllVms == nil {
		return nil, false
	}
	return o.AllVms, true
}

// HasAllVms returns a boolean if a field has been set.
func (o *ReadVmsStateRequest) HasAllVms() bool {
	if o != nil && o.AllVms != nil {
		return true
	}

	return false
}

// SetAllVms gets a reference to the given bool and assigns it to the AllVms field.
func (o *ReadVmsStateRequest) SetAllVms(v bool) {
	o.AllVms = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadVmsStateRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsStateRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadVmsStateRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadVmsStateRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadVmsStateRequest) GetFilters() FiltersVmsState {
	if o == nil || o.Filters == nil {
		var ret FiltersVmsState
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsStateRequest) GetFiltersOk() (*FiltersVmsState, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadVmsStateRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersVmsState and assigns it to the Filters field.
func (o *ReadVmsStateRequest) SetFilters(v FiltersVmsState) {
	o.Filters = &v
}

func (o ReadVmsStateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllVms != nil {
		toSerialize["AllVms"] = o.AllVms
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadVmsStateRequest struct {
	value *ReadVmsStateRequest
	isSet bool
}

func (v NullableReadVmsStateRequest) Get() *ReadVmsStateRequest {
	return v.value
}

func (v *NullableReadVmsStateRequest) Set(val *ReadVmsStateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVmsStateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVmsStateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVmsStateRequest(val *ReadVmsStateRequest) *NullableReadVmsStateRequest {
	return &NullableReadVmsStateRequest{value: val, isSet: true}
}

func (v NullableReadVmsStateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVmsStateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
