/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadFlexibleGpuCatalogRequest struct for ReadFlexibleGpuCatalogRequest
type ReadFlexibleGpuCatalogRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewReadFlexibleGpuCatalogRequest instantiates a new ReadFlexibleGpuCatalogRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadFlexibleGpuCatalogRequest() *ReadFlexibleGpuCatalogRequest {
	this := ReadFlexibleGpuCatalogRequest{}
	return &this
}

// NewReadFlexibleGpuCatalogRequestWithDefaults instantiates a new ReadFlexibleGpuCatalogRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadFlexibleGpuCatalogRequestWithDefaults() *ReadFlexibleGpuCatalogRequest {
	this := ReadFlexibleGpuCatalogRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadFlexibleGpuCatalogRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFlexibleGpuCatalogRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadFlexibleGpuCatalogRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadFlexibleGpuCatalogRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o ReadFlexibleGpuCatalogRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableReadFlexibleGpuCatalogRequest struct {
	value *ReadFlexibleGpuCatalogRequest
	isSet bool
}

func (v NullableReadFlexibleGpuCatalogRequest) Get() *ReadFlexibleGpuCatalogRequest {
	return v.value
}

func (v *NullableReadFlexibleGpuCatalogRequest) Set(val *ReadFlexibleGpuCatalogRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadFlexibleGpuCatalogRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadFlexibleGpuCatalogRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadFlexibleGpuCatalogRequest(val *ReadFlexibleGpuCatalogRequest) *NullableReadFlexibleGpuCatalogRequest {
	return &NullableReadFlexibleGpuCatalogRequest{value: val, isSet: true}
}

func (v NullableReadFlexibleGpuCatalogRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadFlexibleGpuCatalogRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


