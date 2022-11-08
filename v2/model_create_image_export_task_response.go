/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateImageExportTaskResponse struct for CreateImageExportTaskResponse
type CreateImageExportTaskResponse struct {
	ImageExportTask *ImageExportTask `json:"ImageExportTask,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateImageExportTaskResponse instantiates a new CreateImageExportTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageExportTaskResponse() *CreateImageExportTaskResponse {
	this := CreateImageExportTaskResponse{}
	return &this
}

// NewCreateImageExportTaskResponseWithDefaults instantiates a new CreateImageExportTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageExportTaskResponseWithDefaults() *CreateImageExportTaskResponse {
	this := CreateImageExportTaskResponse{}
	return &this
}

// GetImageExportTask returns the ImageExportTask field value if set, zero value otherwise.
func (o *CreateImageExportTaskResponse) GetImageExportTask() ImageExportTask {
	if o == nil || o.ImageExportTask == nil {
		var ret ImageExportTask
		return ret
	}
	return *o.ImageExportTask
}

// GetImageExportTaskOk returns a tuple with the ImageExportTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageExportTaskResponse) GetImageExportTaskOk() (*ImageExportTask, bool) {
	if o == nil || o.ImageExportTask == nil {
		return nil, false
	}
	return o.ImageExportTask, true
}

// HasImageExportTask returns a boolean if a field has been set.
func (o *CreateImageExportTaskResponse) HasImageExportTask() bool {
	if o != nil && o.ImageExportTask != nil {
		return true
	}

	return false
}

// SetImageExportTask gets a reference to the given ImageExportTask and assigns it to the ImageExportTask field.
func (o *CreateImageExportTaskResponse) SetImageExportTask(v ImageExportTask) {
	o.ImageExportTask = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateImageExportTaskResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageExportTaskResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateImageExportTaskResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateImageExportTaskResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateImageExportTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageExportTask != nil {
		toSerialize["ImageExportTask"] = o.ImageExportTask
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateImageExportTaskResponse struct {
	value *CreateImageExportTaskResponse
	isSet bool
}

func (v NullableCreateImageExportTaskResponse) Get() *CreateImageExportTaskResponse {
	return v.value
}

func (v *NullableCreateImageExportTaskResponse) Set(val *CreateImageExportTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageExportTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageExportTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageExportTaskResponse(val *CreateImageExportTaskResponse) *NullableCreateImageExportTaskResponse {
	return &NullableCreateImageExportTaskResponse{value: val, isSet: true}
}

func (v NullableCreateImageExportTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageExportTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
