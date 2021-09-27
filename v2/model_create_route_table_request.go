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

// CreateRouteTableRequest struct for CreateRouteTableRequest
type CreateRouteTableRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net for which you want to create a route table.
	NetId string `json:"NetId"`
}

// NewCreateRouteTableRequest instantiates a new CreateRouteTableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRouteTableRequest(netId string) *CreateRouteTableRequest {
	this := CreateRouteTableRequest{}
	this.NetId = netId
	return &this
}

// NewCreateRouteTableRequestWithDefaults instantiates a new CreateRouteTableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRouteTableRequestWithDefaults() *CreateRouteTableRequest {
	this := CreateRouteTableRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateRouteTableRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteTableRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateRouteTableRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateRouteTableRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetId returns the NetId field value
func (o *CreateRouteTableRequest) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *CreateRouteTableRequest) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *CreateRouteTableRequest) SetNetId(v string) {
	o.NetId = v
}

func (o CreateRouteTableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NetId"] = o.NetId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRouteTableRequest struct {
	value *CreateRouteTableRequest
	isSet bool
}

func (v NullableCreateRouteTableRequest) Get() *CreateRouteTableRequest {
	return v.value
}

func (v *NullableCreateRouteTableRequest) Set(val *CreateRouteTableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRouteTableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRouteTableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRouteTableRequest(val *CreateRouteTableRequest) *NullableCreateRouteTableRequest {
	return &NullableCreateRouteTableRequest{value: val, isSet: true}
}

func (v NullableCreateRouteTableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRouteTableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
