/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// Ca Information about the Client Certificate Authority (CA).
type Ca struct {
	// The fingerprint of the CA.
	CaFingerprint *string `json:"CaFingerprint,omitempty"`
	// The ID of the CA.
	CaId *string `json:"CaId,omitempty"`
	// The description of the CA.
	Description *string `json:"Description,omitempty"`
}

// NewCa instantiates a new Ca object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCa() *Ca {
	this := Ca{}
	return &this
}

// NewCaWithDefaults instantiates a new Ca object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaWithDefaults() *Ca {
	this := Ca{}
	return &this
}

// GetCaFingerprint returns the CaFingerprint field value if set, zero value otherwise.
func (o *Ca) GetCaFingerprint() string {
	if o == nil || o.CaFingerprint == nil {
		var ret string
		return ret
	}
	return *o.CaFingerprint
}

// GetCaFingerprintOk returns a tuple with the CaFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ca) GetCaFingerprintOk() (*string, bool) {
	if o == nil || o.CaFingerprint == nil {
		return nil, false
	}
	return o.CaFingerprint, true
}

// HasCaFingerprint returns a boolean if a field has been set.
func (o *Ca) HasCaFingerprint() bool {
	if o != nil && o.CaFingerprint != nil {
		return true
	}

	return false
}

// SetCaFingerprint gets a reference to the given string and assigns it to the CaFingerprint field.
func (o *Ca) SetCaFingerprint(v string) {
	o.CaFingerprint = &v
}

// GetCaId returns the CaId field value if set, zero value otherwise.
func (o *Ca) GetCaId() string {
	if o == nil || o.CaId == nil {
		var ret string
		return ret
	}
	return *o.CaId
}

// GetCaIdOk returns a tuple with the CaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ca) GetCaIdOk() (*string, bool) {
	if o == nil || o.CaId == nil {
		return nil, false
	}
	return o.CaId, true
}

// HasCaId returns a boolean if a field has been set.
func (o *Ca) HasCaId() bool {
	if o != nil && o.CaId != nil {
		return true
	}

	return false
}

// SetCaId gets a reference to the given string and assigns it to the CaId field.
func (o *Ca) SetCaId(v string) {
	o.CaId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Ca) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ca) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ca) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ca) SetDescription(v string) {
	o.Description = &v
}

func (o Ca) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CaFingerprint != nil {
		toSerialize["CaFingerprint"] = o.CaFingerprint
	}
	if o.CaId != nil {
		toSerialize["CaId"] = o.CaId
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCa struct {
	value *Ca
	isSet bool
}

func (v NullableCa) Get() *Ca {
	return v.value
}

func (v *NullableCa) Set(val *Ca) {
	v.value = val
	v.isSet = true
}

func (v NullableCa) IsSet() bool {
	return v.isSet
}

func (v *NullableCa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCa(val *Ca) *NullableCa {
	return &NullableCa{value: val, isSet: true}
}

func (v NullableCa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
