/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// SendResetPasswordEmailRequest struct for SendResetPasswordEmailRequest
type SendResetPasswordEmailRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The email address provided for the account.
	Email string `json:"Email"`
}

// NewSendResetPasswordEmailRequest instantiates a new SendResetPasswordEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendResetPasswordEmailRequest(email string) *SendResetPasswordEmailRequest {
	this := SendResetPasswordEmailRequest{}
	this.Email = email
	return &this
}

// NewSendResetPasswordEmailRequestWithDefaults instantiates a new SendResetPasswordEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendResetPasswordEmailRequestWithDefaults() *SendResetPasswordEmailRequest {
	this := SendResetPasswordEmailRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *SendResetPasswordEmailRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendResetPasswordEmailRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *SendResetPasswordEmailRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *SendResetPasswordEmailRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetEmail returns the Email field value
func (o *SendResetPasswordEmailRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SendResetPasswordEmailRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SendResetPasswordEmailRequest) SetEmail(v string) {
	o.Email = v
}

func (o SendResetPasswordEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableSendResetPasswordEmailRequest struct {
	value *SendResetPasswordEmailRequest
	isSet bool
}

func (v NullableSendResetPasswordEmailRequest) Get() *SendResetPasswordEmailRequest {
	return v.value
}

func (v *NullableSendResetPasswordEmailRequest) Set(val *SendResetPasswordEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendResetPasswordEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendResetPasswordEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendResetPasswordEmailRequest(val *SendResetPasswordEmailRequest) *NullableSendResetPasswordEmailRequest {
	return &NullableSendResetPasswordEmailRequest{value: val, isSet: true}
}

func (v NullableSendResetPasswordEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendResetPasswordEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
