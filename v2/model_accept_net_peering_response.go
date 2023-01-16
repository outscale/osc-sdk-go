/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// AcceptNetPeeringResponse struct for AcceptNetPeeringResponse
type AcceptNetPeeringResponse struct {
	NetPeering      *NetPeering      `json:"NetPeering,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewAcceptNetPeeringResponse instantiates a new AcceptNetPeeringResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptNetPeeringResponse() *AcceptNetPeeringResponse {
	this := AcceptNetPeeringResponse{}
	return &this
}

// NewAcceptNetPeeringResponseWithDefaults instantiates a new AcceptNetPeeringResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptNetPeeringResponseWithDefaults() *AcceptNetPeeringResponse {
	this := AcceptNetPeeringResponse{}
	return &this
}

// GetNetPeering returns the NetPeering field value if set, zero value otherwise.
func (o *AcceptNetPeeringResponse) GetNetPeering() NetPeering {
	if o == nil || o.NetPeering == nil {
		var ret NetPeering
		return ret
	}
	return *o.NetPeering
}

// GetNetPeeringOk returns a tuple with the NetPeering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptNetPeeringResponse) GetNetPeeringOk() (*NetPeering, bool) {
	if o == nil || o.NetPeering == nil {
		return nil, false
	}
	return o.NetPeering, true
}

// HasNetPeering returns a boolean if a field has been set.
func (o *AcceptNetPeeringResponse) HasNetPeering() bool {
	if o != nil && o.NetPeering != nil {
		return true
	}

	return false
}

// SetNetPeering gets a reference to the given NetPeering and assigns it to the NetPeering field.
func (o *AcceptNetPeeringResponse) SetNetPeering(v NetPeering) {
	o.NetPeering = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *AcceptNetPeeringResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptNetPeeringResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *AcceptNetPeeringResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *AcceptNetPeeringResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o AcceptNetPeeringResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetPeering != nil {
		toSerialize["NetPeering"] = o.NetPeering
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptNetPeeringResponse struct {
	value *AcceptNetPeeringResponse
	isSet bool
}

func (v NullableAcceptNetPeeringResponse) Get() *AcceptNetPeeringResponse {
	return v.value
}

func (v *NullableAcceptNetPeeringResponse) Set(val *AcceptNetPeeringResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptNetPeeringResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptNetPeeringResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptNetPeeringResponse(val *AcceptNetPeeringResponse) *NullableAcceptNetPeeringResponse {
	return &NullableAcceptNetPeeringResponse{value: val, isSet: true}
}

func (v NullableAcceptNetPeeringResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptNetPeeringResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
