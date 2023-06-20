/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.27
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadUsersResponse struct for ReadUsersResponse
type ReadUsersResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// A list of EIM users.
	Users *[]User `json:"Users,omitempty"`
}

// NewReadUsersResponse instantiates a new ReadUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadUsersResponse() *ReadUsersResponse {
	this := ReadUsersResponse{}
	return &this
}

// NewReadUsersResponseWithDefaults instantiates a new ReadUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadUsersResponseWithDefaults() *ReadUsersResponse {
	this := ReadUsersResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadUsersResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadUsersResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadUsersResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadUsersResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ReadUsersResponse) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadUsersResponse) GetUsersOk() (*[]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ReadUsersResponse) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *ReadUsersResponse) SetUsers(v []User) {
	o.Users = &v
}

func (o ReadUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableReadUsersResponse struct {
	value *ReadUsersResponse
	isSet bool
}

func (v NullableReadUsersResponse) Get() *ReadUsersResponse {
	return v.value
}

func (v *NullableReadUsersResponse) Set(val *ReadUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadUsersResponse(val *ReadUsersResponse) *NullableReadUsersResponse {
	return &NullableReadUsersResponse{value: val, isSet: true}
}

func (v NullableReadUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
