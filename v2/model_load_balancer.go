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

// LoadBalancer Information about the load balancer.
type LoadBalancer struct {
	AccessLog *AccessLog `json:"AccessLog,omitempty"`
	// The stickiness policies defined for the load balancer.
	ApplicationStickyCookiePolicies *[]ApplicationStickyCookiePolicy `json:"ApplicationStickyCookiePolicies,omitempty"`
	// One or more public IPs of back-end VMs.
	BackendIps *[]string `json:"BackendIps,omitempty"`
	// One or more IDs of back-end VMs for the load balancer.
	BackendVmIds *[]string `json:"BackendVmIds,omitempty"`
	// The DNS name of the load balancer.
	DnsName     *string      `json:"DnsName,omitempty"`
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty"`
	// The listeners for the load balancer.
	Listeners *[]Listener `json:"Listeners,omitempty"`
	// The name of the load balancer.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty"`
	// The policies defined for the load balancer.
	LoadBalancerStickyCookiePolicies *[]LoadBalancerStickyCookiePolicy `json:"LoadBalancerStickyCookiePolicies,omitempty"`
	// The type of load balancer. Valid only for load balancers in a Net.<br /> If `LoadBalancerType` is `internet-facing`, the load balancer has a public DNS name that resolves to a public IP.<br /> If `LoadBalancerType` is `internal`, the load balancer has a public DNS name that resolves to a private IP.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty"`
	// The ID of the Net for the load balancer.
	NetId *string `json:"NetId,omitempty"`
	// (internet-facing only) The public IP associated with the load balancer.
	PublicIp *string `json:"PublicIp,omitempty"`
	// Whether secure cookies are enabled for the load balancer.
	SecuredCookies *bool `json:"SecuredCookies,omitempty"`
	// One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.
	SecurityGroups      *[]string            `json:"SecurityGroups,omitempty"`
	SourceSecurityGroup *SourceSecurityGroup `json:"SourceSecurityGroup,omitempty"`
	// The ID of the Subnet in which the load balancer was created.
	Subnets *[]string `json:"Subnets,omitempty"`
	// The ID of the Subregion in which the load balancer was created.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
	// One or more tags associated with the load balancer.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewLoadBalancer instantiates a new LoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancer() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerWithDefaults() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// GetAccessLog returns the AccessLog field value if set, zero value otherwise.
func (o *LoadBalancer) GetAccessLog() AccessLog {
	if o == nil || o.AccessLog == nil {
		var ret AccessLog
		return ret
	}
	return *o.AccessLog
}

// GetAccessLogOk returns a tuple with the AccessLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetAccessLogOk() (*AccessLog, bool) {
	if o == nil || o.AccessLog == nil {
		return nil, false
	}
	return o.AccessLog, true
}

// HasAccessLog returns a boolean if a field has been set.
func (o *LoadBalancer) HasAccessLog() bool {
	if o != nil && o.AccessLog != nil {
		return true
	}

	return false
}

// SetAccessLog gets a reference to the given AccessLog and assigns it to the AccessLog field.
func (o *LoadBalancer) SetAccessLog(v AccessLog) {
	o.AccessLog = &v
}

// GetApplicationStickyCookiePolicies returns the ApplicationStickyCookiePolicies field value if set, zero value otherwise.
func (o *LoadBalancer) GetApplicationStickyCookiePolicies() []ApplicationStickyCookiePolicy {
	if o == nil || o.ApplicationStickyCookiePolicies == nil {
		var ret []ApplicationStickyCookiePolicy
		return ret
	}
	return *o.ApplicationStickyCookiePolicies
}

// GetApplicationStickyCookiePoliciesOk returns a tuple with the ApplicationStickyCookiePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetApplicationStickyCookiePoliciesOk() (*[]ApplicationStickyCookiePolicy, bool) {
	if o == nil || o.ApplicationStickyCookiePolicies == nil {
		return nil, false
	}
	return o.ApplicationStickyCookiePolicies, true
}

// HasApplicationStickyCookiePolicies returns a boolean if a field has been set.
func (o *LoadBalancer) HasApplicationStickyCookiePolicies() bool {
	if o != nil && o.ApplicationStickyCookiePolicies != nil {
		return true
	}

	return false
}

// SetApplicationStickyCookiePolicies gets a reference to the given []ApplicationStickyCookiePolicy and assigns it to the ApplicationStickyCookiePolicies field.
func (o *LoadBalancer) SetApplicationStickyCookiePolicies(v []ApplicationStickyCookiePolicy) {
	o.ApplicationStickyCookiePolicies = &v
}

// GetBackendIps returns the BackendIps field value if set, zero value otherwise.
func (o *LoadBalancer) GetBackendIps() []string {
	if o == nil || o.BackendIps == nil {
		var ret []string
		return ret
	}
	return *o.BackendIps
}

// GetBackendIpsOk returns a tuple with the BackendIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetBackendIpsOk() (*[]string, bool) {
	if o == nil || o.BackendIps == nil {
		return nil, false
	}
	return o.BackendIps, true
}

// HasBackendIps returns a boolean if a field has been set.
func (o *LoadBalancer) HasBackendIps() bool {
	if o != nil && o.BackendIps != nil {
		return true
	}

	return false
}

// SetBackendIps gets a reference to the given []string and assigns it to the BackendIps field.
func (o *LoadBalancer) SetBackendIps(v []string) {
	o.BackendIps = &v
}

// GetBackendVmIds returns the BackendVmIds field value if set, zero value otherwise.
func (o *LoadBalancer) GetBackendVmIds() []string {
	if o == nil || o.BackendVmIds == nil {
		var ret []string
		return ret
	}
	return *o.BackendVmIds
}

// GetBackendVmIdsOk returns a tuple with the BackendVmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetBackendVmIdsOk() (*[]string, bool) {
	if o == nil || o.BackendVmIds == nil {
		return nil, false
	}
	return o.BackendVmIds, true
}

// HasBackendVmIds returns a boolean if a field has been set.
func (o *LoadBalancer) HasBackendVmIds() bool {
	if o != nil && o.BackendVmIds != nil {
		return true
	}

	return false
}

// SetBackendVmIds gets a reference to the given []string and assigns it to the BackendVmIds field.
func (o *LoadBalancer) SetBackendVmIds(v []string) {
	o.BackendVmIds = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *LoadBalancer) GetDnsName() string {
	if o == nil || o.DnsName == nil {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetDnsNameOk() (*string, bool) {
	if o == nil || o.DnsName == nil {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *LoadBalancer) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *LoadBalancer) SetDnsName(v string) {
	o.DnsName = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *LoadBalancer) GetHealthCheck() HealthCheck {
	if o == nil || o.HealthCheck == nil {
		var ret HealthCheck
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetHealthCheckOk() (*HealthCheck, bool) {
	if o == nil || o.HealthCheck == nil {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *LoadBalancer) HasHealthCheck() bool {
	if o != nil && o.HealthCheck != nil {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given HealthCheck and assigns it to the HealthCheck field.
func (o *LoadBalancer) SetHealthCheck(v HealthCheck) {
	o.HealthCheck = &v
}

// GetListeners returns the Listeners field value if set, zero value otherwise.
func (o *LoadBalancer) GetListeners() []Listener {
	if o == nil || o.Listeners == nil {
		var ret []Listener
		return ret
	}
	return *o.Listeners
}

// GetListenersOk returns a tuple with the Listeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetListenersOk() (*[]Listener, bool) {
	if o == nil || o.Listeners == nil {
		return nil, false
	}
	return o.Listeners, true
}

// HasListeners returns a boolean if a field has been set.
func (o *LoadBalancer) HasListeners() bool {
	if o != nil && o.Listeners != nil {
		return true
	}

	return false
}

// SetListeners gets a reference to the given []Listener and assigns it to the Listeners field.
func (o *LoadBalancer) SetListeners(v []Listener) {
	o.Listeners = &v
}

// GetLoadBalancerName returns the LoadBalancerName field value if set, zero value otherwise.
func (o *LoadBalancer) GetLoadBalancerName() string {
	if o == nil || o.LoadBalancerName == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil || o.LoadBalancerName == nil {
		return nil, false
	}
	return o.LoadBalancerName, true
}

// HasLoadBalancerName returns a boolean if a field has been set.
func (o *LoadBalancer) HasLoadBalancerName() bool {
	if o != nil && o.LoadBalancerName != nil {
		return true
	}

	return false
}

// SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.
func (o *LoadBalancer) SetLoadBalancerName(v string) {
	o.LoadBalancerName = &v
}

// GetLoadBalancerStickyCookiePolicies returns the LoadBalancerStickyCookiePolicies field value if set, zero value otherwise.
func (o *LoadBalancer) GetLoadBalancerStickyCookiePolicies() []LoadBalancerStickyCookiePolicy {
	if o == nil || o.LoadBalancerStickyCookiePolicies == nil {
		var ret []LoadBalancerStickyCookiePolicy
		return ret
	}
	return *o.LoadBalancerStickyCookiePolicies
}

// GetLoadBalancerStickyCookiePoliciesOk returns a tuple with the LoadBalancerStickyCookiePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetLoadBalancerStickyCookiePoliciesOk() (*[]LoadBalancerStickyCookiePolicy, bool) {
	if o == nil || o.LoadBalancerStickyCookiePolicies == nil {
		return nil, false
	}
	return o.LoadBalancerStickyCookiePolicies, true
}

// HasLoadBalancerStickyCookiePolicies returns a boolean if a field has been set.
func (o *LoadBalancer) HasLoadBalancerStickyCookiePolicies() bool {
	if o != nil && o.LoadBalancerStickyCookiePolicies != nil {
		return true
	}

	return false
}

// SetLoadBalancerStickyCookiePolicies gets a reference to the given []LoadBalancerStickyCookiePolicy and assigns it to the LoadBalancerStickyCookiePolicies field.
func (o *LoadBalancer) SetLoadBalancerStickyCookiePolicies(v []LoadBalancerStickyCookiePolicy) {
	o.LoadBalancerStickyCookiePolicies = &v
}

// GetLoadBalancerType returns the LoadBalancerType field value if set, zero value otherwise.
func (o *LoadBalancer) GetLoadBalancerType() string {
	if o == nil || o.LoadBalancerType == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerType
}

// GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetLoadBalancerTypeOk() (*string, bool) {
	if o == nil || o.LoadBalancerType == nil {
		return nil, false
	}
	return o.LoadBalancerType, true
}

// HasLoadBalancerType returns a boolean if a field has been set.
func (o *LoadBalancer) HasLoadBalancerType() bool {
	if o != nil && o.LoadBalancerType != nil {
		return true
	}

	return false
}

// SetLoadBalancerType gets a reference to the given string and assigns it to the LoadBalancerType field.
func (o *LoadBalancer) SetLoadBalancerType(v string) {
	o.LoadBalancerType = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *LoadBalancer) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *LoadBalancer) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *LoadBalancer) SetNetId(v string) {
	o.NetId = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *LoadBalancer) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *LoadBalancer) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *LoadBalancer) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetSecuredCookies returns the SecuredCookies field value if set, zero value otherwise.
func (o *LoadBalancer) GetSecuredCookies() bool {
	if o == nil || o.SecuredCookies == nil {
		var ret bool
		return ret
	}
	return *o.SecuredCookies
}

// GetSecuredCookiesOk returns a tuple with the SecuredCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetSecuredCookiesOk() (*bool, bool) {
	if o == nil || o.SecuredCookies == nil {
		return nil, false
	}
	return o.SecuredCookies, true
}

// HasSecuredCookies returns a boolean if a field has been set.
func (o *LoadBalancer) HasSecuredCookies() bool {
	if o != nil && o.SecuredCookies != nil {
		return true
	}

	return false
}

// SetSecuredCookies gets a reference to the given bool and assigns it to the SecuredCookies field.
func (o *LoadBalancer) SetSecuredCookies(v bool) {
	o.SecuredCookies = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *LoadBalancer) GetSecurityGroups() []string {
	if o == nil || o.SecurityGroups == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *LoadBalancer) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *LoadBalancer) SetSecurityGroups(v []string) {
	o.SecurityGroups = &v
}

// GetSourceSecurityGroup returns the SourceSecurityGroup field value if set, zero value otherwise.
func (o *LoadBalancer) GetSourceSecurityGroup() SourceSecurityGroup {
	if o == nil || o.SourceSecurityGroup == nil {
		var ret SourceSecurityGroup
		return ret
	}
	return *o.SourceSecurityGroup
}

// GetSourceSecurityGroupOk returns a tuple with the SourceSecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetSourceSecurityGroupOk() (*SourceSecurityGroup, bool) {
	if o == nil || o.SourceSecurityGroup == nil {
		return nil, false
	}
	return o.SourceSecurityGroup, true
}

// HasSourceSecurityGroup returns a boolean if a field has been set.
func (o *LoadBalancer) HasSourceSecurityGroup() bool {
	if o != nil && o.SourceSecurityGroup != nil {
		return true
	}

	return false
}

// SetSourceSecurityGroup gets a reference to the given SourceSecurityGroup and assigns it to the SourceSecurityGroup field.
func (o *LoadBalancer) SetSourceSecurityGroup(v SourceSecurityGroup) {
	o.SourceSecurityGroup = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *LoadBalancer) GetSubnets() []string {
	if o == nil || o.Subnets == nil {
		var ret []string
		return ret
	}
	return *o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetSubnetsOk() (*[]string, bool) {
	if o == nil || o.Subnets == nil {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *LoadBalancer) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *LoadBalancer) SetSubnets(v []string) {
	o.Subnets = &v
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *LoadBalancer) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetSubregionNamesOk() (*[]string, bool) {
	if o == nil || o.SubregionNames == nil {
		return nil, false
	}
	return o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *LoadBalancer) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *LoadBalancer) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LoadBalancer) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LoadBalancer) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *LoadBalancer) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessLog != nil {
		toSerialize["AccessLog"] = o.AccessLog
	}
	if o.ApplicationStickyCookiePolicies != nil {
		toSerialize["ApplicationStickyCookiePolicies"] = o.ApplicationStickyCookiePolicies
	}
	if o.BackendIps != nil {
		toSerialize["BackendIps"] = o.BackendIps
	}
	if o.BackendVmIds != nil {
		toSerialize["BackendVmIds"] = o.BackendVmIds
	}
	if o.DnsName != nil {
		toSerialize["DnsName"] = o.DnsName
	}
	if o.HealthCheck != nil {
		toSerialize["HealthCheck"] = o.HealthCheck
	}
	if o.Listeners != nil {
		toSerialize["Listeners"] = o.Listeners
	}
	if o.LoadBalancerName != nil {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	if o.LoadBalancerStickyCookiePolicies != nil {
		toSerialize["LoadBalancerStickyCookiePolicies"] = o.LoadBalancerStickyCookiePolicies
	}
	if o.LoadBalancerType != nil {
		toSerialize["LoadBalancerType"] = o.LoadBalancerType
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.SecuredCookies != nil {
		toSerialize["SecuredCookies"] = o.SecuredCookies
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.SourceSecurityGroup != nil {
		toSerialize["SourceSecurityGroup"] = o.SourceSecurityGroup
	}
	if o.Subnets != nil {
		toSerialize["Subnets"] = o.Subnets
	}
	if o.SubregionNames != nil {
		toSerialize["SubregionNames"] = o.SubregionNames
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableLoadBalancer struct {
	value *LoadBalancer
	isSet bool
}

func (v NullableLoadBalancer) Get() *LoadBalancer {
	return v.value
}

func (v *NullableLoadBalancer) Set(val *LoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancer(val *LoadBalancer) *NullableLoadBalancer {
	return &NullableLoadBalancer{value: val, isSet: true}
}

func (v NullableLoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
