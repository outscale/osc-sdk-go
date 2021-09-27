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

// LinkRouteTableRequest struct for LinkRouteTableRequest
type LinkRouteTableRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the route table.
	RouteTableId string `json:"RouteTableId"`
	// The ID of the Subnet.
	SubnetId string `json:"SubnetId"`
}

// NewLinkRouteTableRequest instantiates a new LinkRouteTableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkRouteTableRequest(routeTableId string, subnetId string) *LinkRouteTableRequest {
	this := LinkRouteTableRequest{}
	this.RouteTableId = routeTableId
	this.SubnetId = subnetId
	return &this
}

// NewLinkRouteTableRequestWithDefaults instantiates a new LinkRouteTableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkRouteTableRequestWithDefaults() *LinkRouteTableRequest {
	this := LinkRouteTableRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkRouteTableRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkRouteTableRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkRouteTableRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkRouteTableRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetRouteTableId returns the RouteTableId field value
func (o *LinkRouteTableRequest) GetRouteTableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteTableId
}

// GetRouteTableIdOk returns a tuple with the RouteTableId field value
// and a boolean to check if the value has been set.
func (o *LinkRouteTableRequest) GetRouteTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteTableId, true
}

// SetRouteTableId sets field value
func (o *LinkRouteTableRequest) SetRouteTableId(v string) {
	o.RouteTableId = v
}

// GetSubnetId returns the SubnetId field value
func (o *LinkRouteTableRequest) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *LinkRouteTableRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *LinkRouteTableRequest) SetSubnetId(v string) {
	o.SubnetId = v
}

func (o LinkRouteTableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["RouteTableId"] = o.RouteTableId
	}
	if true {
		toSerialize["SubnetId"] = o.SubnetId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkRouteTableRequest struct {
	value *LinkRouteTableRequest
	isSet bool
}

func (v NullableLinkRouteTableRequest) Get() *LinkRouteTableRequest {
	return v.value
}

func (v *NullableLinkRouteTableRequest) Set(val *LinkRouteTableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkRouteTableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkRouteTableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkRouteTableRequest(val *LinkRouteTableRequest) *NullableLinkRouteTableRequest {
	return &NullableLinkRouteTableRequest{value: val, isSet: true}
}

func (v NullableLinkRouteTableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkRouteTableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
