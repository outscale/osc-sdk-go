/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.27
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteCaRequest struct for DeleteCaRequest
type DeleteCaRequest struct {
	// The ID of the CA you want to delete.
	CaId string `json:"CaId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewDeleteCaRequest instantiates a new DeleteCaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCaRequest(caId string) *DeleteCaRequest {
	this := DeleteCaRequest{}
	this.CaId = caId
	return &this
}

// NewDeleteCaRequestWithDefaults instantiates a new DeleteCaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCaRequestWithDefaults() *DeleteCaRequest {
	this := DeleteCaRequest{}
	return &this
}

// GetCaId returns the CaId field value
func (o *DeleteCaRequest) GetCaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaId
}

// GetCaIdOk returns a tuple with the CaId field value
// and a boolean to check if the value has been set.
func (o *DeleteCaRequest) GetCaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaId, true
}

// SetCaId sets field value
func (o *DeleteCaRequest) SetCaId(v string) {
	o.CaId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteCaRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCaRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteCaRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteCaRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o DeleteCaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CaId"] = o.CaId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteCaRequest struct {
	value *DeleteCaRequest
	isSet bool
}

func (v NullableDeleteCaRequest) Get() *DeleteCaRequest {
	return v.value
}

func (v *NullableDeleteCaRequest) Set(val *DeleteCaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCaRequest(val *DeleteCaRequest) *NullableDeleteCaRequest {
	return &NullableDeleteCaRequest{value: val, isSet: true}
}

func (v NullableDeleteCaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
