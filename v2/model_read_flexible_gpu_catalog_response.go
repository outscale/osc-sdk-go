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

// ReadFlexibleGpuCatalogResponse struct for ReadFlexibleGpuCatalogResponse
type ReadFlexibleGpuCatalogResponse struct {
	// Information about one or more fGPUs available in the public catalog.
	FlexibleGpuCatalog *[]FlexibleGpuCatalog `json:"FlexibleGpuCatalog,omitempty"`
	ResponseContext    *ResponseContext      `json:"ResponseContext,omitempty"`
}

// NewReadFlexibleGpuCatalogResponse instantiates a new ReadFlexibleGpuCatalogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadFlexibleGpuCatalogResponse() *ReadFlexibleGpuCatalogResponse {
	this := ReadFlexibleGpuCatalogResponse{}
	return &this
}

// NewReadFlexibleGpuCatalogResponseWithDefaults instantiates a new ReadFlexibleGpuCatalogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadFlexibleGpuCatalogResponseWithDefaults() *ReadFlexibleGpuCatalogResponse {
	this := ReadFlexibleGpuCatalogResponse{}
	return &this
}

// GetFlexibleGpuCatalog returns the FlexibleGpuCatalog field value if set, zero value otherwise.
func (o *ReadFlexibleGpuCatalogResponse) GetFlexibleGpuCatalog() []FlexibleGpuCatalog {
	if o == nil || o.FlexibleGpuCatalog == nil {
		var ret []FlexibleGpuCatalog
		return ret
	}
	return *o.FlexibleGpuCatalog
}

// GetFlexibleGpuCatalogOk returns a tuple with the FlexibleGpuCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFlexibleGpuCatalogResponse) GetFlexibleGpuCatalogOk() (*[]FlexibleGpuCatalog, bool) {
	if o == nil || o.FlexibleGpuCatalog == nil {
		return nil, false
	}
	return o.FlexibleGpuCatalog, true
}

// HasFlexibleGpuCatalog returns a boolean if a field has been set.
func (o *ReadFlexibleGpuCatalogResponse) HasFlexibleGpuCatalog() bool {
	if o != nil && o.FlexibleGpuCatalog != nil {
		return true
	}

	return false
}

// SetFlexibleGpuCatalog gets a reference to the given []FlexibleGpuCatalog and assigns it to the FlexibleGpuCatalog field.
func (o *ReadFlexibleGpuCatalogResponse) SetFlexibleGpuCatalog(v []FlexibleGpuCatalog) {
	o.FlexibleGpuCatalog = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadFlexibleGpuCatalogResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFlexibleGpuCatalogResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadFlexibleGpuCatalogResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadFlexibleGpuCatalogResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadFlexibleGpuCatalogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlexibleGpuCatalog != nil {
		toSerialize["FlexibleGpuCatalog"] = o.FlexibleGpuCatalog
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadFlexibleGpuCatalogResponse struct {
	value *ReadFlexibleGpuCatalogResponse
	isSet bool
}

func (v NullableReadFlexibleGpuCatalogResponse) Get() *ReadFlexibleGpuCatalogResponse {
	return v.value
}

func (v *NullableReadFlexibleGpuCatalogResponse) Set(val *ReadFlexibleGpuCatalogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadFlexibleGpuCatalogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadFlexibleGpuCatalogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadFlexibleGpuCatalogResponse(val *ReadFlexibleGpuCatalogResponse) *NullableReadFlexibleGpuCatalogResponse {
	return &NullableReadFlexibleGpuCatalogResponse{value: val, isSet: true}
}

func (v NullableReadFlexibleGpuCatalogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadFlexibleGpuCatalogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
