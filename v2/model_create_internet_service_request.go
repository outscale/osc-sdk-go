/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateInternetServiceRequest struct for CreateInternetServiceRequest
type CreateInternetServiceRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewCreateInternetServiceRequest instantiates a new CreateInternetServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInternetServiceRequest() *CreateInternetServiceRequest {
	this := CreateInternetServiceRequest{}
	return &this
}

// NewCreateInternetServiceRequestWithDefaults instantiates a new CreateInternetServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInternetServiceRequestWithDefaults() *CreateInternetServiceRequest {
	this := CreateInternetServiceRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateInternetServiceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInternetServiceRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateInternetServiceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateInternetServiceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o CreateInternetServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableCreateInternetServiceRequest struct {
	value *CreateInternetServiceRequest
	isSet bool
}

func (v NullableCreateInternetServiceRequest) Get() *CreateInternetServiceRequest {
	return v.value
}

func (v *NullableCreateInternetServiceRequest) Set(val *CreateInternetServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInternetServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInternetServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInternetServiceRequest(val *CreateInternetServiceRequest) *NullableCreateInternetServiceRequest {
	return &NullableCreateInternetServiceRequest{value: val, isSet: true}
}

func (v NullableCreateInternetServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInternetServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
