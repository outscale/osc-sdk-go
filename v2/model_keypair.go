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

// Keypair Information about the keypair.
type Keypair struct {
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	KeypairFingerprint *string `json:"KeypairFingerprint,omitempty"`
	// The name of the keypair.
	KeypairName *string `json:"KeypairName,omitempty"`
}

// NewKeypair instantiates a new Keypair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeypair() *Keypair {
	this := Keypair{}
	return &this
}

// NewKeypairWithDefaults instantiates a new Keypair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeypairWithDefaults() *Keypair {
	this := Keypair{}
	return &this
}

// GetKeypairFingerprint returns the KeypairFingerprint field value if set, zero value otherwise.
func (o *Keypair) GetKeypairFingerprint() string {
	if o == nil || o.KeypairFingerprint == nil {
		var ret string
		return ret
	}
	return *o.KeypairFingerprint
}

// GetKeypairFingerprintOk returns a tuple with the KeypairFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypair) GetKeypairFingerprintOk() (*string, bool) {
	if o == nil || o.KeypairFingerprint == nil {
		return nil, false
	}
	return o.KeypairFingerprint, true
}

// HasKeypairFingerprint returns a boolean if a field has been set.
func (o *Keypair) HasKeypairFingerprint() bool {
	if o != nil && o.KeypairFingerprint != nil {
		return true
	}

	return false
}

// SetKeypairFingerprint gets a reference to the given string and assigns it to the KeypairFingerprint field.
func (o *Keypair) SetKeypairFingerprint(v string) {
	o.KeypairFingerprint = &v
}

// GetKeypairName returns the KeypairName field value if set, zero value otherwise.
func (o *Keypair) GetKeypairName() string {
	if o == nil || o.KeypairName == nil {
		var ret string
		return ret
	}
	return *o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypair) GetKeypairNameOk() (*string, bool) {
	if o == nil || o.KeypairName == nil {
		return nil, false
	}
	return o.KeypairName, true
}

// HasKeypairName returns a boolean if a field has been set.
func (o *Keypair) HasKeypairName() bool {
	if o != nil && o.KeypairName != nil {
		return true
	}

	return false
}

// SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.
func (o *Keypair) SetKeypairName(v string) {
	o.KeypairName = &v
}

func (o Keypair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeypairFingerprint != nil {
		toSerialize["KeypairFingerprint"] = o.KeypairFingerprint
	}
	if o.KeypairName != nil {
		toSerialize["KeypairName"] = o.KeypairName
	}
	return json.Marshal(toSerialize)
}

type NullableKeypair struct {
	value *Keypair
	isSet bool
}

func (v NullableKeypair) Get() *Keypair {
	return v.value
}

func (v *NullableKeypair) Set(val *Keypair) {
	v.value = val
	v.isSet = true
}

func (v NullableKeypair) IsSet() bool {
	return v.isSet
}

func (v *NullableKeypair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeypair(val *Keypair) *NullableKeypair {
	return &NullableKeypair{value: val, isSet: true}
}

func (v NullableKeypair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeypair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
