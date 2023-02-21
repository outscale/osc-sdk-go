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

// ResetAccountPasswordRequest struct for ResetAccountPasswordRequest
type ResetAccountPasswordRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The new password for the account.
	Password string `json:"Password"`
	// The token you received at the email address provided for the account.
	Token string `json:"Token"`
}

// NewResetAccountPasswordRequest instantiates a new ResetAccountPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetAccountPasswordRequest(password string, token string) *ResetAccountPasswordRequest {
	this := ResetAccountPasswordRequest{}
	this.Password = password
	this.Token = token
	return &this
}

// NewResetAccountPasswordRequestWithDefaults instantiates a new ResetAccountPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetAccountPasswordRequestWithDefaults() *ResetAccountPasswordRequest {
	this := ResetAccountPasswordRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ResetAccountPasswordRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetAccountPasswordRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ResetAccountPasswordRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ResetAccountPasswordRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPassword returns the Password field value
func (o *ResetAccountPasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ResetAccountPasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ResetAccountPasswordRequest) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *ResetAccountPasswordRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ResetAccountPasswordRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ResetAccountPasswordRequest) SetToken(v string) {
	o.Token = v
}

func (o ResetAccountPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Password"] = o.Password
	}
	if true {
		toSerialize["Token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableResetAccountPasswordRequest struct {
	value *ResetAccountPasswordRequest
	isSet bool
}

func (v NullableResetAccountPasswordRequest) Get() *ResetAccountPasswordRequest {
	return v.value
}

func (v *NullableResetAccountPasswordRequest) Set(val *ResetAccountPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResetAccountPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResetAccountPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetAccountPasswordRequest(val *ResetAccountPasswordRequest) *NullableResetAccountPasswordRequest {
	return &NullableResetAccountPasswordRequest{value: val, isSet: true}
}

func (v NullableResetAccountPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetAccountPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
