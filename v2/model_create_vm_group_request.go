/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateVmGroupRequest struct for CreateVmGroupRequest
type CreateVmGroupRequest struct {
	// A description for the VM group.
	Description *string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The positioning strategy of VMs on hypervisors. By default, or if set to `no-strategy` our orchestrator determines the most adequate position for your VMs. If set to `attract`, your VMs are deployed on the same hypervisor, which improves network performance. If set to `repulse`, your VMs are deployed on a different hypervisor, which improves fault tolerance.
	PositioningStrategy *string `json:"PositioningStrategy,omitempty"`
	// One or more IDs of security groups for the VM group.
	SecurityGroupIds []string `json:"SecurityGroupIds"`
	// The ID of the Subnet in which you want to create the VM group.
	SubnetId string `json:"SubnetId"`
	// One or more tags to add to the VM group.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The number of VMs deployed in the VM group.
	VmCount int32 `json:"VmCount"`
	// The name of the VM group.
	VmGroupName string `json:"VmGroupName"`
	// The ID of the VM template used to launch VMs in the VM group.
	VmTemplateId string `json:"VmTemplateId"`
}

// NewCreateVmGroupRequest instantiates a new CreateVmGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVmGroupRequest(securityGroupIds []string, subnetId string, vmCount int32, vmGroupName string, vmTemplateId string) *CreateVmGroupRequest {
	this := CreateVmGroupRequest{}
	var positioningStrategy string = "no-strategy"
	this.PositioningStrategy = &positioningStrategy
	this.SecurityGroupIds = securityGroupIds
	this.SubnetId = subnetId
	this.VmCount = vmCount
	this.VmGroupName = vmGroupName
	this.VmTemplateId = vmTemplateId
	return &this
}

// NewCreateVmGroupRequestWithDefaults instantiates a new CreateVmGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVmGroupRequestWithDefaults() *CreateVmGroupRequest {
	this := CreateVmGroupRequest{}
	var positioningStrategy string = "no-strategy"
	this.PositioningStrategy = &positioningStrategy
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateVmGroupRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateVmGroupRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateVmGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateVmGroupRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateVmGroupRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateVmGroupRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPositioningStrategy returns the PositioningStrategy field value if set, zero value otherwise.
func (o *CreateVmGroupRequest) GetPositioningStrategy() string {
	if o == nil || o.PositioningStrategy == nil {
		var ret string
		return ret
	}
	return *o.PositioningStrategy
}

// GetPositioningStrategyOk returns a tuple with the PositioningStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetPositioningStrategyOk() (*string, bool) {
	if o == nil || o.PositioningStrategy == nil {
		return nil, false
	}
	return o.PositioningStrategy, true
}

// HasPositioningStrategy returns a boolean if a field has been set.
func (o *CreateVmGroupRequest) HasPositioningStrategy() bool {
	if o != nil && o.PositioningStrategy != nil {
		return true
	}

	return false
}

// SetPositioningStrategy gets a reference to the given string and assigns it to the PositioningStrategy field.
func (o *CreateVmGroupRequest) SetPositioningStrategy(v string) {
	o.PositioningStrategy = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value
func (o *CreateVmGroupRequest) GetSecurityGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityGroupIds, true
}

// SetSecurityGroupIds sets field value
func (o *CreateVmGroupRequest) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = v
}

// GetSubnetId returns the SubnetId field value
func (o *CreateVmGroupRequest) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *CreateVmGroupRequest) SetSubnetId(v string) {
	o.SubnetId = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateVmGroupRequest) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateVmGroupRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *CreateVmGroupRequest) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVmCount returns the VmCount field value
func (o *CreateVmGroupRequest) GetVmCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetVmCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmCount, true
}

// SetVmCount sets field value
func (o *CreateVmGroupRequest) SetVmCount(v int32) {
	o.VmCount = v
}

// GetVmGroupName returns the VmGroupName field value
func (o *CreateVmGroupRequest) GetVmGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmGroupName
}

// GetVmGroupNameOk returns a tuple with the VmGroupName field value
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetVmGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmGroupName, true
}

// SetVmGroupName sets field value
func (o *CreateVmGroupRequest) SetVmGroupName(v string) {
	o.VmGroupName = v
}

// GetVmTemplateId returns the VmTemplateId field value
func (o *CreateVmGroupRequest) GetVmTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmTemplateId
}

// GetVmTemplateIdOk returns a tuple with the VmTemplateId field value
// and a boolean to check if the value has been set.
func (o *CreateVmGroupRequest) GetVmTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmTemplateId, true
}

// SetVmTemplateId sets field value
func (o *CreateVmGroupRequest) SetVmTemplateId(v string) {
	o.VmTemplateId = v
}

func (o CreateVmGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.PositioningStrategy != nil {
		toSerialize["PositioningStrategy"] = o.PositioningStrategy
	}
	if true {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	if true {
		toSerialize["SubnetId"] = o.SubnetId
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if true {
		toSerialize["VmCount"] = o.VmCount
	}
	if true {
		toSerialize["VmGroupName"] = o.VmGroupName
	}
	if true {
		toSerialize["VmTemplateId"] = o.VmTemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVmGroupRequest struct {
	value *CreateVmGroupRequest
	isSet bool
}

func (v NullableCreateVmGroupRequest) Get() *CreateVmGroupRequest {
	return v.value
}

func (v *NullableCreateVmGroupRequest) Set(val *CreateVmGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVmGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVmGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVmGroupRequest(val *CreateVmGroupRequest) *NullableCreateVmGroupRequest {
	return &NullableCreateVmGroupRequest{value: val, isSet: true}
}

func (v NullableCreateVmGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVmGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
