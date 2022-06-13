/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersLoadBalancer One or more filters.
type FiltersLoadBalancer struct {
	// The names of the load balancers.
	LoadBalancerNames *[]string `json:"LoadBalancerNames,omitempty"`
}

// NewFiltersLoadBalancer instantiates a new FiltersLoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersLoadBalancer() *FiltersLoadBalancer {
	this := FiltersLoadBalancer{}
	return &this
}

// NewFiltersLoadBalancerWithDefaults instantiates a new FiltersLoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersLoadBalancerWithDefaults() *FiltersLoadBalancer {
	this := FiltersLoadBalancer{}
	return &this
}

// GetLoadBalancerNames returns the LoadBalancerNames field value if set, zero value otherwise.
func (o *FiltersLoadBalancer) GetLoadBalancerNames() []string {
	if o == nil || o.LoadBalancerNames == nil {
		var ret []string
		return ret
	}
	return *o.LoadBalancerNames
}

// GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersLoadBalancer) GetLoadBalancerNamesOk() (*[]string, bool) {
	if o == nil || o.LoadBalancerNames == nil {
		return nil, false
	}
	return o.LoadBalancerNames, true
}

// HasLoadBalancerNames returns a boolean if a field has been set.
func (o *FiltersLoadBalancer) HasLoadBalancerNames() bool {
	if o != nil && o.LoadBalancerNames != nil {
		return true
	}

	return false
}

// SetLoadBalancerNames gets a reference to the given []string and assigns it to the LoadBalancerNames field.
func (o *FiltersLoadBalancer) SetLoadBalancerNames(v []string) {
	o.LoadBalancerNames = &v
}

func (o FiltersLoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoadBalancerNames != nil {
		toSerialize["LoadBalancerNames"] = o.LoadBalancerNames
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersLoadBalancer struct {
	value *FiltersLoadBalancer
	isSet bool
}

func (v NullableFiltersLoadBalancer) Get() *FiltersLoadBalancer {
	return v.value
}

func (v *NullableFiltersLoadBalancer) Set(val *FiltersLoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersLoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersLoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersLoadBalancer(val *FiltersLoadBalancer) *NullableFiltersLoadBalancer {
	return &NullableFiltersLoadBalancer{value: val, isSet: true}
}

func (v NullableFiltersLoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersLoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
