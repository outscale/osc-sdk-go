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

// DeleteLoadBalancerListenersResponse struct for DeleteLoadBalancerListenersResponse
type DeleteLoadBalancerListenersResponse struct {
	LoadBalancer    *LoadBalancer    `json:"LoadBalancer,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewDeleteLoadBalancerListenersResponse instantiates a new DeleteLoadBalancerListenersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteLoadBalancerListenersResponse() *DeleteLoadBalancerListenersResponse {
	this := DeleteLoadBalancerListenersResponse{}
	return &this
}

// NewDeleteLoadBalancerListenersResponseWithDefaults instantiates a new DeleteLoadBalancerListenersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteLoadBalancerListenersResponseWithDefaults() *DeleteLoadBalancerListenersResponse {
	this := DeleteLoadBalancerListenersResponse{}
	return &this
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise.
func (o *DeleteLoadBalancerListenersResponse) GetLoadBalancer() LoadBalancer {
	if o == nil || o.LoadBalancer == nil {
		var ret LoadBalancer
		return ret
	}
	return *o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerListenersResponse) GetLoadBalancerOk() (*LoadBalancer, bool) {
	if o == nil || o.LoadBalancer == nil {
		return nil, false
	}
	return o.LoadBalancer, true
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *DeleteLoadBalancerListenersResponse) HasLoadBalancer() bool {
	if o != nil && o.LoadBalancer != nil {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given LoadBalancer and assigns it to the LoadBalancer field.
func (o *DeleteLoadBalancerListenersResponse) SetLoadBalancer(v LoadBalancer) {
	o.LoadBalancer = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteLoadBalancerListenersResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerListenersResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteLoadBalancerListenersResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteLoadBalancerListenersResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o DeleteLoadBalancerListenersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoadBalancer != nil {
		toSerialize["LoadBalancer"] = o.LoadBalancer
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteLoadBalancerListenersResponse struct {
	value *DeleteLoadBalancerListenersResponse
	isSet bool
}

func (v NullableDeleteLoadBalancerListenersResponse) Get() *DeleteLoadBalancerListenersResponse {
	return v.value
}

func (v *NullableDeleteLoadBalancerListenersResponse) Set(val *DeleteLoadBalancerListenersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteLoadBalancerListenersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteLoadBalancerListenersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteLoadBalancerListenersResponse(val *DeleteLoadBalancerListenersResponse) *NullableDeleteLoadBalancerListenersResponse {
	return &NullableDeleteLoadBalancerListenersResponse{value: val, isSet: true}
}

func (v NullableDeleteLoadBalancerListenersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteLoadBalancerListenersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
