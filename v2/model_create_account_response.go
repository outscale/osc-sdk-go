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

// CreateAccountResponse struct for CreateAccountResponse
type CreateAccountResponse struct {
	Account         *Account         `json:"Account,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateAccountResponse instantiates a new CreateAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountResponse() *CreateAccountResponse {
	this := CreateAccountResponse{}
	return &this
}

// NewCreateAccountResponseWithDefaults instantiates a new CreateAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountResponseWithDefaults() *CreateAccountResponse {
	this := CreateAccountResponse{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetAccount() Account {
	if o == nil || o.Account == nil {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetAccountOk() (*Account, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *CreateAccountResponse) SetAccount(v Account) {
	o.Account = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateAccountResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAccountResponse struct {
	value *CreateAccountResponse
	isSet bool
}

func (v NullableCreateAccountResponse) Get() *CreateAccountResponse {
	return v.value
}

func (v *NullableCreateAccountResponse) Set(val *CreateAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountResponse(val *CreateAccountResponse) *NullableCreateAccountResponse {
	return &NullableCreateAccountResponse{value: val, isSet: true}
}

func (v NullableCreateAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
