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

// LinkLoadBalancerBackendMachinesResponse struct for LinkLoadBalancerBackendMachinesResponse
type LinkLoadBalancerBackendMachinesResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewLinkLoadBalancerBackendMachinesResponse instantiates a new LinkLoadBalancerBackendMachinesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkLoadBalancerBackendMachinesResponse() *LinkLoadBalancerBackendMachinesResponse {
	this := LinkLoadBalancerBackendMachinesResponse{}
	return &this
}

// NewLinkLoadBalancerBackendMachinesResponseWithDefaults instantiates a new LinkLoadBalancerBackendMachinesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkLoadBalancerBackendMachinesResponseWithDefaults() *LinkLoadBalancerBackendMachinesResponse {
	this := LinkLoadBalancerBackendMachinesResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *LinkLoadBalancerBackendMachinesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLoadBalancerBackendMachinesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *LinkLoadBalancerBackendMachinesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *LinkLoadBalancerBackendMachinesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o LinkLoadBalancerBackendMachinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableLinkLoadBalancerBackendMachinesResponse struct {
	value *LinkLoadBalancerBackendMachinesResponse
	isSet bool
}

func (v NullableLinkLoadBalancerBackendMachinesResponse) Get() *LinkLoadBalancerBackendMachinesResponse {
	return v.value
}

func (v *NullableLinkLoadBalancerBackendMachinesResponse) Set(val *LinkLoadBalancerBackendMachinesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkLoadBalancerBackendMachinesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkLoadBalancerBackendMachinesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkLoadBalancerBackendMachinesResponse(val *LinkLoadBalancerBackendMachinesResponse) *NullableLinkLoadBalancerBackendMachinesResponse {
	return &NullableLinkLoadBalancerBackendMachinesResponse{value: val, isSet: true}
}

func (v NullableLinkLoadBalancerBackendMachinesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkLoadBalancerBackendMachinesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
