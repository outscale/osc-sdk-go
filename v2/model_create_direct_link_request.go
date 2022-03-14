/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateDirectLinkRequest struct for CreateDirectLinkRequest
type CreateDirectLinkRequest struct {
	// The bandwidth of the DirectLink (`1Gbps` \\| `10Gbps`).
	Bandwidth string `json:"Bandwidth"`
	// The name of the DirectLink.
	DirectLinkName string `json:"DirectLinkName"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The code of the requested location for the DirectLink, returned by the [ReadLocations](#readlocations) method.
	Location string `json:"Location"`
}

// NewCreateDirectLinkRequest instantiates a new CreateDirectLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDirectLinkRequest(bandwidth string, directLinkName string, location string) *CreateDirectLinkRequest {
	this := CreateDirectLinkRequest{}
	this.Bandwidth = bandwidth
	this.DirectLinkName = directLinkName
	this.Location = location
	return &this
}

// NewCreateDirectLinkRequestWithDefaults instantiates a new CreateDirectLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDirectLinkRequestWithDefaults() *CreateDirectLinkRequest {
	this := CreateDirectLinkRequest{}
	return &this
}

// GetBandwidth returns the Bandwidth field value
func (o *CreateDirectLinkRequest) GetBandwidth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkRequest) GetBandwidthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bandwidth, true
}

// SetBandwidth sets field value
func (o *CreateDirectLinkRequest) SetBandwidth(v string) {
	o.Bandwidth = v
}

// GetDirectLinkName returns the DirectLinkName field value
func (o *CreateDirectLinkRequest) GetDirectLinkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectLinkName
}

// GetDirectLinkNameOk returns a tuple with the DirectLinkName field value
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkRequest) GetDirectLinkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectLinkName, true
}

// SetDirectLinkName sets field value
func (o *CreateDirectLinkRequest) SetDirectLinkName(v string) {
	o.DirectLinkName = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateDirectLinkRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateDirectLinkRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateDirectLinkRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLocation returns the Location field value
func (o *CreateDirectLinkRequest) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkRequest) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *CreateDirectLinkRequest) SetLocation(v string) {
	o.Location = v
}

func (o CreateDirectLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Bandwidth"] = o.Bandwidth
	}
	if true {
		toSerialize["DirectLinkName"] = o.DirectLinkName
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDirectLinkRequest struct {
	value *CreateDirectLinkRequest
	isSet bool
}

func (v NullableCreateDirectLinkRequest) Get() *CreateDirectLinkRequest {
	return v.value
}

func (v *NullableCreateDirectLinkRequest) Set(val *CreateDirectLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDirectLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDirectLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDirectLinkRequest(val *CreateDirectLinkRequest) *NullableCreateDirectLinkRequest {
	return &NullableCreateDirectLinkRequest{value: val, isSet: true}
}

func (v NullableCreateDirectLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDirectLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
