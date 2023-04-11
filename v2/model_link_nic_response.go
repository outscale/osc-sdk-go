/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkNicResponse struct for LinkNicResponse
type LinkNicResponse struct {
	// The ID of the NIC attachment.
	LinkNicId       *string          `json:"LinkNicId,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewLinkNicResponse instantiates a new LinkNicResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkNicResponse() *LinkNicResponse {
	this := LinkNicResponse{}
	return &this
}

// NewLinkNicResponseWithDefaults instantiates a new LinkNicResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkNicResponseWithDefaults() *LinkNicResponse {
	this := LinkNicResponse{}
	return &this
}

// GetLinkNicId returns the LinkNicId field value if set, zero value otherwise.
func (o *LinkNicResponse) GetLinkNicId() string {
	if o == nil || o.LinkNicId == nil {
		var ret string
		return ret
	}
	return *o.LinkNicId
}

// GetLinkNicIdOk returns a tuple with the LinkNicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNicResponse) GetLinkNicIdOk() (*string, bool) {
	if o == nil || o.LinkNicId == nil {
		return nil, false
	}
	return o.LinkNicId, true
}

// HasLinkNicId returns a boolean if a field has been set.
func (o *LinkNicResponse) HasLinkNicId() bool {
	if o != nil && o.LinkNicId != nil {
		return true
	}

	return false
}

// SetLinkNicId gets a reference to the given string and assigns it to the LinkNicId field.
func (o *LinkNicResponse) SetLinkNicId(v string) {
	o.LinkNicId = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *LinkNicResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNicResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *LinkNicResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *LinkNicResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o LinkNicResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkNicId != nil {
		toSerialize["LinkNicId"] = o.LinkNicId
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableLinkNicResponse struct {
	value *LinkNicResponse
	isSet bool
}

func (v NullableLinkNicResponse) Get() *LinkNicResponse {
	return v.value
}

func (v *NullableLinkNicResponse) Set(val *LinkNicResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkNicResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkNicResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkNicResponse(val *LinkNicResponse) *NullableLinkNicResponse {
	return &NullableLinkNicResponse{value: val, isSet: true}
}

func (v NullableLinkNicResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkNicResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
