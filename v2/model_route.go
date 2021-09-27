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

// Route Information about the route.
type Route struct {
	// The method used to create the route.
	CreationMethod *string `json:"CreationMethod,omitempty"`
	// The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).
	DestinationIpRange *string `json:"DestinationIpRange,omitempty"`
	// The ID of the OUTSCALE service.
	DestinationServiceId *string `json:"DestinationServiceId,omitempty"`
	// The ID of the Internet service or virtual gateway attached to the Net.
	GatewayId *string `json:"GatewayId,omitempty"`
	// The ID of a NAT service attached to the Net.
	NatServiceId *string `json:"NatServiceId,omitempty"`
	// The ID of the Net access point.
	NetAccessPointId *string `json:"NetAccessPointId,omitempty"`
	// The ID of the Net peering connection.
	NetPeeringId *string `json:"NetPeeringId,omitempty"`
	// The ID of the NIC.
	NicId *string `json:"NicId,omitempty"`
	// The state of a route in the route table (`active` \\| `blackhole`). The `blackhole` state indicates that the target of the route is not available.
	State *string `json:"State,omitempty"`
	// The account ID of the owner of the VM.
	VmAccountId *string `json:"VmAccountId,omitempty"`
	// The ID of a VM specified in a route in the table.
	VmId *string `json:"VmId,omitempty"`
}

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute() *Route {
	this := Route{}
	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	return &this
}

// GetCreationMethod returns the CreationMethod field value if set, zero value otherwise.
func (o *Route) GetCreationMethod() string {
	if o == nil || o.CreationMethod == nil {
		var ret string
		return ret
	}
	return *o.CreationMethod
}

// GetCreationMethodOk returns a tuple with the CreationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetCreationMethodOk() (*string, bool) {
	if o == nil || o.CreationMethod == nil {
		return nil, false
	}
	return o.CreationMethod, true
}

// HasCreationMethod returns a boolean if a field has been set.
func (o *Route) HasCreationMethod() bool {
	if o != nil && o.CreationMethod != nil {
		return true
	}

	return false
}

// SetCreationMethod gets a reference to the given string and assigns it to the CreationMethod field.
func (o *Route) SetCreationMethod(v string) {
	o.CreationMethod = &v
}

// GetDestinationIpRange returns the DestinationIpRange field value if set, zero value otherwise.
func (o *Route) GetDestinationIpRange() string {
	if o == nil || o.DestinationIpRange == nil {
		var ret string
		return ret
	}
	return *o.DestinationIpRange
}

// GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetDestinationIpRangeOk() (*string, bool) {
	if o == nil || o.DestinationIpRange == nil {
		return nil, false
	}
	return o.DestinationIpRange, true
}

// HasDestinationIpRange returns a boolean if a field has been set.
func (o *Route) HasDestinationIpRange() bool {
	if o != nil && o.DestinationIpRange != nil {
		return true
	}

	return false
}

// SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.
func (o *Route) SetDestinationIpRange(v string) {
	o.DestinationIpRange = &v
}

// GetDestinationServiceId returns the DestinationServiceId field value if set, zero value otherwise.
func (o *Route) GetDestinationServiceId() string {
	if o == nil || o.DestinationServiceId == nil {
		var ret string
		return ret
	}
	return *o.DestinationServiceId
}

// GetDestinationServiceIdOk returns a tuple with the DestinationServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetDestinationServiceIdOk() (*string, bool) {
	if o == nil || o.DestinationServiceId == nil {
		return nil, false
	}
	return o.DestinationServiceId, true
}

// HasDestinationServiceId returns a boolean if a field has been set.
func (o *Route) HasDestinationServiceId() bool {
	if o != nil && o.DestinationServiceId != nil {
		return true
	}

	return false
}

// SetDestinationServiceId gets a reference to the given string and assigns it to the DestinationServiceId field.
func (o *Route) SetDestinationServiceId(v string) {
	o.DestinationServiceId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *Route) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *Route) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *Route) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetNatServiceId returns the NatServiceId field value if set, zero value otherwise.
func (o *Route) GetNatServiceId() string {
	if o == nil || o.NatServiceId == nil {
		var ret string
		return ret
	}
	return *o.NatServiceId
}

// GetNatServiceIdOk returns a tuple with the NatServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetNatServiceIdOk() (*string, bool) {
	if o == nil || o.NatServiceId == nil {
		return nil, false
	}
	return o.NatServiceId, true
}

// HasNatServiceId returns a boolean if a field has been set.
func (o *Route) HasNatServiceId() bool {
	if o != nil && o.NatServiceId != nil {
		return true
	}

	return false
}

// SetNatServiceId gets a reference to the given string and assigns it to the NatServiceId field.
func (o *Route) SetNatServiceId(v string) {
	o.NatServiceId = &v
}

// GetNetAccessPointId returns the NetAccessPointId field value if set, zero value otherwise.
func (o *Route) GetNetAccessPointId() string {
	if o == nil || o.NetAccessPointId == nil {
		var ret string
		return ret
	}
	return *o.NetAccessPointId
}

// GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetNetAccessPointIdOk() (*string, bool) {
	if o == nil || o.NetAccessPointId == nil {
		return nil, false
	}
	return o.NetAccessPointId, true
}

// HasNetAccessPointId returns a boolean if a field has been set.
func (o *Route) HasNetAccessPointId() bool {
	if o != nil && o.NetAccessPointId != nil {
		return true
	}

	return false
}

// SetNetAccessPointId gets a reference to the given string and assigns it to the NetAccessPointId field.
func (o *Route) SetNetAccessPointId(v string) {
	o.NetAccessPointId = &v
}

// GetNetPeeringId returns the NetPeeringId field value if set, zero value otherwise.
func (o *Route) GetNetPeeringId() string {
	if o == nil || o.NetPeeringId == nil {
		var ret string
		return ret
	}
	return *o.NetPeeringId
}

// GetNetPeeringIdOk returns a tuple with the NetPeeringId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetNetPeeringIdOk() (*string, bool) {
	if o == nil || o.NetPeeringId == nil {
		return nil, false
	}
	return o.NetPeeringId, true
}

// HasNetPeeringId returns a boolean if a field has been set.
func (o *Route) HasNetPeeringId() bool {
	if o != nil && o.NetPeeringId != nil {
		return true
	}

	return false
}

// SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.
func (o *Route) SetNetPeeringId(v string) {
	o.NetPeeringId = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *Route) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetNicIdOk() (*string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *Route) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *Route) SetNicId(v string) {
	o.NicId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Route) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Route) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Route) SetState(v string) {
	o.State = &v
}

// GetVmAccountId returns the VmAccountId field value if set, zero value otherwise.
func (o *Route) GetVmAccountId() string {
	if o == nil || o.VmAccountId == nil {
		var ret string
		return ret
	}
	return *o.VmAccountId
}

// GetVmAccountIdOk returns a tuple with the VmAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetVmAccountIdOk() (*string, bool) {
	if o == nil || o.VmAccountId == nil {
		return nil, false
	}
	return o.VmAccountId, true
}

// HasVmAccountId returns a boolean if a field has been set.
func (o *Route) HasVmAccountId() bool {
	if o != nil && o.VmAccountId != nil {
		return true
	}

	return false
}

// SetVmAccountId gets a reference to the given string and assigns it to the VmAccountId field.
func (o *Route) SetVmAccountId(v string) {
	o.VmAccountId = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *Route) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *Route) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *Route) SetVmId(v string) {
	o.VmId = &v
}

func (o Route) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationMethod != nil {
		toSerialize["CreationMethod"] = o.CreationMethod
	}
	if o.DestinationIpRange != nil {
		toSerialize["DestinationIpRange"] = o.DestinationIpRange
	}
	if o.DestinationServiceId != nil {
		toSerialize["DestinationServiceId"] = o.DestinationServiceId
	}
	if o.GatewayId != nil {
		toSerialize["GatewayId"] = o.GatewayId
	}
	if o.NatServiceId != nil {
		toSerialize["NatServiceId"] = o.NatServiceId
	}
	if o.NetAccessPointId != nil {
		toSerialize["NetAccessPointId"] = o.NetAccessPointId
	}
	if o.NetPeeringId != nil {
		toSerialize["NetPeeringId"] = o.NetPeeringId
	}
	if o.NicId != nil {
		toSerialize["NicId"] = o.NicId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.VmAccountId != nil {
		toSerialize["VmAccountId"] = o.VmAccountId
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
