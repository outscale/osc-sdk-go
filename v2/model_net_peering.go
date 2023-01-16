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

// NetPeering Information about the Net peering.
type NetPeering struct {
	AccepterNet *AccepterNet `json:"AccepterNet,omitempty"`
	// The ID of the Net peering.
	NetPeeringId *string          `json:"NetPeeringId,omitempty"`
	SourceNet    *SourceNet       `json:"SourceNet,omitempty"`
	State        *NetPeeringState `json:"State,omitempty"`
	// One or more tags associated with the Net peering.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewNetPeering instantiates a new NetPeering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetPeering() *NetPeering {
	this := NetPeering{}
	return &this
}

// NewNetPeeringWithDefaults instantiates a new NetPeering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetPeeringWithDefaults() *NetPeering {
	this := NetPeering{}
	return &this
}

// GetAccepterNet returns the AccepterNet field value if set, zero value otherwise.
func (o *NetPeering) GetAccepterNet() AccepterNet {
	if o == nil || o.AccepterNet == nil {
		var ret AccepterNet
		return ret
	}
	return *o.AccepterNet
}

// GetAccepterNetOk returns a tuple with the AccepterNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetAccepterNetOk() (*AccepterNet, bool) {
	if o == nil || o.AccepterNet == nil {
		return nil, false
	}
	return o.AccepterNet, true
}

// HasAccepterNet returns a boolean if a field has been set.
func (o *NetPeering) HasAccepterNet() bool {
	if o != nil && o.AccepterNet != nil {
		return true
	}

	return false
}

// SetAccepterNet gets a reference to the given AccepterNet and assigns it to the AccepterNet field.
func (o *NetPeering) SetAccepterNet(v AccepterNet) {
	o.AccepterNet = &v
}

// GetNetPeeringId returns the NetPeeringId field value if set, zero value otherwise.
func (o *NetPeering) GetNetPeeringId() string {
	if o == nil || o.NetPeeringId == nil {
		var ret string
		return ret
	}
	return *o.NetPeeringId
}

// GetNetPeeringIdOk returns a tuple with the NetPeeringId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetNetPeeringIdOk() (*string, bool) {
	if o == nil || o.NetPeeringId == nil {
		return nil, false
	}
	return o.NetPeeringId, true
}

// HasNetPeeringId returns a boolean if a field has been set.
func (o *NetPeering) HasNetPeeringId() bool {
	if o != nil && o.NetPeeringId != nil {
		return true
	}

	return false
}

// SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.
func (o *NetPeering) SetNetPeeringId(v string) {
	o.NetPeeringId = &v
}

// GetSourceNet returns the SourceNet field value if set, zero value otherwise.
func (o *NetPeering) GetSourceNet() SourceNet {
	if o == nil || o.SourceNet == nil {
		var ret SourceNet
		return ret
	}
	return *o.SourceNet
}

// GetSourceNetOk returns a tuple with the SourceNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetSourceNetOk() (*SourceNet, bool) {
	if o == nil || o.SourceNet == nil {
		return nil, false
	}
	return o.SourceNet, true
}

// HasSourceNet returns a boolean if a field has been set.
func (o *NetPeering) HasSourceNet() bool {
	if o != nil && o.SourceNet != nil {
		return true
	}

	return false
}

// SetSourceNet gets a reference to the given SourceNet and assigns it to the SourceNet field.
func (o *NetPeering) SetSourceNet(v SourceNet) {
	o.SourceNet = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NetPeering) GetState() NetPeeringState {
	if o == nil || o.State == nil {
		var ret NetPeeringState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetStateOk() (*NetPeeringState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NetPeering) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given NetPeeringState and assigns it to the State field.
func (o *NetPeering) SetState(v NetPeeringState) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NetPeering) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPeering) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NetPeering) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *NetPeering) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o NetPeering) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccepterNet != nil {
		toSerialize["AccepterNet"] = o.AccepterNet
	}
	if o.NetPeeringId != nil {
		toSerialize["NetPeeringId"] = o.NetPeeringId
	}
	if o.SourceNet != nil {
		toSerialize["SourceNet"] = o.SourceNet
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableNetPeering struct {
	value *NetPeering
	isSet bool
}

func (v NullableNetPeering) Get() *NetPeering {
	return v.value
}

func (v *NullableNetPeering) Set(val *NetPeering) {
	v.value = val
	v.isSet = true
}

func (v NullableNetPeering) IsSet() bool {
	return v.isSet
}

func (v *NullableNetPeering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetPeering(val *NetPeering) *NullableNetPeering {
	return &NullableNetPeering{value: val, isSet: true}
}

func (v NullableNetPeering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetPeering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
