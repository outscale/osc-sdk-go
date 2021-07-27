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

// UpdateAccountResponse struct for UpdateAccountResponse
type UpdateAccountResponse struct {
	Account         *Account         `json:"Account,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateAccountResponse instantiates a new UpdateAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountResponse() *UpdateAccountResponse {
	this := UpdateAccountResponse{}
	return &this
}

// NewUpdateAccountResponseWithDefaults instantiates a new UpdateAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountResponseWithDefaults() *UpdateAccountResponse {
	this := UpdateAccountResponse{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *UpdateAccountResponse) GetAccount() Account {
	if o == nil || o.Account == nil {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountResponse) GetAccountOk() (*Account, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *UpdateAccountResponse) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *UpdateAccountResponse) SetAccount(v Account) {
	o.Account = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateAccountResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateAccountResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateAccountResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAccountResponse struct {
	value *UpdateAccountResponse
	isSet bool
}

func (v NullableUpdateAccountResponse) Get() *UpdateAccountResponse {
	return v.value
}

func (v *NullableUpdateAccountResponse) Set(val *UpdateAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountResponse(val *UpdateAccountResponse) *NullableUpdateAccountResponse {
	return &NullableUpdateAccountResponse{value: val, isSet: true}
}

func (v NullableUpdateAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
