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

// SecurityGroupRule Information about the security group rule.
type SecurityGroupRule struct {
	// The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
	FromPortRange *int32 `json:"FromPortRange,omitempty"`
	// The IP protocol name (`tcp`, `udp`, `icmp`, or `-1` for all protocols). By default, `-1`. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	IpProtocol *string `json:"IpProtocol,omitempty"`
	// One or more IP ranges for the security group rules, in CIDR notation (for example, `10.0.0.0/16`).
	IpRanges *[]string `json:"IpRanges,omitempty"`
	// Information about one or more source or destination security groups.
	SecurityGroupsMembers *[]SecurityGroupsMember `json:"SecurityGroupsMembers,omitempty"`
	// One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices).
	ServiceIds *[]string `json:"ServiceIds,omitempty"`
	// The end of the port range for the TCP and UDP protocols, or an ICMP code number.
	ToPortRange *int32 `json:"ToPortRange,omitempty"`
}

// NewSecurityGroupRule instantiates a new SecurityGroupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupRule() *SecurityGroupRule {
	this := SecurityGroupRule{}
	return &this
}

// NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule {
	this := SecurityGroupRule{}
	return &this
}

// GetFromPortRange returns the FromPortRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetFromPortRange() int32 {
	if o == nil || o.FromPortRange == nil {
		var ret int32
		return ret
	}
	return *o.FromPortRange
}

// GetFromPortRangeOk returns a tuple with the FromPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetFromPortRangeOk() (*int32, bool) {
	if o == nil || o.FromPortRange == nil {
		return nil, false
	}
	return o.FromPortRange, true
}

// HasFromPortRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasFromPortRange() bool {
	if o != nil && o.FromPortRange != nil {
		return true
	}

	return false
}

// SetFromPortRange gets a reference to the given int32 and assigns it to the FromPortRange field.
func (o *SecurityGroupRule) SetFromPortRange(v int32) {
	o.FromPortRange = &v
}

// GetIpProtocol returns the IpProtocol field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIpProtocol() string {
	if o == nil || o.IpProtocol == nil {
		var ret string
		return ret
	}
	return *o.IpProtocol
}

// GetIpProtocolOk returns a tuple with the IpProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIpProtocolOk() (*string, bool) {
	if o == nil || o.IpProtocol == nil {
		return nil, false
	}
	return o.IpProtocol, true
}

// HasIpProtocol returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIpProtocol() bool {
	if o != nil && o.IpProtocol != nil {
		return true
	}

	return false
}

// SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.
func (o *SecurityGroupRule) SetIpProtocol(v string) {
	o.IpProtocol = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIpRangesOk() (*[]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *SecurityGroupRule) SetIpRanges(v []string) {
	o.IpRanges = &v
}

// GetSecurityGroupsMembers returns the SecurityGroupsMembers field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetSecurityGroupsMembers() []SecurityGroupsMember {
	if o == nil || o.SecurityGroupsMembers == nil {
		var ret []SecurityGroupsMember
		return ret
	}
	return *o.SecurityGroupsMembers
}

// GetSecurityGroupsMembersOk returns a tuple with the SecurityGroupsMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetSecurityGroupsMembersOk() (*[]SecurityGroupsMember, bool) {
	if o == nil || o.SecurityGroupsMembers == nil {
		return nil, false
	}
	return o.SecurityGroupsMembers, true
}

// HasSecurityGroupsMembers returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasSecurityGroupsMembers() bool {
	if o != nil && o.SecurityGroupsMembers != nil {
		return true
	}

	return false
}

// SetSecurityGroupsMembers gets a reference to the given []SecurityGroupsMember and assigns it to the SecurityGroupsMembers field.
func (o *SecurityGroupRule) SetSecurityGroupsMembers(v []SecurityGroupsMember) {
	o.SecurityGroupsMembers = &v
}

// GetServiceIds returns the ServiceIds field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetServiceIds() []string {
	if o == nil || o.ServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.ServiceIds
}

// GetServiceIdsOk returns a tuple with the ServiceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetServiceIdsOk() (*[]string, bool) {
	if o == nil || o.ServiceIds == nil {
		return nil, false
	}
	return o.ServiceIds, true
}

// HasServiceIds returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasServiceIds() bool {
	if o != nil && o.ServiceIds != nil {
		return true
	}

	return false
}

// SetServiceIds gets a reference to the given []string and assigns it to the ServiceIds field.
func (o *SecurityGroupRule) SetServiceIds(v []string) {
	o.ServiceIds = &v
}

// GetToPortRange returns the ToPortRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetToPortRange() int32 {
	if o == nil || o.ToPortRange == nil {
		var ret int32
		return ret
	}
	return *o.ToPortRange
}

// GetToPortRangeOk returns a tuple with the ToPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetToPortRangeOk() (*int32, bool) {
	if o == nil || o.ToPortRange == nil {
		return nil, false
	}
	return o.ToPortRange, true
}

// HasToPortRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasToPortRange() bool {
	if o != nil && o.ToPortRange != nil {
		return true
	}

	return false
}

// SetToPortRange gets a reference to the given int32 and assigns it to the ToPortRange field.
func (o *SecurityGroupRule) SetToPortRange(v int32) {
	o.ToPortRange = &v
}

func (o SecurityGroupRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromPortRange != nil {
		toSerialize["FromPortRange"] = o.FromPortRange
	}
	if o.IpProtocol != nil {
		toSerialize["IpProtocol"] = o.IpProtocol
	}
	if o.IpRanges != nil {
		toSerialize["IpRanges"] = o.IpRanges
	}
	if o.SecurityGroupsMembers != nil {
		toSerialize["SecurityGroupsMembers"] = o.SecurityGroupsMembers
	}
	if o.ServiceIds != nil {
		toSerialize["ServiceIds"] = o.ServiceIds
	}
	if o.ToPortRange != nil {
		toSerialize["ToPortRange"] = o.ToPortRange
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityGroupRule struct {
	value *SecurityGroupRule
	isSet bool
}

func (v NullableSecurityGroupRule) Get() *SecurityGroupRule {
	return v.value
}

func (v *NullableSecurityGroupRule) Set(val *SecurityGroupRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupRule(val *SecurityGroupRule) *NullableSecurityGroupRule {
	return &NullableSecurityGroupRule{value: val, isSet: true}
}

func (v NullableSecurityGroupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
