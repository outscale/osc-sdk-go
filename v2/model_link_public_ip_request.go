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

// LinkPublicIpRequest struct for LinkPublicIpRequest
type LinkPublicIpRequest struct {
	// If true, allows the public IP to be associated with the VM or NIC that you specify even if it is already associated with another VM or NIC. If false, prevents the EIP from being associated with the VM or NIC that you specify if it is already associated with another VM or NIC. (By default, true in the public Cloud, false in a Net.)
	AllowRelink *bool `json:"AllowRelink,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// (Net only) The ID of the NIC. This parameter is required if the VM has more than one NIC attached. Otherwise, you need to specify the `VmId` parameter instead. You cannot specify both parameters at the same time.
	NicId *string `json:"NicId,omitempty"`
	// (Net only) The primary or secondary private IP of the specified NIC. By default, the primary private IP.
	PrivateIp *string `json:"PrivateIp,omitempty"`
	// The public IP. This parameter is required unless you use the `PublicIpId` parameter.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The allocation ID of the public IP. This parameter is required unless you use the `PublicIp` parameter.
	PublicIpId *string `json:"PublicIpId,omitempty"`
	// The ID of the VM.<br /> - In the public Cloud, this parameter is required.<br /> - In a Net, this parameter is required if the VM has only one NIC. Otherwise, you need to specify the `NicId` parameter instead. You cannot specify both parameters at the same time.
	VmId *string `json:"VmId,omitempty"`
}

// NewLinkPublicIpRequest instantiates a new LinkPublicIpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkPublicIpRequest() *LinkPublicIpRequest {
	this := LinkPublicIpRequest{}
	return &this
}

// NewLinkPublicIpRequestWithDefaults instantiates a new LinkPublicIpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkPublicIpRequestWithDefaults() *LinkPublicIpRequest {
	this := LinkPublicIpRequest{}
	return &this
}

// GetAllowRelink returns the AllowRelink field value if set, zero value otherwise.
func (o *LinkPublicIpRequest) GetAllowRelink() bool {
	if o == nil || o.AllowRelink == nil {
		var ret bool
		return ret
	}
	return *o.AllowRelink
}

// GetAllowRelinkOk returns a tuple with the AllowRelink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpRequest) GetAllowRelinkOk() (*bool, bool) {
	if o == nil || o.AllowRelink == nil {
		return nil, false
	}
	return o.AllowRelink, true
}

// HasAllowRelink returns a boolean if a field has been set.
func (o *LinkPublicIpRequest) HasAllowRelink() bool {
	if o != nil && o.AllowRelink != nil {
		return true
	}

	return false
}

// SetAllowRelink gets a reference to the given bool and assigns it to the AllowRelink field.
func (o *LinkPublicIpRequest) SetAllowRelink(v bool) {
	o.AllowRelink = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkPublicIpRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkPublicIpRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkPublicIpRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *LinkPublicIpRequest) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpRequest) GetNicIdOk() (*string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *LinkPublicIpRequest) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *LinkPublicIpRequest) SetNicId(v string) {
	o.NicId = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *LinkPublicIpRequest) GetPrivateIp() string {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpRequest) GetPrivateIpOk() (*string, bool) {
	if o == nil || o.PrivateIp == nil {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *LinkPublicIpRequest) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *LinkPublicIpRequest) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *LinkPublicIpRequest) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpRequest) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *LinkPublicIpRequest) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *LinkPublicIpRequest) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicIpId returns the PublicIpId field value if set, zero value otherwise.
func (o *LinkPublicIpRequest) GetPublicIpId() string {
	if o == nil || o.PublicIpId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpId
}

// GetPublicIpIdOk returns a tuple with the PublicIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpRequest) GetPublicIpIdOk() (*string, bool) {
	if o == nil || o.PublicIpId == nil {
		return nil, false
	}
	return o.PublicIpId, true
}

// HasPublicIpId returns a boolean if a field has been set.
func (o *LinkPublicIpRequest) HasPublicIpId() bool {
	if o != nil && o.PublicIpId != nil {
		return true
	}

	return false
}

// SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.
func (o *LinkPublicIpRequest) SetPublicIpId(v string) {
	o.PublicIpId = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *LinkPublicIpRequest) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpRequest) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *LinkPublicIpRequest) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *LinkPublicIpRequest) SetVmId(v string) {
	o.VmId = &v
}

func (o LinkPublicIpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowRelink != nil {
		toSerialize["AllowRelink"] = o.AllowRelink
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.NicId != nil {
		toSerialize["NicId"] = o.NicId
	}
	if o.PrivateIp != nil {
		toSerialize["PrivateIp"] = o.PrivateIp
	}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.PublicIpId != nil {
		toSerialize["PublicIpId"] = o.PublicIpId
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkPublicIpRequest struct {
	value *LinkPublicIpRequest
	isSet bool
}

func (v NullableLinkPublicIpRequest) Get() *LinkPublicIpRequest {
	return v.value
}

func (v *NullableLinkPublicIpRequest) Set(val *LinkPublicIpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkPublicIpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkPublicIpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkPublicIpRequest(val *LinkPublicIpRequest) *NullableLinkPublicIpRequest {
	return &NullableLinkPublicIpRequest{value: val, isSet: true}
}

func (v NullableLinkPublicIpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkPublicIpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
