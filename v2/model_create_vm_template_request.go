/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateVmTemplateRequest struct for CreateVmTemplateRequest
type CreateVmTemplateRequest struct {
	// The number of vCores to use for each VM.
	CpuCores int32 `json:"CpuCores"`
	// The processor generation to use for each VM (for example, `v4`).
	CpuGeneration string `json:"CpuGeneration"`
	// The performance of the VMs (`medium` \\| `high` \\|  `highest`).
	CpuPerformance *string `json:"CpuPerformance,omitempty"`
	// A description for the VM template.
	Description *string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the OMI to use for each VM. You can find a list of OMIs by calling the [ReadImages](#readimages) method.
	ImageId string `json:"ImageId"`
	// The name of the keypair to use for each VM.
	KeypairName *string `json:"KeypairName,omitempty"`
	// The amount of RAM to use for each VM.
	Ram int32 `json:"Ram"`
	// One or more tags to add to the VM template.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The name of the VM template.
	VmTemplateName string `json:"VmTemplateName"`
}

// NewCreateVmTemplateRequest instantiates a new CreateVmTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVmTemplateRequest(cpuCores int32, cpuGeneration string, imageId string, ram int32, vmTemplateName string) *CreateVmTemplateRequest {
	this := CreateVmTemplateRequest{}
	this.CpuCores = cpuCores
	this.CpuGeneration = cpuGeneration
	var cpuPerformance string = "high"
	this.CpuPerformance = &cpuPerformance
	this.ImageId = imageId
	this.Ram = ram
	this.VmTemplateName = vmTemplateName
	return &this
}

// NewCreateVmTemplateRequestWithDefaults instantiates a new CreateVmTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVmTemplateRequestWithDefaults() *CreateVmTemplateRequest {
	this := CreateVmTemplateRequest{}
	var cpuPerformance string = "high"
	this.CpuPerformance = &cpuPerformance
	return &this
}

// GetCpuCores returns the CpuCores field value
func (o *CreateVmTemplateRequest) GetCpuCores() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetCpuCoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuCores, true
}

// SetCpuCores sets field value
func (o *CreateVmTemplateRequest) SetCpuCores(v int32) {
	o.CpuCores = v
}

// GetCpuGeneration returns the CpuGeneration field value
func (o *CreateVmTemplateRequest) GetCpuGeneration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CpuGeneration
}

// GetCpuGenerationOk returns a tuple with the CpuGeneration field value
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetCpuGenerationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuGeneration, true
}

// SetCpuGeneration sets field value
func (o *CreateVmTemplateRequest) SetCpuGeneration(v string) {
	o.CpuGeneration = v
}

// GetCpuPerformance returns the CpuPerformance field value if set, zero value otherwise.
func (o *CreateVmTemplateRequest) GetCpuPerformance() string {
	if o == nil || o.CpuPerformance == nil {
		var ret string
		return ret
	}
	return *o.CpuPerformance
}

// GetCpuPerformanceOk returns a tuple with the CpuPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetCpuPerformanceOk() (*string, bool) {
	if o == nil || o.CpuPerformance == nil {
		return nil, false
	}
	return o.CpuPerformance, true
}

// HasCpuPerformance returns a boolean if a field has been set.
func (o *CreateVmTemplateRequest) HasCpuPerformance() bool {
	if o != nil && o.CpuPerformance != nil {
		return true
	}

	return false
}

// SetCpuPerformance gets a reference to the given string and assigns it to the CpuPerformance field.
func (o *CreateVmTemplateRequest) SetCpuPerformance(v string) {
	o.CpuPerformance = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateVmTemplateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateVmTemplateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateVmTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateVmTemplateRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateVmTemplateRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateVmTemplateRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetImageId returns the ImageId field value
func (o *CreateVmTemplateRequest) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *CreateVmTemplateRequest) SetImageId(v string) {
	o.ImageId = v
}

// GetKeypairName returns the KeypairName field value if set, zero value otherwise.
func (o *CreateVmTemplateRequest) GetKeypairName() string {
	if o == nil || o.KeypairName == nil {
		var ret string
		return ret
	}
	return *o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetKeypairNameOk() (*string, bool) {
	if o == nil || o.KeypairName == nil {
		return nil, false
	}
	return o.KeypairName, true
}

// HasKeypairName returns a boolean if a field has been set.
func (o *CreateVmTemplateRequest) HasKeypairName() bool {
	if o != nil && o.KeypairName != nil {
		return true
	}

	return false
}

// SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.
func (o *CreateVmTemplateRequest) SetKeypairName(v string) {
	o.KeypairName = &v
}

// GetRam returns the Ram field value
func (o *CreateVmTemplateRequest) GetRam() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetRamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *CreateVmTemplateRequest) SetRam(v int32) {
	o.Ram = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateVmTemplateRequest) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateVmTemplateRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *CreateVmTemplateRequest) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVmTemplateName returns the VmTemplateName field value
func (o *CreateVmTemplateRequest) GetVmTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmTemplateName
}

// GetVmTemplateNameOk returns a tuple with the VmTemplateName field value
// and a boolean to check if the value has been set.
func (o *CreateVmTemplateRequest) GetVmTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmTemplateName, true
}

// SetVmTemplateName sets field value
func (o *CreateVmTemplateRequest) SetVmTemplateName(v string) {
	o.VmTemplateName = v
}

func (o CreateVmTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CpuCores"] = o.CpuCores
	}
	if true {
		toSerialize["CpuGeneration"] = o.CpuGeneration
	}
	if o.CpuPerformance != nil {
		toSerialize["CpuPerformance"] = o.CpuPerformance
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["ImageId"] = o.ImageId
	}
	if o.KeypairName != nil {
		toSerialize["KeypairName"] = o.KeypairName
	}
	if true {
		toSerialize["Ram"] = o.Ram
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if true {
		toSerialize["VmTemplateName"] = o.VmTemplateName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVmTemplateRequest struct {
	value *CreateVmTemplateRequest
	isSet bool
}

func (v NullableCreateVmTemplateRequest) Get() *CreateVmTemplateRequest {
	return v.value
}

func (v *NullableCreateVmTemplateRequest) Set(val *CreateVmTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVmTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVmTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVmTemplateRequest(val *CreateVmTemplateRequest) *NullableCreateVmTemplateRequest {
	return &NullableCreateVmTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateVmTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVmTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
