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

// OsuExportSnapshotExportTask Information about the snapshot export task.
type OsuExportSnapshotExportTask struct {
	// The format of the export disk (`qcow2` \\| `raw`).
	DiskImageFormat string `json:"DiskImageFormat"`
	// The name of the OOS bucket the snapshot is exported to.
	OsuBucket string `json:"OsuBucket"`
	// The prefix for the key of the OOS object corresponding to the snapshot.
	OsuPrefix *string `json:"OsuPrefix,omitempty"`
}

// NewOsuExportSnapshotExportTask instantiates a new OsuExportSnapshotExportTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsuExportSnapshotExportTask(diskImageFormat string, osuBucket string) *OsuExportSnapshotExportTask {
	this := OsuExportSnapshotExportTask{}
	this.DiskImageFormat = diskImageFormat
	this.OsuBucket = osuBucket
	return &this
}

// NewOsuExportSnapshotExportTaskWithDefaults instantiates a new OsuExportSnapshotExportTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsuExportSnapshotExportTaskWithDefaults() *OsuExportSnapshotExportTask {
	this := OsuExportSnapshotExportTask{}
	return &this
}

// GetDiskImageFormat returns the DiskImageFormat field value
func (o *OsuExportSnapshotExportTask) GetDiskImageFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskImageFormat
}

// GetDiskImageFormatOk returns a tuple with the DiskImageFormat field value
// and a boolean to check if the value has been set.
func (o *OsuExportSnapshotExportTask) GetDiskImageFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskImageFormat, true
}

// SetDiskImageFormat sets field value
func (o *OsuExportSnapshotExportTask) SetDiskImageFormat(v string) {
	o.DiskImageFormat = v
}

// GetOsuBucket returns the OsuBucket field value
func (o *OsuExportSnapshotExportTask) GetOsuBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsuBucket
}

// GetOsuBucketOk returns a tuple with the OsuBucket field value
// and a boolean to check if the value has been set.
func (o *OsuExportSnapshotExportTask) GetOsuBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsuBucket, true
}

// SetOsuBucket sets field value
func (o *OsuExportSnapshotExportTask) SetOsuBucket(v string) {
	o.OsuBucket = v
}

// GetOsuPrefix returns the OsuPrefix field value if set, zero value otherwise.
func (o *OsuExportSnapshotExportTask) GetOsuPrefix() string {
	if o == nil || o.OsuPrefix == nil {
		var ret string
		return ret
	}
	return *o.OsuPrefix
}

// GetOsuPrefixOk returns a tuple with the OsuPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsuExportSnapshotExportTask) GetOsuPrefixOk() (*string, bool) {
	if o == nil || o.OsuPrefix == nil {
		return nil, false
	}
	return o.OsuPrefix, true
}

// HasOsuPrefix returns a boolean if a field has been set.
func (o *OsuExportSnapshotExportTask) HasOsuPrefix() bool {
	if o != nil && o.OsuPrefix != nil {
		return true
	}

	return false
}

// SetOsuPrefix gets a reference to the given string and assigns it to the OsuPrefix field.
func (o *OsuExportSnapshotExportTask) SetOsuPrefix(v string) {
	o.OsuPrefix = &v
}

func (o OsuExportSnapshotExportTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DiskImageFormat"] = o.DiskImageFormat
	}
	if true {
		toSerialize["OsuBucket"] = o.OsuBucket
	}
	if o.OsuPrefix != nil {
		toSerialize["OsuPrefix"] = o.OsuPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableOsuExportSnapshotExportTask struct {
	value *OsuExportSnapshotExportTask
	isSet bool
}

func (v NullableOsuExportSnapshotExportTask) Get() *OsuExportSnapshotExportTask {
	return v.value
}

func (v *NullableOsuExportSnapshotExportTask) Set(val *OsuExportSnapshotExportTask) {
	v.value = val
	v.isSet = true
}

func (v NullableOsuExportSnapshotExportTask) IsSet() bool {
	return v.isSet
}

func (v *NullableOsuExportSnapshotExportTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsuExportSnapshotExportTask(val *OsuExportSnapshotExportTask) *NullableOsuExportSnapshotExportTask {
	return &NullableOsuExportSnapshotExportTask{value: val, isSet: true}
}

func (v NullableOsuExportSnapshotExportTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsuExportSnapshotExportTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
