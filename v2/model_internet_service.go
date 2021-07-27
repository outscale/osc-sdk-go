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

// InternetService Information about the Internet service.
type InternetService struct {
	// The ID of the Internet service.
	InternetServiceId *string `json:"InternetServiceId,omitempty"`
	// The ID of the Net attached to the Internet service.
	NetId *string `json:"NetId,omitempty"`
	// The state of the attachment of the Net to the Internet service (always `available`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the Internet service.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewInternetService instantiates a new InternetService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternetService() *InternetService {
	this := InternetService{}
	return &this
}

// NewInternetServiceWithDefaults instantiates a new InternetService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternetServiceWithDefaults() *InternetService {
	this := InternetService{}
	return &this
}

// GetInternetServiceId returns the InternetServiceId field value if set, zero value otherwise.
func (o *InternetService) GetInternetServiceId() string {
	if o == nil || o.InternetServiceId == nil {
		var ret string
		return ret
	}
	return *o.InternetServiceId
}

// GetInternetServiceIdOk returns a tuple with the InternetServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternetService) GetInternetServiceIdOk() (*string, bool) {
	if o == nil || o.InternetServiceId == nil {
		return nil, false
	}
	return o.InternetServiceId, true
}

// HasInternetServiceId returns a boolean if a field has been set.
func (o *InternetService) HasInternetServiceId() bool {
	if o != nil && o.InternetServiceId != nil {
		return true
	}

	return false
}

// SetInternetServiceId gets a reference to the given string and assigns it to the InternetServiceId field.
func (o *InternetService) SetInternetServiceId(v string) {
	o.InternetServiceId = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *InternetService) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternetService) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *InternetService) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *InternetService) SetNetId(v string) {
	o.NetId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *InternetService) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternetService) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *InternetService) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *InternetService) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InternetService) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternetService) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InternetService) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *InternetService) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o InternetService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InternetServiceId != nil {
		toSerialize["InternetServiceId"] = o.InternetServiceId
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableInternetService struct {
	value *InternetService
	isSet bool
}

func (v NullableInternetService) Get() *InternetService {
	return v.value
}

func (v *NullableInternetService) Set(val *InternetService) {
	v.value = val
	v.isSet = true
}

func (v NullableInternetService) IsSet() bool {
	return v.isSet
}

func (v *NullableInternetService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternetService(val *InternetService) *NullableInternetService {
	return &NullableInternetService{value: val, isSet: true}
}

func (v NullableInternetService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternetService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
