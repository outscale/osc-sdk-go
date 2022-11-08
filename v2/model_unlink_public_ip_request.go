/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UnlinkPublicIpRequest struct for UnlinkPublicIpRequest
type UnlinkPublicIpRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID representing the association of the public IP with the VM or the NIC. This parameter is required unless you use the `PublicIp` parameter.
	LinkPublicIpId *string `json:"LinkPublicIpId,omitempty"`
	// The public IP. This parameter is required unless you use the `LinkPublicIpId` parameter.
	PublicIp *string `json:"PublicIp,omitempty"`
}

// NewUnlinkPublicIpRequest instantiates a new UnlinkPublicIpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkPublicIpRequest() *UnlinkPublicIpRequest {
	this := UnlinkPublicIpRequest{}
	return &this
}

// NewUnlinkPublicIpRequestWithDefaults instantiates a new UnlinkPublicIpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkPublicIpRequestWithDefaults() *UnlinkPublicIpRequest {
	this := UnlinkPublicIpRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkPublicIpRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkPublicIpRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkPublicIpRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkPublicIpRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLinkPublicIpId returns the LinkPublicIpId field value if set, zero value otherwise.
func (o *UnlinkPublicIpRequest) GetLinkPublicIpId() string {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret
	}
	return *o.LinkPublicIpId
}

// GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkPublicIpRequest) GetLinkPublicIpIdOk() (*string, bool) {
	if o == nil || o.LinkPublicIpId == nil {
		return nil, false
	}
	return o.LinkPublicIpId, true
}

// HasLinkPublicIpId returns a boolean if a field has been set.
func (o *UnlinkPublicIpRequest) HasLinkPublicIpId() bool {
	if o != nil && o.LinkPublicIpId != nil {
		return true
	}

	return false
}

// SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.
func (o *UnlinkPublicIpRequest) SetLinkPublicIpId(v string) {
	o.LinkPublicIpId = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *UnlinkPublicIpRequest) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkPublicIpRequest) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *UnlinkPublicIpRequest) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *UnlinkPublicIpRequest) SetPublicIp(v string) {
	o.PublicIp = &v
}

func (o UnlinkPublicIpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.LinkPublicIpId != nil {
		toSerialize["LinkPublicIpId"] = o.LinkPublicIpId
	}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	return json.Marshal(toSerialize)
}

type NullableUnlinkPublicIpRequest struct {
	value *UnlinkPublicIpRequest
	isSet bool
}

func (v NullableUnlinkPublicIpRequest) Get() *UnlinkPublicIpRequest {
	return v.value
}

func (v *NullableUnlinkPublicIpRequest) Set(val *UnlinkPublicIpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkPublicIpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkPublicIpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkPublicIpRequest(val *UnlinkPublicIpRequest) *NullableUnlinkPublicIpRequest {
	return &NullableUnlinkPublicIpRequest{value: val, isSet: true}
}

func (v NullableUnlinkPublicIpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlinkPublicIpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
