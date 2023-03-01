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

// ReadInternetServicesResponse struct for ReadInternetServicesResponse
type ReadInternetServicesResponse struct {
	// Information about one or more Internet services.
	InternetServices *[]InternetService `json:"InternetServices,omitempty"`
	ResponseContext  *ResponseContext   `json:"ResponseContext,omitempty"`
}

// NewReadInternetServicesResponse instantiates a new ReadInternetServicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadInternetServicesResponse() *ReadInternetServicesResponse {
	this := ReadInternetServicesResponse{}
	return &this
}

// NewReadInternetServicesResponseWithDefaults instantiates a new ReadInternetServicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadInternetServicesResponseWithDefaults() *ReadInternetServicesResponse {
	this := ReadInternetServicesResponse{}
	return &this
}

// GetInternetServices returns the InternetServices field value if set, zero value otherwise.
func (o *ReadInternetServicesResponse) GetInternetServices() []InternetService {
	if o == nil || o.InternetServices == nil {
		var ret []InternetService
		return ret
	}
	return *o.InternetServices
}

// GetInternetServicesOk returns a tuple with the InternetServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadInternetServicesResponse) GetInternetServicesOk() (*[]InternetService, bool) {
	if o == nil || o.InternetServices == nil {
		return nil, false
	}
	return o.InternetServices, true
}

// HasInternetServices returns a boolean if a field has been set.
func (o *ReadInternetServicesResponse) HasInternetServices() bool {
	if o != nil && o.InternetServices != nil {
		return true
	}

	return false
}

// SetInternetServices gets a reference to the given []InternetService and assigns it to the InternetServices field.
func (o *ReadInternetServicesResponse) SetInternetServices(v []InternetService) {
	o.InternetServices = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadInternetServicesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadInternetServicesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadInternetServicesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadInternetServicesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadInternetServicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InternetServices != nil {
		toSerialize["InternetServices"] = o.InternetServices
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadInternetServicesResponse struct {
	value *ReadInternetServicesResponse
	isSet bool
}

func (v NullableReadInternetServicesResponse) Get() *ReadInternetServicesResponse {
	return v.value
}

func (v *NullableReadInternetServicesResponse) Set(val *ReadInternetServicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadInternetServicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadInternetServicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadInternetServicesResponse(val *ReadInternetServicesResponse) *NullableReadInternetServicesResponse {
	return &NullableReadInternetServicesResponse{value: val, isSet: true}
}

func (v NullableReadInternetServicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadInternetServicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
