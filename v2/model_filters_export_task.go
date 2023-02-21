/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.25
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersExportTask One or more filters.
type FiltersExportTask struct {
	// The IDs of the export tasks.
	TaskIds *[]string `json:"TaskIds,omitempty"`
}

// NewFiltersExportTask instantiates a new FiltersExportTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersExportTask() *FiltersExportTask {
	this := FiltersExportTask{}
	return &this
}

// NewFiltersExportTaskWithDefaults instantiates a new FiltersExportTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersExportTaskWithDefaults() *FiltersExportTask {
	this := FiltersExportTask{}
	return &this
}

// GetTaskIds returns the TaskIds field value if set, zero value otherwise.
func (o *FiltersExportTask) GetTaskIds() []string {
	if o == nil || o.TaskIds == nil {
		var ret []string
		return ret
	}
	return *o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersExportTask) GetTaskIdsOk() (*[]string, bool) {
	if o == nil || o.TaskIds == nil {
		return nil, false
	}
	return o.TaskIds, true
}

// HasTaskIds returns a boolean if a field has been set.
func (o *FiltersExportTask) HasTaskIds() bool {
	if o != nil && o.TaskIds != nil {
		return true
	}

	return false
}

// SetTaskIds gets a reference to the given []string and assigns it to the TaskIds field.
func (o *FiltersExportTask) SetTaskIds(v []string) {
	o.TaskIds = &v
}

func (o FiltersExportTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaskIds != nil {
		toSerialize["TaskIds"] = o.TaskIds
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersExportTask struct {
	value *FiltersExportTask
	isSet bool
}

func (v NullableFiltersExportTask) Get() *FiltersExportTask {
	return v.value
}

func (v *NullableFiltersExportTask) Set(val *FiltersExportTask) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersExportTask) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersExportTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersExportTask(val *FiltersExportTask) *NullableFiltersExportTask {
	return &NullableFiltersExportTask{value: val, isSet: true}
}

func (v NullableFiltersExportTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersExportTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
