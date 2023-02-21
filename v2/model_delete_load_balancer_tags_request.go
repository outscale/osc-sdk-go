/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteLoadBalancerTagsRequest struct for DeleteLoadBalancerTagsRequest
type DeleteLoadBalancerTagsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more load balancer names.
	LoadBalancerNames []string `json:"LoadBalancerNames"`
	// One or more tags to delete from the load balancers.
	Tags []ResourceLoadBalancerTag `json:"Tags"`
}

// NewDeleteLoadBalancerTagsRequest instantiates a new DeleteLoadBalancerTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteLoadBalancerTagsRequest(loadBalancerNames []string, tags []ResourceLoadBalancerTag) *DeleteLoadBalancerTagsRequest {
	this := DeleteLoadBalancerTagsRequest{}
	this.LoadBalancerNames = loadBalancerNames
	this.Tags = tags
	return &this
}

// NewDeleteLoadBalancerTagsRequestWithDefaults instantiates a new DeleteLoadBalancerTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteLoadBalancerTagsRequestWithDefaults() *DeleteLoadBalancerTagsRequest {
	this := DeleteLoadBalancerTagsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteLoadBalancerTagsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerTagsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteLoadBalancerTagsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteLoadBalancerTagsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLoadBalancerNames returns the LoadBalancerNames field value
func (o *DeleteLoadBalancerTagsRequest) GetLoadBalancerNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LoadBalancerNames
}

// GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field value
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerTagsRequest) GetLoadBalancerNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerNames, true
}

// SetLoadBalancerNames sets field value
func (o *DeleteLoadBalancerTagsRequest) SetLoadBalancerNames(v []string) {
	o.LoadBalancerNames = v
}

// GetTags returns the Tags field value
func (o *DeleteLoadBalancerTagsRequest) GetTags() []ResourceLoadBalancerTag {
	if o == nil {
		var ret []ResourceLoadBalancerTag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *DeleteLoadBalancerTagsRequest) GetTagsOk() (*[]ResourceLoadBalancerTag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *DeleteLoadBalancerTagsRequest) SetTags(v []ResourceLoadBalancerTag) {
	o.Tags = v
}

func (o DeleteLoadBalancerTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["LoadBalancerNames"] = o.LoadBalancerNames
	}
	if true {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteLoadBalancerTagsRequest struct {
	value *DeleteLoadBalancerTagsRequest
	isSet bool
}

func (v NullableDeleteLoadBalancerTagsRequest) Get() *DeleteLoadBalancerTagsRequest {
	return v.value
}

func (v *NullableDeleteLoadBalancerTagsRequest) Set(val *DeleteLoadBalancerTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteLoadBalancerTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteLoadBalancerTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteLoadBalancerTagsRequest(val *DeleteLoadBalancerTagsRequest) *NullableDeleteLoadBalancerTagsRequest {
	return &NullableDeleteLoadBalancerTagsRequest{value: val, isSet: true}
}

func (v NullableDeleteLoadBalancerTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteLoadBalancerTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
