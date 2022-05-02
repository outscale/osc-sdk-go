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

// DeleteSecurityGroupRuleResponse struct for DeleteSecurityGroupRuleResponse
type DeleteSecurityGroupRuleResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	SecurityGroup   *SecurityGroup   `json:"SecurityGroup,omitempty"`
}

// NewDeleteSecurityGroupRuleResponse instantiates a new DeleteSecurityGroupRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSecurityGroupRuleResponse() *DeleteSecurityGroupRuleResponse {
	this := DeleteSecurityGroupRuleResponse{}
	return &this
}

// NewDeleteSecurityGroupRuleResponseWithDefaults instantiates a new DeleteSecurityGroupRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSecurityGroupRuleResponseWithDefaults() *DeleteSecurityGroupRuleResponse {
	this := DeleteSecurityGroupRuleResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteSecurityGroupRuleResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSecurityGroup returns the SecurityGroup field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleResponse) GetSecurityGroup() SecurityGroup {
	if o == nil || o.SecurityGroup == nil {
		var ret SecurityGroup
		return ret
	}
	return *o.SecurityGroup
}

// GetSecurityGroupOk returns a tuple with the SecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleResponse) GetSecurityGroupOk() (*SecurityGroup, bool) {
	if o == nil || o.SecurityGroup == nil {
		return nil, false
	}
	return o.SecurityGroup, true
}

// HasSecurityGroup returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleResponse) HasSecurityGroup() bool {
	if o != nil && o.SecurityGroup != nil {
		return true
	}

	return false
}

// SetSecurityGroup gets a reference to the given SecurityGroup and assigns it to the SecurityGroup field.
func (o *DeleteSecurityGroupRuleResponse) SetSecurityGroup(v SecurityGroup) {
	o.SecurityGroup = &v
}

func (o DeleteSecurityGroupRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.SecurityGroup != nil {
		toSerialize["SecurityGroup"] = o.SecurityGroup
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteSecurityGroupRuleResponse struct {
	value *DeleteSecurityGroupRuleResponse
	isSet bool
}

func (v NullableDeleteSecurityGroupRuleResponse) Get() *DeleteSecurityGroupRuleResponse {
	return v.value
}

func (v *NullableDeleteSecurityGroupRuleResponse) Set(val *DeleteSecurityGroupRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSecurityGroupRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSecurityGroupRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSecurityGroupRuleResponse(val *DeleteSecurityGroupRuleResponse) *NullableDeleteSecurityGroupRuleResponse {
	return &NullableDeleteSecurityGroupRuleResponse{value: val, isSet: true}
}

func (v NullableDeleteSecurityGroupRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSecurityGroupRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
