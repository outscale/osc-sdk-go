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

// ReadImagesRequest struct for ReadImagesRequest
type ReadImagesRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool         `json:"DryRun,omitempty"`
	Filters *FiltersImage `json:"Filters,omitempty"`
}

// NewReadImagesRequest instantiates a new ReadImagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadImagesRequest() *ReadImagesRequest {
	this := ReadImagesRequest{}
	return &this
}

// NewReadImagesRequestWithDefaults instantiates a new ReadImagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadImagesRequestWithDefaults() *ReadImagesRequest {
	this := ReadImagesRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadImagesRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadImagesRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadImagesRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadImagesRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadImagesRequest) GetFilters() FiltersImage {
	if o == nil || o.Filters == nil {
		var ret FiltersImage
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadImagesRequest) GetFiltersOk() (*FiltersImage, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadImagesRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersImage and assigns it to the Filters field.
func (o *ReadImagesRequest) SetFilters(v FiltersImage) {
	o.Filters = &v
}

func (o ReadImagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadImagesRequest struct {
	value *ReadImagesRequest
	isSet bool
}

func (v NullableReadImagesRequest) Get() *ReadImagesRequest {
	return v.value
}

func (v *NullableReadImagesRequest) Set(val *ReadImagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadImagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadImagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadImagesRequest(val *ReadImagesRequest) *NullableReadImagesRequest {
	return &NullableReadImagesRequest{value: val, isSet: true}
}

func (v NullableReadImagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadImagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
