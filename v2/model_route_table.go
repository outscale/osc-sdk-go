/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// RouteTable Information about the route table.
type RouteTable struct {
	// One or more associations between the route table and Subnets.
	LinkRouteTables *[]LinkRouteTable `json:"LinkRouteTables,omitempty"`
	// The ID of the Net for the route table.
	NetId *string `json:"NetId,omitempty"`
	// Information about virtual gateways propagating routes.
	RoutePropagatingVirtualGateways *[]RoutePropagatingVirtualGateway `json:"RoutePropagatingVirtualGateways,omitempty"`
	// The ID of the route table.
	RouteTableId *string `json:"RouteTableId,omitempty"`
	// One or more routes in the route table.
	Routes *[]Route `json:"Routes,omitempty"`
	// One or more tags associated with the route table.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewRouteTable instantiates a new RouteTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteTable() *RouteTable {
	this := RouteTable{}
	return &this
}

// NewRouteTableWithDefaults instantiates a new RouteTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteTableWithDefaults() *RouteTable {
	this := RouteTable{}
	return &this
}

// GetLinkRouteTables returns the LinkRouteTables field value if set, zero value otherwise.
func (o *RouteTable) GetLinkRouteTables() []LinkRouteTable {
	if o == nil || o.LinkRouteTables == nil {
		var ret []LinkRouteTable
		return ret
	}
	return *o.LinkRouteTables
}

// GetLinkRouteTablesOk returns a tuple with the LinkRouteTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTable) GetLinkRouteTablesOk() (*[]LinkRouteTable, bool) {
	if o == nil || o.LinkRouteTables == nil {
		return nil, false
	}
	return o.LinkRouteTables, true
}

// HasLinkRouteTables returns a boolean if a field has been set.
func (o *RouteTable) HasLinkRouteTables() bool {
	if o != nil && o.LinkRouteTables != nil {
		return true
	}

	return false
}

// SetLinkRouteTables gets a reference to the given []LinkRouteTable and assigns it to the LinkRouteTables field.
func (o *RouteTable) SetLinkRouteTables(v []LinkRouteTable) {
	o.LinkRouteTables = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *RouteTable) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTable) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *RouteTable) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *RouteTable) SetNetId(v string) {
	o.NetId = &v
}

// GetRoutePropagatingVirtualGateways returns the RoutePropagatingVirtualGateways field value if set, zero value otherwise.
func (o *RouteTable) GetRoutePropagatingVirtualGateways() []RoutePropagatingVirtualGateway {
	if o == nil || o.RoutePropagatingVirtualGateways == nil {
		var ret []RoutePropagatingVirtualGateway
		return ret
	}
	return *o.RoutePropagatingVirtualGateways
}

// GetRoutePropagatingVirtualGatewaysOk returns a tuple with the RoutePropagatingVirtualGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTable) GetRoutePropagatingVirtualGatewaysOk() (*[]RoutePropagatingVirtualGateway, bool) {
	if o == nil || o.RoutePropagatingVirtualGateways == nil {
		return nil, false
	}
	return o.RoutePropagatingVirtualGateways, true
}

// HasRoutePropagatingVirtualGateways returns a boolean if a field has been set.
func (o *RouteTable) HasRoutePropagatingVirtualGateways() bool {
	if o != nil && o.RoutePropagatingVirtualGateways != nil {
		return true
	}

	return false
}

// SetRoutePropagatingVirtualGateways gets a reference to the given []RoutePropagatingVirtualGateway and assigns it to the RoutePropagatingVirtualGateways field.
func (o *RouteTable) SetRoutePropagatingVirtualGateways(v []RoutePropagatingVirtualGateway) {
	o.RoutePropagatingVirtualGateways = &v
}

// GetRouteTableId returns the RouteTableId field value if set, zero value otherwise.
func (o *RouteTable) GetRouteTableId() string {
	if o == nil || o.RouteTableId == nil {
		var ret string
		return ret
	}
	return *o.RouteTableId
}

// GetRouteTableIdOk returns a tuple with the RouteTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTable) GetRouteTableIdOk() (*string, bool) {
	if o == nil || o.RouteTableId == nil {
		return nil, false
	}
	return o.RouteTableId, true
}

// HasRouteTableId returns a boolean if a field has been set.
func (o *RouteTable) HasRouteTableId() bool {
	if o != nil && o.RouteTableId != nil {
		return true
	}

	return false
}

// SetRouteTableId gets a reference to the given string and assigns it to the RouteTableId field.
func (o *RouteTable) SetRouteTableId(v string) {
	o.RouteTableId = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *RouteTable) GetRoutes() []Route {
	if o == nil || o.Routes == nil {
		var ret []Route
		return ret
	}
	return *o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTable) GetRoutesOk() (*[]Route, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *RouteTable) HasRoutes() bool {
	if o != nil && o.Routes != nil {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *RouteTable) SetRoutes(v []Route) {
	o.Routes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RouteTable) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTable) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RouteTable) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *RouteTable) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o RouteTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkRouteTables != nil {
		toSerialize["LinkRouteTables"] = o.LinkRouteTables
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.RoutePropagatingVirtualGateways != nil {
		toSerialize["RoutePropagatingVirtualGateways"] = o.RoutePropagatingVirtualGateways
	}
	if o.RouteTableId != nil {
		toSerialize["RouteTableId"] = o.RouteTableId
	}
	if o.Routes != nil {
		toSerialize["Routes"] = o.Routes
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableRouteTable struct {
	value *RouteTable
	isSet bool
}

func (v NullableRouteTable) Get() *RouteTable {
	return v.value
}

func (v *NullableRouteTable) Set(val *RouteTable) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTable) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTable(val *RouteTable) *NullableRouteTable {
	return &NullableRouteTable{value: val, isSet: true}
}

func (v NullableRouteTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
