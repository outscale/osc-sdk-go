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

// OsuApiKey Information about the OOS API key.
type OsuApiKey struct {
	// The API key of the OOS account that enables you to access the bucket.
	ApiKeyId *string `json:"ApiKeyId,omitempty"`
	// The secret key of the OOS account that enables you to access the bucket.
	SecretKey *string `json:"SecretKey,omitempty"`
}

// NewOsuApiKey instantiates a new OsuApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsuApiKey() *OsuApiKey {
	this := OsuApiKey{}
	return &this
}

// NewOsuApiKeyWithDefaults instantiates a new OsuApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsuApiKeyWithDefaults() *OsuApiKey {
	this := OsuApiKey{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *OsuApiKey) GetApiKeyId() string {
	if o == nil || o.ApiKeyId == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsuApiKey) GetApiKeyIdOk() (*string, bool) {
	if o == nil || o.ApiKeyId == nil {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *OsuApiKey) HasApiKeyId() bool {
	if o != nil && o.ApiKeyId != nil {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *OsuApiKey) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *OsuApiKey) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsuApiKey) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *OsuApiKey) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *OsuApiKey) SetSecretKey(v string) {
	o.SecretKey = &v
}

func (o OsuApiKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKeyId != nil {
		toSerialize["ApiKeyId"] = o.ApiKeyId
	}
	if o.SecretKey != nil {
		toSerialize["SecretKey"] = o.SecretKey
	}
	return json.Marshal(toSerialize)
}

type NullableOsuApiKey struct {
	value *OsuApiKey
	isSet bool
}

func (v NullableOsuApiKey) Get() *OsuApiKey {
	return v.value
}

func (v *NullableOsuApiKey) Set(val *OsuApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableOsuApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableOsuApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsuApiKey(val *OsuApiKey) *NullableOsuApiKey {
	return &NullableOsuApiKey{value: val, isSet: true}
}

func (v NullableOsuApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsuApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
