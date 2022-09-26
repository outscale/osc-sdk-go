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

// ReadSnapshotExportTasksResponse struct for ReadSnapshotExportTasksResponse
type ReadSnapshotExportTasksResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more snapshot export tasks.
	SnapshotExportTasks *[]SnapshotExportTask `json:"SnapshotExportTasks,omitempty"`
}

// NewReadSnapshotExportTasksResponse instantiates a new ReadSnapshotExportTasksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadSnapshotExportTasksResponse() *ReadSnapshotExportTasksResponse {
	this := ReadSnapshotExportTasksResponse{}
	return &this
}

// NewReadSnapshotExportTasksResponseWithDefaults instantiates a new ReadSnapshotExportTasksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadSnapshotExportTasksResponseWithDefaults() *ReadSnapshotExportTasksResponse {
	this := ReadSnapshotExportTasksResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadSnapshotExportTasksResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSnapshotExportTasksResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadSnapshotExportTasksResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadSnapshotExportTasksResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSnapshotExportTasks returns the SnapshotExportTasks field value if set, zero value otherwise.
func (o *ReadSnapshotExportTasksResponse) GetSnapshotExportTasks() []SnapshotExportTask {
	if o == nil || o.SnapshotExportTasks == nil {
		var ret []SnapshotExportTask
		return ret
	}
	return *o.SnapshotExportTasks
}

// GetSnapshotExportTasksOk returns a tuple with the SnapshotExportTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSnapshotExportTasksResponse) GetSnapshotExportTasksOk() (*[]SnapshotExportTask, bool) {
	if o == nil || o.SnapshotExportTasks == nil {
		return nil, false
	}
	return o.SnapshotExportTasks, true
}

// HasSnapshotExportTasks returns a boolean if a field has been set.
func (o *ReadSnapshotExportTasksResponse) HasSnapshotExportTasks() bool {
	if o != nil && o.SnapshotExportTasks != nil {
		return true
	}

	return false
}

// SetSnapshotExportTasks gets a reference to the given []SnapshotExportTask and assigns it to the SnapshotExportTasks field.
func (o *ReadSnapshotExportTasksResponse) SetSnapshotExportTasks(v []SnapshotExportTask) {
	o.SnapshotExportTasks = &v
}

func (o ReadSnapshotExportTasksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.SnapshotExportTasks != nil {
		toSerialize["SnapshotExportTasks"] = o.SnapshotExportTasks
	}
	return json.Marshal(toSerialize)
}

type NullableReadSnapshotExportTasksResponse struct {
	value *ReadSnapshotExportTasksResponse
	isSet bool
}

func (v NullableReadSnapshotExportTasksResponse) Get() *ReadSnapshotExportTasksResponse {
	return v.value
}

func (v *NullableReadSnapshotExportTasksResponse) Set(val *ReadSnapshotExportTasksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadSnapshotExportTasksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadSnapshotExportTasksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadSnapshotExportTasksResponse(val *ReadSnapshotExportTasksResponse) *NullableReadSnapshotExportTasksResponse {
	return &NullableReadSnapshotExportTasksResponse{value: val, isSet: true}
}

func (v NullableReadSnapshotExportTasksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadSnapshotExportTasksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
