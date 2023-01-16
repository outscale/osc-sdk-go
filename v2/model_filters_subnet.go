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

// FiltersSubnet One or more filters.
type FiltersSubnet struct {
	// The number of available IPs.
	AvailableIpsCounts *[]int32 `json:"AvailableIpsCounts,omitempty"`
	// The IP ranges in the Subnets, in CIDR notation (for example, `10.0.0.0/16`).
	IpRanges *[]string `json:"IpRanges,omitempty"`
	// The IDs of the Nets in which the Subnets are.
	NetIds *[]string `json:"NetIds,omitempty"`
	// The states of the Subnets (`pending` \\| `available` \\| `deleted`).
	States *[]string `json:"States,omitempty"`
	// The IDs of the Subnets.
	SubnetIds *[]string `json:"SubnetIds,omitempty"`
	// The names of the Subregions in which the Subnets are located.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
	// The keys of the tags associated with the Subnets.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Subnets.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Subnets, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersSubnet instantiates a new FiltersSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersSubnet() *FiltersSubnet {
	this := FiltersSubnet{}
	return &this
}

// NewFiltersSubnetWithDefaults instantiates a new FiltersSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersSubnetWithDefaults() *FiltersSubnet {
	this := FiltersSubnet{}
	return &this
}

// GetAvailableIpsCounts returns the AvailableIpsCounts field value if set, zero value otherwise.
func (o *FiltersSubnet) GetAvailableIpsCounts() []int32 {
	if o == nil || o.AvailableIpsCounts == nil {
		var ret []int32
		return ret
	}
	return *o.AvailableIpsCounts
}

// GetAvailableIpsCountsOk returns a tuple with the AvailableIpsCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetAvailableIpsCountsOk() (*[]int32, bool) {
	if o == nil || o.AvailableIpsCounts == nil {
		return nil, false
	}
	return o.AvailableIpsCounts, true
}

// HasAvailableIpsCounts returns a boolean if a field has been set.
func (o *FiltersSubnet) HasAvailableIpsCounts() bool {
	if o != nil && o.AvailableIpsCounts != nil {
		return true
	}

	return false
}

// SetAvailableIpsCounts gets a reference to the given []int32 and assigns it to the AvailableIpsCounts field.
func (o *FiltersSubnet) SetAvailableIpsCounts(v []int32) {
	o.AvailableIpsCounts = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *FiltersSubnet) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetIpRangesOk() (*[]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *FiltersSubnet) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *FiltersSubnet) SetIpRanges(v []string) {
	o.IpRanges = &v
}

// GetNetIds returns the NetIds field value if set, zero value otherwise.
func (o *FiltersSubnet) GetNetIds() []string {
	if o == nil || o.NetIds == nil {
		var ret []string
		return ret
	}
	return *o.NetIds
}

// GetNetIdsOk returns a tuple with the NetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetNetIdsOk() (*[]string, bool) {
	if o == nil || o.NetIds == nil {
		return nil, false
	}
	return o.NetIds, true
}

// HasNetIds returns a boolean if a field has been set.
func (o *FiltersSubnet) HasNetIds() bool {
	if o != nil && o.NetIds != nil {
		return true
	}

	return false
}

// SetNetIds gets a reference to the given []string and assigns it to the NetIds field.
func (o *FiltersSubnet) SetNetIds(v []string) {
	o.NetIds = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersSubnet) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetStatesOk() (*[]string, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersSubnet) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersSubnet) SetStates(v []string) {
	o.States = &v
}

// GetSubnetIds returns the SubnetIds field value if set, zero value otherwise.
func (o *FiltersSubnet) GetSubnetIds() []string {
	if o == nil || o.SubnetIds == nil {
		var ret []string
		return ret
	}
	return *o.SubnetIds
}

// GetSubnetIdsOk returns a tuple with the SubnetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetSubnetIdsOk() (*[]string, bool) {
	if o == nil || o.SubnetIds == nil {
		return nil, false
	}
	return o.SubnetIds, true
}

// HasSubnetIds returns a boolean if a field has been set.
func (o *FiltersSubnet) HasSubnetIds() bool {
	if o != nil && o.SubnetIds != nil {
		return true
	}

	return false
}

// SetSubnetIds gets a reference to the given []string and assigns it to the SubnetIds field.
func (o *FiltersSubnet) SetSubnetIds(v []string) {
	o.SubnetIds = &v
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *FiltersSubnet) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetSubregionNamesOk() (*[]string, bool) {
	if o == nil || o.SubregionNames == nil {
		return nil, false
	}
	return o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *FiltersSubnet) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *FiltersSubnet) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersSubnet) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersSubnet) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersSubnet) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersSubnet) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersSubnet) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersSubnet) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersSubnet) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubnet) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersSubnet) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersSubnet) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersSubnet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailableIpsCounts != nil {
		toSerialize["AvailableIpsCounts"] = o.AvailableIpsCounts
	}
	if o.IpRanges != nil {
		toSerialize["IpRanges"] = o.IpRanges
	}
	if o.NetIds != nil {
		toSerialize["NetIds"] = o.NetIds
	}
	if o.States != nil {
		toSerialize["States"] = o.States
	}
	if o.SubnetIds != nil {
		toSerialize["SubnetIds"] = o.SubnetIds
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
	return json.Marshal(toSerialize)
}

type NullableFiltersSubnet struct {
	value *FiltersSubnet
	isSet bool
}

func (v NullableFiltersSubnet) Get() *FiltersSubnet {
	return v.value
}

func (v *NullableFiltersSubnet) Set(val *FiltersSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersSubnet(val *FiltersSubnet) *NullableFiltersSubnet {
	return &NullableFiltersSubnet{value: val, isSet: true}
}

func (v NullableFiltersSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
