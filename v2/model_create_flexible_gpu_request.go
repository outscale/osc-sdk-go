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

// CreateFlexibleGpuRequest struct for CreateFlexibleGpuRequest
type CreateFlexibleGpuRequest struct {
	// If true, the fGPU is deleted when the VM is terminated.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The processor generation that the fGPU must be compatible with. If not specified, the oldest possible processor generation is selected (as provided by [ReadFlexibleGpuCatalog](#readflexiblegpucatalog) for the specified model of fGPU).
	Generation *string `json:"Generation,omitempty"`
	// The model of fGPU you want to allocate. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).
	ModelName string `json:"ModelName"`
	// The Subregion in which you want to create the fGPU.
	SubregionName string `json:"SubregionName"`
}

// NewCreateFlexibleGpuRequest instantiates a new CreateFlexibleGpuRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFlexibleGpuRequest(modelName string, subregionName string) *CreateFlexibleGpuRequest {
	this := CreateFlexibleGpuRequest{}
	var deleteOnVmDeletion bool = false
	this.DeleteOnVmDeletion = &deleteOnVmDeletion
	this.ModelName = modelName
	this.SubregionName = subregionName
	return &this
}

// NewCreateFlexibleGpuRequestWithDefaults instantiates a new CreateFlexibleGpuRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFlexibleGpuRequestWithDefaults() *CreateFlexibleGpuRequest {
	this := CreateFlexibleGpuRequest{}
	var deleteOnVmDeletion bool = false
	this.DeleteOnVmDeletion = &deleteOnVmDeletion
	return &this
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *CreateFlexibleGpuRequest) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuRequest) GetDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *CreateFlexibleGpuRequest) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *CreateFlexibleGpuRequest) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateFlexibleGpuRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateFlexibleGpuRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateFlexibleGpuRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetGeneration returns the Generation field value if set, zero value otherwise.
func (o *CreateFlexibleGpuRequest) GetGeneration() string {
	if o == nil || o.Generation == nil {
		var ret string
		return ret
	}
	return *o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuRequest) GetGenerationOk() (*string, bool) {
	if o == nil || o.Generation == nil {
		return nil, false
	}
	return o.Generation, true
}

// HasGeneration returns a boolean if a field has been set.
func (o *CreateFlexibleGpuRequest) HasGeneration() bool {
	if o != nil && o.Generation != nil {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given string and assigns it to the Generation field.
func (o *CreateFlexibleGpuRequest) SetGeneration(v string) {
	o.Generation = &v
}

// GetModelName returns the ModelName field value
func (o *CreateFlexibleGpuRequest) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuRequest) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *CreateFlexibleGpuRequest) SetModelName(v string) {
	o.ModelName = v
}

// GetSubregionName returns the SubregionName field value
func (o *CreateFlexibleGpuRequest) GetSubregionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuRequest) GetSubregionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubregionName, true
}

// SetSubregionName sets field value
func (o *CreateFlexibleGpuRequest) SetSubregionName(v string) {
	o.SubregionName = v
}

func (o CreateFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteOnVmDeletion != nil {
		toSerialize["DeleteOnVmDeletion"] = o.DeleteOnVmDeletion
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Generation != nil {
		toSerialize["Generation"] = o.Generation
	}
	if true {
		toSerialize["ModelName"] = o.ModelName
	}
	if true {
		toSerialize["SubregionName"] = o.SubregionName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFlexibleGpuRequest struct {
	value *CreateFlexibleGpuRequest
	isSet bool
}

func (v NullableCreateFlexibleGpuRequest) Get() *CreateFlexibleGpuRequest {
	return v.value
}

func (v *NullableCreateFlexibleGpuRequest) Set(val *CreateFlexibleGpuRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFlexibleGpuRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFlexibleGpuRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFlexibleGpuRequest(val *CreateFlexibleGpuRequest) *NullableCreateFlexibleGpuRequest {
	return &NullableCreateFlexibleGpuRequest{value: val, isSet: true}
}

func (v NullableCreateFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFlexibleGpuRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
