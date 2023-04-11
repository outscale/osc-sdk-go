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
)

// ReadSnapshotsResponse struct for ReadSnapshotsResponse
type ReadSnapshotsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more snapshots and their permissions.
	Snapshots *[]Snapshot `json:"Snapshots,omitempty"`
}

// NewReadSnapshotsResponse instantiates a new ReadSnapshotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadSnapshotsResponse() *ReadSnapshotsResponse {
	this := ReadSnapshotsResponse{}
	return &this
}

// NewReadSnapshotsResponseWithDefaults instantiates a new ReadSnapshotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadSnapshotsResponseWithDefaults() *ReadSnapshotsResponse {
	this := ReadSnapshotsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadSnapshotsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSnapshotsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadSnapshotsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadSnapshotsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise.
func (o *ReadSnapshotsResponse) GetSnapshots() []Snapshot {
	if o == nil || o.Snapshots == nil {
		var ret []Snapshot
		return ret
	}
	return *o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSnapshotsResponse) GetSnapshotsOk() (*[]Snapshot, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return o.Snapshots, true
}

// HasSnapshots returns a boolean if a field has been set.
func (o *ReadSnapshotsResponse) HasSnapshots() bool {
	if o != nil && o.Snapshots != nil {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given []Snapshot and assigns it to the Snapshots field.
func (o *ReadSnapshotsResponse) SetSnapshots(v []Snapshot) {
	o.Snapshots = &v
}

func (o ReadSnapshotsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Snapshots != nil {
		toSerialize["Snapshots"] = o.Snapshots
	}
	return json.Marshal(toSerialize)
}

type NullableReadSnapshotsResponse struct {
	value *ReadSnapshotsResponse
	isSet bool
}

func (v NullableReadSnapshotsResponse) Get() *ReadSnapshotsResponse {
	return v.value
}

func (v *NullableReadSnapshotsResponse) Set(val *ReadSnapshotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadSnapshotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadSnapshotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadSnapshotsResponse(val *ReadSnapshotsResponse) *NullableReadSnapshotsResponse {
	return &NullableReadSnapshotsResponse{value: val, isSet: true}
}

func (v NullableReadSnapshotsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadSnapshotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
