/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateNatServiceResponse struct for CreateNatServiceResponse
type CreateNatServiceResponse struct {
	NatService      *NatService      `json:"NatService,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateNatServiceResponse instantiates a new CreateNatServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNatServiceResponse() *CreateNatServiceResponse {
	this := CreateNatServiceResponse{}
	return &this
}

// NewCreateNatServiceResponseWithDefaults instantiates a new CreateNatServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNatServiceResponseWithDefaults() *CreateNatServiceResponse {
	this := CreateNatServiceResponse{}
	return &this
}

// GetNatService returns the NatService field value if set, zero value otherwise.
func (o *CreateNatServiceResponse) GetNatService() NatService {
	if o == nil || o.NatService == nil {
		var ret NatService
		return ret
	}
	return *o.NatService
}

// GetNatServiceOk returns a tuple with the NatService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNatServiceResponse) GetNatServiceOk() (*NatService, bool) {
	if o == nil || o.NatService == nil {
		return nil, false
	}
	return o.NatService, true
}

// HasNatService returns a boolean if a field has been set.
func (o *CreateNatServiceResponse) HasNatService() bool {
	if o != nil && o.NatService != nil {
		return true
	}

	return false
}

// SetNatService gets a reference to the given NatService and assigns it to the NatService field.
func (o *CreateNatServiceResponse) SetNatService(v NatService) {
	o.NatService = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateNatServiceResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNatServiceResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateNatServiceResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateNatServiceResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateNatServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NatService != nil {
		toSerialize["NatService"] = o.NatService
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNatServiceResponse struct {
	value *CreateNatServiceResponse
	isSet bool
}

func (v NullableCreateNatServiceResponse) Get() *CreateNatServiceResponse {
	return v.value
}

func (v *NullableCreateNatServiceResponse) Set(val *CreateNatServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNatServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNatServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNatServiceResponse(val *CreateNatServiceResponse) *NullableCreateNatServiceResponse {
	return &NullableCreateNatServiceResponse{value: val, isSet: true}
}

func (v NullableCreateNatServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNatServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
