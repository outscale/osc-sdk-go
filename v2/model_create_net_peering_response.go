/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateNetPeeringResponse struct for CreateNetPeeringResponse
type CreateNetPeeringResponse struct {
	NetPeering      *NetPeering      `json:"NetPeering,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateNetPeeringResponse instantiates a new CreateNetPeeringResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetPeeringResponse() *CreateNetPeeringResponse {
	this := CreateNetPeeringResponse{}
	return &this
}

// NewCreateNetPeeringResponseWithDefaults instantiates a new CreateNetPeeringResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetPeeringResponseWithDefaults() *CreateNetPeeringResponse {
	this := CreateNetPeeringResponse{}
	return &this
}

// GetNetPeering returns the NetPeering field value if set, zero value otherwise.
func (o *CreateNetPeeringResponse) GetNetPeering() NetPeering {
	if o == nil || o.NetPeering == nil {
		var ret NetPeering
		return ret
	}
	return *o.NetPeering
}

// GetNetPeeringOk returns a tuple with the NetPeering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetPeeringResponse) GetNetPeeringOk() (*NetPeering, bool) {
	if o == nil || o.NetPeering == nil {
		return nil, false
	}
	return o.NetPeering, true
}

// HasNetPeering returns a boolean if a field has been set.
func (o *CreateNetPeeringResponse) HasNetPeering() bool {
	if o != nil && o.NetPeering != nil {
		return true
	}

	return false
}

// SetNetPeering gets a reference to the given NetPeering and assigns it to the NetPeering field.
func (o *CreateNetPeeringResponse) SetNetPeering(v NetPeering) {
	o.NetPeering = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateNetPeeringResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetPeeringResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateNetPeeringResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateNetPeeringResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateNetPeeringResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetPeering != nil {
		toSerialize["NetPeering"] = o.NetPeering
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetPeeringResponse struct {
	value *CreateNetPeeringResponse
	isSet bool
}

func (v NullableCreateNetPeeringResponse) Get() *CreateNetPeeringResponse {
	return v.value
}

func (v *NullableCreateNetPeeringResponse) Set(val *CreateNetPeeringResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetPeeringResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetPeeringResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetPeeringResponse(val *CreateNetPeeringResponse) *NullableCreateNetPeeringResponse {
	return &NullableCreateNetPeeringResponse{value: val, isSet: true}
}

func (v NullableCreateNetPeeringResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetPeeringResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
