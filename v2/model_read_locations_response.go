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

// ReadLocationsResponse struct for ReadLocationsResponse
type ReadLocationsResponse struct {
	// Information about one or more locations.
	Locations       *[]Location      `json:"Locations,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadLocationsResponse instantiates a new ReadLocationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadLocationsResponse() *ReadLocationsResponse {
	this := ReadLocationsResponse{}
	return &this
}

// NewReadLocationsResponseWithDefaults instantiates a new ReadLocationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadLocationsResponseWithDefaults() *ReadLocationsResponse {
	this := ReadLocationsResponse{}
	return &this
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ReadLocationsResponse) GetLocations() []Location {
	if o == nil || o.Locations == nil {
		var ret []Location
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLocationsResponse) GetLocationsOk() (*[]Location, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ReadLocationsResponse) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []Location and assigns it to the Locations field.
func (o *ReadLocationsResponse) SetLocations(v []Location) {
	o.Locations = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadLocationsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLocationsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadLocationsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadLocationsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadLocationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Locations != nil {
		toSerialize["Locations"] = o.Locations
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadLocationsResponse struct {
	value *ReadLocationsResponse
	isSet bool
}

func (v NullableReadLocationsResponse) Get() *ReadLocationsResponse {
	return v.value
}

func (v *NullableReadLocationsResponse) Set(val *ReadLocationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadLocationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadLocationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadLocationsResponse(val *ReadLocationsResponse) *NullableReadLocationsResponse {
	return &NullableReadLocationsResponse{value: val, isSet: true}
}

func (v NullableReadLocationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadLocationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
