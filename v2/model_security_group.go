/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// SecurityGroup Information about the security group.
type SecurityGroup struct {
	// The account ID of a user that has been granted permission.
	AccountId *string `json:"AccountId,omitempty"`
	// The description of the security group.
	Description *string `json:"Description,omitempty"`
	// The inbound rules associated with the security group.
	InboundRules *[]SecurityGroupRule `json:"InboundRules,omitempty"`
	// The ID of the Net for the security group.
	NetId *string `json:"NetId,omitempty"`
	// The outbound rules associated with the security group.
	OutboundRules *[]SecurityGroupRule `json:"OutboundRules,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty"`
	// The name of the security group.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty"`
	// One or more tags associated with the security group.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewSecurityGroup instantiates a new SecurityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroup() *SecurityGroup {
	this := SecurityGroup{}
	return &this
}

// NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupWithDefaults() *SecurityGroup {
	this := SecurityGroup{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SecurityGroup) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SecurityGroup) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SecurityGroup) SetAccountId(v string) {
	o.AccountId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityGroup) SetDescription(v string) {
	o.Description = &v
}

// GetInboundRules returns the InboundRules field value if set, zero value otherwise.
func (o *SecurityGroup) GetInboundRules() []SecurityGroupRule {
	if o == nil || o.InboundRules == nil {
		var ret []SecurityGroupRule
		return ret
	}
	return *o.InboundRules
}

// GetInboundRulesOk returns a tuple with the InboundRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetInboundRulesOk() (*[]SecurityGroupRule, bool) {
	if o == nil || o.InboundRules == nil {
		return nil, false
	}
	return o.InboundRules, true
}

// HasInboundRules returns a boolean if a field has been set.
func (o *SecurityGroup) HasInboundRules() bool {
	if o != nil && o.InboundRules != nil {
		return true
	}

	return false
}

// SetInboundRules gets a reference to the given []SecurityGroupRule and assigns it to the InboundRules field.
func (o *SecurityGroup) SetInboundRules(v []SecurityGroupRule) {
	o.InboundRules = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *SecurityGroup) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *SecurityGroup) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *SecurityGroup) SetNetId(v string) {
	o.NetId = &v
}

// GetOutboundRules returns the OutboundRules field value if set, zero value otherwise.
func (o *SecurityGroup) GetOutboundRules() []SecurityGroupRule {
	if o == nil || o.OutboundRules == nil {
		var ret []SecurityGroupRule
		return ret
	}
	return *o.OutboundRules
}

// GetOutboundRulesOk returns a tuple with the OutboundRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetOutboundRulesOk() (*[]SecurityGroupRule, bool) {
	if o == nil || o.OutboundRules == nil {
		return nil, false
	}
	return o.OutboundRules, true
}

// HasOutboundRules returns a boolean if a field has been set.
func (o *SecurityGroup) HasOutboundRules() bool {
	if o != nil && o.OutboundRules != nil {
		return true
	}

	return false
}

// SetOutboundRules gets a reference to the given []SecurityGroupRule and assigns it to the OutboundRules field.
func (o *SecurityGroup) SetOutboundRules(v []SecurityGroupRule) {
	o.OutboundRules = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *SecurityGroup) GetSecurityGroupId() string {
	if o == nil || o.SecurityGroupId == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || o.SecurityGroupId == nil {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *SecurityGroup) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId != nil {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *SecurityGroup) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetSecurityGroupName returns the SecurityGroupName field value if set, zero value otherwise.
func (o *SecurityGroup) GetSecurityGroupName() string {
	if o == nil || o.SecurityGroupName == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupName
}

// GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetSecurityGroupNameOk() (*string, bool) {
	if o == nil || o.SecurityGroupName == nil {
		return nil, false
	}
	return o.SecurityGroupName, true
}

// HasSecurityGroupName returns a boolean if a field has been set.
func (o *SecurityGroup) HasSecurityGroupName() bool {
	if o != nil && o.SecurityGroupName != nil {
		return true
	}

	return false
}

// SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.
func (o *SecurityGroup) SetSecurityGroupName(v string) {
	o.SecurityGroupName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityGroup) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityGroup) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *SecurityGroup) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o SecurityGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InboundRules != nil {
		toSerialize["InboundRules"] = o.InboundRules
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.OutboundRules != nil {
		toSerialize["OutboundRules"] = o.OutboundRules
	}
	if o.SecurityGroupId != nil {
		toSerialize["SecurityGroupId"] = o.SecurityGroupId
	}
	if o.SecurityGroupName != nil {
		toSerialize["SecurityGroupName"] = o.SecurityGroupName
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityGroup struct {
	value *SecurityGroup
	isSet bool
}

func (v NullableSecurityGroup) Get() *SecurityGroup {
	return v.value
}

func (v *NullableSecurityGroup) Set(val *SecurityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroup(val *SecurityGroup) *NullableSecurityGroup {
	return &NullableSecurityGroup{value: val, isSet: true}
}

func (v NullableSecurityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
