/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersVolume One or more filters.
type FiltersVolume struct {
	// The dates and times of creation of the volumes, in ISO 8601 date-time format (for example, `2020-06-30T00:00:00.000Z`).
	CreationDates *[]string `json:"CreationDates,omitempty"`
	// Whether the volumes are deleted or not when terminating the VMs.
	LinkVolumeDeleteOnVmDeletion *bool `json:"LinkVolumeDeleteOnVmDeletion,omitempty"`
	// The VM device names.
	LinkVolumeDeviceNames *[]string `json:"LinkVolumeDeviceNames,omitempty"`
	// The dates and times of creation of the volumes, in ISO 8601 date-time format (for example, `2020-06-30T00:00:00.000Z`).
	LinkVolumeLinkDates *[]string `json:"LinkVolumeLinkDates,omitempty"`
	// The attachment states of the volumes (`attaching` \\| `detaching` \\| `attached` \\| `detached`).
	LinkVolumeLinkStates *[]string `json:"LinkVolumeLinkStates,omitempty"`
	// One or more IDs of VMs.
	LinkVolumeVmIds *[]string `json:"LinkVolumeVmIds,omitempty"`
	// The snapshots from which the volumes were created.
	SnapshotIds *[]string `json:"SnapshotIds,omitempty"`
	// The names of the Subregions in which the volumes were created.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
	// The keys of the tags associated with the volumes.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the volumes.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the volumes, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
	// The IDs of the volumes.
	VolumeIds *[]string `json:"VolumeIds,omitempty"`
	// The sizes of the volumes, in gibibytes (GiB).
	VolumeSizes *[]int32 `json:"VolumeSizes,omitempty"`
	// The states of the volumes (`creating` \\| `available` \\| `in-use` \\| `updating` \\| `deleting` \\| `error`).
	VolumeStates *[]string `json:"VolumeStates,omitempty"`
	// The types of the volumes (`standard` \\| `gp2` \\| `io1`).
	VolumeTypes *[]string `json:"VolumeTypes,omitempty"`
}

// NewFiltersVolume instantiates a new FiltersVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersVolume() *FiltersVolume {
	this := FiltersVolume{}
	return &this
}

// NewFiltersVolumeWithDefaults instantiates a new FiltersVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersVolumeWithDefaults() *FiltersVolume {
	this := FiltersVolume{}
	return &this
}

// GetCreationDates returns the CreationDates field value if set, zero value otherwise.
func (o *FiltersVolume) GetCreationDates() []string {
	if o == nil || o.CreationDates == nil {
		var ret []string
		return ret
	}
	return *o.CreationDates
}

// GetCreationDatesOk returns a tuple with the CreationDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetCreationDatesOk() (*[]string, bool) {
	if o == nil || o.CreationDates == nil {
		return nil, false
	}
	return o.CreationDates, true
}

// HasCreationDates returns a boolean if a field has been set.
func (o *FiltersVolume) HasCreationDates() bool {
	if o != nil && o.CreationDates != nil {
		return true
	}

	return false
}

// SetCreationDates gets a reference to the given []string and assigns it to the CreationDates field.
func (o *FiltersVolume) SetCreationDates(v []string) {
	o.CreationDates = &v
}

// GetLinkVolumeDeleteOnVmDeletion returns the LinkVolumeDeleteOnVmDeletion field value if set, zero value otherwise.
func (o *FiltersVolume) GetLinkVolumeDeleteOnVmDeletion() bool {
	if o == nil || o.LinkVolumeDeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.LinkVolumeDeleteOnVmDeletion
}

// GetLinkVolumeDeleteOnVmDeletionOk returns a tuple with the LinkVolumeDeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetLinkVolumeDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.LinkVolumeDeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.LinkVolumeDeleteOnVmDeletion, true
}

// HasLinkVolumeDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *FiltersVolume) HasLinkVolumeDeleteOnVmDeletion() bool {
	if o != nil && o.LinkVolumeDeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetLinkVolumeDeleteOnVmDeletion gets a reference to the given bool and assigns it to the LinkVolumeDeleteOnVmDeletion field.
func (o *FiltersVolume) SetLinkVolumeDeleteOnVmDeletion(v bool) {
	o.LinkVolumeDeleteOnVmDeletion = &v
}

// GetLinkVolumeDeviceNames returns the LinkVolumeDeviceNames field value if set, zero value otherwise.
func (o *FiltersVolume) GetLinkVolumeDeviceNames() []string {
	if o == nil || o.LinkVolumeDeviceNames == nil {
		var ret []string
		return ret
	}
	return *o.LinkVolumeDeviceNames
}

// GetLinkVolumeDeviceNamesOk returns a tuple with the LinkVolumeDeviceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetLinkVolumeDeviceNamesOk() (*[]string, bool) {
	if o == nil || o.LinkVolumeDeviceNames == nil {
		return nil, false
	}
	return o.LinkVolumeDeviceNames, true
}

// HasLinkVolumeDeviceNames returns a boolean if a field has been set.
func (o *FiltersVolume) HasLinkVolumeDeviceNames() bool {
	if o != nil && o.LinkVolumeDeviceNames != nil {
		return true
	}

	return false
}

// SetLinkVolumeDeviceNames gets a reference to the given []string and assigns it to the LinkVolumeDeviceNames field.
func (o *FiltersVolume) SetLinkVolumeDeviceNames(v []string) {
	o.LinkVolumeDeviceNames = &v
}

// GetLinkVolumeLinkDates returns the LinkVolumeLinkDates field value if set, zero value otherwise.
func (o *FiltersVolume) GetLinkVolumeLinkDates() []string {
	if o == nil || o.LinkVolumeLinkDates == nil {
		var ret []string
		return ret
	}
	return *o.LinkVolumeLinkDates
}

// GetLinkVolumeLinkDatesOk returns a tuple with the LinkVolumeLinkDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetLinkVolumeLinkDatesOk() (*[]string, bool) {
	if o == nil || o.LinkVolumeLinkDates == nil {
		return nil, false
	}
	return o.LinkVolumeLinkDates, true
}

// HasLinkVolumeLinkDates returns a boolean if a field has been set.
func (o *FiltersVolume) HasLinkVolumeLinkDates() bool {
	if o != nil && o.LinkVolumeLinkDates != nil {
		return true
	}

	return false
}

// SetLinkVolumeLinkDates gets a reference to the given []string and assigns it to the LinkVolumeLinkDates field.
func (o *FiltersVolume) SetLinkVolumeLinkDates(v []string) {
	o.LinkVolumeLinkDates = &v
}

// GetLinkVolumeLinkStates returns the LinkVolumeLinkStates field value if set, zero value otherwise.
func (o *FiltersVolume) GetLinkVolumeLinkStates() []string {
	if o == nil || o.LinkVolumeLinkStates == nil {
		var ret []string
		return ret
	}
	return *o.LinkVolumeLinkStates
}

// GetLinkVolumeLinkStatesOk returns a tuple with the LinkVolumeLinkStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetLinkVolumeLinkStatesOk() (*[]string, bool) {
	if o == nil || o.LinkVolumeLinkStates == nil {
		return nil, false
	}
	return o.LinkVolumeLinkStates, true
}

// HasLinkVolumeLinkStates returns a boolean if a field has been set.
func (o *FiltersVolume) HasLinkVolumeLinkStates() bool {
	if o != nil && o.LinkVolumeLinkStates != nil {
		return true
	}

	return false
}

// SetLinkVolumeLinkStates gets a reference to the given []string and assigns it to the LinkVolumeLinkStates field.
func (o *FiltersVolume) SetLinkVolumeLinkStates(v []string) {
	o.LinkVolumeLinkStates = &v
}

// GetLinkVolumeVmIds returns the LinkVolumeVmIds field value if set, zero value otherwise.
func (o *FiltersVolume) GetLinkVolumeVmIds() []string {
	if o == nil || o.LinkVolumeVmIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkVolumeVmIds
}

// GetLinkVolumeVmIdsOk returns a tuple with the LinkVolumeVmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetLinkVolumeVmIdsOk() (*[]string, bool) {
	if o == nil || o.LinkVolumeVmIds == nil {
		return nil, false
	}
	return o.LinkVolumeVmIds, true
}

// HasLinkVolumeVmIds returns a boolean if a field has been set.
func (o *FiltersVolume) HasLinkVolumeVmIds() bool {
	if o != nil && o.LinkVolumeVmIds != nil {
		return true
	}

	return false
}

// SetLinkVolumeVmIds gets a reference to the given []string and assigns it to the LinkVolumeVmIds field.
func (o *FiltersVolume) SetLinkVolumeVmIds(v []string) {
	o.LinkVolumeVmIds = &v
}

// GetSnapshotIds returns the SnapshotIds field value if set, zero value otherwise.
func (o *FiltersVolume) GetSnapshotIds() []string {
	if o == nil || o.SnapshotIds == nil {
		var ret []string
		return ret
	}
	return *o.SnapshotIds
}

// GetSnapshotIdsOk returns a tuple with the SnapshotIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetSnapshotIdsOk() (*[]string, bool) {
	if o == nil || o.SnapshotIds == nil {
		return nil, false
	}
	return o.SnapshotIds, true
}

// HasSnapshotIds returns a boolean if a field has been set.
func (o *FiltersVolume) HasSnapshotIds() bool {
	if o != nil && o.SnapshotIds != nil {
		return true
	}

	return false
}

// SetSnapshotIds gets a reference to the given []string and assigns it to the SnapshotIds field.
func (o *FiltersVolume) SetSnapshotIds(v []string) {
	o.SnapshotIds = &v
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *FiltersVolume) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetSubregionNamesOk() (*[]string, bool) {
	if o == nil || o.SubregionNames == nil {
		return nil, false
	}
	return o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *FiltersVolume) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *FiltersVolume) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersVolume) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersVolume) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersVolume) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersVolume) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersVolume) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersVolume) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersVolume) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersVolume) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersVolume) SetTags(v []string) {
	o.Tags = &v
}

// GetVolumeIds returns the VolumeIds field value if set, zero value otherwise.
func (o *FiltersVolume) GetVolumeIds() []string {
	if o == nil || o.VolumeIds == nil {
		var ret []string
		return ret
	}
	return *o.VolumeIds
}

// GetVolumeIdsOk returns a tuple with the VolumeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetVolumeIdsOk() (*[]string, bool) {
	if o == nil || o.VolumeIds == nil {
		return nil, false
	}
	return o.VolumeIds, true
}

// HasVolumeIds returns a boolean if a field has been set.
func (o *FiltersVolume) HasVolumeIds() bool {
	if o != nil && o.VolumeIds != nil {
		return true
	}

	return false
}

// SetVolumeIds gets a reference to the given []string and assigns it to the VolumeIds field.
func (o *FiltersVolume) SetVolumeIds(v []string) {
	o.VolumeIds = &v
}

// GetVolumeSizes returns the VolumeSizes field value if set, zero value otherwise.
func (o *FiltersVolume) GetVolumeSizes() []int32 {
	if o == nil || o.VolumeSizes == nil {
		var ret []int32
		return ret
	}
	return *o.VolumeSizes
}

// GetVolumeSizesOk returns a tuple with the VolumeSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetVolumeSizesOk() (*[]int32, bool) {
	if o == nil || o.VolumeSizes == nil {
		return nil, false
	}
	return o.VolumeSizes, true
}

// HasVolumeSizes returns a boolean if a field has been set.
func (o *FiltersVolume) HasVolumeSizes() bool {
	if o != nil && o.VolumeSizes != nil {
		return true
	}

	return false
}

// SetVolumeSizes gets a reference to the given []int32 and assigns it to the VolumeSizes field.
func (o *FiltersVolume) SetVolumeSizes(v []int32) {
	o.VolumeSizes = &v
}

// GetVolumeStates returns the VolumeStates field value if set, zero value otherwise.
func (o *FiltersVolume) GetVolumeStates() []string {
	if o == nil || o.VolumeStates == nil {
		var ret []string
		return ret
	}
	return *o.VolumeStates
}

// GetVolumeStatesOk returns a tuple with the VolumeStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetVolumeStatesOk() (*[]string, bool) {
	if o == nil || o.VolumeStates == nil {
		return nil, false
	}
	return o.VolumeStates, true
}

// HasVolumeStates returns a boolean if a field has been set.
func (o *FiltersVolume) HasVolumeStates() bool {
	if o != nil && o.VolumeStates != nil {
		return true
	}

	return false
}

// SetVolumeStates gets a reference to the given []string and assigns it to the VolumeStates field.
func (o *FiltersVolume) SetVolumeStates(v []string) {
	o.VolumeStates = &v
}

// GetVolumeTypes returns the VolumeTypes field value if set, zero value otherwise.
func (o *FiltersVolume) GetVolumeTypes() []string {
	if o == nil || o.VolumeTypes == nil {
		var ret []string
		return ret
	}
	return *o.VolumeTypes
}

// GetVolumeTypesOk returns a tuple with the VolumeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVolume) GetVolumeTypesOk() (*[]string, bool) {
	if o == nil || o.VolumeTypes == nil {
		return nil, false
	}
	return o.VolumeTypes, true
}

// HasVolumeTypes returns a boolean if a field has been set.
func (o *FiltersVolume) HasVolumeTypes() bool {
	if o != nil && o.VolumeTypes != nil {
		return true
	}

	return false
}

// SetVolumeTypes gets a reference to the given []string and assigns it to the VolumeTypes field.
func (o *FiltersVolume) SetVolumeTypes(v []string) {
	o.VolumeTypes = &v
}

func (o FiltersVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationDates != nil {
		toSerialize["CreationDates"] = o.CreationDates
	}
	if o.LinkVolumeDeleteOnVmDeletion != nil {
		toSerialize["LinkVolumeDeleteOnVmDeletion"] = o.LinkVolumeDeleteOnVmDeletion
	}
	if o.LinkVolumeDeviceNames != nil {
		toSerialize["LinkVolumeDeviceNames"] = o.LinkVolumeDeviceNames
	}
	if o.LinkVolumeLinkDates != nil {
		toSerialize["LinkVolumeLinkDates"] = o.LinkVolumeLinkDates
	}
	if o.LinkVolumeLinkStates != nil {
		toSerialize["LinkVolumeLinkStates"] = o.LinkVolumeLinkStates
	}
	if o.LinkVolumeVmIds != nil {
		toSerialize["LinkVolumeVmIds"] = o.LinkVolumeVmIds
	}
	if o.SnapshotIds != nil {
		toSerialize["SnapshotIds"] = o.SnapshotIds
	}
	if o.SubregionNames != nil {
		toSerialize["SubregionNames"] = o.SubregionNames
	}
	if o.TagKeys != nil {
		toSerialize["TagKeys"] = o.TagKeys
	}
	if o.TagValues != nil {
		toSerialize["TagValues"] = o.TagValues
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.VolumeIds != nil {
		toSerialize["VolumeIds"] = o.VolumeIds
	}
	if o.VolumeSizes != nil {
		toSerialize["VolumeSizes"] = o.VolumeSizes
	}
	if o.VolumeStates != nil {
		toSerialize["VolumeStates"] = o.VolumeStates
	}
	if o.VolumeTypes != nil {
		toSerialize["VolumeTypes"] = o.VolumeTypes
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersVolume struct {
	value *FiltersVolume
	isSet bool
}

func (v NullableFiltersVolume) Get() *FiltersVolume {
	return v.value
}

func (v *NullableFiltersVolume) Set(val *FiltersVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersVolume(val *FiltersVolume) *NullableFiltersVolume {
	return &NullableFiltersVolume{value: val, isSet: true}
}

func (v NullableFiltersVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
