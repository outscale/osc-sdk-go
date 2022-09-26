/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ImageExportTask Information about the OMI export task.
type ImageExportTask struct {
	// If the OMI export task fails, an error message appears.
	Comment *string `json:"Comment,omitempty"`
	// The ID of the OMI to be exported.
	ImageId   *string                   `json:"ImageId,omitempty"`
	OsuExport *OsuExportImageExportTask `json:"OsuExport,omitempty"`
	// The progress of the OMI export task, as a percentage.
	Progress *int32 `json:"Progress,omitempty"`
	// The state of the OMI export task (`pending/queued` \\| `pending` \\| `completed` \\| `failed` \\| `cancelled`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the image export task.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the OMI export task.
	TaskId *string `json:"TaskId,omitempty"`
}

// NewImageExportTask instantiates a new ImageExportTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageExportTask() *ImageExportTask {
	this := ImageExportTask{}
	return &this
}

// NewImageExportTaskWithDefaults instantiates a new ImageExportTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageExportTaskWithDefaults() *ImageExportTask {
	this := ImageExportTask{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ImageExportTask) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ImageExportTask) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ImageExportTask) SetComment(v string) {
	o.Comment = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ImageExportTask) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ImageExportTask) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ImageExportTask) SetImageId(v string) {
	o.ImageId = &v
}

// GetOsuExport returns the OsuExport field value if set, zero value otherwise.
func (o *ImageExportTask) GetOsuExport() OsuExportImageExportTask {
	if o == nil || o.OsuExport == nil {
		var ret OsuExportImageExportTask
		return ret
	}
	return *o.OsuExport
}

// GetOsuExportOk returns a tuple with the OsuExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetOsuExportOk() (*OsuExportImageExportTask, bool) {
	if o == nil || o.OsuExport == nil {
		return nil, false
	}
	return o.OsuExport, true
}

// HasOsuExport returns a boolean if a field has been set.
func (o *ImageExportTask) HasOsuExport() bool {
	if o != nil && o.OsuExport != nil {
		return true
	}

	return false
}

// SetOsuExport gets a reference to the given OsuExportImageExportTask and assigns it to the OsuExport field.
func (o *ImageExportTask) SetOsuExport(v OsuExportImageExportTask) {
	o.OsuExport = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ImageExportTask) GetProgress() int32 {
	if o == nil || o.Progress == nil {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetProgressOk() (*int32, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ImageExportTask) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *ImageExportTask) SetProgress(v int32) {
	o.Progress = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ImageExportTask) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ImageExportTask) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ImageExportTask) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ImageExportTask) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ImageExportTask) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *ImageExportTask) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *ImageExportTask) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *ImageExportTask) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *ImageExportTask) SetTaskId(v string) {
	o.TaskId = &v
}

func (o ImageExportTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["Comment"] = o.Comment
	}
	if o.ImageId != nil {
		toSerialize["ImageId"] = o.ImageId
	}
	if o.OsuExport != nil {
		toSerialize["OsuExport"] = o.OsuExport
	}
	if o.Progress != nil {
		toSerialize["Progress"] = o.Progress
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

type NullableImageExportTask struct {
	value *ImageExportTask
	isSet bool
}

func (v NullableImageExportTask) Get() *ImageExportTask {
	return v.value
}

func (v *NullableImageExportTask) Set(val *ImageExportTask) {
	v.value = val
	v.isSet = true
}

func (v NullableImageExportTask) IsSet() bool {
	return v.isSet
}

func (v *NullableImageExportTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageExportTask(val *ImageExportTask) *NullableImageExportTask {
	return &NullableImageExportTask{value: val, isSet: true}
}

func (v NullableImageExportTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageExportTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
