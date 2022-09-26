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

// UpdateApiAccessRuleResponse struct for UpdateApiAccessRuleResponse
type UpdateApiAccessRuleResponse struct {
	ApiAccessRule   *ApiAccessRule   `json:"ApiAccessRule,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateApiAccessRuleResponse instantiates a new UpdateApiAccessRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiAccessRuleResponse() *UpdateApiAccessRuleResponse {
	this := UpdateApiAccessRuleResponse{}
	return &this
}

// NewUpdateApiAccessRuleResponseWithDefaults instantiates a new UpdateApiAccessRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiAccessRuleResponseWithDefaults() *UpdateApiAccessRuleResponse {
	this := UpdateApiAccessRuleResponse{}
	return &this
}

// GetApiAccessRule returns the ApiAccessRule field value if set, zero value otherwise.
func (o *UpdateApiAccessRuleResponse) GetApiAccessRule() ApiAccessRule {
	if o == nil || o.ApiAccessRule == nil {
		var ret ApiAccessRule
		return ret
	}
	return *o.ApiAccessRule
}

// GetApiAccessRuleOk returns a tuple with the ApiAccessRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessRuleResponse) GetApiAccessRuleOk() (*ApiAccessRule, bool) {
	if o == nil || o.ApiAccessRule == nil {
		return nil, false
	}
	return o.ApiAccessRule, true
}

// HasApiAccessRule returns a boolean if a field has been set.
func (o *UpdateApiAccessRuleResponse) HasApiAccessRule() bool {
	if o != nil && o.ApiAccessRule != nil {
		return true
	}

	return false
}

// SetApiAccessRule gets a reference to the given ApiAccessRule and assigns it to the ApiAccessRule field.
func (o *UpdateApiAccessRuleResponse) SetApiAccessRule(v ApiAccessRule) {
	o.ApiAccessRule = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateApiAccessRuleResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiAccessRuleResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateApiAccessRuleResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateApiAccessRuleResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateApiAccessRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiAccessRule != nil {
		toSerialize["ApiAccessRule"] = o.ApiAccessRule
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateApiAccessRuleResponse struct {
	value *UpdateApiAccessRuleResponse
	isSet bool
}

func (v NullableUpdateApiAccessRuleResponse) Get() *UpdateApiAccessRuleResponse {
	return v.value
}

func (v *NullableUpdateApiAccessRuleResponse) Set(val *UpdateApiAccessRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiAccessRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiAccessRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiAccessRuleResponse(val *UpdateApiAccessRuleResponse) *NullableUpdateApiAccessRuleResponse {
	return &NullableUpdateApiAccessRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateApiAccessRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiAccessRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
