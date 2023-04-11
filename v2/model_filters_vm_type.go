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

// FiltersVmType One or more filters.
type FiltersVmType struct {
	// This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.
	BsuOptimized *bool `json:"BsuOptimized,omitempty"`
	// The amounts of memory, in gibibytes (GiB).
	MemorySizes *[]float32 `json:"MemorySizes,omitempty"`
	// The numbers of vCores.
	VcoreCounts *[]int32 `json:"VcoreCounts,omitempty"`
	// The names of the VM types. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).
	VmTypeNames *[]string `json:"VmTypeNames,omitempty"`
	// The maximum number of ephemeral storage disks.
	VolumeCounts *[]int32 `json:"VolumeCounts,omitempty"`
	// The size of one ephemeral storage disk, in gibibytes (GiB).
	VolumeSizes *[]int32 `json:"VolumeSizes,omitempty"`
}

// NewFiltersVmType instantiates a new FiltersVmType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersVmType() *FiltersVmType {
	this := FiltersVmType{}
	return &this
}

// NewFiltersVmTypeWithDefaults instantiates a new FiltersVmType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersVmTypeWithDefaults() *FiltersVmType {
	this := FiltersVmType{}
	return &this
}

// GetBsuOptimized returns the BsuOptimized field value if set, zero value otherwise.
func (o *FiltersVmType) GetBsuOptimized() bool {
	if o == nil || o.BsuOptimized == nil {
		var ret bool
		return ret
	}
	return *o.BsuOptimized
}

// GetBsuOptimizedOk returns a tuple with the BsuOptimized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmType) GetBsuOptimizedOk() (*bool, bool) {
	if o == nil || o.BsuOptimized == nil {
		return nil, false
	}
	return o.BsuOptimized, true
}

// HasBsuOptimized returns a boolean if a field has been set.
func (o *FiltersVmType) HasBsuOptimized() bool {
	if o != nil && o.BsuOptimized != nil {
		return true
	}

	return false
}

// SetBsuOptimized gets a reference to the given bool and assigns it to the BsuOptimized field.
func (o *FiltersVmType) SetBsuOptimized(v bool) {
	o.BsuOptimized = &v
}

// GetMemorySizes returns the MemorySizes field value if set, zero value otherwise.
func (o *FiltersVmType) GetMemorySizes() []float32 {
	if o == nil || o.MemorySizes == nil {
		var ret []float32
		return ret
	}
	return *o.MemorySizes
}

// GetMemorySizesOk returns a tuple with the MemorySizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmType) GetMemorySizesOk() (*[]float32, bool) {
	if o == nil || o.MemorySizes == nil {
		return nil, false
	}
	return o.MemorySizes, true
}

// HasMemorySizes returns a boolean if a field has been set.
func (o *FiltersVmType) HasMemorySizes() bool {
	if o != nil && o.MemorySizes != nil {
		return true
	}

	return false
}

// SetMemorySizes gets a reference to the given []float32 and assigns it to the MemorySizes field.
func (o *FiltersVmType) SetMemorySizes(v []float32) {
	o.MemorySizes = &v
}

// GetVcoreCounts returns the VcoreCounts field value if set, zero value otherwise.
func (o *FiltersVmType) GetVcoreCounts() []int32 {
	if o == nil || o.VcoreCounts == nil {
		var ret []int32
		return ret
	}
	return *o.VcoreCounts
}

// GetVcoreCountsOk returns a tuple with the VcoreCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmType) GetVcoreCountsOk() (*[]int32, bool) {
	if o == nil || o.VcoreCounts == nil {
		return nil, false
	}
	return o.VcoreCounts, true
}

// HasVcoreCounts returns a boolean if a field has been set.
func (o *FiltersVmType) HasVcoreCounts() bool {
	if o != nil && o.VcoreCounts != nil {
		return true
	}

	return false
}

// SetVcoreCounts gets a reference to the given []int32 and assigns it to the VcoreCounts field.
func (o *FiltersVmType) SetVcoreCounts(v []int32) {
	o.VcoreCounts = &v
}

// GetVmTypeNames returns the VmTypeNames field value if set, zero value otherwise.
func (o *FiltersVmType) GetVmTypeNames() []string {
	if o == nil || o.VmTypeNames == nil {
		var ret []string
		return ret
	}
	return *o.VmTypeNames
}

// GetVmTypeNamesOk returns a tuple with the VmTypeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmType) GetVmTypeNamesOk() (*[]string, bool) {
	if o == nil || o.VmTypeNames == nil {
		return nil, false
	}
	return o.VmTypeNames, true
}

// HasVmTypeNames returns a boolean if a field has been set.
func (o *FiltersVmType) HasVmTypeNames() bool {
	if o != nil && o.VmTypeNames != nil {
		return true
	}

	return false
}

// SetVmTypeNames gets a reference to the given []string and assigns it to the VmTypeNames field.
func (o *FiltersVmType) SetVmTypeNames(v []string) {
	o.VmTypeNames = &v
}

// GetVolumeCounts returns the VolumeCounts field value if set, zero value otherwise.
func (o *FiltersVmType) GetVolumeCounts() []int32 {
	if o == nil || o.VolumeCounts == nil {
		var ret []int32
		return ret
	}
	return *o.VolumeCounts
}

// GetVolumeCountsOk returns a tuple with the VolumeCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmType) GetVolumeCountsOk() (*[]int32, bool) {
	if o == nil || o.VolumeCounts == nil {
		return nil, false
	}
	return o.VolumeCounts, true
}

// HasVolumeCounts returns a boolean if a field has been set.
func (o *FiltersVmType) HasVolumeCounts() bool {
	if o != nil && o.VolumeCounts != nil {
		return true
	}

	return false
}

// SetVolumeCounts gets a reference to the given []int32 and assigns it to the VolumeCounts field.
func (o *FiltersVmType) SetVolumeCounts(v []int32) {
	o.VolumeCounts = &v
}

// GetVolumeSizes returns the VolumeSizes field value if set, zero value otherwise.
func (o *FiltersVmType) GetVolumeSizes() []int32 {
	if o == nil || o.VolumeSizes == nil {
		var ret []int32
		return ret
	}
	return *o.VolumeSizes
}

// GetVolumeSizesOk returns a tuple with the VolumeSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmType) GetVolumeSizesOk() (*[]int32, bool) {
	if o == nil || o.VolumeSizes == nil {
		return nil, false
	}
	return o.VolumeSizes, true
}

// HasVolumeSizes returns a boolean if a field has been set.
func (o *FiltersVmType) HasVolumeSizes() bool {
	if o != nil && o.VolumeSizes != nil {
		return true
	}

	return false
}

// SetVolumeSizes gets a reference to the given []int32 and assigns it to the VolumeSizes field.
func (o *FiltersVmType) SetVolumeSizes(v []int32) {
	o.VolumeSizes = &v
}

func (o FiltersVmType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BsuOptimized != nil {
		toSerialize["BsuOptimized"] = o.BsuOptimized
	}
	if o.MemorySizes != nil {
		toSerialize["MemorySizes"] = o.MemorySizes
	}
	if o.VcoreCounts != nil {
		toSerialize["VcoreCounts"] = o.VcoreCounts
	}
	if o.VmTypeNames != nil {
		toSerialize["VmTypeNames"] = o.VmTypeNames
	}
	if o.VolumeCounts != nil {
		toSerialize["VolumeCounts"] = o.VolumeCounts
	}
	if o.VolumeSizes != nil {
		toSerialize["VolumeSizes"] = o.VolumeSizes
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersVmType struct {
	value *FiltersVmType
	isSet bool
}

func (v NullableFiltersVmType) Get() *FiltersVmType {
	return v.value
}

func (v *NullableFiltersVmType) Set(val *FiltersVmType) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersVmType) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersVmType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersVmType(val *FiltersVmType) *NullableFiltersVmType {
	return &NullableFiltersVmType{value: val, isSet: true}
}

func (v NullableFiltersVmType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersVmType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
