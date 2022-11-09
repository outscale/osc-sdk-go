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

// DeleteAccessKeyRequest struct for DeleteAccessKeyRequest
type DeleteAccessKeyRequest struct {
	// The ID of the access key you want to delete.
	AccessKeyId string `json:"AccessKeyId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewDeleteAccessKeyRequest instantiates a new DeleteAccessKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAccessKeyRequest(accessKeyId string) *DeleteAccessKeyRequest {
	this := DeleteAccessKeyRequest{}
	this.AccessKeyId = accessKeyId
	return &this
}

// NewDeleteAccessKeyRequestWithDefaults instantiates a new DeleteAccessKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAccessKeyRequestWithDefaults() *DeleteAccessKeyRequest {
	this := DeleteAccessKeyRequest{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *DeleteAccessKeyRequest) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *DeleteAccessKeyRequest) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *DeleteAccessKeyRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteAccessKeyRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAccessKeyRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteAccessKeyRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteAccessKeyRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o DeleteAccessKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AccessKeyId"] = o.AccessKeyId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAccessKeyRequest struct {
	value *DeleteAccessKeyRequest
	isSet bool
}

func (v NullableDeleteAccessKeyRequest) Get() *DeleteAccessKeyRequest {
	return v.value
}

func (v *NullableDeleteAccessKeyRequest) Set(val *DeleteAccessKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAccessKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAccessKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAccessKeyRequest(val *DeleteAccessKeyRequest) *NullableDeleteAccessKeyRequest {
	return &NullableDeleteAccessKeyRequest{value: val, isSet: true}
}

func (v NullableDeleteAccessKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAccessKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
