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

// FlexibleGpuCatalog Information about the flexible GPU (fGPU) that is available in the public catalog.
type FlexibleGpuCatalog struct {
	// The generations of VMs that the fGPU is compatible with.
	Generations *[]string `json:"Generations,omitempty"`
	// The maximum number of VM vCores that the fGPU is compatible with.
	MaxCpu *int32 `json:"MaxCpu,omitempty"`
	// The maximum amount of VM memory that the fGPU is compatible with.
	MaxRam *int32 `json:"MaxRam,omitempty"`
	// The model of fGPU.
	ModelName *string `json:"ModelName,omitempty"`
	// The amount of video RAM (VRAM) of the fGPU.
	VRam *int32 `json:"VRam,omitempty"`
}

// NewFlexibleGpuCatalog instantiates a new FlexibleGpuCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexibleGpuCatalog() *FlexibleGpuCatalog {
	this := FlexibleGpuCatalog{}
	return &this
}

// NewFlexibleGpuCatalogWithDefaults instantiates a new FlexibleGpuCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexibleGpuCatalogWithDefaults() *FlexibleGpuCatalog {
	this := FlexibleGpuCatalog{}
	return &this
}

// GetGenerations returns the Generations field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetGenerations() []string {
	if o == nil || o.Generations == nil {
		var ret []string
		return ret
	}
	return *o.Generations
}

// GetGenerationsOk returns a tuple with the Generations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetGenerationsOk() (*[]string, bool) {
	if o == nil || o.Generations == nil {
		return nil, false
	}
	return o.Generations, true
}

// HasGenerations returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasGenerations() bool {
	if o != nil && o.Generations != nil {
		return true
	}

	return false
}

// SetGenerations gets a reference to the given []string and assigns it to the Generations field.
func (o *FlexibleGpuCatalog) SetGenerations(v []string) {
	o.Generations = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetMaxCpu() int32 {
	if o == nil || o.MaxCpu == nil {
		var ret int32
		return ret
	}
	return *o.MaxCpu
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetMaxCpuOk() (*int32, bool) {
	if o == nil || o.MaxCpu == nil {
		return nil, false
	}
	return o.MaxCpu, true
}

// HasMaxCpu returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasMaxCpu() bool {
	if o != nil && o.MaxCpu != nil {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given int32 and assigns it to the MaxCpu field.
func (o *FlexibleGpuCatalog) SetMaxCpu(v int32) {
	o.MaxCpu = &v
}

// GetMaxRam returns the MaxRam field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetMaxRam() int32 {
	if o == nil || o.MaxRam == nil {
		var ret int32
		return ret
	}
	return *o.MaxRam
}

// GetMaxRamOk returns a tuple with the MaxRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetMaxRamOk() (*int32, bool) {
	if o == nil || o.MaxRam == nil {
		return nil, false
	}
	return o.MaxRam, true
}

// HasMaxRam returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasMaxRam() bool {
	if o != nil && o.MaxRam != nil {
		return true
	}

	return false
}

// SetMaxRam gets a reference to the given int32 and assigns it to the MaxRam field.
func (o *FlexibleGpuCatalog) SetMaxRam(v int32) {
	o.MaxRam = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasModelName() bool {
	if o != nil && o.ModelName != nil {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *FlexibleGpuCatalog) SetModelName(v string) {
	o.ModelName = &v
}

// GetVRam returns the VRam field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetVRam() int32 {
	if o == nil || o.VRam == nil {
		var ret int32
		return ret
	}
	return *o.VRam
}

// GetVRamOk returns a tuple with the VRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetVRamOk() (*int32, bool) {
	if o == nil || o.VRam == nil {
		return nil, false
	}
	return o.VRam, true
}

// HasVRam returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasVRam() bool {
	if o != nil && o.VRam != nil {
		return true
	}

	return false
}

// SetVRam gets a reference to the given int32 and assigns it to the VRam field.
func (o *FlexibleGpuCatalog) SetVRam(v int32) {
	o.VRam = &v
}

func (o FlexibleGpuCatalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Generations != nil {
		toSerialize["Generations"] = o.Generations
	}
	if o.MaxCpu != nil {
		toSerialize["MaxCpu"] = o.MaxCpu
	}
	if o.MaxRam != nil {
		toSerialize["MaxRam"] = o.MaxRam
	}
	if o.ModelName != nil {
		toSerialize["ModelName"] = o.ModelName
	}
	if o.VRam != nil {
		toSerialize["VRam"] = o.VRam
	}
	return json.Marshal(toSerialize)
}

type NullableFlexibleGpuCatalog struct {
	value *FlexibleGpuCatalog
	isSet bool
}

func (v NullableFlexibleGpuCatalog) Get() *FlexibleGpuCatalog {
	return v.value
}

func (v *NullableFlexibleGpuCatalog) Set(val *FlexibleGpuCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleGpuCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleGpuCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleGpuCatalog(val *FlexibleGpuCatalog) *NullableFlexibleGpuCatalog {
	return &NullableFlexibleGpuCatalog{value: val, isSet: true}
}

func (v NullableFlexibleGpuCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleGpuCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
