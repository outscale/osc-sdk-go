/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.27
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateSnapshotResponse struct for UpdateSnapshotResponse
type UpdateSnapshotResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	Snapshot        *Snapshot        `json:"Snapshot,omitempty"`
}

// NewUpdateSnapshotResponse instantiates a new UpdateSnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSnapshotResponse() *UpdateSnapshotResponse {
	this := UpdateSnapshotResponse{}
	return &this
}

// NewUpdateSnapshotResponseWithDefaults instantiates a new UpdateSnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSnapshotResponseWithDefaults() *UpdateSnapshotResponse {
	this := UpdateSnapshotResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateSnapshotResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateSnapshotResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateSnapshotResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *UpdateSnapshotResponse) GetSnapshot() Snapshot {
	if o == nil || o.Snapshot == nil {
		var ret Snapshot
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotResponse) GetSnapshotOk() (*Snapshot, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *UpdateSnapshotResponse) HasSnapshot() bool {
	if o != nil && o.Snapshot != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given Snapshot and assigns it to the Snapshot field.
func (o *UpdateSnapshotResponse) SetSnapshot(v Snapshot) {
	o.Snapshot = &v
}

func (o UpdateSnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Snapshot != nil {
		toSerialize["Snapshot"] = o.Snapshot
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSnapshotResponse struct {
	value *UpdateSnapshotResponse
	isSet bool
}

func (v NullableUpdateSnapshotResponse) Get() *UpdateSnapshotResponse {
	return v.value
}

func (v *NullableUpdateSnapshotResponse) Set(val *UpdateSnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSnapshotResponse(val *UpdateSnapshotResponse) *NullableUpdateSnapshotResponse {
	return &NullableUpdateSnapshotResponse{value: val, isSet: true}
}

func (v NullableUpdateSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
