/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateLoadBalancerPolicyResponse struct for CreateLoadBalancerPolicyResponse
type CreateLoadBalancerPolicyResponse struct {
	LoadBalancer *LoadBalancer `json:"LoadBalancer,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateLoadBalancerPolicyResponse instantiates a new CreateLoadBalancerPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerPolicyResponse() *CreateLoadBalancerPolicyResponse {
	this := CreateLoadBalancerPolicyResponse{}
	return &this
}

// NewCreateLoadBalancerPolicyResponseWithDefaults instantiates a new CreateLoadBalancerPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerPolicyResponseWithDefaults() *CreateLoadBalancerPolicyResponse {
	this := CreateLoadBalancerPolicyResponse{}
	return &this
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise.
func (o *CreateLoadBalancerPolicyResponse) GetLoadBalancer() LoadBalancer {
	if o == nil || o.LoadBalancer == nil {
		var ret LoadBalancer
		return ret
	}
	return *o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPolicyResponse) GetLoadBalancerOk() (*LoadBalancer, bool) {
	if o == nil || o.LoadBalancer == nil {
		return nil, false
	}
	return o.LoadBalancer, true
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *CreateLoadBalancerPolicyResponse) HasLoadBalancer() bool {
	if o != nil && o.LoadBalancer != nil {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given LoadBalancer and assigns it to the LoadBalancer field.
func (o *CreateLoadBalancerPolicyResponse) SetLoadBalancer(v LoadBalancer) {
	o.LoadBalancer = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateLoadBalancerPolicyResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPolicyResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateLoadBalancerPolicyResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateLoadBalancerPolicyResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateLoadBalancerPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoadBalancer != nil {
		toSerialize["LoadBalancer"] = o.LoadBalancer
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLoadBalancerPolicyResponse struct {
	value *CreateLoadBalancerPolicyResponse
	isSet bool
}

func (v NullableCreateLoadBalancerPolicyResponse) Get() *CreateLoadBalancerPolicyResponse {
	return v.value
}

func (v *NullableCreateLoadBalancerPolicyResponse) Set(val *CreateLoadBalancerPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLoadBalancerPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLoadBalancerPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLoadBalancerPolicyResponse(val *CreateLoadBalancerPolicyResponse) *NullableCreateLoadBalancerPolicyResponse {
	return &NullableCreateLoadBalancerPolicyResponse{value: val, isSet: true}
}

func (v NullableCreateLoadBalancerPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLoadBalancerPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


