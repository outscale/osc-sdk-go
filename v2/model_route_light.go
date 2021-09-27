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

// RouteLight Information about the route.
type RouteLight struct {
	// The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).
	DestinationIpRange *string `json:"DestinationIpRange,omitempty"`
	// The type of route (always `static`).
	RouteType *string `json:"RouteType,omitempty"`
	// The current state of the static route (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
}

// NewRouteLight instantiates a new RouteLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteLight() *RouteLight {
	this := RouteLight{}
	return &this
}

// NewRouteLightWithDefaults instantiates a new RouteLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteLightWithDefaults() *RouteLight {
	this := RouteLight{}
	return &this
}

// GetDestinationIpRange returns the DestinationIpRange field value if set, zero value otherwise.
func (o *RouteLight) GetDestinationIpRange() string {
	if o == nil || o.DestinationIpRange == nil {
		var ret string
		return ret
	}
	return *o.DestinationIpRange
}

// GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetDestinationIpRangeOk() (*string, bool) {
	if o == nil || o.DestinationIpRange == nil {
		return nil, false
	}
	return o.DestinationIpRange, true
}

// HasDestinationIpRange returns a boolean if a field has been set.
func (o *RouteLight) HasDestinationIpRange() bool {
	if o != nil && o.DestinationIpRange != nil {
		return true
	}

	return false
}

// SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.
func (o *RouteLight) SetDestinationIpRange(v string) {
	o.DestinationIpRange = &v
}

// GetRouteType returns the RouteType field value if set, zero value otherwise.
func (o *RouteLight) GetRouteType() string {
	if o == nil || o.RouteType == nil {
		var ret string
		return ret
	}
	return *o.RouteType
}

// GetRouteTypeOk returns a tuple with the RouteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetRouteTypeOk() (*string, bool) {
	if o == nil || o.RouteType == nil {
		return nil, false
	}
	return o.RouteType, true
}

// HasRouteType returns a boolean if a field has been set.
func (o *RouteLight) HasRouteType() bool {
	if o != nil && o.RouteType != nil {
		return true
	}

	return false
}

// SetRouteType gets a reference to the given string and assigns it to the RouteType field.
func (o *RouteLight) SetRouteType(v string) {
	o.RouteType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RouteLight) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteLight) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RouteLight) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RouteLight) SetState(v string) {
	o.State = &v
}

func (o RouteLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationIpRange != nil {
		toSerialize["DestinationIpRange"] = o.DestinationIpRange
	}
	if o.RouteType != nil {
		toSerialize["RouteType"] = o.RouteType
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableRouteLight struct {
	value *RouteLight
	isSet bool
}

func (v NullableRouteLight) Get() *RouteLight {
	return v.value
}

func (v *NullableRouteLight) Set(val *RouteLight) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteLight) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteLight(val *RouteLight) *NullableRouteLight {
	return &NullableRouteLight{value: val, isSet: true}
}

func (v NullableRouteLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
