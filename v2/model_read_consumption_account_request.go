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

// ReadConsumptionAccountRequest struct for ReadConsumptionAccountRequest
type ReadConsumptionAccountRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The beginning of the time period, in ISO 8601 date-time format (for example, `2017-06-14` or `2017-06-14T00:00:00Z`).
	FromDate string `json:"FromDate"`
	// The end of the time period, in ISO 8601 date-time format (for example, `2017-06-30` or `2017-06-30T00:00:00Z`).
	ToDate string `json:"ToDate"`
}

// NewReadConsumptionAccountRequest instantiates a new ReadConsumptionAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadConsumptionAccountRequest(fromDate string, toDate string) *ReadConsumptionAccountRequest {
	this := ReadConsumptionAccountRequest{}
	this.FromDate = fromDate
	this.ToDate = toDate
	return &this
}

// NewReadConsumptionAccountRequestWithDefaults instantiates a new ReadConsumptionAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadConsumptionAccountRequestWithDefaults() *ReadConsumptionAccountRequest {
	this := ReadConsumptionAccountRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadConsumptionAccountRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadConsumptionAccountRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadConsumptionAccountRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFromDate returns the FromDate field value
func (o *ReadConsumptionAccountRequest) GetFromDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountRequest) GetFromDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromDate, true
}

// SetFromDate sets field value
func (o *ReadConsumptionAccountRequest) SetFromDate(v string) {
	o.FromDate = v
}

// GetToDate returns the ToDate field value
func (o *ReadConsumptionAccountRequest) GetToDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value
// and a boolean to check if the value has been set.
func (o *ReadConsumptionAccountRequest) GetToDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToDate, true
}

// SetToDate sets field value
func (o *ReadConsumptionAccountRequest) SetToDate(v string) {
	o.ToDate = v
}

func (o ReadConsumptionAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["FromDate"] = o.FromDate
	}
	if true {
		toSerialize["ToDate"] = o.ToDate
	}
	return json.Marshal(toSerialize)
}

type NullableReadConsumptionAccountRequest struct {
	value *ReadConsumptionAccountRequest
	isSet bool
}

func (v NullableReadConsumptionAccountRequest) Get() *ReadConsumptionAccountRequest {
	return v.value
}

func (v *NullableReadConsumptionAccountRequest) Set(val *ReadConsumptionAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadConsumptionAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadConsumptionAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadConsumptionAccountRequest(val *ReadConsumptionAccountRequest) *NullableReadConsumptionAccountRequest {
	return &NullableReadConsumptionAccountRequest{value: val, isSet: true}
}

func (v NullableReadConsumptionAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadConsumptionAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
