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

// SnapshotExportTask Information about the snapshot export task.
type SnapshotExportTask struct {
	// If the snapshot export task fails, an error message appears.
	Comment   *string                      `json:"Comment,omitempty"`
	OsuExport *OsuExportSnapshotExportTask `json:"OsuExport,omitempty"`
	// The progress of the snapshot export task, as a percentage.
	Progress *int32 `json:"Progress,omitempty"`
	// The ID of the snapshot to be exported.
	SnapshotId *string `json:"SnapshotId,omitempty"`
	// The state of the snapshot export task (`pending` \\| `active` \\| `completed` \\| `failed`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the snapshot export task.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the snapshot export task.
	TaskId *string `json:"TaskId,omitempty"`
}

// NewSnapshotExportTask instantiates a new SnapshotExportTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotExportTask() *SnapshotExportTask {
	this := SnapshotExportTask{}
	return &this
}

// NewSnapshotExportTaskWithDefaults instantiates a new SnapshotExportTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotExportTaskWithDefaults() *SnapshotExportTask {
	this := SnapshotExportTask{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SnapshotExportTask) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotExportTask) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SnapshotExportTask) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SnapshotExportTask) SetComment(v string) {
	o.Comment = &v
}

// GetOsuExport returns the OsuExport field value if set, zero value otherwise.
func (o *SnapshotExportTask) GetOsuExport() OsuExportSnapshotExportTask {
	if o == nil || o.OsuExport == nil {
		var ret OsuExportSnapshotExportTask
		return ret
	}
	return *o.OsuExport
}

// GetOsuExportOk returns a tuple with the OsuExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotExportTask) GetOsuExportOk() (*OsuExportSnapshotExportTask, bool) {
	if o == nil || o.OsuExport == nil {
		return nil, false
	}
	return o.OsuExport, true
}

// HasOsuExport returns a boolean if a field has been set.
func (o *SnapshotExportTask) HasOsuExport() bool {
	if o != nil && o.OsuExport != nil {
		return true
	}

	return false
}

// SetOsuExport gets a reference to the given OsuExportSnapshotExportTask and assigns it to the OsuExport field.
func (o *SnapshotExportTask) SetOsuExport(v OsuExportSnapshotExportTask) {
	o.OsuExport = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *SnapshotExportTask) GetProgress() int32 {
	if o == nil || o.Progress == nil {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotExportTask) GetProgressOk() (*int32, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *SnapshotExportTask) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *SnapshotExportTask) SetProgress(v int32) {
	o.Progress = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *SnapshotExportTask) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotExportTask) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *SnapshotExportTask) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *SnapshotExportTask) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SnapshotExportTask) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotExportTask) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SnapshotExportTask) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SnapshotExportTask) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SnapshotExportTask) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotExportTask) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SnapshotExportTask) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *SnapshotExportTask) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *SnapshotExportTask) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotExportTask) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *SnapshotExportTask) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *SnapshotExportTask) SetTaskId(v string) {
	o.TaskId = &v
}

func (o SnapshotExportTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["Comment"] = o.Comment
	}
	if o.OsuExport != nil {
		toSerialize["OsuExport"] = o.OsuExport
	}
	if o.Progress != nil {
		toSerialize["Progress"] = o.Progress
	}
	if o.SnapshotId != nil {
		toSerialize["SnapshotId"] = o.SnapshotId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.TaskId != nil {
		toSerialize["TaskId"] = o.TaskId
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotExportTask struct {
	value *SnapshotExportTask
	isSet bool
}

func (v NullableSnapshotExportTask) Get() *SnapshotExportTask {
	return v.value
}

func (v *NullableSnapshotExportTask) Set(val *SnapshotExportTask) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotExportTask) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotExportTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotExportTask(val *SnapshotExportTask) *NullableSnapshotExportTask {
	return &NullableSnapshotExportTask{value: val, isSet: true}
}

func (v NullableSnapshotExportTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotExportTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
