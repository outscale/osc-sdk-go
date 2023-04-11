/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
	"time"
)

// Snapshot Information about the snapshot.
type Snapshot struct {
	// The account alias of the owner of the snapshot.
	AccountAlias *string `json:"AccountAlias,omitempty"`
	// The account ID of the owner of the snapshot.
	AccountId *string `json:"AccountId,omitempty"`
	// The date and time of creation of the snapshot.
	CreationDate *time.Time `json:"CreationDate,omitempty"`
	// The description of the snapshot.
	Description               *string                `json:"Description,omitempty"`
	PermissionsToCreateVolume *PermissionsOnResource `json:"PermissionsToCreateVolume,omitempty"`
	// The progress of the snapshot, as a percentage.
	Progress *int32 `json:"Progress,omitempty"`
	// The ID of the snapshot.
	SnapshotId *string `json:"SnapshotId,omitempty"`
	// The state of the snapshot (`in-queue` \\| `completed` \\| `error`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the snapshot.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the volume used to create the snapshot.
	VolumeId *string `json:"VolumeId,omitempty"`
	// The size of the volume used to create the snapshot, in gibibytes (GiB).
	VolumeSize *int32 `json:"VolumeSize,omitempty"`
}

// NewSnapshot instantiates a new Snapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshot() *Snapshot {
	this := Snapshot{}
	return &this
}

// NewSnapshotWithDefaults instantiates a new Snapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotWithDefaults() *Snapshot {
	this := Snapshot{}
	return &this
}

// GetAccountAlias returns the AccountAlias field value if set, zero value otherwise.
func (o *Snapshot) GetAccountAlias() string {
	if o == nil || o.AccountAlias == nil {
		var ret string
		return ret
	}
	return *o.AccountAlias
}

// GetAccountAliasOk returns a tuple with the AccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetAccountAliasOk() (*string, bool) {
	if o == nil || o.AccountAlias == nil {
		return nil, false
	}
	return o.AccountAlias, true
}

// HasAccountAlias returns a boolean if a field has been set.
func (o *Snapshot) HasAccountAlias() bool {
	if o != nil && o.AccountAlias != nil {
		return true
	}

	return false
}

// SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.
func (o *Snapshot) SetAccountAlias(v string) {
	o.AccountAlias = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Snapshot) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Snapshot) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Snapshot) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Snapshot) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Snapshot) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Snapshot) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Snapshot) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Snapshot) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Snapshot) SetDescription(v string) {
	o.Description = &v
}

// GetPermissionsToCreateVolume returns the PermissionsToCreateVolume field value if set, zero value otherwise.
func (o *Snapshot) GetPermissionsToCreateVolume() PermissionsOnResource {
	if o == nil || o.PermissionsToCreateVolume == nil {
		var ret PermissionsOnResource
		return ret
	}
	return *o.PermissionsToCreateVolume
}

// GetPermissionsToCreateVolumeOk returns a tuple with the PermissionsToCreateVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetPermissionsToCreateVolumeOk() (*PermissionsOnResource, bool) {
	if o == nil || o.PermissionsToCreateVolume == nil {
		return nil, false
	}
	return o.PermissionsToCreateVolume, true
}

// HasPermissionsToCreateVolume returns a boolean if a field has been set.
func (o *Snapshot) HasPermissionsToCreateVolume() bool {
	if o != nil && o.PermissionsToCreateVolume != nil {
		return true
	}

	return false
}

// SetPermissionsToCreateVolume gets a reference to the given PermissionsOnResource and assigns it to the PermissionsToCreateVolume field.
func (o *Snapshot) SetPermissionsToCreateVolume(v PermissionsOnResource) {
	o.PermissionsToCreateVolume = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *Snapshot) GetProgress() int32 {
	if o == nil || o.Progress == nil {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetProgressOk() (*int32, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *Snapshot) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *Snapshot) SetProgress(v int32) {
	o.Progress = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *Snapshot) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *Snapshot) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *Snapshot) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Snapshot) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Snapshot) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Snapshot) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Snapshot) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Snapshot) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *Snapshot) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *Snapshot) GetVolumeId() string {
	if o == nil || o.VolumeId == nil {
		var ret string
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetVolumeIdOk() (*string, bool) {
	if o == nil || o.VolumeId == nil {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *Snapshot) HasVolumeId() bool {
	if o != nil && o.VolumeId != nil {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *Snapshot) SetVolumeId(v string) {
	o.VolumeId = &v
}

// GetVolumeSize returns the VolumeSize field value if set, zero value otherwise.
func (o *Snapshot) GetVolumeSize() int32 {
	if o == nil || o.VolumeSize == nil {
		var ret int32
		return ret
	}
	return *o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetVolumeSizeOk() (*int32, bool) {
	if o == nil || o.VolumeSize == nil {
		return nil, false
	}
	return o.VolumeSize, true
}

// HasVolumeSize returns a boolean if a field has been set.
func (o *Snapshot) HasVolumeSize() bool {
	if o != nil && o.VolumeSize != nil {
		return true
	}

	return false
}

// SetVolumeSize gets a reference to the given int32 and assigns it to the VolumeSize field.
func (o *Snapshot) SetVolumeSize(v int32) {
	o.VolumeSize = &v
}

func (o Snapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountAlias != nil {
		toSerialize["AccountAlias"] = o.AccountAlias
	}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.CreationDate != nil {
		toSerialize["CreationDate"] = o.CreationDate
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.PermissionsToCreateVolume != nil {
		toSerialize["PermissionsToCreateVolume"] = o.PermissionsToCreateVolume
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
	if o.VolumeId != nil {
		toSerialize["VolumeId"] = o.VolumeId
	}
	if o.VolumeSize != nil {
		toSerialize["VolumeSize"] = o.VolumeSize
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshot struct {
	value *Snapshot
	isSet bool
}

func (v NullableSnapshot) Get() *Snapshot {
	return v.value
}

func (v *NullableSnapshot) Set(val *Snapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshot(val *Snapshot) *NullableSnapshot {
	return &NullableSnapshot{value: val, isSet: true}
}

func (v NullableSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
