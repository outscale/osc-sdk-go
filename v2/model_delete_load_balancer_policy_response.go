/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteLoadBalancerPolicyResponse struct for DeleteLoadBalancerPolicyResponse
type DeleteLoadBalancerPolicyResponse struct {
	LoadBalancer    *LoadBalancer    `json:"LoadBalancer,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewDeleteLoadBalancerPolicyResponse instantiates a new DeleteLoadBalancerPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteLoadBalancerPolicyResponse() *DeleteLoadBalancerPolicyResponse {
	this := DeleteLoadBalancerPolicyResponse{}
	return &this
}

// NewDeleteLoadBalancerPolicyResponseWithDefaults instantiates a new DeleteLoadBalancerPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteLoadBalancerPolicyResponseWithDefaults() *DeleteLoadBalancerPolicyResponse {
	this := DeleteLoadBalancerPolicyResponse{}
	return &this
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise.
func (o *DeleteLoadBalancerPolicyResponse) GetLoadBalancer() LoadBalancer {
	if o == nil || o.LoadBalancer == nil {
		var ret LoadBalancer
		return ret
	}
	return *o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerPolicyResponse) GetLoadBalancerOk() (*LoadBalancer, bool) {
	if o == nil || o.LoadBalancer == nil {
		return nil, false
	}
	return o.LoadBalancer, true
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *DeleteLoadBalancerPolicyResponse) HasLoadBalancer() bool {
	if o != nil && o.LoadBalancer != nil {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given LoadBalancer and assigns it to the LoadBalancer field.
func (o *DeleteLoadBalancerPolicyResponse) SetLoadBalancer(v LoadBalancer) {
	o.LoadBalancer = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteLoadBalancerPolicyResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerPolicyResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteLoadBalancerPolicyResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteLoadBalancerPolicyResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o DeleteLoadBalancerPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoadBalancer != nil {
		toSerialize["LoadBalancer"] = o.LoadBalancer
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteLoadBalancerPolicyResponse struct {
	value *DeleteLoadBalancerPolicyResponse
	isSet bool
}

func (v NullableDeleteLoadBalancerPolicyResponse) Get() *DeleteLoadBalancerPolicyResponse {
	return v.value
}

func (v *NullableDeleteLoadBalancerPolicyResponse) Set(val *DeleteLoadBalancerPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteLoadBalancerPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteLoadBalancerPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteLoadBalancerPolicyResponse(val *DeleteLoadBalancerPolicyResponse) *NullableDeleteLoadBalancerPolicyResponse {
	return &NullableDeleteLoadBalancerPolicyResponse{value: val, isSet: true}
}

func (v NullableDeleteLoadBalancerPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteLoadBalancerPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
