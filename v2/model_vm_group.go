/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> Throttling: To protect against overloads, the number of identical requests allowed in a given time period is limited.<br /> Brute force: To protect against brute force attacks, the number of failed authentication attempts in a given time period is limited.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).<br /> # Authentication Schemes ### Access Key/Secret Key The main way to authenticate your requests to the OUTSCALE API is to use an access key and a secret key.<br /> The mechanism behind this is based on AWS Signature Version 4, whose technical implementation details are described in [Signature of API Requests](https://docs.outscale.com/en/userguide/Signature-of-API-Requests.html).<br /><br /> In practice, the way to specify your access key and secret key depends on the tool or SDK you want to use to interact with the API.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify your access key, secret key, and the Region of your account. > 2. You then specify the `--profile` option when executing OSC CLI commands. >  > For more information, see [Installing and Configuring OSC CLI](https://docs.outscale.com/en/userguide/Installing-and-Configuring-OSC-CLI.html).  See the code samples in each section of this documentation for specific examples in different programming languages.<br /> For more information about access keys, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html). ### Login/Password For certain API actions, you can also use basic authentication with the login (email address) and password of your TINA account.<br /> This is useful only in special circumstances, for example if you do not know your access key/secret key and want to retrieve them programmatically.<br /> In most cases, however, you can use the Cockpit web interface to retrieve them.<br />  > For example, if you use OSC CLI: > 1. You need to create an **~/.osc/config.json** file to specify the Region of your account, but you leave the access key value and secret key value empty (`&quot;&quot;`). > 2. You then specify the `--profile`, `--authentication-method`, `--login`, and `--password` options when executing OSC CLI commands.  See the code samples in each section of this documentation for specific examples in different programming languages. ### No Authentication A few API actions do not require any authentication. They are indicated as such in this documentation.<br /> ### Other Security Mechanisms In parallel with the authentication schemes, you can add other security mechanisms to your OUTSCALE account, for example to restrict API requests by IP or other criteria.<br /> For more information, see [Managing Your API Accesses](https://docs.outscale.com/en/userguide/Managing-Your-API-Accesses.html).
 *
 * API version: 1.28.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// VmGroup Information about the VM group.
type VmGroup struct {
	// The date and time of creation of the VM group.
	CreationDate *string `json:"CreationDate,omitempty"`
	// The description of the VM group.
	Description *string `json:"Description,omitempty"`
	// The positioning strategy of the VMs on hypervisors. By default, or if set to `no-strategy`, TINA determines the most adequate position for the VMs. If set to `attract`, the VMs are deployed on the same hypervisor, which improves network performance. If set to `repulse`, the VMs are deployed on a different hypervisor, which improves fault tolerance.
	PositioningStrategy *string `json:"PositioningStrategy,omitempty"`
	// One or more IDs of security groups for the VM group.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The state of the VM group (`pending` \\| `available` \\| `scaling up` \\| `scaling down` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
	// The ID of the Subnet for the VM group.
	SubnetId *string `json:"SubnetId,omitempty"`
	// One or more tags associated with the VM group.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The number of VMs in the VM group.
	VmCount *int32 `json:"VmCount,omitempty"`
	// The ID of the VM group.
	VmGroupId *string `json:"VmGroupId,omitempty"`
	// The name of the VM group.
	VmGroupName *string `json:"VmGroupName,omitempty"`
	// The IDs of the VMs in the VM group.
	VmIds *[]string `json:"VmIds,omitempty"`
	// The ID of the VM template used by the VM group.
	VmTemplateId *string `json:"VmTemplateId,omitempty"`
}

// NewVmGroup instantiates a new VmGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmGroup() *VmGroup {
	this := VmGroup{}
	return &this
}

// NewVmGroupWithDefaults instantiates a new VmGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmGroupWithDefaults() *VmGroup {
	this := VmGroup{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *VmGroup) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *VmGroup) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *VmGroup) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VmGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VmGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VmGroup) SetDescription(v string) {
	o.Description = &v
}

// GetPositioningStrategy returns the PositioningStrategy field value if set, zero value otherwise.
func (o *VmGroup) GetPositioningStrategy() string {
	if o == nil || o.PositioningStrategy == nil {
		var ret string
		return ret
	}
	return *o.PositioningStrategy
}

// GetPositioningStrategyOk returns a tuple with the PositioningStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetPositioningStrategyOk() (*string, bool) {
	if o == nil || o.PositioningStrategy == nil {
		return nil, false
	}
	return o.PositioningStrategy, true
}

// HasPositioningStrategy returns a boolean if a field has been set.
func (o *VmGroup) HasPositioningStrategy() bool {
	if o != nil && o.PositioningStrategy != nil {
		return true
	}

	return false
}

// SetPositioningStrategy gets a reference to the given string and assigns it to the PositioningStrategy field.
func (o *VmGroup) SetPositioningStrategy(v string) {
	o.PositioningStrategy = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *VmGroup) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *VmGroup) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *VmGroup) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VmGroup) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VmGroup) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VmGroup) SetState(v string) {
	o.State = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *VmGroup) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetSubnetIdOk() (*string, bool) {
	if o == nil || o.SubnetId == nil {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *VmGroup) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *VmGroup) SetSubnetId(v string) {
	o.SubnetId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VmGroup) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VmGroup) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *VmGroup) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *VmGroup) GetVmCount() int32 {
	if o == nil || o.VmCount == nil {
		var ret int32
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetVmCountOk() (*int32, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *VmGroup) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int32 and assigns it to the VmCount field.
func (o *VmGroup) SetVmCount(v int32) {
	o.VmCount = &v
}

// GetVmGroupId returns the VmGroupId field value if set, zero value otherwise.
func (o *VmGroup) GetVmGroupId() string {
	if o == nil || o.VmGroupId == nil {
		var ret string
		return ret
	}
	return *o.VmGroupId
}

// GetVmGroupIdOk returns a tuple with the VmGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetVmGroupIdOk() (*string, bool) {
	if o == nil || o.VmGroupId == nil {
		return nil, false
	}
	return o.VmGroupId, true
}

// HasVmGroupId returns a boolean if a field has been set.
func (o *VmGroup) HasVmGroupId() bool {
	if o != nil && o.VmGroupId != nil {
		return true
	}

	return false
}

// SetVmGroupId gets a reference to the given string and assigns it to the VmGroupId field.
func (o *VmGroup) SetVmGroupId(v string) {
	o.VmGroupId = &v
}

// GetVmGroupName returns the VmGroupName field value if set, zero value otherwise.
func (o *VmGroup) GetVmGroupName() string {
	if o == nil || o.VmGroupName == nil {
		var ret string
		return ret
	}
	return *o.VmGroupName
}

// GetVmGroupNameOk returns a tuple with the VmGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetVmGroupNameOk() (*string, bool) {
	if o == nil || o.VmGroupName == nil {
		return nil, false
	}
	return o.VmGroupName, true
}

// HasVmGroupName returns a boolean if a field has been set.
func (o *VmGroup) HasVmGroupName() bool {
	if o != nil && o.VmGroupName != nil {
		return true
	}

	return false
}

// SetVmGroupName gets a reference to the given string and assigns it to the VmGroupName field.
func (o *VmGroup) SetVmGroupName(v string) {
	o.VmGroupName = &v
}

// GetVmIds returns the VmIds field value if set, zero value otherwise.
func (o *VmGroup) GetVmIds() []string {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret
	}
	return *o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetVmIdsOk() (*[]string, bool) {
	if o == nil || o.VmIds == nil {
		return nil, false
	}
	return o.VmIds, true
}

// HasVmIds returns a boolean if a field has been set.
func (o *VmGroup) HasVmIds() bool {
	if o != nil && o.VmIds != nil {
		return true
	}

	return false
}

// SetVmIds gets a reference to the given []string and assigns it to the VmIds field.
func (o *VmGroup) SetVmIds(v []string) {
	o.VmIds = &v
}

// GetVmTemplateId returns the VmTemplateId field value if set, zero value otherwise.
func (o *VmGroup) GetVmTemplateId() string {
	if o == nil || o.VmTemplateId == nil {
		var ret string
		return ret
	}
	return *o.VmTemplateId
}

// GetVmTemplateIdOk returns a tuple with the VmTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmGroup) GetVmTemplateIdOk() (*string, bool) {
	if o == nil || o.VmTemplateId == nil {
		return nil, false
	}
	return o.VmTemplateId, true
}

// HasVmTemplateId returns a boolean if a field has been set.
func (o *VmGroup) HasVmTemplateId() bool {
	if o != nil && o.VmTemplateId != nil {
		return true
	}

	return false
}

// SetVmTemplateId gets a reference to the given string and assigns it to the VmTemplateId field.
func (o *VmGroup) SetVmTemplateId(v string) {
	o.VmTemplateId = &v
}

func (o VmGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationDate != nil {
		toSerialize["CreationDate"] = o.CreationDate
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.PositioningStrategy != nil {
		toSerialize["PositioningStrategy"] = o.PositioningStrategy
	}
	if o.SecurityGroupIds != nil {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SubnetId != nil {
		toSerialize["SubnetId"] = o.SubnetId
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	if o.VmGroupId != nil {
		toSerialize["VmGroupId"] = o.VmGroupId
	}
	if o.VmGroupName != nil {
		toSerialize["VmGroupName"] = o.VmGroupName
	}
	if o.VmIds != nil {
		toSerialize["VmIds"] = o.VmIds
	}
	if o.VmTemplateId != nil {
		toSerialize["VmTemplateId"] = o.VmTemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableVmGroup struct {
	value *VmGroup
	isSet bool
}

func (v NullableVmGroup) Get() *VmGroup {
	return v.value
}

func (v *NullableVmGroup) Set(val *VmGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableVmGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableVmGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmGroup(val *VmGroup) *NullableVmGroup {
	return &NullableVmGroup{value: val, isSet: true}
}

func (v NullableVmGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
