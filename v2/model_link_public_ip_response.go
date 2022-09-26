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

// LinkPublicIpResponse struct for LinkPublicIpResponse
type LinkPublicIpResponse struct {
	// (Net only) The ID representing the association of the public IP with the VM or the NIC.
	LinkPublicIpId  *string          `json:"LinkPublicIpId,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewLinkPublicIpResponse instantiates a new LinkPublicIpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkPublicIpResponse() *LinkPublicIpResponse {
	this := LinkPublicIpResponse{}
	return &this
}

// NewLinkPublicIpResponseWithDefaults instantiates a new LinkPublicIpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkPublicIpResponseWithDefaults() *LinkPublicIpResponse {
	this := LinkPublicIpResponse{}
	return &this
}

// GetLinkPublicIpId returns the LinkPublicIpId field value if set, zero value otherwise.
func (o *LinkPublicIpResponse) GetLinkPublicIpId() string {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret
	}
	return *o.LinkPublicIpId
}

// GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpResponse) GetLinkPublicIpIdOk() (*string, bool) {
	if o == nil || o.LinkPublicIpId == nil {
		return nil, false
	}
	return o.LinkPublicIpId, true
}

// HasLinkPublicIpId returns a boolean if a field has been set.
func (o *LinkPublicIpResponse) HasLinkPublicIpId() bool {
	if o != nil && o.LinkPublicIpId != nil {
		return true
	}

	return false
}

// SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.
func (o *LinkPublicIpResponse) SetLinkPublicIpId(v string) {
	o.LinkPublicIpId = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *LinkPublicIpResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *LinkPublicIpResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *LinkPublicIpResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o LinkPublicIpResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkPublicIpId != nil {
		toSerialize["LinkPublicIpId"] = o.LinkPublicIpId
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableLinkPublicIpResponse struct {
	value *LinkPublicIpResponse
	isSet bool
}

func (v NullableLinkPublicIpResponse) Get() *LinkPublicIpResponse {
	return v.value
}

func (v *NullableLinkPublicIpResponse) Set(val *LinkPublicIpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkPublicIpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkPublicIpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkPublicIpResponse(val *LinkPublicIpResponse) *NullableLinkPublicIpResponse {
	return &NullableLinkPublicIpResponse{value: val, isSet: true}
}

func (v NullableLinkPublicIpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkPublicIpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
