/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.5
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersVmGroup One or more filters.
type FiltersVmGroup struct {
	// The descriptions of the VM groups.
	Descriptions *[]string `json:"Descriptions,omitempty"`
	// The IDs of the security groups.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The IDs of the Subnets.
	SubnetIds *[]string `json:"SubnetIds,omitempty"`
	// The keys of the tags associated with the VM groups.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the VM groups.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the VMs, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
	// The number of VMs in the VM group.
	VmCounts *[]int32 `json:"VmCounts,omitempty"`
	// The IDs of the VM groups.
	VmGroupIds *[]string `json:"VmGroupIds,omitempty"`
	// The names of the VM groups.
	VmGroupNames *[]string `json:"VmGroupNames,omitempty"`
	// The IDs of the VM templates.
	VmTemplateIds *[]string `json:"VmTemplateIds,omitempty"`
}

// NewFiltersVmGroup instantiates a new FiltersVmGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersVmGroup() *FiltersVmGroup {
	this := FiltersVmGroup{}
	return &this
}

// NewFiltersVmGroupWithDefaults instantiates a new FiltersVmGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersVmGroupWithDefaults() *FiltersVmGroup {
	this := FiltersVmGroup{}
	return &this
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetDescriptionsOk() (*[]string, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *FiltersVmGroup) SetDescriptions(v []string) {
	o.Descriptions = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *FiltersVmGroup) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSubnetIds returns the SubnetIds field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetSubnetIds() []string {
	if o == nil || o.SubnetIds == nil {
		var ret []string
		return ret
	}
	return *o.SubnetIds
}

// GetSubnetIdsOk returns a tuple with the SubnetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetSubnetIdsOk() (*[]string, bool) {
	if o == nil || o.SubnetIds == nil {
		return nil, false
	}
	return o.SubnetIds, true
}

// HasSubnetIds returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasSubnetIds() bool {
	if o != nil && o.SubnetIds != nil {
		return true
	}

	return false
}

// SetSubnetIds gets a reference to the given []string and assigns it to the SubnetIds field.
func (o *FiltersVmGroup) SetSubnetIds(v []string) {
	o.SubnetIds = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersVmGroup) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersVmGroup) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersVmGroup) SetTags(v []string) {
	o.Tags = &v
}

// GetVmCounts returns the VmCounts field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetVmCounts() []int32 {
	if o == nil || o.VmCounts == nil {
		var ret []int32
		return ret
	}
	return *o.VmCounts
}

// GetVmCountsOk returns a tuple with the VmCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetVmCountsOk() (*[]int32, bool) {
	if o == nil || o.VmCounts == nil {
		return nil, false
	}
	return o.VmCounts, true
}

// HasVmCounts returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasVmCounts() bool {
	if o != nil && o.VmCounts != nil {
		return true
	}

	return false
}

// SetVmCounts gets a reference to the given []int32 and assigns it to the VmCounts field.
func (o *FiltersVmGroup) SetVmCounts(v []int32) {
	o.VmCounts = &v
}

// GetVmGroupIds returns the VmGroupIds field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetVmGroupIds() []string {
	if o == nil || o.VmGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.VmGroupIds
}

// GetVmGroupIdsOk returns a tuple with the VmGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetVmGroupIdsOk() (*[]string, bool) {
	if o == nil || o.VmGroupIds == nil {
		return nil, false
	}
	return o.VmGroupIds, true
}

// HasVmGroupIds returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasVmGroupIds() bool {
	if o != nil && o.VmGroupIds != nil {
		return true
	}

	return false
}

// SetVmGroupIds gets a reference to the given []string and assigns it to the VmGroupIds field.
func (o *FiltersVmGroup) SetVmGroupIds(v []string) {
	o.VmGroupIds = &v
}

// GetVmGroupNames returns the VmGroupNames field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetVmGroupNames() []string {
	if o == nil || o.VmGroupNames == nil {
		var ret []string
		return ret
	}
	return *o.VmGroupNames
}

// GetVmGroupNamesOk returns a tuple with the VmGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetVmGroupNamesOk() (*[]string, bool) {
	if o == nil || o.VmGroupNames == nil {
		return nil, false
	}
	return o.VmGroupNames, true
}

// HasVmGroupNames returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasVmGroupNames() bool {
	if o != nil && o.VmGroupNames != nil {
		return true
	}

	return false
}

// SetVmGroupNames gets a reference to the given []string and assigns it to the VmGroupNames field.
func (o *FiltersVmGroup) SetVmGroupNames(v []string) {
	o.VmGroupNames = &v
}

// GetVmTemplateIds returns the VmTemplateIds field value if set, zero value otherwise.
func (o *FiltersVmGroup) GetVmTemplateIds() []string {
	if o == nil || o.VmTemplateIds == nil {
		var ret []string
		return ret
	}
	return *o.VmTemplateIds
}

// GetVmTemplateIdsOk returns a tuple with the VmTemplateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVmGroup) GetVmTemplateIdsOk() (*[]string, bool) {
	if o == nil || o.VmTemplateIds == nil {
		return nil, false
	}
	return o.VmTemplateIds, true
}

// HasVmTemplateIds returns a boolean if a field has been set.
func (o *FiltersVmGroup) HasVmTemplateIds() bool {
	if o != nil && o.VmTemplateIds != nil {
		return true
	}

	return false
}

// SetVmTemplateIds gets a reference to the given []string and assigns it to the VmTemplateIds field.
func (o *FiltersVmGroup) SetVmTemplateIds(v []string) {
	o.VmTemplateIds = &v
}

func (o FiltersVmGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Descriptions != nil {
		toSerialize["Descriptions"] = o.Descriptions
	}
	if o.SecurityGroupIds != nil {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	if o.SubnetIds != nil {
		toSerialize["SubnetIds"] = o.SubnetIds
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
	if o.VmCounts != nil {
		toSerialize["VmCounts"] = o.VmCounts
	}
	if o.VmGroupIds != nil {
		toSerialize["VmGroupIds"] = o.VmGroupIds
	}
	if o.VmGroupNames != nil {
		toSerialize["VmGroupNames"] = o.VmGroupNames
	}
	if o.VmTemplateIds != nil {
		toSerialize["VmTemplateIds"] = o.VmTemplateIds
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersVmGroup struct {
	value *FiltersVmGroup
	isSet bool
}

func (v NullableFiltersVmGroup) Get() *FiltersVmGroup {
	return v.value
}

func (v *NullableFiltersVmGroup) Set(val *FiltersVmGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersVmGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersVmGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersVmGroup(val *FiltersVmGroup) *NullableFiltersVmGroup {
	return &NullableFiltersVmGroup{value: val, isSet: true}
}

func (v NullableFiltersVmGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersVmGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
