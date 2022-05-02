/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateNetAccessPointRequest struct for CreateNetAccessPointRequest
type CreateNetAccessPointRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net.
	NetId string `json:"NetId"`
	// One or more IDs of route tables to use for the connection.
	RouteTableIds *[]string `json:"RouteTableIds,omitempty"`
	// The name of the service (in the format `com.outscale.region.service`).
	ServiceName string `json:"ServiceName"`
}

// NewCreateNetAccessPointRequest instantiates a new CreateNetAccessPointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetAccessPointRequest(netId string, serviceName string) *CreateNetAccessPointRequest {
	this := CreateNetAccessPointRequest{}
	this.NetId = netId
	this.ServiceName = serviceName
	return &this
}

// NewCreateNetAccessPointRequestWithDefaults instantiates a new CreateNetAccessPointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetAccessPointRequestWithDefaults() *CreateNetAccessPointRequest {
	this := CreateNetAccessPointRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateNetAccessPointRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetAccessPointRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateNetAccessPointRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateNetAccessPointRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetId returns the NetId field value
func (o *CreateNetAccessPointRequest) GetNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value
// and a boolean to check if the value has been set.
func (o *CreateNetAccessPointRequest) GetNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetId, true
}

// SetNetId sets field value
func (o *CreateNetAccessPointRequest) SetNetId(v string) {
	o.NetId = v
}

// GetRouteTableIds returns the RouteTableIds field value if set, zero value otherwise.
func (o *CreateNetAccessPointRequest) GetRouteTableIds() []string {
	if o == nil || o.RouteTableIds == nil {
		var ret []string
		return ret
	}
	return *o.RouteTableIds
}

// GetRouteTableIdsOk returns a tuple with the RouteTableIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetAccessPointRequest) GetRouteTableIdsOk() (*[]string, bool) {
	if o == nil || o.RouteTableIds == nil {
		return nil, false
	}
	return o.RouteTableIds, true
}

// HasRouteTableIds returns a boolean if a field has been set.
func (o *CreateNetAccessPointRequest) HasRouteTableIds() bool {
	if o != nil && o.RouteTableIds != nil {
		return true
	}

	return false
}

// SetRouteTableIds gets a reference to the given []string and assigns it to the RouteTableIds field.
func (o *CreateNetAccessPointRequest) SetRouteTableIds(v []string) {
	o.RouteTableIds = &v
}

// GetServiceName returns the ServiceName field value
func (o *CreateNetAccessPointRequest) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *CreateNetAccessPointRequest) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *CreateNetAccessPointRequest) SetServiceName(v string) {
	o.ServiceName = v
}

func (o CreateNetAccessPointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NetId"] = o.NetId
	}
	if o.RouteTableIds != nil {
		toSerialize["RouteTableIds"] = o.RouteTableIds
	}
	if true {
		toSerialize["ServiceName"] = o.ServiceName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetAccessPointRequest struct {
	value *CreateNetAccessPointRequest
	isSet bool
}

func (v NullableCreateNetAccessPointRequest) Get() *CreateNetAccessPointRequest {
	return v.value
}

func (v *NullableCreateNetAccessPointRequest) Set(val *CreateNetAccessPointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetAccessPointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetAccessPointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetAccessPointRequest(val *CreateNetAccessPointRequest) *NullableCreateNetAccessPointRequest {
	return &NullableCreateNetAccessPointRequest{value: val, isSet: true}
}

func (v NullableCreateNetAccessPointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetAccessPointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
