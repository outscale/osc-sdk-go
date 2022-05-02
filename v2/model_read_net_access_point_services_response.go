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

// ReadNetAccessPointServicesResponse struct for ReadNetAccessPointServicesResponse
type ReadNetAccessPointServicesResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// The names of the services you can use for Net access points.
	Services *[]Service `json:"Services,omitempty"`
}

// NewReadNetAccessPointServicesResponse instantiates a new ReadNetAccessPointServicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadNetAccessPointServicesResponse() *ReadNetAccessPointServicesResponse {
	this := ReadNetAccessPointServicesResponse{}
	return &this
}

// NewReadNetAccessPointServicesResponseWithDefaults instantiates a new ReadNetAccessPointServicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadNetAccessPointServicesResponseWithDefaults() *ReadNetAccessPointServicesResponse {
	this := ReadNetAccessPointServicesResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNetAccessPointServicesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetAccessPointServicesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNetAccessPointServicesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNetAccessPointServicesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ReadNetAccessPointServicesResponse) GetServices() []Service {
	if o == nil || o.Services == nil {
		var ret []Service
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetAccessPointServicesResponse) GetServicesOk() (*[]Service, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ReadNetAccessPointServicesResponse) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []Service and assigns it to the Services field.
func (o *ReadNetAccessPointServicesResponse) SetServices(v []Service) {
	o.Services = &v
}

func (o ReadNetAccessPointServicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableReadNetAccessPointServicesResponse struct {
	value *ReadNetAccessPointServicesResponse
	isSet bool
}

func (v NullableReadNetAccessPointServicesResponse) Get() *ReadNetAccessPointServicesResponse {
	return v.value
}

func (v *NullableReadNetAccessPointServicesResponse) Set(val *ReadNetAccessPointServicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadNetAccessPointServicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadNetAccessPointServicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadNetAccessPointServicesResponse(val *ReadNetAccessPointServicesResponse) *NullableReadNetAccessPointServicesResponse {
	return &NullableReadNetAccessPointServicesResponse{value: val, isSet: true}
}

func (v NullableReadNetAccessPointServicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadNetAccessPointServicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
