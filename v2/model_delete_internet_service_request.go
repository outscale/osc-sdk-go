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

// DeleteInternetServiceRequest struct for DeleteInternetServiceRequest
type DeleteInternetServiceRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Internet service you want to delete.
	InternetServiceId string `json:"InternetServiceId"`
}

// NewDeleteInternetServiceRequest instantiates a new DeleteInternetServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInternetServiceRequest(internetServiceId string) *DeleteInternetServiceRequest {
	this := DeleteInternetServiceRequest{}
	this.InternetServiceId = internetServiceId
	return &this
}

// NewDeleteInternetServiceRequestWithDefaults instantiates a new DeleteInternetServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInternetServiceRequestWithDefaults() *DeleteInternetServiceRequest {
	this := DeleteInternetServiceRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteInternetServiceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInternetServiceRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteInternetServiceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteInternetServiceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetInternetServiceId returns the InternetServiceId field value
func (o *DeleteInternetServiceRequest) GetInternetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InternetServiceId
}

// GetInternetServiceIdOk returns a tuple with the InternetServiceId field value
// and a boolean to check if the value has been set.
func (o *DeleteInternetServiceRequest) GetInternetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternetServiceId, true
}

// SetInternetServiceId sets field value
func (o *DeleteInternetServiceRequest) SetInternetServiceId(v string) {
	o.InternetServiceId = v
}

func (o DeleteInternetServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["InternetServiceId"] = o.InternetServiceId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteInternetServiceRequest struct {
	value *DeleteInternetServiceRequest
	isSet bool
}

func (v NullableDeleteInternetServiceRequest) Get() *DeleteInternetServiceRequest {
	return v.value
}

func (v *NullableDeleteInternetServiceRequest) Set(val *DeleteInternetServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInternetServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInternetServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInternetServiceRequest(val *DeleteInternetServiceRequest) *NullableDeleteInternetServiceRequest {
	return &NullableDeleteInternetServiceRequest{value: val, isSet: true}
}

func (v NullableDeleteInternetServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteInternetServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
