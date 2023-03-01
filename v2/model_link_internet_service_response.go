/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkInternetServiceResponse struct for LinkInternetServiceResponse
type LinkInternetServiceResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewLinkInternetServiceResponse instantiates a new LinkInternetServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkInternetServiceResponse() *LinkInternetServiceResponse {
	this := LinkInternetServiceResponse{}
	return &this
}

// NewLinkInternetServiceResponseWithDefaults instantiates a new LinkInternetServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkInternetServiceResponseWithDefaults() *LinkInternetServiceResponse {
	this := LinkInternetServiceResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *LinkInternetServiceResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkInternetServiceResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *LinkInternetServiceResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *LinkInternetServiceResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o LinkInternetServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableLinkInternetServiceResponse struct {
	value *LinkInternetServiceResponse
	isSet bool
}

func (v NullableLinkInternetServiceResponse) Get() *LinkInternetServiceResponse {
	return v.value
}

func (v *NullableLinkInternetServiceResponse) Set(val *LinkInternetServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkInternetServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkInternetServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkInternetServiceResponse(val *LinkInternetServiceResponse) *NullableLinkInternetServiceResponse {
	return &NullableLinkInternetServiceResponse{value: val, isSet: true}
}

func (v NullableLinkInternetServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkInternetServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
