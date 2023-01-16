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

// FiltersDhcpOptions One or more filters.
type FiltersDhcpOptions struct {
	// If true, lists all default DHCP options set. If false, lists all non-default DHCP options set.
	Default *bool `json:"Default,omitempty"`
	// The IDs of the DHCP options sets.
	DhcpOptionsSetIds *[]string `json:"DhcpOptionsSetIds,omitempty"`
	// The IPs of the domain name servers used for the DHCP options sets.
	DomainNameServers *[]string `json:"DomainNameServers,omitempty"`
	// The domain names used for the DHCP options sets.
	DomainNames *[]string `json:"DomainNames,omitempty"`
	// The IPs of the log servers used for the DHCP options sets.
	LogServers *[]string `json:"LogServers,omitempty"`
	// The IPs of the Network Time Protocol (NTP) servers used for the DHCP options sets.
	NtpServers *[]string `json:"NtpServers,omitempty"`
	// The keys of the tags associated with the DHCP options sets.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the DHCP options sets.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the DHCP options sets, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersDhcpOptions instantiates a new FiltersDhcpOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersDhcpOptions() *FiltersDhcpOptions {
	this := FiltersDhcpOptions{}
	return &this
}

// NewFiltersDhcpOptionsWithDefaults instantiates a new FiltersDhcpOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersDhcpOptionsWithDefaults() *FiltersDhcpOptions {
	this := FiltersDhcpOptions{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *FiltersDhcpOptions) SetDefault(v bool) {
	o.Default = &v
}

// GetDhcpOptionsSetIds returns the DhcpOptionsSetIds field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetDhcpOptionsSetIds() []string {
	if o == nil || o.DhcpOptionsSetIds == nil {
		var ret []string
		return ret
	}
	return *o.DhcpOptionsSetIds
}

// GetDhcpOptionsSetIdsOk returns a tuple with the DhcpOptionsSetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetDhcpOptionsSetIdsOk() (*[]string, bool) {
	if o == nil || o.DhcpOptionsSetIds == nil {
		return nil, false
	}
	return o.DhcpOptionsSetIds, true
}

// HasDhcpOptionsSetIds returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasDhcpOptionsSetIds() bool {
	if o != nil && o.DhcpOptionsSetIds != nil {
		return true
	}

	return false
}

// SetDhcpOptionsSetIds gets a reference to the given []string and assigns it to the DhcpOptionsSetIds field.
func (o *FiltersDhcpOptions) SetDhcpOptionsSetIds(v []string) {
	o.DhcpOptionsSetIds = &v
}

// GetDomainNameServers returns the DomainNameServers field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetDomainNameServers() []string {
	if o == nil || o.DomainNameServers == nil {
		var ret []string
		return ret
	}
	return *o.DomainNameServers
}

// GetDomainNameServersOk returns a tuple with the DomainNameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetDomainNameServersOk() (*[]string, bool) {
	if o == nil || o.DomainNameServers == nil {
		return nil, false
	}
	return o.DomainNameServers, true
}

// HasDomainNameServers returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasDomainNameServers() bool {
	if o != nil && o.DomainNameServers != nil {
		return true
	}

	return false
}

// SetDomainNameServers gets a reference to the given []string and assigns it to the DomainNameServers field.
func (o *FiltersDhcpOptions) SetDomainNameServers(v []string) {
	o.DomainNameServers = &v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetDomainNames() []string {
	if o == nil || o.DomainNames == nil {
		var ret []string
		return ret
	}
	return *o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetDomainNamesOk() (*[]string, bool) {
	if o == nil || o.DomainNames == nil {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasDomainNames() bool {
	if o != nil && o.DomainNames != nil {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *FiltersDhcpOptions) SetDomainNames(v []string) {
	o.DomainNames = &v
}

// GetLogServers returns the LogServers field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetLogServers() []string {
	if o == nil || o.LogServers == nil {
		var ret []string
		return ret
	}
	return *o.LogServers
}

// GetLogServersOk returns a tuple with the LogServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetLogServersOk() (*[]string, bool) {
	if o == nil || o.LogServers == nil {
		return nil, false
	}
	return o.LogServers, true
}

// HasLogServers returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasLogServers() bool {
	if o != nil && o.LogServers != nil {
		return true
	}

	return false
}

// SetLogServers gets a reference to the given []string and assigns it to the LogServers field.
func (o *FiltersDhcpOptions) SetLogServers(v []string) {
	o.LogServers = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetNtpServers() []string {
	if o == nil || o.NtpServers == nil {
		var ret []string
		return ret
	}
	return *o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *FiltersDhcpOptions) SetNtpServers(v []string) {
	o.NtpServers = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersDhcpOptions) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersDhcpOptions) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersDhcpOptions) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDhcpOptions) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersDhcpOptions) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersDhcpOptions) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersDhcpOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil {
		toSerialize["Default"] = o.Default
	}
	if o.DhcpOptionsSetIds != nil {
		toSerialize["DhcpOptionsSetIds"] = o.DhcpOptionsSetIds
	}
	if o.DomainNameServers != nil {
		toSerialize["DomainNameServers"] = o.DomainNameServers
	}
	if o.DomainNames != nil {
		toSerialize["DomainNames"] = o.DomainNames
	}
	if o.LogServers != nil {
		toSerialize["LogServers"] = o.LogServers
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
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

type NullableFiltersDhcpOptions struct {
	value *FiltersDhcpOptions
	isSet bool
}

func (v NullableFiltersDhcpOptions) Get() *FiltersDhcpOptions {
	return v.value
}

func (v *NullableFiltersDhcpOptions) Set(val *FiltersDhcpOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersDhcpOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersDhcpOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersDhcpOptions(val *FiltersDhcpOptions) *NullableFiltersDhcpOptions {
	return &NullableFiltersDhcpOptions{value: val, isSet: true}
}

func (v NullableFiltersDhcpOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersDhcpOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
