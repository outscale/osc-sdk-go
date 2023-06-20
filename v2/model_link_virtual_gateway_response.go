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

// LinkVirtualGatewayResponse struct for LinkVirtualGatewayResponse
type LinkVirtualGatewayResponse struct {
	NetToVirtualGatewayLink *NetToVirtualGatewayLink `json:"NetToVirtualGatewayLink,omitempty"`
	ResponseContext         *ResponseContext         `json:"ResponseContext,omitempty"`
}

// NewLinkVirtualGatewayResponse instantiates a new LinkVirtualGatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkVirtualGatewayResponse() *LinkVirtualGatewayResponse {
	this := LinkVirtualGatewayResponse{}
	return &this
}

// NewLinkVirtualGatewayResponseWithDefaults instantiates a new LinkVirtualGatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkVirtualGatewayResponseWithDefaults() *LinkVirtualGatewayResponse {
	this := LinkVirtualGatewayResponse{}
	return &this
}

// GetNetToVirtualGatewayLink returns the NetToVirtualGatewayLink field value if set, zero value otherwise.
func (o *LinkVirtualGatewayResponse) GetNetToVirtualGatewayLink() NetToVirtualGatewayLink {
	if o == nil || o.NetToVirtualGatewayLink == nil {
		var ret NetToVirtualGatewayLink
		return ret
	}
	return *o.NetToVirtualGatewayLink
}

// GetNetToVirtualGatewayLinkOk returns a tuple with the NetToVirtualGatewayLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkVirtualGatewayResponse) GetNetToVirtualGatewayLinkOk() (*NetToVirtualGatewayLink, bool) {
	if o == nil || o.NetToVirtualGatewayLink == nil {
		return nil, false
	}
	return o.NetToVirtualGatewayLink, true
}

// HasNetToVirtualGatewayLink returns a boolean if a field has been set.
func (o *LinkVirtualGatewayResponse) HasNetToVirtualGatewayLink() bool {
	if o != nil && o.NetToVirtualGatewayLink != nil {
		return true
	}

	return false
}

// SetNetToVirtualGatewayLink gets a reference to the given NetToVirtualGatewayLink and assigns it to the NetToVirtualGatewayLink field.
func (o *LinkVirtualGatewayResponse) SetNetToVirtualGatewayLink(v NetToVirtualGatewayLink) {
	o.NetToVirtualGatewayLink = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *LinkVirtualGatewayResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkVirtualGatewayResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *LinkVirtualGatewayResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *LinkVirtualGatewayResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o LinkVirtualGatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetToVirtualGatewayLink != nil {
		toSerialize["NetToVirtualGatewayLink"] = o.NetToVirtualGatewayLink
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableLinkVirtualGatewayResponse struct {
	value *LinkVirtualGatewayResponse
	isSet bool
}

func (v NullableLinkVirtualGatewayResponse) Get() *LinkVirtualGatewayResponse {
	return v.value
}

func (v *NullableLinkVirtualGatewayResponse) Set(val *LinkVirtualGatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkVirtualGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkVirtualGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkVirtualGatewayResponse(val *LinkVirtualGatewayResponse) *NullableLinkVirtualGatewayResponse {
	return &NullableLinkVirtualGatewayResponse{value: val, isSet: true}
}

func (v NullableLinkVirtualGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkVirtualGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
