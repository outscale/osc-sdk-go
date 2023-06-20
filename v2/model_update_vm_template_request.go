/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.27
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateVmTemplateRequest struct for UpdateVmTemplateRequest
type UpdateVmTemplateRequest struct {
	// A new description for the VM template.
	Description *string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// New tags for your VM template.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the VM template you want to update.
	VmTemplateId string `json:"VmTemplateId"`
	// A new name for your VM template.
	VmTemplateName *string `json:"VmTemplateName,omitempty"`
}

// NewUpdateVmTemplateRequest instantiates a new UpdateVmTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVmTemplateRequest(vmTemplateId string) *UpdateVmTemplateRequest {
	this := UpdateVmTemplateRequest{}
	this.VmTemplateId = vmTemplateId
	return &this
}

// NewUpdateVmTemplateRequestWithDefaults instantiates a new UpdateVmTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVmTemplateRequestWithDefaults() *UpdateVmTemplateRequest {
	this := UpdateVmTemplateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVmTemplateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVmTemplateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVmTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateVmTemplateRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmTemplateRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateVmTemplateRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateVmTemplateRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateVmTemplateRequest) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmTemplateRequest) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateVmTemplateRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *UpdateVmTemplateRequest) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVmTemplateId returns the VmTemplateId field value
func (o *UpdateVmTemplateRequest) GetVmTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmTemplateId
}

// GetVmTemplateIdOk returns a tuple with the VmTemplateId field value
// and a boolean to check if the value has been set.
func (o *UpdateVmTemplateRequest) GetVmTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmTemplateId, true
}

// SetVmTemplateId sets field value
func (o *UpdateVmTemplateRequest) SetVmTemplateId(v string) {
	o.VmTemplateId = v
}

// GetVmTemplateName returns the VmTemplateName field value if set, zero value otherwise.
func (o *UpdateVmTemplateRequest) GetVmTemplateName() string {
	if o == nil || o.VmTemplateName == nil {
		var ret string
		return ret
	}
	return *o.VmTemplateName
}

// GetVmTemplateNameOk returns a tuple with the VmTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVmTemplateRequest) GetVmTemplateNameOk() (*string, bool) {
	if o == nil || o.VmTemplateName == nil {
		return nil, false
	}
	return o.VmTemplateName, true
}

// HasVmTemplateName returns a boolean if a field has been set.
func (o *UpdateVmTemplateRequest) HasVmTemplateName() bool {
	if o != nil && o.VmTemplateName != nil {
		return true
	}

	return false
}

// SetVmTemplateName gets a reference to the given string and assigns it to the VmTemplateName field.
func (o *UpdateVmTemplateRequest) SetVmTemplateName(v string) {
	o.VmTemplateName = &v
}

func (o UpdateVmTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if true {
		toSerialize["VmTemplateId"] = o.VmTemplateId
	}
	if o.VmTemplateName != nil {
		toSerialize["VmTemplateName"] = o.VmTemplateName
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateVmTemplateRequest struct {
	value *UpdateVmTemplateRequest
	isSet bool
}

func (v NullableUpdateVmTemplateRequest) Get() *UpdateVmTemplateRequest {
	return v.value
}

func (v *NullableUpdateVmTemplateRequest) Set(val *UpdateVmTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVmTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVmTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVmTemplateRequest(val *UpdateVmTemplateRequest) *NullableUpdateVmTemplateRequest {
	return &NullableUpdateVmTemplateRequest{value: val, isSet: true}
}

func (v NullableUpdateVmTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVmTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
