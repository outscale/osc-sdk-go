/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateApiAccessPolicyRequest struct for UpdateApiAccessPolicyRequest
type UpdateApiAccessPolicyRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The maximum possible lifetime for your access keys, in seconds (between `0` and `3153600000`, both included). If set to `O`, your access keys can have unlimited lifetimes, but a trusted session cannot be activated. Otherwise, all your access keys must have an expiration date. This value must be greater than the remaining lifetime of each access key of your account.
	MaxAccessKeyExpirationSeconds int64 `json:"MaxAccessKeyExpirationSeconds"`
	// If true, a trusted session is activated, provided that you specify the `MaxAccessKeyExpirationSeconds` parameter with a value greater than `0`.
	RequireTrustedEnv bool `json:"RequireTrustedEnv"`
}

// NewUpdateApiAccessPolicyRequest instantiates a new UpdateApiAccessPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiAccessPolicyRequest(maxAccessKeyExpirationSeconds int64, requireTrustedEnv bool) *UpdateApiAccessPolicyRequest {
	this := UpdateApiAccessPolicyRequest{}
	this.MaxAccessKeyExpirationSeconds = maxAccessKeyExpirationSeconds
	this.RequireTrustedEnv = requireTrustedEnv
	return &this
}

// NewUpdateApiAccessPolicyRequestWithDefaults instantiates a new UpdateApiAccessPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiAccessPolicyRequestWithDefaults() *UpdateApiAccessPolicyRequest {
	this := UpdateApiAccessPolicyRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateApiAccessPolicyRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessPolicyRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateApiAccessPolicyRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateApiAccessPolicyRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetMaxAccessKeyExpirationSeconds returns the MaxAccessKeyExpirationSeconds field value
func (o *UpdateApiAccessPolicyRequest) GetMaxAccessKeyExpirationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxAccessKeyExpirationSeconds
}

// GetMaxAccessKeyExpirationSecondsOk returns a tuple with the MaxAccessKeyExpirationSeconds field value
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessPolicyRequest) GetMaxAccessKeyExpirationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAccessKeyExpirationSeconds, true
}

// SetMaxAccessKeyExpirationSeconds sets field value
func (o *UpdateApiAccessPolicyRequest) SetMaxAccessKeyExpirationSeconds(v int64) {
	o.MaxAccessKeyExpirationSeconds = v
}

// GetRequireTrustedEnv returns the RequireTrustedEnv field value
func (o *UpdateApiAccessPolicyRequest) GetRequireTrustedEnv() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireTrustedEnv
}

// GetRequireTrustedEnvOk returns a tuple with the RequireTrustedEnv field value
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessPolicyRequest) GetRequireTrustedEnvOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireTrustedEnv, true
}

// SetRequireTrustedEnv sets field value
func (o *UpdateApiAccessPolicyRequest) SetRequireTrustedEnv(v bool) {
	o.RequireTrustedEnv = v
}

func (o UpdateApiAccessPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["MaxAccessKeyExpirationSeconds"] = o.MaxAccessKeyExpirationSeconds
	}
	if true {
		toSerialize["RequireTrustedEnv"] = o.RequireTrustedEnv
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateApiAccessPolicyRequest struct {
	value *UpdateApiAccessPolicyRequest
	isSet bool
}

func (v NullableUpdateApiAccessPolicyRequest) Get() *UpdateApiAccessPolicyRequest {
	return v.value
}

func (v *NullableUpdateApiAccessPolicyRequest) Set(val *UpdateApiAccessPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiAccessPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiAccessPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiAccessPolicyRequest(val *UpdateApiAccessPolicyRequest) *NullableUpdateApiAccessPolicyRequest {
	return &NullableUpdateApiAccessPolicyRequest{value: val, isSet: true}
}

func (v NullableUpdateApiAccessPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiAccessPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
