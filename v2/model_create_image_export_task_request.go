/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.19
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateImageExportTaskRequest struct for CreateImageExportTaskRequest
type CreateImageExportTaskRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the OMI to export.
	ImageId   string            `json:"ImageId"`
	OsuExport OsuExportToCreate `json:"OsuExport"`
}

// NewCreateImageExportTaskRequest instantiates a new CreateImageExportTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageExportTaskRequest(imageId string, osuExport OsuExportToCreate) *CreateImageExportTaskRequest {
	this := CreateImageExportTaskRequest{}
	this.ImageId = imageId
	this.OsuExport = osuExport
	return &this
}

// NewCreateImageExportTaskRequestWithDefaults instantiates a new CreateImageExportTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageExportTaskRequestWithDefaults() *CreateImageExportTaskRequest {
	this := CreateImageExportTaskRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateImageExportTaskRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageExportTaskRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateImageExportTaskRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateImageExportTaskRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetImageId returns the ImageId field value
func (o *CreateImageExportTaskRequest) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *CreateImageExportTaskRequest) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *CreateImageExportTaskRequest) SetImageId(v string) {
	o.ImageId = v
}

// GetOsuExport returns the OsuExport field value
func (o *CreateImageExportTaskRequest) GetOsuExport() OsuExportToCreate {
	if o == nil {
		var ret OsuExportToCreate
		return ret
	}

	return o.OsuExport
}

// GetOsuExportOk returns a tuple with the OsuExport field value
// and a boolean to check if the value has been set.
func (o *CreateImageExportTaskRequest) GetOsuExportOk() (*OsuExportToCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsuExport, true
}

// SetOsuExport sets field value
func (o *CreateImageExportTaskRequest) SetOsuExport(v OsuExportToCreate) {
	o.OsuExport = v
}

func (o CreateImageExportTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["ImageId"] = o.ImageId
	}
	if true {
		toSerialize["OsuExport"] = o.OsuExport
	}
	return json.Marshal(toSerialize)
}

type NullableCreateImageExportTaskRequest struct {
	value *CreateImageExportTaskRequest
	isSet bool
}

func (v NullableCreateImageExportTaskRequest) Get() *CreateImageExportTaskRequest {
	return v.value
}

func (v *NullableCreateImageExportTaskRequest) Set(val *CreateImageExportTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageExportTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageExportTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageExportTaskRequest(val *CreateImageExportTaskRequest) *NullableCreateImageExportTaskRequest {
	return &NullableCreateImageExportTaskRequest{value: val, isSet: true}
}

func (v NullableCreateImageExportTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageExportTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
