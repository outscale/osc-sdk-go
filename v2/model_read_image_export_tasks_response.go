/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadImageExportTasksResponse struct for ReadImageExportTasksResponse
type ReadImageExportTasksResponse struct {
	// Information about one or more image export tasks.
	ImageExportTasks *[]ImageExportTask `json:"ImageExportTasks,omitempty"`
	ResponseContext  *ResponseContext   `json:"ResponseContext,omitempty"`
}

// NewReadImageExportTasksResponse instantiates a new ReadImageExportTasksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadImageExportTasksResponse() *ReadImageExportTasksResponse {
	this := ReadImageExportTasksResponse{}
	return &this
}

// NewReadImageExportTasksResponseWithDefaults instantiates a new ReadImageExportTasksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadImageExportTasksResponseWithDefaults() *ReadImageExportTasksResponse {
	this := ReadImageExportTasksResponse{}
	return &this
}

// GetImageExportTasks returns the ImageExportTasks field value if set, zero value otherwise.
func (o *ReadImageExportTasksResponse) GetImageExportTasks() []ImageExportTask {
	if o == nil || o.ImageExportTasks == nil {
		var ret []ImageExportTask
		return ret
	}
	return *o.ImageExportTasks
}

// GetImageExportTasksOk returns a tuple with the ImageExportTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadImageExportTasksResponse) GetImageExportTasksOk() (*[]ImageExportTask, bool) {
	if o == nil || o.ImageExportTasks == nil {
		return nil, false
	}
	return o.ImageExportTasks, true
}

// HasImageExportTasks returns a boolean if a field has been set.
func (o *ReadImageExportTasksResponse) HasImageExportTasks() bool {
	if o != nil && o.ImageExportTasks != nil {
		return true
	}

	return false
}

// SetImageExportTasks gets a reference to the given []ImageExportTask and assigns it to the ImageExportTasks field.
func (o *ReadImageExportTasksResponse) SetImageExportTasks(v []ImageExportTask) {
	o.ImageExportTasks = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadImageExportTasksResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadImageExportTasksResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadImageExportTasksResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadImageExportTasksResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadImageExportTasksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageExportTasks != nil {
		toSerialize["ImageExportTasks"] = o.ImageExportTasks
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadImageExportTasksResponse struct {
	value *ReadImageExportTasksResponse
	isSet bool
}

func (v NullableReadImageExportTasksResponse) Get() *ReadImageExportTasksResponse {
	return v.value
}

func (v *NullableReadImageExportTasksResponse) Set(val *ReadImageExportTasksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadImageExportTasksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadImageExportTasksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadImageExportTasksResponse(val *ReadImageExportTasksResponse) *NullableReadImageExportTasksResponse {
	return &NullableReadImageExportTasksResponse{value: val, isSet: true}
}

func (v NullableReadImageExportTasksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadImageExportTasksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
