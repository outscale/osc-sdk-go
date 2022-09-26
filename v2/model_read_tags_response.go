/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadTagsResponse struct for ReadTagsResponse
type ReadTagsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more tags.
	Tags *[]Tag `json:"Tags,omitempty"`
}

// NewReadTagsResponse instantiates a new ReadTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadTagsResponse() *ReadTagsResponse {
	this := ReadTagsResponse{}
	return &this
}

// NewReadTagsResponseWithDefaults instantiates a new ReadTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadTagsResponseWithDefaults() *ReadTagsResponse {
	this := ReadTagsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadTagsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadTagsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadTagsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadTagsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ReadTagsResponse) GetTags() []Tag {
	if o == nil || o.Tags == nil {
		var ret []Tag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadTagsResponse) GetTagsOk() (*[]Tag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ReadTagsResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *ReadTagsResponse) SetTags(v []Tag) {
	o.Tags = &v
}

func (o ReadTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableReadTagsResponse struct {
	value *ReadTagsResponse
	isSet bool
}

func (v NullableReadTagsResponse) Get() *ReadTagsResponse {
	return v.value
}

func (v *NullableReadTagsResponse) Set(val *ReadTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadTagsResponse(val *ReadTagsResponse) *NullableReadTagsResponse {
	return &NullableReadTagsResponse{value: val, isSet: true}
}

func (v NullableReadTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
