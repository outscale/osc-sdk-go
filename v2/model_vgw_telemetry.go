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

// VgwTelemetry Information about the current state of a VPN tunnel.
type VgwTelemetry struct {
	// The number of routes accepted through BGP (Border Gateway Protocol) route exchanges.
	AcceptedRouteCount *int32 `json:"AcceptedRouteCount,omitempty"`
	// The date and time (UTC) of the latest state update.
	LastStateChangeDate *string `json:"LastStateChangeDate,omitempty"`
	// The IP on the OUTSCALE side of the tunnel.
	OutsideIpAddress *string `json:"OutsideIpAddress,omitempty"`
	// The state of the IPSEC tunnel (`UP` \\| `DOWN`).
	State *string `json:"State,omitempty"`
	// A description of the current state of the tunnel.
	StateDescription *string `json:"StateDescription,omitempty"`
}

// NewVgwTelemetry instantiates a new VgwTelemetry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVgwTelemetry() *VgwTelemetry {
	this := VgwTelemetry{}
	return &this
}

// NewVgwTelemetryWithDefaults instantiates a new VgwTelemetry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVgwTelemetryWithDefaults() *VgwTelemetry {
	this := VgwTelemetry{}
	return &this
}

// GetAcceptedRouteCount returns the AcceptedRouteCount field value if set, zero value otherwise.
func (o *VgwTelemetry) GetAcceptedRouteCount() int32 {
	if o == nil || o.AcceptedRouteCount == nil {
		var ret int32
		return ret
	}
	return *o.AcceptedRouteCount
}

// GetAcceptedRouteCountOk returns a tuple with the AcceptedRouteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VgwTelemetry) GetAcceptedRouteCountOk() (*int32, bool) {
	if o == nil || o.AcceptedRouteCount == nil {
		return nil, false
	}
	return o.AcceptedRouteCount, true
}

// HasAcceptedRouteCount returns a boolean if a field has been set.
func (o *VgwTelemetry) HasAcceptedRouteCount() bool {
	if o != nil && o.AcceptedRouteCount != nil {
		return true
	}

	return false
}

// SetAcceptedRouteCount gets a reference to the given int32 and assigns it to the AcceptedRouteCount field.
func (o *VgwTelemetry) SetAcceptedRouteCount(v int32) {
	o.AcceptedRouteCount = &v
}

// GetLastStateChangeDate returns the LastStateChangeDate field value if set, zero value otherwise.
func (o *VgwTelemetry) GetLastStateChangeDate() string {
	if o == nil || o.LastStateChangeDate == nil {
		var ret string
		return ret
	}
	return *o.LastStateChangeDate
}

// GetLastStateChangeDateOk returns a tuple with the LastStateChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VgwTelemetry) GetLastStateChangeDateOk() (*string, bool) {
	if o == nil || o.LastStateChangeDate == nil {
		return nil, false
	}
	return o.LastStateChangeDate, true
}

// HasLastStateChangeDate returns a boolean if a field has been set.
func (o *VgwTelemetry) HasLastStateChangeDate() bool {
	if o != nil && o.LastStateChangeDate != nil {
		return true
	}

	return false
}

// SetLastStateChangeDate gets a reference to the given string and assigns it to the LastStateChangeDate field.
func (o *VgwTelemetry) SetLastStateChangeDate(v string) {
	o.LastStateChangeDate = &v
}

// GetOutsideIpAddress returns the OutsideIpAddress field value if set, zero value otherwise.
func (o *VgwTelemetry) GetOutsideIpAddress() string {
	if o == nil || o.OutsideIpAddress == nil {
		var ret string
		return ret
	}
	return *o.OutsideIpAddress
}

// GetOutsideIpAddressOk returns a tuple with the OutsideIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VgwTelemetry) GetOutsideIpAddressOk() (*string, bool) {
	if o == nil || o.OutsideIpAddress == nil {
		return nil, false
	}
	return o.OutsideIpAddress, true
}

// HasOutsideIpAddress returns a boolean if a field has been set.
func (o *VgwTelemetry) HasOutsideIpAddress() bool {
	if o != nil && o.OutsideIpAddress != nil {
		return true
	}

	return false
}

// SetOutsideIpAddress gets a reference to the given string and assigns it to the OutsideIpAddress field.
func (o *VgwTelemetry) SetOutsideIpAddress(v string) {
	o.OutsideIpAddress = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VgwTelemetry) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VgwTelemetry) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VgwTelemetry) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VgwTelemetry) SetState(v string) {
	o.State = &v
}

// GetStateDescription returns the StateDescription field value if set, zero value otherwise.
func (o *VgwTelemetry) GetStateDescription() string {
	if o == nil || o.StateDescription == nil {
		var ret string
		return ret
	}
	return *o.StateDescription
}

// GetStateDescriptionOk returns a tuple with the StateDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VgwTelemetry) GetStateDescriptionOk() (*string, bool) {
	if o == nil || o.StateDescription == nil {
		return nil, false
	}
	return o.StateDescription, true
}

// HasStateDescription returns a boolean if a field has been set.
func (o *VgwTelemetry) HasStateDescription() bool {
	if o != nil && o.StateDescription != nil {
		return true
	}

	return false
}

// SetStateDescription gets a reference to the given string and assigns it to the StateDescription field.
func (o *VgwTelemetry) SetStateDescription(v string) {
	o.StateDescription = &v
}

func (o VgwTelemetry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptedRouteCount != nil {
		toSerialize["AcceptedRouteCount"] = o.AcceptedRouteCount
	}
	if o.LastStateChangeDate != nil {
		toSerialize["LastStateChangeDate"] = o.LastStateChangeDate
	}
	if o.OutsideIpAddress != nil {
		toSerialize["OutsideIpAddress"] = o.OutsideIpAddress
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.StateDescription != nil {
		toSerialize["StateDescription"] = o.StateDescription
	}
	return json.Marshal(toSerialize)
}

type NullableVgwTelemetry struct {
	value *VgwTelemetry
	isSet bool
}

func (v NullableVgwTelemetry) Get() *VgwTelemetry {
	return v.value
}

func (v *NullableVgwTelemetry) Set(val *VgwTelemetry) {
	v.value = val
	v.isSet = true
}

func (v NullableVgwTelemetry) IsSet() bool {
	return v.isSet
}

func (v *NullableVgwTelemetry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVgwTelemetry(val *VgwTelemetry) *NullableVgwTelemetry {
	return &NullableVgwTelemetry{value: val, isSet: true}
}

func (v NullableVgwTelemetry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVgwTelemetry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
