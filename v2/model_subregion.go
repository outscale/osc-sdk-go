/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// Subregion Information about the Subregion.
type Subregion struct {
	// The location code of the Subregion.
	LocationCode *string `json:"LocationCode,omitempty"`
	// The name of the Region containing the Subregion.
	RegionName *string `json:"RegionName,omitempty"`
	// The state of the Subregion (`available` \\| `information` \\| `impaired` \\| `unavailable`).
	State *string `json:"State,omitempty"`
	// The name of the Subregion.
	SubregionName *string `json:"SubregionName,omitempty"`
}

// NewSubregion instantiates a new Subregion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubregion() *Subregion {
	this := Subregion{}
	return &this
}

// NewSubregionWithDefaults instantiates a new Subregion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubregionWithDefaults() *Subregion {
	this := Subregion{}
	return &this
}

// GetLocationCode returns the LocationCode field value if set, zero value otherwise.
func (o *Subregion) GetLocationCode() string {
	if o == nil || o.LocationCode == nil {
		var ret string
		return ret
	}
	return *o.LocationCode
}

// GetLocationCodeOk returns a tuple with the LocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subregion) GetLocationCodeOk() (*string, bool) {
	if o == nil || o.LocationCode == nil {
		return nil, false
	}
	return o.LocationCode, true
}

// HasLocationCode returns a boolean if a field has been set.
func (o *Subregion) HasLocationCode() bool {
	if o != nil && o.LocationCode != nil {
		return true
	}

	return false
}

// SetLocationCode gets a reference to the given string and assigns it to the LocationCode field.
func (o *Subregion) SetLocationCode(v string) {
	o.LocationCode = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *Subregion) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subregion) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *Subregion) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *Subregion) SetRegionName(v string) {
	o.RegionName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Subregion) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subregion) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Subregion) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Subregion) SetState(v string) {
	o.State = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *Subregion) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subregion) GetSubregionNameOk() (*string, bool) {
	if o == nil || o.SubregionName == nil {
		return nil, false
	}
	return o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *Subregion) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *Subregion) SetSubregionName(v string) {
	o.SubregionName = &v
}

func (o Subregion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocationCode != nil {
		toSerialize["LocationCode"] = o.LocationCode
	}
	if o.RegionName != nil {
		toSerialize["RegionName"] = o.RegionName
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SubregionName != nil {
		toSerialize["SubregionName"] = o.SubregionName
	}
	return json.Marshal(toSerialize)
}

type NullableSubregion struct {
	value *Subregion
	isSet bool
}

func (v NullableSubregion) Get() *Subregion {
	return v.value
}

func (v *NullableSubregion) Set(val *Subregion) {
	v.value = val
	v.isSet = true
}

func (v NullableSubregion) IsSet() bool {
	return v.isSet
}

func (v *NullableSubregion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubregion(val *Subregion) *NullableSubregion {
	return &NullableSubregion{value: val, isSet: true}
}

func (v NullableSubregion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubregion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
