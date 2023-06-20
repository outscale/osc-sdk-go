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

// AccessKey Information about the access key.
type AccessKey struct {
	// The ID of the access key.
	AccessKeyId *string `json:"AccessKeyId,omitempty"`
	// The date and time (UTC) of creation of the access key.
	CreationDate *string `json:"CreationDate,omitempty"`
	// The date (UTC) at which the access key expires.
	ExpirationDate *string `json:"ExpirationDate,omitempty"`
	// The date and time (UTC) of the last modification of the access key.
	LastModificationDate *string `json:"LastModificationDate,omitempty"`
	// The state of the access key (`ACTIVE` if the key is valid for API calls, or `INACTIVE` if not).
	State *string `json:"State,omitempty"`
}

// NewAccessKey instantiates a new AccessKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKey() *AccessKey {
	this := AccessKey{}
	return &this
}

// NewAccessKeyWithDefaults instantiates a new AccessKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyWithDefaults() *AccessKey {
	this := AccessKey{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AccessKey) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AccessKey) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AccessKey) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *AccessKey) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *AccessKey) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *AccessKey) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *AccessKey) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *AccessKey) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *AccessKey) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetLastModificationDate returns the LastModificationDate field value if set, zero value otherwise.
func (o *AccessKey) GetLastModificationDate() string {
	if o == nil || o.LastModificationDate == nil {
		var ret string
		return ret
	}
	return *o.LastModificationDate
}

// GetLastModificationDateOk returns a tuple with the LastModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetLastModificationDateOk() (*string, bool) {
	if o == nil || o.LastModificationDate == nil {
		return nil, false
	}
	return o.LastModificationDate, true
}

// HasLastModificationDate returns a boolean if a field has been set.
func (o *AccessKey) HasLastModificationDate() bool {
	if o != nil && o.LastModificationDate != nil {
		return true
	}

	return false
}

// SetLastModificationDate gets a reference to the given string and assigns it to the LastModificationDate field.
func (o *AccessKey) SetLastModificationDate(v string) {
	o.LastModificationDate = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AccessKey) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AccessKey) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AccessKey) SetState(v string) {
	o.State = &v
}

func (o AccessKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKeyId != nil {
		toSerialize["AccessKeyId"] = o.AccessKeyId
	}
	if o.CreationDate != nil {
		toSerialize["CreationDate"] = o.CreationDate
	}
	if o.ExpirationDate != nil {
		toSerialize["ExpirationDate"] = o.ExpirationDate
	}
	if o.LastModificationDate != nil {
		toSerialize["LastModificationDate"] = o.LastModificationDate
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableAccessKey struct {
	value *AccessKey
	isSet bool
}

func (v NullableAccessKey) Get() *AccessKey {
	return v.value
}

func (v *NullableAccessKey) Set(val *AccessKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKey(val *AccessKey) *NullableAccessKey {
	return &NullableAccessKey{value: val, isSet: true}
}

func (v NullableAccessKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
