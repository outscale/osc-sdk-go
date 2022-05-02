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

// DirectLink Information about the DirectLink.
type DirectLink struct {
	// The account ID of the owner of the DirectLink.
	AccountId *string `json:"AccountId,omitempty"`
	// The physical link bandwidth (either 1 Gbps or 10 Gbps).
	Bandwidth *string `json:"Bandwidth,omitempty"`
	// The ID of the DirectLink (for example, dxcon-xxxxxxxx).
	DirectLinkId *string `json:"DirectLinkId,omitempty"`
	// The name of the DirectLink.
	DirectLinkName *string `json:"DirectLinkName,omitempty"`
	// The datacenter where the DirectLink is located.
	Location *string `json:"Location,omitempty"`
	// The Region in which the DirectLink has been created.
	RegionName *string `json:"RegionName,omitempty"`
	// The state of the DirectLink.<br /> * `requested`: The DirectLink is requested but the request has not been validated yet.<br /> * `pending`: The DirectLink request has been validated. It remains in the `pending` state until you establish the physical link.<br /> * `available`: The physical link is established and the connection is ready to use.<br /> * `deleting`: The deletion process is in progress.<br /> * `deleted`: The DirectLink is deleted.
	State *string `json:"State,omitempty"`
}

// NewDirectLink instantiates a new DirectLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectLink() *DirectLink {
	this := DirectLink{}
	return &this
}

// NewDirectLinkWithDefaults instantiates a new DirectLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectLinkWithDefaults() *DirectLink {
	this := DirectLink{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *DirectLink) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLink) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *DirectLink) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *DirectLink) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *DirectLink) GetBandwidth() string {
	if o == nil || o.Bandwidth == nil {
		var ret string
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLink) GetBandwidthOk() (*string, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *DirectLink) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given string and assigns it to the Bandwidth field.
func (o *DirectLink) SetBandwidth(v string) {
	o.Bandwidth = &v
}

// GetDirectLinkId returns the DirectLinkId field value if set, zero value otherwise.
func (o *DirectLink) GetDirectLinkId() string {
	if o == nil || o.DirectLinkId == nil {
		var ret string
		return ret
	}
	return *o.DirectLinkId
}

// GetDirectLinkIdOk returns a tuple with the DirectLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLink) GetDirectLinkIdOk() (*string, bool) {
	if o == nil || o.DirectLinkId == nil {
		return nil, false
	}
	return o.DirectLinkId, true
}

// HasDirectLinkId returns a boolean if a field has been set.
func (o *DirectLink) HasDirectLinkId() bool {
	if o != nil && o.DirectLinkId != nil {
		return true
	}

	return false
}

// SetDirectLinkId gets a reference to the given string and assigns it to the DirectLinkId field.
func (o *DirectLink) SetDirectLinkId(v string) {
	o.DirectLinkId = &v
}

// GetDirectLinkName returns the DirectLinkName field value if set, zero value otherwise.
func (o *DirectLink) GetDirectLinkName() string {
	if o == nil || o.DirectLinkName == nil {
		var ret string
		return ret
	}
	return *o.DirectLinkName
}

// GetDirectLinkNameOk returns a tuple with the DirectLinkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLink) GetDirectLinkNameOk() (*string, bool) {
	if o == nil || o.DirectLinkName == nil {
		return nil, false
	}
	return o.DirectLinkName, true
}

// HasDirectLinkName returns a boolean if a field has been set.
func (o *DirectLink) HasDirectLinkName() bool {
	if o != nil && o.DirectLinkName != nil {
		return true
	}

	return false
}

// SetDirectLinkName gets a reference to the given string and assigns it to the DirectLinkName field.
func (o *DirectLink) SetDirectLinkName(v string) {
	o.DirectLinkName = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *DirectLink) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLink) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *DirectLink) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *DirectLink) SetLocation(v string) {
	o.Location = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *DirectLink) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLink) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *DirectLink) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *DirectLink) SetRegionName(v string) {
	o.RegionName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DirectLink) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectLink) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DirectLink) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DirectLink) SetState(v string) {
	o.State = &v
}

func (o DirectLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.Bandwidth != nil {
		toSerialize["Bandwidth"] = o.Bandwidth
	}
	if o.DirectLinkId != nil {
		toSerialize["DirectLinkId"] = o.DirectLinkId
	}
	if o.DirectLinkName != nil {
		toSerialize["DirectLinkName"] = o.DirectLinkName
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.RegionName != nil {
		toSerialize["RegionName"] = o.RegionName
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableDirectLink struct {
	value *DirectLink
	isSet bool
}

func (v NullableDirectLink) Get() *DirectLink {
	return v.value
}

func (v *NullableDirectLink) Set(val *DirectLink) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectLink) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectLink(val *DirectLink) *NullableDirectLink {
	return &NullableDirectLink{value: val, isSet: true}
}

func (v NullableDirectLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
