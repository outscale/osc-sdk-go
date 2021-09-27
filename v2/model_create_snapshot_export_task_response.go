/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateSnapshotExportTaskResponse struct for CreateSnapshotExportTaskResponse
type CreateSnapshotExportTaskResponse struct {
	ResponseContext    *ResponseContext    `json:"ResponseContext,omitempty"`
	SnapshotExportTask *SnapshotExportTask `json:"SnapshotExportTask,omitempty"`
}

// NewCreateSnapshotExportTaskResponse instantiates a new CreateSnapshotExportTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSnapshotExportTaskResponse() *CreateSnapshotExportTaskResponse {
	this := CreateSnapshotExportTaskResponse{}
	return &this
}

// NewCreateSnapshotExportTaskResponseWithDefaults instantiates a new CreateSnapshotExportTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSnapshotExportTaskResponseWithDefaults() *CreateSnapshotExportTaskResponse {
	this := CreateSnapshotExportTaskResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateSnapshotExportTaskResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotExportTaskResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateSnapshotExportTaskResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateSnapshotExportTaskResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSnapshotExportTask returns the SnapshotExportTask field value if set, zero value otherwise.
func (o *CreateSnapshotExportTaskResponse) GetSnapshotExportTask() SnapshotExportTask {
	if o == nil || o.SnapshotExportTask == nil {
		var ret SnapshotExportTask
		return ret
	}
	return *o.SnapshotExportTask
}

// GetSnapshotExportTaskOk returns a tuple with the SnapshotExportTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotExportTaskResponse) GetSnapshotExportTaskOk() (*SnapshotExportTask, bool) {
	if o == nil || o.SnapshotExportTask == nil {
		return nil, false
	}
	return o.SnapshotExportTask, true
}

// HasSnapshotExportTask returns a boolean if a field has been set.
func (o *CreateSnapshotExportTaskResponse) HasSnapshotExportTask() bool {
	if o != nil && o.SnapshotExportTask != nil {
		return true
	}

	return false
}

// SetSnapshotExportTask gets a reference to the given SnapshotExportTask and assigns it to the SnapshotExportTask field.
func (o *CreateSnapshotExportTaskResponse) SetSnapshotExportTask(v SnapshotExportTask) {
	o.SnapshotExportTask = &v
}

func (o CreateSnapshotExportTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.SnapshotExportTask != nil {
		toSerialize["SnapshotExportTask"] = o.SnapshotExportTask
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSnapshotExportTaskResponse struct {
	value *CreateSnapshotExportTaskResponse
	isSet bool
}

func (v NullableCreateSnapshotExportTaskResponse) Get() *CreateSnapshotExportTaskResponse {
	return v.value
}

func (v *NullableCreateSnapshotExportTaskResponse) Set(val *CreateSnapshotExportTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSnapshotExportTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSnapshotExportTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSnapshotExportTaskResponse(val *CreateSnapshotExportTaskResponse) *NullableCreateSnapshotExportTaskResponse {
	return &NullableCreateSnapshotExportTaskResponse{value: val, isSet: true}
}

func (v NullableCreateSnapshotExportTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSnapshotExportTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
