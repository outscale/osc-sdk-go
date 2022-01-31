/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersVpnConnection One or more filters.
type FiltersVpnConnection struct {
	// The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.
	BgpAsns *[]int32 `json:"BgpAsns,omitempty"`
	// The IDs of the client gateways.
	ClientGatewayIds *[]string `json:"ClientGatewayIds,omitempty"`
	// The types of the VPN connections (only `ipsec.1` is supported).
	ConnectionTypes *[]string `json:"ConnectionTypes,omitempty"`
	// The destination IP ranges.
	RouteDestinationIpRanges *[]string `json:"RouteDestinationIpRanges,omitempty"`
	// The states of the VPN connections (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	States *[]string `json:"States,omitempty"`
	// If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute).
	StaticRoutesOnly *bool `json:"StaticRoutesOnly,omitempty"`
	// The keys of the tags associated with the VPN connections.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the VPN connections.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the VPN connections, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
	// The IDs of the virtual gateways.
	VirtualGatewayIds *[]string `json:"VirtualGatewayIds,omitempty"`
	// The IDs of the VPN connections.
	VpnConnectionIds *[]string `json:"VpnConnectionIds,omitempty"`
}

// NewFiltersVpnConnection instantiates a new FiltersVpnConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersVpnConnection() *FiltersVpnConnection {
	this := FiltersVpnConnection{}
	return &this
}

// NewFiltersVpnConnectionWithDefaults instantiates a new FiltersVpnConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersVpnConnectionWithDefaults() *FiltersVpnConnection {
	this := FiltersVpnConnection{}
	return &this
}

// GetBgpAsns returns the BgpAsns field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetBgpAsns() []int32 {
	if o == nil || o.BgpAsns == nil {
		var ret []int32
		return ret
	}
	return *o.BgpAsns
}

// GetBgpAsnsOk returns a tuple with the BgpAsns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetBgpAsnsOk() (*[]int32, bool) {
	if o == nil || o.BgpAsns == nil {
		return nil, false
	}
	return o.BgpAsns, true
}

// HasBgpAsns returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasBgpAsns() bool {
	if o != nil && o.BgpAsns != nil {
		return true
	}

	return false
}

// SetBgpAsns gets a reference to the given []int32 and assigns it to the BgpAsns field.
func (o *FiltersVpnConnection) SetBgpAsns(v []int32) {
	o.BgpAsns = &v
}

// GetClientGatewayIds returns the ClientGatewayIds field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetClientGatewayIds() []string {
	if o == nil || o.ClientGatewayIds == nil {
		var ret []string
		return ret
	}
	return *o.ClientGatewayIds
}

// GetClientGatewayIdsOk returns a tuple with the ClientGatewayIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetClientGatewayIdsOk() (*[]string, bool) {
	if o == nil || o.ClientGatewayIds == nil {
		return nil, false
	}
	return o.ClientGatewayIds, true
}

// HasClientGatewayIds returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasClientGatewayIds() bool {
	if o != nil && o.ClientGatewayIds != nil {
		return true
	}

	return false
}

// SetClientGatewayIds gets a reference to the given []string and assigns it to the ClientGatewayIds field.
func (o *FiltersVpnConnection) SetClientGatewayIds(v []string) {
	o.ClientGatewayIds = &v
}

// GetConnectionTypes returns the ConnectionTypes field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetConnectionTypes() []string {
	if o == nil || o.ConnectionTypes == nil {
		var ret []string
		return ret
	}
	return *o.ConnectionTypes
}

// GetConnectionTypesOk returns a tuple with the ConnectionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetConnectionTypesOk() (*[]string, bool) {
	if o == nil || o.ConnectionTypes == nil {
		return nil, false
	}
	return o.ConnectionTypes, true
}

// HasConnectionTypes returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasConnectionTypes() bool {
	if o != nil && o.ConnectionTypes != nil {
		return true
	}

	return false
}

// SetConnectionTypes gets a reference to the given []string and assigns it to the ConnectionTypes field.
func (o *FiltersVpnConnection) SetConnectionTypes(v []string) {
	o.ConnectionTypes = &v
}

// GetRouteDestinationIpRanges returns the RouteDestinationIpRanges field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetRouteDestinationIpRanges() []string {
	if o == nil || o.RouteDestinationIpRanges == nil {
		var ret []string
		return ret
	}
	return *o.RouteDestinationIpRanges
}

// GetRouteDestinationIpRangesOk returns a tuple with the RouteDestinationIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetRouteDestinationIpRangesOk() (*[]string, bool) {
	if o == nil || o.RouteDestinationIpRanges == nil {
		return nil, false
	}
	return o.RouteDestinationIpRanges, true
}

// HasRouteDestinationIpRanges returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasRouteDestinationIpRanges() bool {
	if o != nil && o.RouteDestinationIpRanges != nil {
		return true
	}

	return false
}

// SetRouteDestinationIpRanges gets a reference to the given []string and assigns it to the RouteDestinationIpRanges field.
func (o *FiltersVpnConnection) SetRouteDestinationIpRanges(v []string) {
	o.RouteDestinationIpRanges = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetStatesOk() (*[]string, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersVpnConnection) SetStates(v []string) {
	o.States = &v
}

// GetStaticRoutesOnly returns the StaticRoutesOnly field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetStaticRoutesOnly() bool {
	if o == nil || o.StaticRoutesOnly == nil {
		var ret bool
		return ret
	}
	return *o.StaticRoutesOnly
}

// GetStaticRoutesOnlyOk returns a tuple with the StaticRoutesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetStaticRoutesOnlyOk() (*bool, bool) {
	if o == nil || o.StaticRoutesOnly == nil {
		return nil, false
	}
	return o.StaticRoutesOnly, true
}

// HasStaticRoutesOnly returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasStaticRoutesOnly() bool {
	if o != nil && o.StaticRoutesOnly != nil {
		return true
	}

	return false
}

// SetStaticRoutesOnly gets a reference to the given bool and assigns it to the StaticRoutesOnly field.
func (o *FiltersVpnConnection) SetStaticRoutesOnly(v bool) {
	o.StaticRoutesOnly = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersVpnConnection) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersVpnConnection) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersVpnConnection) SetTags(v []string) {
	o.Tags = &v
}

// GetVirtualGatewayIds returns the VirtualGatewayIds field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetVirtualGatewayIds() []string {
	if o == nil || o.VirtualGatewayIds == nil {
		var ret []string
		return ret
	}
	return *o.VirtualGatewayIds
}

// GetVirtualGatewayIdsOk returns a tuple with the VirtualGatewayIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetVirtualGatewayIdsOk() (*[]string, bool) {
	if o == nil || o.VirtualGatewayIds == nil {
		return nil, false
	}
	return o.VirtualGatewayIds, true
}

// HasVirtualGatewayIds returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasVirtualGatewayIds() bool {
	if o != nil && o.VirtualGatewayIds != nil {
		return true
	}

	return false
}

// SetVirtualGatewayIds gets a reference to the given []string and assigns it to the VirtualGatewayIds field.
func (o *FiltersVpnConnection) SetVirtualGatewayIds(v []string) {
	o.VirtualGatewayIds = &v
}

// GetVpnConnectionIds returns the VpnConnectionIds field value if set, zero value otherwise.
func (o *FiltersVpnConnection) GetVpnConnectionIds() []string {
	if o == nil || o.VpnConnectionIds == nil {
		var ret []string
		return ret
	}
	return *o.VpnConnectionIds
}

// GetVpnConnectionIdsOk returns a tuple with the VpnConnectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVpnConnection) GetVpnConnectionIdsOk() (*[]string, bool) {
	if o == nil || o.VpnConnectionIds == nil {
		return nil, false
	}
	return o.VpnConnectionIds, true
}

// HasVpnConnectionIds returns a boolean if a field has been set.
func (o *FiltersVpnConnection) HasVpnConnectionIds() bool {
	if o != nil && o.VpnConnectionIds != nil {
		return true
	}

	return false
}

// SetVpnConnectionIds gets a reference to the given []string and assigns it to the VpnConnectionIds field.
func (o *FiltersVpnConnection) SetVpnConnectionIds(v []string) {
	o.VpnConnectionIds = &v
}

func (o FiltersVpnConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BgpAsns != nil {
		toSerialize["BgpAsns"] = o.BgpAsns
	}
	if o.ClientGatewayIds != nil {
		toSerialize["ClientGatewayIds"] = o.ClientGatewayIds
	}
	if o.ConnectionTypes != nil {
		toSerialize["ConnectionTypes"] = o.ConnectionTypes
	}
	if o.RouteDestinationIpRanges != nil {
		toSerialize["RouteDestinationIpRanges"] = o.RouteDestinationIpRanges
	}
	if o.States != nil {
		toSerialize["States"] = o.States
	}
	if o.StaticRoutesOnly != nil {
		toSerialize["StaticRoutesOnly"] = o.StaticRoutesOnly
	}
	if o.TagKeys != nil {
		toSerialize["TagKeys"] = o.TagKeys
	}
	if o.TagValues != nil {
		toSerialize["TagValues"] = o.TagValues
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.VirtualGatewayIds != nil {
		toSerialize["VirtualGatewayIds"] = o.VirtualGatewayIds
	}
	if o.VpnConnectionIds != nil {
		toSerialize["VpnConnectionIds"] = o.VpnConnectionIds
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersVpnConnection struct {
	value *FiltersVpnConnection
	isSet bool
}

func (v NullableFiltersVpnConnection) Get() *FiltersVpnConnection {
	return v.value
}

func (v *NullableFiltersVpnConnection) Set(val *FiltersVpnConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersVpnConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersVpnConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersVpnConnection(val *FiltersVpnConnection) *NullableFiltersVpnConnection {
	return &NullableFiltersVpnConnection{value: val, isSet: true}
}

func (v NullableFiltersVpnConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersVpnConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
