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

// ReadSecurityGroupsResponse struct for ReadSecurityGroupsResponse
type ReadSecurityGroupsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more security groups.
	SecurityGroups *[]SecurityGroup `json:"SecurityGroups,omitempty"`
}

// NewReadSecurityGroupsResponse instantiates a new ReadSecurityGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadSecurityGroupsResponse() *ReadSecurityGroupsResponse {
	this := ReadSecurityGroupsResponse{}
	return &this
}

// NewReadSecurityGroupsResponseWithDefaults instantiates a new ReadSecurityGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadSecurityGroupsResponseWithDefaults() *ReadSecurityGroupsResponse {
	this := ReadSecurityGroupsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadSecurityGroupsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSecurityGroupsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadSecurityGroupsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadSecurityGroupsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *ReadSecurityGroupsResponse) GetSecurityGroups() []SecurityGroup {
	if o == nil || o.SecurityGroups == nil {
		var ret []SecurityGroup
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSecurityGroupsResponse) GetSecurityGroupsOk() (*[]SecurityGroup, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *ReadSecurityGroupsResponse) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []SecurityGroup and assigns it to the SecurityGroups field.
func (o *ReadSecurityGroupsResponse) SetSecurityGroups(v []SecurityGroup) {
	o.SecurityGroups = &v
}

func (o ReadSecurityGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	return json.Marshal(toSerialize)
}

type NullableReadSecurityGroupsResponse struct {
	value *ReadSecurityGroupsResponse
	isSet bool
}

func (v NullableReadSecurityGroupsResponse) Get() *ReadSecurityGroupsResponse {
	return v.value
}

func (v *NullableReadSecurityGroupsResponse) Set(val *ReadSecurityGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadSecurityGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadSecurityGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadSecurityGroupsResponse(val *ReadSecurityGroupsResponse) *NullableReadSecurityGroupsResponse {
	return &NullableReadSecurityGroupsResponse{value: val, isSet: true}
}

func (v NullableReadSecurityGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadSecurityGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
