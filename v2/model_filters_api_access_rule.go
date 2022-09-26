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

// FiltersApiAccessRule One or more filters.
type FiltersApiAccessRule struct {
	// One or more IDs of API access rules.
	ApiAccessRuleIds *[]string `json:"ApiAccessRuleIds,omitempty"`
	// One or more IDs of Client Certificate Authorities (CAs).
	CaIds *[]string `json:"CaIds,omitempty"`
	// One or more Client Certificate Common Names (CNs).
	Cns *[]string `json:"Cns,omitempty"`
	// One or more descriptions of API access rules.
	Descriptions *[]string `json:"Descriptions,omitempty"`
	// One or more IP ranges, in CIDR notation (for example, `192.0.2.0/16`).
	IpRanges *[]string `json:"IpRanges,omitempty"`
}

// NewFiltersApiAccessRule instantiates a new FiltersApiAccessRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersApiAccessRule() *FiltersApiAccessRule {
	this := FiltersApiAccessRule{}
	return &this
}

// NewFiltersApiAccessRuleWithDefaults instantiates a new FiltersApiAccessRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersApiAccessRuleWithDefaults() *FiltersApiAccessRule {
	this := FiltersApiAccessRule{}
	return &this
}

// GetApiAccessRuleIds returns the ApiAccessRuleIds field value if set, zero value otherwise.
func (o *FiltersApiAccessRule) GetApiAccessRuleIds() []string {
	if o == nil || o.ApiAccessRuleIds == nil {
		var ret []string
		return ret
	}
	return *o.ApiAccessRuleIds
}

// GetApiAccessRuleIdsOk returns a tuple with the ApiAccessRuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiAccessRule) GetApiAccessRuleIdsOk() (*[]string, bool) {
	if o == nil || o.ApiAccessRuleIds == nil {
		return nil, false
	}
	return o.ApiAccessRuleIds, true
}

// HasApiAccessRuleIds returns a boolean if a field has been set.
func (o *FiltersApiAccessRule) HasApiAccessRuleIds() bool {
	if o != nil && o.ApiAccessRuleIds != nil {
		return true
	}

	return false
}

// SetApiAccessRuleIds gets a reference to the given []string and assigns it to the ApiAccessRuleIds field.
func (o *FiltersApiAccessRule) SetApiAccessRuleIds(v []string) {
	o.ApiAccessRuleIds = &v
}

// GetCaIds returns the CaIds field value if set, zero value otherwise.
func (o *FiltersApiAccessRule) GetCaIds() []string {
	if o == nil || o.CaIds == nil {
		var ret []string
		return ret
	}
	return *o.CaIds
}

// GetCaIdsOk returns a tuple with the CaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiAccessRule) GetCaIdsOk() (*[]string, bool) {
	if o == nil || o.CaIds == nil {
		return nil, false
	}
	return o.CaIds, true
}

// HasCaIds returns a boolean if a field has been set.
func (o *FiltersApiAccessRule) HasCaIds() bool {
	if o != nil && o.CaIds != nil {
		return true
	}

	return false
}

// SetCaIds gets a reference to the given []string and assigns it to the CaIds field.
func (o *FiltersApiAccessRule) SetCaIds(v []string) {
	o.CaIds = &v
}

// GetCns returns the Cns field value if set, zero value otherwise.
func (o *FiltersApiAccessRule) GetCns() []string {
	if o == nil || o.Cns == nil {
		var ret []string
		return ret
	}
	return *o.Cns
}

// GetCnsOk returns a tuple with the Cns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiAccessRule) GetCnsOk() (*[]string, bool) {
	if o == nil || o.Cns == nil {
		return nil, false
	}
	return o.Cns, true
}

// HasCns returns a boolean if a field has been set.
func (o *FiltersApiAccessRule) HasCns() bool {
	if o != nil && o.Cns != nil {
		return true
	}

	return false
}

// SetCns gets a reference to the given []string and assigns it to the Cns field.
func (o *FiltersApiAccessRule) SetCns(v []string) {
	o.Cns = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *FiltersApiAccessRule) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiAccessRule) GetDescriptionsOk() (*[]string, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *FiltersApiAccessRule) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *FiltersApiAccessRule) SetDescriptions(v []string) {
	o.Descriptions = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *FiltersApiAccessRule) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiAccessRule) GetIpRangesOk() (*[]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *FiltersApiAccessRule) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *FiltersApiAccessRule) SetIpRanges(v []string) {
	o.IpRanges = &v
}

func (o FiltersApiAccessRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiAccessRuleIds != nil {
		toSerialize["ApiAccessRuleIds"] = o.ApiAccessRuleIds
	}
	if o.CaIds != nil {
		toSerialize["CaIds"] = o.CaIds
	}
	if o.Cns != nil {
		toSerialize["Cns"] = o.Cns
	}
	if o.Descriptions != nil {
		toSerialize["Descriptions"] = o.Descriptions
	}
	if o.IpRanges != nil {
		toSerialize["IpRanges"] = o.IpRanges
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersApiAccessRule struct {
	value *FiltersApiAccessRule
	isSet bool
}

func (v NullableFiltersApiAccessRule) Get() *FiltersApiAccessRule {
	return v.value
}

func (v *NullableFiltersApiAccessRule) Set(val *FiltersApiAccessRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersApiAccessRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersApiAccessRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersApiAccessRule(val *FiltersApiAccessRule) *NullableFiltersApiAccessRule {
	return &NullableFiltersApiAccessRule{value: val, isSet: true}
}

func (v NullableFiltersApiAccessRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersApiAccessRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
