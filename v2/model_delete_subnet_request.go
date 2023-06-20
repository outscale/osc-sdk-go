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

// DeleteSubnetRequest struct for DeleteSubnetRequest
type DeleteSubnetRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Subnet you want to delete.
	SubnetId string `json:"SubnetId"`
}

// NewDeleteSubnetRequest instantiates a new DeleteSubnetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSubnetRequest(subnetId string) *DeleteSubnetRequest {
	this := DeleteSubnetRequest{}
	this.SubnetId = subnetId
	return &this
}

// NewDeleteSubnetRequestWithDefaults instantiates a new DeleteSubnetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSubnetRequestWithDefaults() *DeleteSubnetRequest {
	this := DeleteSubnetRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteSubnetRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSubnetRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteSubnetRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteSubnetRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetSubnetId returns the SubnetId field value
func (o *DeleteSubnetRequest) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *DeleteSubnetRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *DeleteSubnetRequest) SetSubnetId(v string) {
	o.SubnetId = v
}

func (o DeleteSubnetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["SubnetId"] = o.SubnetId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteSubnetRequest struct {
	value *DeleteSubnetRequest
	isSet bool
}

func (v NullableDeleteSubnetRequest) Get() *DeleteSubnetRequest {
	return v.value
}

func (v *NullableDeleteSubnetRequest) Set(val *DeleteSubnetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSubnetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSubnetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSubnetRequest(val *DeleteSubnetRequest) *NullableDeleteSubnetRequest {
	return &NullableDeleteSubnetRequest{value: val, isSet: true}
}

func (v NullableDeleteSubnetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSubnetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
