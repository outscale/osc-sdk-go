/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadDirectLinksResponse struct for ReadDirectLinksResponse
type ReadDirectLinksResponse struct {
	// Information about one or more DirectLinks.
	DirectLinks     *[]DirectLink    `json:"DirectLinks,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadDirectLinksResponse instantiates a new ReadDirectLinksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadDirectLinksResponse() *ReadDirectLinksResponse {
	this := ReadDirectLinksResponse{}
	return &this
}

// NewReadDirectLinksResponseWithDefaults instantiates a new ReadDirectLinksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadDirectLinksResponseWithDefaults() *ReadDirectLinksResponse {
	this := ReadDirectLinksResponse{}
	return &this
}

// GetDirectLinks returns the DirectLinks field value if set, zero value otherwise.
func (o *ReadDirectLinksResponse) GetDirectLinks() []DirectLink {
	if o == nil || o.DirectLinks == nil {
		var ret []DirectLink
		return ret
	}
	return *o.DirectLinks
}

// GetDirectLinksOk returns a tuple with the DirectLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadDirectLinksResponse) GetDirectLinksOk() (*[]DirectLink, bool) {
	if o == nil || o.DirectLinks == nil {
		return nil, false
	}
	return o.DirectLinks, true
}

// HasDirectLinks returns a boolean if a field has been set.
func (o *ReadDirectLinksResponse) HasDirectLinks() bool {
	if o != nil && o.DirectLinks != nil {
		return true
	}

	return false
}

// SetDirectLinks gets a reference to the given []DirectLink and assigns it to the DirectLinks field.
func (o *ReadDirectLinksResponse) SetDirectLinks(v []DirectLink) {
	o.DirectLinks = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadDirectLinksResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadDirectLinksResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadDirectLinksResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadDirectLinksResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadDirectLinksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DirectLinks != nil {
		toSerialize["DirectLinks"] = o.DirectLinks
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadDirectLinksResponse struct {
	value *ReadDirectLinksResponse
	isSet bool
}

func (v NullableReadDirectLinksResponse) Get() *ReadDirectLinksResponse {
	return v.value
}

func (v *NullableReadDirectLinksResponse) Set(val *ReadDirectLinksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadDirectLinksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadDirectLinksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadDirectLinksResponse(val *ReadDirectLinksResponse) *NullableReadDirectLinksResponse {
	return &NullableReadDirectLinksResponse{value: val, isSet: true}
}

func (v NullableReadDirectLinksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadDirectLinksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
