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

// ReadVmGroupsResponse struct for ReadVmGroupsResponse
type ReadVmGroupsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more VM groups.
	VmGroups *[]VmGroup `json:"VmGroups,omitempty"`
}

// NewReadVmGroupsResponse instantiates a new ReadVmGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVmGroupsResponse() *ReadVmGroupsResponse {
	this := ReadVmGroupsResponse{}
	return &this
}

// NewReadVmGroupsResponseWithDefaults instantiates a new ReadVmGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVmGroupsResponseWithDefaults() *ReadVmGroupsResponse {
	this := ReadVmGroupsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVmGroupsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmGroupsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVmGroupsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVmGroupsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVmGroups returns the VmGroups field value if set, zero value otherwise.
func (o *ReadVmGroupsResponse) GetVmGroups() []VmGroup {
	if o == nil || o.VmGroups == nil {
		var ret []VmGroup
		return ret
	}
	return *o.VmGroups
}

// GetVmGroupsOk returns a tuple with the VmGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmGroupsResponse) GetVmGroupsOk() (*[]VmGroup, bool) {
	if o == nil || o.VmGroups == nil {
		return nil, false
	}
	return o.VmGroups, true
}

// HasVmGroups returns a boolean if a field has been set.
func (o *ReadVmGroupsResponse) HasVmGroups() bool {
	if o != nil && o.VmGroups != nil {
		return true
	}

	return false
}

// SetVmGroups gets a reference to the given []VmGroup and assigns it to the VmGroups field.
func (o *ReadVmGroupsResponse) SetVmGroups(v []VmGroup) {
	o.VmGroups = &v
}

func (o ReadVmGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.VmGroups != nil {
		toSerialize["VmGroups"] = o.VmGroups
	}
	return json.Marshal(toSerialize)
}

type NullableReadVmGroupsResponse struct {
	value *ReadVmGroupsResponse
	isSet bool
}

func (v NullableReadVmGroupsResponse) Get() *ReadVmGroupsResponse {
	return v.value
}

func (v *NullableReadVmGroupsResponse) Set(val *ReadVmGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVmGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVmGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVmGroupsResponse(val *ReadVmGroupsResponse) *NullableReadVmGroupsResponse {
	return &NullableReadVmGroupsResponse{value: val, isSet: true}
}

func (v NullableReadVmGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVmGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
