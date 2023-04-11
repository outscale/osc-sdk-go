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

// ReadLoadBalancersResponse struct for ReadLoadBalancersResponse
type ReadLoadBalancersResponse struct {
	// Information about one or more load balancers.
	LoadBalancers   *[]LoadBalancer  `json:"LoadBalancers,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadLoadBalancersResponse instantiates a new ReadLoadBalancersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadLoadBalancersResponse() *ReadLoadBalancersResponse {
	this := ReadLoadBalancersResponse{}
	return &this
}

// NewReadLoadBalancersResponseWithDefaults instantiates a new ReadLoadBalancersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadLoadBalancersResponseWithDefaults() *ReadLoadBalancersResponse {
	this := ReadLoadBalancersResponse{}
	return &this
}

// GetLoadBalancers returns the LoadBalancers field value if set, zero value otherwise.
func (o *ReadLoadBalancersResponse) GetLoadBalancers() []LoadBalancer {
	if o == nil || o.LoadBalancers == nil {
		var ret []LoadBalancer
		return ret
	}
	return *o.LoadBalancers
}

// GetLoadBalancersOk returns a tuple with the LoadBalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLoadBalancersResponse) GetLoadBalancersOk() (*[]LoadBalancer, bool) {
	if o == nil || o.LoadBalancers == nil {
		return nil, false
	}
	return o.LoadBalancers, true
}

// HasLoadBalancers returns a boolean if a field has been set.
func (o *ReadLoadBalancersResponse) HasLoadBalancers() bool {
	if o != nil && o.LoadBalancers != nil {
		return true
	}

	return false
}

// SetLoadBalancers gets a reference to the given []LoadBalancer and assigns it to the LoadBalancers field.
func (o *ReadLoadBalancersResponse) SetLoadBalancers(v []LoadBalancer) {
	o.LoadBalancers = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadLoadBalancersResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLoadBalancersResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadLoadBalancersResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadLoadBalancersResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadLoadBalancersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoadBalancers != nil {
		toSerialize["LoadBalancers"] = o.LoadBalancers
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadLoadBalancersResponse struct {
	value *ReadLoadBalancersResponse
	isSet bool
}

func (v NullableReadLoadBalancersResponse) Get() *ReadLoadBalancersResponse {
	return v.value
}

func (v *NullableReadLoadBalancersResponse) Set(val *ReadLoadBalancersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadLoadBalancersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadLoadBalancersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadLoadBalancersResponse(val *ReadLoadBalancersResponse) *NullableReadLoadBalancersResponse {
	return &NullableReadLoadBalancersResponse{value: val, isSet: true}
}

func (v NullableReadLoadBalancersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadLoadBalancersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
