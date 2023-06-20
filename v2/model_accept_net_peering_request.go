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

// AcceptNetPeeringRequest struct for AcceptNetPeeringRequest
type AcceptNetPeeringRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net peering you want to accept.
	NetPeeringId string `json:"NetPeeringId"`
}

// NewAcceptNetPeeringRequest instantiates a new AcceptNetPeeringRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptNetPeeringRequest(netPeeringId string) *AcceptNetPeeringRequest {
	this := AcceptNetPeeringRequest{}
	this.NetPeeringId = netPeeringId
	return &this
}

// NewAcceptNetPeeringRequestWithDefaults instantiates a new AcceptNetPeeringRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptNetPeeringRequestWithDefaults() *AcceptNetPeeringRequest {
	this := AcceptNetPeeringRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AcceptNetPeeringRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptNetPeeringRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AcceptNetPeeringRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AcceptNetPeeringRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetPeeringId returns the NetPeeringId field value
func (o *AcceptNetPeeringRequest) GetNetPeeringId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetPeeringId
}

// GetNetPeeringIdOk returns a tuple with the NetPeeringId field value
// and a boolean to check if the value has been set.
func (o *AcceptNetPeeringRequest) GetNetPeeringIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetPeeringId, true
}

// SetNetPeeringId sets field value
func (o *AcceptNetPeeringRequest) SetNetPeeringId(v string) {
	o.NetPeeringId = v
}

func (o AcceptNetPeeringRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NetPeeringId"] = o.NetPeeringId
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptNetPeeringRequest struct {
	value *AcceptNetPeeringRequest
	isSet bool
}

func (v NullableAcceptNetPeeringRequest) Get() *AcceptNetPeeringRequest {
	return v.value
}

func (v *NullableAcceptNetPeeringRequest) Set(val *AcceptNetPeeringRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptNetPeeringRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptNetPeeringRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptNetPeeringRequest(val *AcceptNetPeeringRequest) *NullableAcceptNetPeeringRequest {
	return &NullableAcceptNetPeeringRequest{value: val, isSet: true}
}

func (v NullableAcceptNetPeeringRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptNetPeeringRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
