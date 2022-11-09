/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LoadBalancerLight Information about the load balancer.
type LoadBalancerLight struct {
	// The name of the load balancer to which the listener is attached.
	LoadBalancerName string `json:"LoadBalancerName"`
	// The port of load balancer on which the load balancer is listening (between `1` and `65535` both included).
	LoadBalancerPort int32 `json:"LoadBalancerPort"`
}

// NewLoadBalancerLight instantiates a new LoadBalancerLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerLight(loadBalancerName string, loadBalancerPort int32) *LoadBalancerLight {
	this := LoadBalancerLight{}
	this.LoadBalancerName = loadBalancerName
	this.LoadBalancerPort = loadBalancerPort
	return &this
}

// NewLoadBalancerLightWithDefaults instantiates a new LoadBalancerLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerLightWithDefaults() *LoadBalancerLight {
	this := LoadBalancerLight{}
	return &this
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *LoadBalancerLight) GetLoadBalancerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerLight) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *LoadBalancerLight) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

// GetLoadBalancerPort returns the LoadBalancerPort field value
func (o *LoadBalancerLight) GetLoadBalancerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoadBalancerPort
}

// GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerLight) GetLoadBalancerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerPort, true
}

// SetLoadBalancerPort sets field value
func (o *LoadBalancerLight) SetLoadBalancerPort(v int32) {
	o.LoadBalancerPort = v
}

func (o LoadBalancerLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	if true {
		toSerialize["LoadBalancerPort"] = o.LoadBalancerPort
	}
	return json.Marshal(toSerialize)
}

type NullableLoadBalancerLight struct {
	value *LoadBalancerLight
	isSet bool
}

func (v NullableLoadBalancerLight) Get() *LoadBalancerLight {
	return v.value
}

func (v *NullableLoadBalancerLight) Set(val *LoadBalancerLight) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerLight) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerLight(val *LoadBalancerLight) *NullableLoadBalancerLight {
	return &NullableLoadBalancerLight{value: val, isSet: true}
}

func (v NullableLoadBalancerLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
