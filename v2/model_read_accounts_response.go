/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadAccountsResponse struct for ReadAccountsResponse
type ReadAccountsResponse struct {
	// The list of the accounts.
	Accounts        *[]Account       `json:"Accounts,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadAccountsResponse instantiates a new ReadAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAccountsResponse() *ReadAccountsResponse {
	this := ReadAccountsResponse{}
	return &this
}

// NewReadAccountsResponseWithDefaults instantiates a new ReadAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAccountsResponseWithDefaults() *ReadAccountsResponse {
	this := ReadAccountsResponse{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ReadAccountsResponse) GetAccounts() []Account {
	if o == nil || o.Accounts == nil {
		var ret []Account
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAccountsResponse) GetAccountsOk() (*[]Account, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *ReadAccountsResponse) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []Account and assigns it to the Accounts field.
func (o *ReadAccountsResponse) SetAccounts(v []Account) {
	o.Accounts = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadAccountsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAccountsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadAccountsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadAccountsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accounts != nil {
		toSerialize["Accounts"] = o.Accounts
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadAccountsResponse struct {
	value *ReadAccountsResponse
	isSet bool
}

func (v NullableReadAccountsResponse) Get() *ReadAccountsResponse {
	return v.value
}

func (v *NullableReadAccountsResponse) Set(val *ReadAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAccountsResponse(val *ReadAccountsResponse) *NullableReadAccountsResponse {
	return &NullableReadAccountsResponse{value: val, isSet: true}
}

func (v NullableReadAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
