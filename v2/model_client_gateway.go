/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ClientGateway Information about the client gateway.
type ClientGateway struct {
	// The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.
	BgpAsn *int32 `json:"BgpAsn,omitempty"`
	// The ID of the client gateway.
	ClientGatewayId *string `json:"ClientGatewayId,omitempty"`
	// The type of communication tunnel used by the client gateway (only `ipsec.1` is supported).
	ConnectionType *string `json:"ConnectionType,omitempty"`
	// The public IPv4 address of the client gateway (must be a fixed address into a NATed network).
	PublicIp *string `json:"PublicIp,omitempty"`
	// The state of the client gateway (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the client gateway.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewClientGateway instantiates a new ClientGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientGateway() *ClientGateway {
	this := ClientGateway{}
	return &this
}

// NewClientGatewayWithDefaults instantiates a new ClientGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientGatewayWithDefaults() *ClientGateway {
	this := ClientGateway{}
	return &this
}

// GetBgpAsn returns the BgpAsn field value if set, zero value otherwise.
func (o *ClientGateway) GetBgpAsn() int32 {
	if o == nil || o.BgpAsn == nil {
		var ret int32
		return ret
	}
	return *o.BgpAsn
}

// GetBgpAsnOk returns a tuple with the BgpAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGateway) GetBgpAsnOk() (*int32, bool) {
	if o == nil || o.BgpAsn == nil {
		return nil, false
	}
	return o.BgpAsn, true
}

// HasBgpAsn returns a boolean if a field has been set.
func (o *ClientGateway) HasBgpAsn() bool {
	if o != nil && o.BgpAsn != nil {
		return true
	}

	return false
}

// SetBgpAsn gets a reference to the given int32 and assigns it to the BgpAsn field.
func (o *ClientGateway) SetBgpAsn(v int32) {
	o.BgpAsn = &v
}

// GetClientGatewayId returns the ClientGatewayId field value if set, zero value otherwise.
func (o *ClientGateway) GetClientGatewayId() string {
	if o == nil || o.ClientGatewayId == nil {
		var ret string
		return ret
	}
	return *o.ClientGatewayId
}

// GetClientGatewayIdOk returns a tuple with the ClientGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGateway) GetClientGatewayIdOk() (*string, bool) {
	if o == nil || o.ClientGatewayId == nil {
		return nil, false
	}
	return o.ClientGatewayId, true
}

// HasClientGatewayId returns a boolean if a field has been set.
func (o *ClientGateway) HasClientGatewayId() bool {
	if o != nil && o.ClientGatewayId != nil {
		return true
	}

	return false
}

// SetClientGatewayId gets a reference to the given string and assigns it to the ClientGatewayId field.
func (o *ClientGateway) SetClientGatewayId(v string) {
	o.ClientGatewayId = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *ClientGateway) GetConnectionType() string {
	if o == nil || o.ConnectionType == nil {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGateway) GetConnectionTypeOk() (*string, bool) {
	if o == nil || o.ConnectionType == nil {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *ClientGateway) HasConnectionType() bool {
	if o != nil && o.ConnectionType != nil {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *ClientGateway) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *ClientGateway) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGateway) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *ClientGateway) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *ClientGateway) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ClientGateway) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGateway) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ClientGateway) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ClientGateway) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ClientGateway) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGateway) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ClientGateway) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *ClientGateway) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o ClientGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BgpAsn != nil {
		toSerialize["BgpAsn"] = o.BgpAsn
	}
	if o.ClientGatewayId != nil {
		toSerialize["ClientGatewayId"] = o.ClientGatewayId
	}
	if o.ConnectionType != nil {
		toSerialize["ConnectionType"] = o.ConnectionType
	}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableClientGateway struct {
	value *ClientGateway
	isSet bool
}

func (v NullableClientGateway) Get() *ClientGateway {
	return v.value
}

func (v *NullableClientGateway) Set(val *ClientGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableClientGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableClientGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientGateway(val *ClientGateway) *NullableClientGateway {
	return &NullableClientGateway{value: val, isSet: true}
}

func (v NullableClientGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
