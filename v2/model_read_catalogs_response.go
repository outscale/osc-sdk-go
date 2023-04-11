/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadCatalogsResponse ReadCatalogsResponse
type ReadCatalogsResponse struct {
	// Information about one or more catalogs.
	Catalogs        *[]Catalogs      `json:"Catalogs,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadCatalogsResponse instantiates a new ReadCatalogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadCatalogsResponse() *ReadCatalogsResponse {
	this := ReadCatalogsResponse{}
	return &this
}

// NewReadCatalogsResponseWithDefaults instantiates a new ReadCatalogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadCatalogsResponseWithDefaults() *ReadCatalogsResponse {
	this := ReadCatalogsResponse{}
	return &this
}

// GetCatalogs returns the Catalogs field value if set, zero value otherwise.
func (o *ReadCatalogsResponse) GetCatalogs() []Catalogs {
	if o == nil || o.Catalogs == nil {
		var ret []Catalogs
		return ret
	}
	return *o.Catalogs
}

// GetCatalogsOk returns a tuple with the Catalogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadCatalogsResponse) GetCatalogsOk() (*[]Catalogs, bool) {
	if o == nil || o.Catalogs == nil {
		return nil, false
	}
	return o.Catalogs, true
}

// HasCatalogs returns a boolean if a field has been set.
func (o *ReadCatalogsResponse) HasCatalogs() bool {
	if o != nil && o.Catalogs != nil {
		return true
	}

	return false
}

// SetCatalogs gets a reference to the given []Catalogs and assigns it to the Catalogs field.
func (o *ReadCatalogsResponse) SetCatalogs(v []Catalogs) {
	o.Catalogs = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadCatalogsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadCatalogsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadCatalogsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadCatalogsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadCatalogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Catalogs != nil {
		toSerialize["Catalogs"] = o.Catalogs
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadCatalogsResponse struct {
	value *ReadCatalogsResponse
	isSet bool
}

func (v NullableReadCatalogsResponse) Get() *ReadCatalogsResponse {
	return v.value
}

func (v *NullableReadCatalogsResponse) Set(val *ReadCatalogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadCatalogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadCatalogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadCatalogsResponse(val *ReadCatalogsResponse) *NullableReadCatalogsResponse {
	return &NullableReadCatalogsResponse{value: val, isSet: true}
}

func (v NullableReadCatalogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadCatalogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}