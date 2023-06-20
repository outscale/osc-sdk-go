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

// PermissionsOnResourceCreation Information about the permissions for the resource.<br /> Specify either the `Additions` or the `Removals` parameter.
type PermissionsOnResourceCreation struct {
	Additions *PermissionsOnResource `json:"Additions,omitempty"`
	Removals  *PermissionsOnResource `json:"Removals,omitempty"`
}

// NewPermissionsOnResourceCreation instantiates a new PermissionsOnResourceCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsOnResourceCreation() *PermissionsOnResourceCreation {
	this := PermissionsOnResourceCreation{}
	return &this
}

// NewPermissionsOnResourceCreationWithDefaults instantiates a new PermissionsOnResourceCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsOnResourceCreationWithDefaults() *PermissionsOnResourceCreation {
	this := PermissionsOnResourceCreation{}
	return &this
}

// GetAdditions returns the Additions field value if set, zero value otherwise.
func (o *PermissionsOnResourceCreation) GetAdditions() PermissionsOnResource {
	if o == nil || o.Additions == nil {
		var ret PermissionsOnResource
		return ret
	}
	return *o.Additions
}

// GetAdditionsOk returns a tuple with the Additions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsOnResourceCreation) GetAdditionsOk() (*PermissionsOnResource, bool) {
	if o == nil || o.Additions == nil {
		return nil, false
	}
	return o.Additions, true
}

// HasAdditions returns a boolean if a field has been set.
func (o *PermissionsOnResourceCreation) HasAdditions() bool {
	if o != nil && o.Additions != nil {
		return true
	}

	return false
}

// SetAdditions gets a reference to the given PermissionsOnResource and assigns it to the Additions field.
func (o *PermissionsOnResourceCreation) SetAdditions(v PermissionsOnResource) {
	o.Additions = &v
}

// GetRemovals returns the Removals field value if set, zero value otherwise.
func (o *PermissionsOnResourceCreation) GetRemovals() PermissionsOnResource {
	if o == nil || o.Removals == nil {
		var ret PermissionsOnResource
		return ret
	}
	return *o.Removals
}

// GetRemovalsOk returns a tuple with the Removals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsOnResourceCreation) GetRemovalsOk() (*PermissionsOnResource, bool) {
	if o == nil || o.Removals == nil {
		return nil, false
	}
	return o.Removals, true
}

// HasRemovals returns a boolean if a field has been set.
func (o *PermissionsOnResourceCreation) HasRemovals() bool {
	if o != nil && o.Removals != nil {
		return true
	}

	return false
}

// SetRemovals gets a reference to the given PermissionsOnResource and assigns it to the Removals field.
func (o *PermissionsOnResourceCreation) SetRemovals(v PermissionsOnResource) {
	o.Removals = &v
}

func (o PermissionsOnResourceCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Additions != nil {
		toSerialize["Additions"] = o.Additions
	}
	if o.Removals != nil {
		toSerialize["Removals"] = o.Removals
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionsOnResourceCreation struct {
	value *PermissionsOnResourceCreation
	isSet bool
}

func (v NullablePermissionsOnResourceCreation) Get() *PermissionsOnResourceCreation {
	return v.value
}

func (v *NullablePermissionsOnResourceCreation) Set(val *PermissionsOnResourceCreation) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsOnResourceCreation) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsOnResourceCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsOnResourceCreation(val *PermissionsOnResourceCreation) *NullablePermissionsOnResourceCreation {
	return &NullablePermissionsOnResourceCreation{value: val, isSet: true}
}

func (v NullablePermissionsOnResourceCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsOnResourceCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
