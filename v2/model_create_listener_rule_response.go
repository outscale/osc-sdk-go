/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateListenerRuleResponse struct for CreateListenerRuleResponse
type CreateListenerRuleResponse struct {
	ListenerRule    *ListenerRule    `json:"ListenerRule,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateListenerRuleResponse instantiates a new CreateListenerRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateListenerRuleResponse() *CreateListenerRuleResponse {
	this := CreateListenerRuleResponse{}
	return &this
}

// NewCreateListenerRuleResponseWithDefaults instantiates a new CreateListenerRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateListenerRuleResponseWithDefaults() *CreateListenerRuleResponse {
	this := CreateListenerRuleResponse{}
	return &this
}

// GetListenerRule returns the ListenerRule field value if set, zero value otherwise.
func (o *CreateListenerRuleResponse) GetListenerRule() ListenerRule {
	if o == nil || o.ListenerRule == nil {
		var ret ListenerRule
		return ret
	}
	return *o.ListenerRule
}

// GetListenerRuleOk returns a tuple with the ListenerRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateListenerRuleResponse) GetListenerRuleOk() (*ListenerRule, bool) {
	if o == nil || o.ListenerRule == nil {
		return nil, false
	}
	return o.ListenerRule, true
}

// HasListenerRule returns a boolean if a field has been set.
func (o *CreateListenerRuleResponse) HasListenerRule() bool {
	if o != nil && o.ListenerRule != nil {
		return true
	}

	return false
}

// SetListenerRule gets a reference to the given ListenerRule and assigns it to the ListenerRule field.
func (o *CreateListenerRuleResponse) SetListenerRule(v ListenerRule) {
	o.ListenerRule = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateListenerRuleResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateListenerRuleResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateListenerRuleResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateListenerRuleResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateListenerRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListenerRule != nil {
		toSerialize["ListenerRule"] = o.ListenerRule
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateListenerRuleResponse struct {
	value *CreateListenerRuleResponse
	isSet bool
}

func (v NullableCreateListenerRuleResponse) Get() *CreateListenerRuleResponse {
	return v.value
}

func (v *NullableCreateListenerRuleResponse) Set(val *CreateListenerRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateListenerRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateListenerRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateListenerRuleResponse(val *CreateListenerRuleResponse) *NullableCreateListenerRuleResponse {
	return &NullableCreateListenerRuleResponse{value: val, isSet: true}
}

func (v NullableCreateListenerRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateListenerRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
