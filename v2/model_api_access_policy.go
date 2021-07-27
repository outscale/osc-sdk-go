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

// ApiAccessPolicy Information about the API access policy.
type ApiAccessPolicy struct {
	// The maximum possible lifetime for your access keys, in seconds. If `0`, your access keys can have unlimited lifetimes.
	MaxAccessKeyExpirationSeconds *int64 `json:"MaxAccessKeyExpirationSeconds,omitempty"`
	// If true, a trusted session is activated, allowing you to bypass Certificate Authorities (CAs) enforcement. For more information, see the `ApiKeyAuth` authentication scheme in the [Authentication](#authentication) section.
	RequireTrustedEnv *bool `json:"RequireTrustedEnv,omitempty"`
}

// NewApiAccessPolicy instantiates a new ApiAccessPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccessPolicy() *ApiAccessPolicy {
	this := ApiAccessPolicy{}
	return &this
}

// NewApiAccessPolicyWithDefaults instantiates a new ApiAccessPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccessPolicyWithDefaults() *ApiAccessPolicy {
	this := ApiAccessPolicy{}
	return &this
}

// GetMaxAccessKeyExpirationSeconds returns the MaxAccessKeyExpirationSeconds field value if set, zero value otherwise.
func (o *ApiAccessPolicy) GetMaxAccessKeyExpirationSeconds() int64 {
	if o == nil || o.MaxAccessKeyExpirationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.MaxAccessKeyExpirationSeconds
}

// GetMaxAccessKeyExpirationSecondsOk returns a tuple with the MaxAccessKeyExpirationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessPolicy) GetMaxAccessKeyExpirationSecondsOk() (*int64, bool) {
	if o == nil || o.MaxAccessKeyExpirationSeconds == nil {
		return nil, false
	}
	return o.MaxAccessKeyExpirationSeconds, true
}

// HasMaxAccessKeyExpirationSeconds returns a boolean if a field has been set.
func (o *ApiAccessPolicy) HasMaxAccessKeyExpirationSeconds() bool {
	if o != nil && o.MaxAccessKeyExpirationSeconds != nil {
		return true
	}

	return false
}

// SetMaxAccessKeyExpirationSeconds gets a reference to the given int64 and assigns it to the MaxAccessKeyExpirationSeconds field.
func (o *ApiAccessPolicy) SetMaxAccessKeyExpirationSeconds(v int64) {
	o.MaxAccessKeyExpirationSeconds = &v
}

// GetRequireTrustedEnv returns the RequireTrustedEnv field value if set, zero value otherwise.
func (o *ApiAccessPolicy) GetRequireTrustedEnv() bool {
	if o == nil || o.RequireTrustedEnv == nil {
		var ret bool
		return ret
	}
	return *o.RequireTrustedEnv
}

// GetRequireTrustedEnvOk returns a tuple with the RequireTrustedEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessPolicy) GetRequireTrustedEnvOk() (*bool, bool) {
	if o == nil || o.RequireTrustedEnv == nil {
		return nil, false
	}
	return o.RequireTrustedEnv, true
}

// HasRequireTrustedEnv returns a boolean if a field has been set.
func (o *ApiAccessPolicy) HasRequireTrustedEnv() bool {
	if o != nil && o.RequireTrustedEnv != nil {
		return true
	}

	return false
}

// SetRequireTrustedEnv gets a reference to the given bool and assigns it to the RequireTrustedEnv field.
func (o *ApiAccessPolicy) SetRequireTrustedEnv(v bool) {
	o.RequireTrustedEnv = &v
}

func (o ApiAccessPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxAccessKeyExpirationSeconds != nil {
		toSerialize["MaxAccessKeyExpirationSeconds"] = o.MaxAccessKeyExpirationSeconds
	}
	if o.RequireTrustedEnv != nil {
		toSerialize["RequireTrustedEnv"] = o.RequireTrustedEnv
	}
	return json.Marshal(toSerialize)
}

type NullableApiAccessPolicy struct {
	value *ApiAccessPolicy
	isSet bool
}

func (v NullableApiAccessPolicy) Get() *ApiAccessPolicy {
	return v.value
}

func (v *NullableApiAccessPolicy) Set(val *ApiAccessPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccessPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccessPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccessPolicy(val *ApiAccessPolicy) *NullableApiAccessPolicy {
	return &NullableApiAccessPolicy{value: val, isSet: true}
}

func (v NullableApiAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccessPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
