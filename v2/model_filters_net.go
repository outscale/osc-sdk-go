/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.22
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersNet One or more filters.
type FiltersNet struct {
	// The IDs of the DHCP options sets.
	DhcpOptionsSetIds *[]string `json:"DhcpOptionsSetIds,omitempty"`
	// The IP ranges for the Nets, in CIDR notation (for example, `10.0.0.0/16`).
	IpRanges *[]string `json:"IpRanges,omitempty"`
	// If true, the Net used is the default one.
	IsDefault *bool `json:"IsDefault,omitempty"`
	// The IDs of the Nets.
	NetIds *[]string `json:"NetIds,omitempty"`
	// The states of the Nets (`pending` \\| `available` \\| `deleted`).
	States *[]string `json:"States,omitempty"`
	// The keys of the tags associated with the Nets.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Nets.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Nets, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersNet instantiates a new FiltersNet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersNet() *FiltersNet {
	this := FiltersNet{}
	return &this
}

// NewFiltersNetWithDefaults instantiates a new FiltersNet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersNetWithDefaults() *FiltersNet {
	this := FiltersNet{}
	return &this
}

// GetDhcpOptionsSetIds returns the DhcpOptionsSetIds field value if set, zero value otherwise.
func (o *FiltersNet) GetDhcpOptionsSetIds() []string {
	if o == nil || o.DhcpOptionsSetIds == nil {
		var ret []string
		return ret
	}
	return *o.DhcpOptionsSetIds
}

// GetDhcpOptionsSetIdsOk returns a tuple with the DhcpOptionsSetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNet) GetDhcpOptionsSetIdsOk() (*[]string, bool) {
	if o == nil || o.DhcpOptionsSetIds == nil {
		return nil, false
	}
	return o.DhcpOptionsSetIds, true
}

// HasDhcpOptionsSetIds returns a boolean if a field has been set.
func (o *FiltersNet) HasDhcpOptionsSetIds() bool {
	if o != nil && o.DhcpOptionsSetIds != nil {
		return true
	}

	return false
}

// SetDhcpOptionsSetIds gets a reference to the given []string and assigns it to the DhcpOptionsSetIds field.
func (o *FiltersNet) SetDhcpOptionsSetIds(v []string) {
	o.DhcpOptionsSetIds = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *FiltersNet) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNet) GetIpRangesOk() (*[]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *FiltersNet) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *FiltersNet) SetIpRanges(v []string) {
	o.IpRanges = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *FiltersNet) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNet) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *FiltersNet) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *FiltersNet) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetNetIds returns the NetIds field value if set, zero value otherwise.
func (o *FiltersNet) GetNetIds() []string {
	if o == nil || o.NetIds == nil {
		var ret []string
		return ret
	}
	return *o.NetIds
}

// GetNetIdsOk returns a tuple with the NetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNet) GetNetIdsOk() (*[]string, bool) {
	if o == nil || o.NetIds == nil {
		return nil, false
	}
	return o.NetIds, true
}

// HasNetIds returns a boolean if a field has been set.
func (o *FiltersNet) HasNetIds() bool {
	if o != nil && o.NetIds != nil {
		return true
	}

	return false
}

// SetNetIds gets a reference to the given []string and assigns it to the NetIds field.
func (o *FiltersNet) SetNetIds(v []string) {
	o.NetIds = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersNet) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNet) GetStatesOk() (*[]string, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersNet) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersNet) SetStates(v []string) {
	o.States = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersNet) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNet) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersNet) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersNet) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersNet) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNet) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersNet) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersNet) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersNet) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNet) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersNet) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersNet) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersNet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DhcpOptionsSetIds != nil {
		toSerialize["DhcpOptionsSetIds"] = o.DhcpOptionsSetIds
	}
	if o.IpRanges != nil {
		toSerialize["IpRanges"] = o.IpRanges
	}
	if o.IsDefault != nil {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if o.NetIds != nil {
		toSerialize["NetIds"] = o.NetIds
	}
	if o.States != nil {
		toSerialize["States"] = o.States
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
	return json.Marshal(toSerialize)
}

type NullableFiltersNet struct {
	value *FiltersNet
	isSet bool
}

func (v NullableFiltersNet) Get() *FiltersNet {
	return v.value
}

func (v *NullableFiltersNet) Set(val *FiltersNet) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersNet) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersNet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersNet(val *FiltersNet) *NullableFiltersNet {
	return &NullableFiltersNet{value: val, isSet: true}
}

func (v NullableFiltersNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersNet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
