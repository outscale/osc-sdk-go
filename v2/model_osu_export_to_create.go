/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.16
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// OsuExportToCreate Information about the OOS export task to create.
type OsuExportToCreate struct {
	// The format of the export disk (`qcow2` \\| `raw`).
	DiskImageFormat string     `json:"DiskImageFormat"`
	OsuApiKey       *OsuApiKey `json:"OsuApiKey,omitempty"`
	// The name of the OOS bucket where you want to export the object.
	OsuBucket string `json:"OsuBucket"`
	// The URL of the manifest file.
	OsuManifestUrl *string `json:"OsuManifestUrl,omitempty"`
	// The prefix for the key of the OOS object.
	OsuPrefix *string `json:"OsuPrefix,omitempty"`
}

// NewOsuExportToCreate instantiates a new OsuExportToCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsuExportToCreate(diskImageFormat string, osuBucket string) *OsuExportToCreate {
	this := OsuExportToCreate{}
	this.DiskImageFormat = diskImageFormat
	this.OsuBucket = osuBucket
	return &this
}

// NewOsuExportToCreateWithDefaults instantiates a new OsuExportToCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsuExportToCreateWithDefaults() *OsuExportToCreate {
	this := OsuExportToCreate{}
	return &this
}

// GetDiskImageFormat returns the DiskImageFormat field value
func (o *OsuExportToCreate) GetDiskImageFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskImageFormat
}

// GetDiskImageFormatOk returns a tuple with the DiskImageFormat field value
// and a boolean to check if the value has been set.
func (o *OsuExportToCreate) GetDiskImageFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskImageFormat, true
}

// SetDiskImageFormat sets field value
func (o *OsuExportToCreate) SetDiskImageFormat(v string) {
	o.DiskImageFormat = v
}

// GetOsuApiKey returns the OsuApiKey field value if set, zero value otherwise.
func (o *OsuExportToCreate) GetOsuApiKey() OsuApiKey {
	if o == nil || o.OsuApiKey == nil {
		var ret OsuApiKey
		return ret
	}
	return *o.OsuApiKey
}

// GetOsuApiKeyOk returns a tuple with the OsuApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsuExportToCreate) GetOsuApiKeyOk() (*OsuApiKey, bool) {
	if o == nil || o.OsuApiKey == nil {
		return nil, false
	}
	return o.OsuApiKey, true
}

// HasOsuApiKey returns a boolean if a field has been set.
func (o *OsuExportToCreate) HasOsuApiKey() bool {
	if o != nil && o.OsuApiKey != nil {
		return true
	}

	return false
}

// SetOsuApiKey gets a reference to the given OsuApiKey and assigns it to the OsuApiKey field.
func (o *OsuExportToCreate) SetOsuApiKey(v OsuApiKey) {
	o.OsuApiKey = &v
}

// GetOsuBucket returns the OsuBucket field value
func (o *OsuExportToCreate) GetOsuBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsuBucket
}

// GetOsuBucketOk returns a tuple with the OsuBucket field value
// and a boolean to check if the value has been set.
func (o *OsuExportToCreate) GetOsuBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsuBucket, true
}

// SetOsuBucket sets field value
func (o *OsuExportToCreate) SetOsuBucket(v string) {
	o.OsuBucket = v
}

// GetOsuManifestUrl returns the OsuManifestUrl field value if set, zero value otherwise.
func (o *OsuExportToCreate) GetOsuManifestUrl() string {
	if o == nil || o.OsuManifestUrl == nil {
		var ret string
		return ret
	}
	return *o.OsuManifestUrl
}

// GetOsuManifestUrlOk returns a tuple with the OsuManifestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsuExportToCreate) GetOsuManifestUrlOk() (*string, bool) {
	if o == nil || o.OsuManifestUrl == nil {
		return nil, false
	}
	return o.OsuManifestUrl, true
}

// HasOsuManifestUrl returns a boolean if a field has been set.
func (o *OsuExportToCreate) HasOsuManifestUrl() bool {
	if o != nil && o.OsuManifestUrl != nil {
		return true
	}

	return false
}

// SetOsuManifestUrl gets a reference to the given string and assigns it to the OsuManifestUrl field.
func (o *OsuExportToCreate) SetOsuManifestUrl(v string) {
	o.OsuManifestUrl = &v
}

// GetOsuPrefix returns the OsuPrefix field value if set, zero value otherwise.
func (o *OsuExportToCreate) GetOsuPrefix() string {
	if o == nil || o.OsuPrefix == nil {
		var ret string
		return ret
	}
	return *o.OsuPrefix
}

// GetOsuPrefixOk returns a tuple with the OsuPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsuExportToCreate) GetOsuPrefixOk() (*string, bool) {
	if o == nil || o.OsuPrefix == nil {
		return nil, false
	}
	return o.OsuPrefix, true
}

// HasOsuPrefix returns a boolean if a field has been set.
func (o *OsuExportToCreate) HasOsuPrefix() bool {
	if o != nil && o.OsuPrefix != nil {
		return true
	}

	return false
}

// SetOsuPrefix gets a reference to the given string and assigns it to the OsuPrefix field.
func (o *OsuExportToCreate) SetOsuPrefix(v string) {
	o.OsuPrefix = &v
}

func (o OsuExportToCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DiskImageFormat"] = o.DiskImageFormat
	}
	if o.OsuApiKey != nil {
		toSerialize["OsuApiKey"] = o.OsuApiKey
	}
	if true {
		toSerialize["OsuBucket"] = o.OsuBucket
	}
	if o.OsuManifestUrl != nil {
		toSerialize["OsuManifestUrl"] = o.OsuManifestUrl
	}
	if o.OsuPrefix != nil {
		toSerialize["OsuPrefix"] = o.OsuPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableOsuExportToCreate struct {
	value *OsuExportToCreate
	isSet bool
}

func (v NullableOsuExportToCreate) Get() *OsuExportToCreate {
	return v.value
}

func (v *NullableOsuExportToCreate) Set(val *OsuExportToCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOsuExportToCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOsuExportToCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsuExportToCreate(val *OsuExportToCreate) *NullableOsuExportToCreate {
	return &NullableOsuExportToCreate{value: val, isSet: true}
}

func (v NullableOsuExportToCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsuExportToCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
