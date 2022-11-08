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

// DeleteKeypairRequest struct for DeleteKeypairRequest
type DeleteKeypairRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The name of the keypair you want to delete.
	KeypairName string `json:"KeypairName"`
}

// NewDeleteKeypairRequest instantiates a new DeleteKeypairRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteKeypairRequest(keypairName string) *DeleteKeypairRequest {
	this := DeleteKeypairRequest{}
	this.KeypairName = keypairName
	return &this
}

// NewDeleteKeypairRequestWithDefaults instantiates a new DeleteKeypairRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteKeypairRequestWithDefaults() *DeleteKeypairRequest {
	this := DeleteKeypairRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteKeypairRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteKeypairRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteKeypairRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteKeypairRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetKeypairName returns the KeypairName field value
func (o *DeleteKeypairRequest) GetKeypairName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value
// and a boolean to check if the value has been set.
func (o *DeleteKeypairRequest) GetKeypairNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeypairName, true
}

// SetKeypairName sets field value
func (o *DeleteKeypairRequest) SetKeypairName(v string) {
	o.KeypairName = v
}

func (o DeleteKeypairRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["KeypairName"] = o.KeypairName
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteKeypairRequest struct {
	value *DeleteKeypairRequest
	isSet bool
}

func (v NullableDeleteKeypairRequest) Get() *DeleteKeypairRequest {
	return v.value
}

func (v *NullableDeleteKeypairRequest) Set(val *DeleteKeypairRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteKeypairRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteKeypairRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteKeypairRequest(val *DeleteKeypairRequest) *NullableDeleteKeypairRequest {
	return &NullableDeleteKeypairRequest{value: val, isSet: true}
}

func (v NullableDeleteKeypairRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteKeypairRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
