/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateImageRequest struct for CreateImageRequest
type CreateImageRequest struct {
	// The architecture of the OMI (by default, `i386`).
	Architecture *string `json:"Architecture,omitempty"`
	// One or more block device mappings.
	BlockDeviceMappings *[]BlockDeviceMappingImage `json:"BlockDeviceMappings,omitempty"`
	// A description for the new OMI.
	Description *string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The pre-signed URL of the OMI manifest file, or the full path to the OMI stored in a bucket. If you specify this parameter, a copy of the OMI is created in your account.
	FileLocation *string `json:"FileLocation,omitempty"`
	// A unique name for the new OMI.<br /> Constraints: 3-128 alphanumeric characters, underscores (_), spaces ( ), parentheses (()), slashes (/), periods (.), or dashes (-).
	ImageName *string `json:"ImageName,omitempty"`
	// If false, the VM shuts down before creating the OMI and then reboots. If true, the VM does not.
	NoReboot *bool `json:"NoReboot,omitempty"`
	// The name of the root device.
	RootDeviceName *string `json:"RootDeviceName,omitempty"`
	// The ID of the OMI you want to copy.
	SourceImageId *string `json:"SourceImageId,omitempty"`
	// The name of the source Region, which must be the same as the Region of your account.
	SourceRegionName *string `json:"SourceRegionName,omitempty"`
	// The ID of the VM from which you want to create the OMI.
	VmId *string `json:"VmId,omitempty"`
}

// NewCreateImageRequest instantiates a new CreateImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageRequest() *CreateImageRequest {
	this := CreateImageRequest{}
	return &this
}

// NewCreateImageRequestWithDefaults instantiates a new CreateImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageRequestWithDefaults() *CreateImageRequest {
	this := CreateImageRequest{}
	return &this
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *CreateImageRequest) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *CreateImageRequest) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *CreateImageRequest) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBlockDeviceMappings returns the BlockDeviceMappings field value if set, zero value otherwise.
func (o *CreateImageRequest) GetBlockDeviceMappings() []BlockDeviceMappingImage {
	if o == nil || o.BlockDeviceMappings == nil {
		var ret []BlockDeviceMappingImage
		return ret
	}
	return *o.BlockDeviceMappings
}

// GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetBlockDeviceMappingsOk() (*[]BlockDeviceMappingImage, bool) {
	if o == nil || o.BlockDeviceMappings == nil {
		return nil, false
	}
	return o.BlockDeviceMappings, true
}

// HasBlockDeviceMappings returns a boolean if a field has been set.
func (o *CreateImageRequest) HasBlockDeviceMappings() bool {
	if o != nil && o.BlockDeviceMappings != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingImage and assigns it to the BlockDeviceMappings field.
func (o *CreateImageRequest) SetBlockDeviceMappings(v []BlockDeviceMappingImage) {
	o.BlockDeviceMappings = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateImageRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateImageRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateImageRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateImageRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateImageRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateImageRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *CreateImageRequest) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *CreateImageRequest) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *CreateImageRequest) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *CreateImageRequest) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *CreateImageRequest) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *CreateImageRequest) SetImageName(v string) {
	o.ImageName = &v
}

// GetNoReboot returns the NoReboot field value if set, zero value otherwise.
func (o *CreateImageRequest) GetNoReboot() bool {
	if o == nil || o.NoReboot == nil {
		var ret bool
		return ret
	}
	return *o.NoReboot
}

// GetNoRebootOk returns a tuple with the NoReboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetNoRebootOk() (*bool, bool) {
	if o == nil || o.NoReboot == nil {
		return nil, false
	}
	return o.NoReboot, true
}

// HasNoReboot returns a boolean if a field has been set.
func (o *CreateImageRequest) HasNoReboot() bool {
	if o != nil && o.NoReboot != nil {
		return true
	}

	return false
}

// SetNoReboot gets a reference to the given bool and assigns it to the NoReboot field.
func (o *CreateImageRequest) SetNoReboot(v bool) {
	o.NoReboot = &v
}

// GetRootDeviceName returns the RootDeviceName field value if set, zero value otherwise.
func (o *CreateImageRequest) GetRootDeviceName() string {
	if o == nil || o.RootDeviceName == nil {
		var ret string
		return ret
	}
	return *o.RootDeviceName
}

// GetRootDeviceNameOk returns a tuple with the RootDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetRootDeviceNameOk() (*string, bool) {
	if o == nil || o.RootDeviceName == nil {
		return nil, false
	}
	return o.RootDeviceName, true
}

// HasRootDeviceName returns a boolean if a field has been set.
func (o *CreateImageRequest) HasRootDeviceName() bool {
	if o != nil && o.RootDeviceName != nil {
		return true
	}

	return false
}

// SetRootDeviceName gets a reference to the given string and assigns it to the RootDeviceName field.
func (o *CreateImageRequest) SetRootDeviceName(v string) {
	o.RootDeviceName = &v
}

// GetSourceImageId returns the SourceImageId field value if set, zero value otherwise.
func (o *CreateImageRequest) GetSourceImageId() string {
	if o == nil || o.SourceImageId == nil {
		var ret string
		return ret
	}
	return *o.SourceImageId
}

// GetSourceImageIdOk returns a tuple with the SourceImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetSourceImageIdOk() (*string, bool) {
	if o == nil || o.SourceImageId == nil {
		return nil, false
	}
	return o.SourceImageId, true
}

// HasSourceImageId returns a boolean if a field has been set.
func (o *CreateImageRequest) HasSourceImageId() bool {
	if o != nil && o.SourceImageId != nil {
		return true
	}

	return false
}

// SetSourceImageId gets a reference to the given string and assigns it to the SourceImageId field.
func (o *CreateImageRequest) SetSourceImageId(v string) {
	o.SourceImageId = &v
}

// GetSourceRegionName returns the SourceRegionName field value if set, zero value otherwise.
func (o *CreateImageRequest) GetSourceRegionName() string {
	if o == nil || o.SourceRegionName == nil {
		var ret string
		return ret
	}
	return *o.SourceRegionName
}

// GetSourceRegionNameOk returns a tuple with the SourceRegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetSourceRegionNameOk() (*string, bool) {
	if o == nil || o.SourceRegionName == nil {
		return nil, false
	}
	return o.SourceRegionName, true
}

// HasSourceRegionName returns a boolean if a field has been set.
func (o *CreateImageRequest) HasSourceRegionName() bool {
	if o != nil && o.SourceRegionName != nil {
		return true
	}

	return false
}

// SetSourceRegionName gets a reference to the given string and assigns it to the SourceRegionName field.
func (o *CreateImageRequest) SetSourceRegionName(v string) {
	o.SourceRegionName = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *CreateImageRequest) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *CreateImageRequest) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *CreateImageRequest) SetVmId(v string) {
	o.VmId = &v
}

func (o CreateImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Architecture != nil {
		toSerialize["Architecture"] = o.Architecture
	}
	if o.BlockDeviceMappings != nil {
		toSerialize["BlockDeviceMappings"] = o.BlockDeviceMappings
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.ImageName != nil {
		toSerialize["ImageName"] = o.ImageName
	}
	if o.NoReboot != nil {
		toSerialize["NoReboot"] = o.NoReboot
	}
	if o.RootDeviceName != nil {
		toSerialize["RootDeviceName"] = o.RootDeviceName
	}
	if o.SourceImageId != nil {
		toSerialize["SourceImageId"] = o.SourceImageId
	}
	if o.SourceRegionName != nil {
		toSerialize["SourceRegionName"] = o.SourceRegionName
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateImageRequest struct {
	value *CreateImageRequest
	isSet bool
}

func (v NullableCreateImageRequest) Get() *CreateImageRequest {
	return v.value
}

func (v *NullableCreateImageRequest) Set(val *CreateImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageRequest(val *CreateImageRequest) *NullableCreateImageRequest {
	return &NullableCreateImageRequest{value: val, isSet: true}
}

func (v NullableCreateImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
