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

// FiltersSecurityGroup One or more filters.
type FiltersSecurityGroup struct {
	// The account IDs of the owners of the security groups.
	AccountIds *[]string `json:"AccountIds,omitempty"`
	// The descriptions of the security groups.
	Descriptions *[]string `json:"Descriptions,omitempty"`
	// The account IDs that have been granted permissions.
	InboundRuleAccountIds *[]string `json:"InboundRuleAccountIds,omitempty"`
	// The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.
	InboundRuleFromPortRanges *[]int32 `json:"InboundRuleFromPortRanges,omitempty"`
	// The IP ranges that have been granted permissions, in CIDR notation (for example, `10.0.0.0/24`).
	InboundRuleIpRanges *[]string `json:"InboundRuleIpRanges,omitempty"`
	// The IP protocols for the permissions (`tcp` \\| `udp` \\| `icmp`, or a protocol number, or `-1` for all protocols).
	InboundRuleProtocols *[]string `json:"InboundRuleProtocols,omitempty"`
	// The IDs of the security groups that have been granted permissions.
	InboundRuleSecurityGroupIds *[]string `json:"InboundRuleSecurityGroupIds,omitempty"`
	// The names of the security groups that have been granted permissions.
	InboundRuleSecurityGroupNames *[]string `json:"InboundRuleSecurityGroupNames,omitempty"`
	// The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.
	InboundRuleToPortRanges *[]int32 `json:"InboundRuleToPortRanges,omitempty"`
	// The IDs of the Nets specified when the security groups were created.
	NetIds *[]string `json:"NetIds,omitempty"`
	// The account IDs that have been granted permissions.
	OutboundRuleAccountIds *[]string `json:"OutboundRuleAccountIds,omitempty"`
	// The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.
	OutboundRuleFromPortRanges *[]int32 `json:"OutboundRuleFromPortRanges,omitempty"`
	// The IP ranges that have been granted permissions, in CIDR notation (for example, `10.0.0.0/24`).
	OutboundRuleIpRanges *[]string `json:"OutboundRuleIpRanges,omitempty"`
	// The IP protocols for the permissions (`tcp` \\| `udp` \\| `icmp`, or a protocol number, or `-1` for all protocols).
	OutboundRuleProtocols *[]string `json:"OutboundRuleProtocols,omitempty"`
	// The IDs of the security groups that have been granted permissions.
	OutboundRuleSecurityGroupIds *[]string `json:"OutboundRuleSecurityGroupIds,omitempty"`
	// The names of the security groups that have been granted permissions.
	OutboundRuleSecurityGroupNames *[]string `json:"OutboundRuleSecurityGroupNames,omitempty"`
	// The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.
	OutboundRuleToPortRanges *[]int32 `json:"OutboundRuleToPortRanges,omitempty"`
	// The IDs of the security groups.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The names of the security groups.
	SecurityGroupNames *[]string `json:"SecurityGroupNames,omitempty"`
	// The keys of the tags associated with the security groups.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the security groups.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the security groups, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersSecurityGroup instantiates a new FiltersSecurityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersSecurityGroup() *FiltersSecurityGroup {
	this := FiltersSecurityGroup{}
	return &this
}

// NewFiltersSecurityGroupWithDefaults instantiates a new FiltersSecurityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersSecurityGroupWithDefaults() *FiltersSecurityGroup {
	this := FiltersSecurityGroup{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *FiltersSecurityGroup) SetAccountIds(v []string) {
	o.AccountIds = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetDescriptionsOk() (*[]string, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *FiltersSecurityGroup) SetDescriptions(v []string) {
	o.Descriptions = &v
}

// GetInboundRuleAccountIds returns the InboundRuleAccountIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetInboundRuleAccountIds() []string {
	if o == nil || o.InboundRuleAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.InboundRuleAccountIds
}

// GetInboundRuleAccountIdsOk returns a tuple with the InboundRuleAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetInboundRuleAccountIdsOk() (*[]string, bool) {
	if o == nil || o.InboundRuleAccountIds == nil {
		return nil, false
	}
	return o.InboundRuleAccountIds, true
}

// HasInboundRuleAccountIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasInboundRuleAccountIds() bool {
	if o != nil && o.InboundRuleAccountIds != nil {
		return true
	}

	return false
}

// SetInboundRuleAccountIds gets a reference to the given []string and assigns it to the InboundRuleAccountIds field.
func (o *FiltersSecurityGroup) SetInboundRuleAccountIds(v []string) {
	o.InboundRuleAccountIds = &v
}

// GetInboundRuleFromPortRanges returns the InboundRuleFromPortRanges field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetInboundRuleFromPortRanges() []int32 {
	if o == nil || o.InboundRuleFromPortRanges == nil {
		var ret []int32
		return ret
	}
	return *o.InboundRuleFromPortRanges
}

// GetInboundRuleFromPortRangesOk returns a tuple with the InboundRuleFromPortRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetInboundRuleFromPortRangesOk() (*[]int32, bool) {
	if o == nil || o.InboundRuleFromPortRanges == nil {
		return nil, false
	}
	return o.InboundRuleFromPortRanges, true
}

// HasInboundRuleFromPortRanges returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasInboundRuleFromPortRanges() bool {
	if o != nil && o.InboundRuleFromPortRanges != nil {
		return true
	}

	return false
}

// SetInboundRuleFromPortRanges gets a reference to the given []int32 and assigns it to the InboundRuleFromPortRanges field.
func (o *FiltersSecurityGroup) SetInboundRuleFromPortRanges(v []int32) {
	o.InboundRuleFromPortRanges = &v
}

// GetInboundRuleIpRanges returns the InboundRuleIpRanges field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetInboundRuleIpRanges() []string {
	if o == nil || o.InboundRuleIpRanges == nil {
		var ret []string
		return ret
	}
	return *o.InboundRuleIpRanges
}

// GetInboundRuleIpRangesOk returns a tuple with the InboundRuleIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetInboundRuleIpRangesOk() (*[]string, bool) {
	if o == nil || o.InboundRuleIpRanges == nil {
		return nil, false
	}
	return o.InboundRuleIpRanges, true
}

// HasInboundRuleIpRanges returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasInboundRuleIpRanges() bool {
	if o != nil && o.InboundRuleIpRanges != nil {
		return true
	}

	return false
}

// SetInboundRuleIpRanges gets a reference to the given []string and assigns it to the InboundRuleIpRanges field.
func (o *FiltersSecurityGroup) SetInboundRuleIpRanges(v []string) {
	o.InboundRuleIpRanges = &v
}

// GetInboundRuleProtocols returns the InboundRuleProtocols field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetInboundRuleProtocols() []string {
	if o == nil || o.InboundRuleProtocols == nil {
		var ret []string
		return ret
	}
	return *o.InboundRuleProtocols
}

// GetInboundRuleProtocolsOk returns a tuple with the InboundRuleProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetInboundRuleProtocolsOk() (*[]string, bool) {
	if o == nil || o.InboundRuleProtocols == nil {
		return nil, false
	}
	return o.InboundRuleProtocols, true
}

// HasInboundRuleProtocols returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasInboundRuleProtocols() bool {
	if o != nil && o.InboundRuleProtocols != nil {
		return true
	}

	return false
}

// SetInboundRuleProtocols gets a reference to the given []string and assigns it to the InboundRuleProtocols field.
func (o *FiltersSecurityGroup) SetInboundRuleProtocols(v []string) {
	o.InboundRuleProtocols = &v
}

// GetInboundRuleSecurityGroupIds returns the InboundRuleSecurityGroupIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetInboundRuleSecurityGroupIds() []string {
	if o == nil || o.InboundRuleSecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.InboundRuleSecurityGroupIds
}

// GetInboundRuleSecurityGroupIdsOk returns a tuple with the InboundRuleSecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetInboundRuleSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.InboundRuleSecurityGroupIds == nil {
		return nil, false
	}
	return o.InboundRuleSecurityGroupIds, true
}

// HasInboundRuleSecurityGroupIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasInboundRuleSecurityGroupIds() bool {
	if o != nil && o.InboundRuleSecurityGroupIds != nil {
		return true
	}

	return false
}

// SetInboundRuleSecurityGroupIds gets a reference to the given []string and assigns it to the InboundRuleSecurityGroupIds field.
func (o *FiltersSecurityGroup) SetInboundRuleSecurityGroupIds(v []string) {
	o.InboundRuleSecurityGroupIds = &v
}

// GetInboundRuleSecurityGroupNames returns the InboundRuleSecurityGroupNames field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetInboundRuleSecurityGroupNames() []string {
	if o == nil || o.InboundRuleSecurityGroupNames == nil {
		var ret []string
		return ret
	}
	return *o.InboundRuleSecurityGroupNames
}

// GetInboundRuleSecurityGroupNamesOk returns a tuple with the InboundRuleSecurityGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetInboundRuleSecurityGroupNamesOk() (*[]string, bool) {
	if o == nil || o.InboundRuleSecurityGroupNames == nil {
		return nil, false
	}
	return o.InboundRuleSecurityGroupNames, true
}

// HasInboundRuleSecurityGroupNames returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasInboundRuleSecurityGroupNames() bool {
	if o != nil && o.InboundRuleSecurityGroupNames != nil {
		return true
	}

	return false
}

// SetInboundRuleSecurityGroupNames gets a reference to the given []string and assigns it to the InboundRuleSecurityGroupNames field.
func (o *FiltersSecurityGroup) SetInboundRuleSecurityGroupNames(v []string) {
	o.InboundRuleSecurityGroupNames = &v
}

// GetInboundRuleToPortRanges returns the InboundRuleToPortRanges field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetInboundRuleToPortRanges() []int32 {
	if o == nil || o.InboundRuleToPortRanges == nil {
		var ret []int32
		return ret
	}
	return *o.InboundRuleToPortRanges
}

// GetInboundRuleToPortRangesOk returns a tuple with the InboundRuleToPortRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetInboundRuleToPortRangesOk() (*[]int32, bool) {
	if o == nil || o.InboundRuleToPortRanges == nil {
		return nil, false
	}
	return o.InboundRuleToPortRanges, true
}

// HasInboundRuleToPortRanges returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasInboundRuleToPortRanges() bool {
	if o != nil && o.InboundRuleToPortRanges != nil {
		return true
	}

	return false
}

// SetInboundRuleToPortRanges gets a reference to the given []int32 and assigns it to the InboundRuleToPortRanges field.
func (o *FiltersSecurityGroup) SetInboundRuleToPortRanges(v []int32) {
	o.InboundRuleToPortRanges = &v
}

// GetNetIds returns the NetIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetNetIds() []string {
	if o == nil || o.NetIds == nil {
		var ret []string
		return ret
	}
	return *o.NetIds
}

// GetNetIdsOk returns a tuple with the NetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetNetIdsOk() (*[]string, bool) {
	if o == nil || o.NetIds == nil {
		return nil, false
	}
	return o.NetIds, true
}

// HasNetIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasNetIds() bool {
	if o != nil && o.NetIds != nil {
		return true
	}

	return false
}

// SetNetIds gets a reference to the given []string and assigns it to the NetIds field.
func (o *FiltersSecurityGroup) SetNetIds(v []string) {
	o.NetIds = &v
}

// GetOutboundRuleAccountIds returns the OutboundRuleAccountIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetOutboundRuleAccountIds() []string {
	if o == nil || o.OutboundRuleAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.OutboundRuleAccountIds
}

// GetOutboundRuleAccountIdsOk returns a tuple with the OutboundRuleAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetOutboundRuleAccountIdsOk() (*[]string, bool) {
	if o == nil || o.OutboundRuleAccountIds == nil {
		return nil, false
	}
	return o.OutboundRuleAccountIds, true
}

// HasOutboundRuleAccountIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasOutboundRuleAccountIds() bool {
	if o != nil && o.OutboundRuleAccountIds != nil {
		return true
	}

	return false
}

// SetOutboundRuleAccountIds gets a reference to the given []string and assigns it to the OutboundRuleAccountIds field.
func (o *FiltersSecurityGroup) SetOutboundRuleAccountIds(v []string) {
	o.OutboundRuleAccountIds = &v
}

// GetOutboundRuleFromPortRanges returns the OutboundRuleFromPortRanges field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetOutboundRuleFromPortRanges() []int32 {
	if o == nil || o.OutboundRuleFromPortRanges == nil {
		var ret []int32
		return ret
	}
	return *o.OutboundRuleFromPortRanges
}

// GetOutboundRuleFromPortRangesOk returns a tuple with the OutboundRuleFromPortRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetOutboundRuleFromPortRangesOk() (*[]int32, bool) {
	if o == nil || o.OutboundRuleFromPortRanges == nil {
		return nil, false
	}
	return o.OutboundRuleFromPortRanges, true
}

// HasOutboundRuleFromPortRanges returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasOutboundRuleFromPortRanges() bool {
	if o != nil && o.OutboundRuleFromPortRanges != nil {
		return true
	}

	return false
}

// SetOutboundRuleFromPortRanges gets a reference to the given []int32 and assigns it to the OutboundRuleFromPortRanges field.
func (o *FiltersSecurityGroup) SetOutboundRuleFromPortRanges(v []int32) {
	o.OutboundRuleFromPortRanges = &v
}

// GetOutboundRuleIpRanges returns the OutboundRuleIpRanges field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetOutboundRuleIpRanges() []string {
	if o == nil || o.OutboundRuleIpRanges == nil {
		var ret []string
		return ret
	}
	return *o.OutboundRuleIpRanges
}

// GetOutboundRuleIpRangesOk returns a tuple with the OutboundRuleIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetOutboundRuleIpRangesOk() (*[]string, bool) {
	if o == nil || o.OutboundRuleIpRanges == nil {
		return nil, false
	}
	return o.OutboundRuleIpRanges, true
}

// HasOutboundRuleIpRanges returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasOutboundRuleIpRanges() bool {
	if o != nil && o.OutboundRuleIpRanges != nil {
		return true
	}

	return false
}

// SetOutboundRuleIpRanges gets a reference to the given []string and assigns it to the OutboundRuleIpRanges field.
func (o *FiltersSecurityGroup) SetOutboundRuleIpRanges(v []string) {
	o.OutboundRuleIpRanges = &v
}

// GetOutboundRuleProtocols returns the OutboundRuleProtocols field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetOutboundRuleProtocols() []string {
	if o == nil || o.OutboundRuleProtocols == nil {
		var ret []string
		return ret
	}
	return *o.OutboundRuleProtocols
}

// GetOutboundRuleProtocolsOk returns a tuple with the OutboundRuleProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetOutboundRuleProtocolsOk() (*[]string, bool) {
	if o == nil || o.OutboundRuleProtocols == nil {
		return nil, false
	}
	return o.OutboundRuleProtocols, true
}

// HasOutboundRuleProtocols returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasOutboundRuleProtocols() bool {
	if o != nil && o.OutboundRuleProtocols != nil {
		return true
	}

	return false
}

// SetOutboundRuleProtocols gets a reference to the given []string and assigns it to the OutboundRuleProtocols field.
func (o *FiltersSecurityGroup) SetOutboundRuleProtocols(v []string) {
	o.OutboundRuleProtocols = &v
}

// GetOutboundRuleSecurityGroupIds returns the OutboundRuleSecurityGroupIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetOutboundRuleSecurityGroupIds() []string {
	if o == nil || o.OutboundRuleSecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.OutboundRuleSecurityGroupIds
}

// GetOutboundRuleSecurityGroupIdsOk returns a tuple with the OutboundRuleSecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetOutboundRuleSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.OutboundRuleSecurityGroupIds == nil {
		return nil, false
	}
	return o.OutboundRuleSecurityGroupIds, true
}

// HasOutboundRuleSecurityGroupIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasOutboundRuleSecurityGroupIds() bool {
	if o != nil && o.OutboundRuleSecurityGroupIds != nil {
		return true
	}

	return false
}

// SetOutboundRuleSecurityGroupIds gets a reference to the given []string and assigns it to the OutboundRuleSecurityGroupIds field.
func (o *FiltersSecurityGroup) SetOutboundRuleSecurityGroupIds(v []string) {
	o.OutboundRuleSecurityGroupIds = &v
}

// GetOutboundRuleSecurityGroupNames returns the OutboundRuleSecurityGroupNames field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetOutboundRuleSecurityGroupNames() []string {
	if o == nil || o.OutboundRuleSecurityGroupNames == nil {
		var ret []string
		return ret
	}
	return *o.OutboundRuleSecurityGroupNames
}

// GetOutboundRuleSecurityGroupNamesOk returns a tuple with the OutboundRuleSecurityGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetOutboundRuleSecurityGroupNamesOk() (*[]string, bool) {
	if o == nil || o.OutboundRuleSecurityGroupNames == nil {
		return nil, false
	}
	return o.OutboundRuleSecurityGroupNames, true
}

// HasOutboundRuleSecurityGroupNames returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasOutboundRuleSecurityGroupNames() bool {
	if o != nil && o.OutboundRuleSecurityGroupNames != nil {
		return true
	}

	return false
}

// SetOutboundRuleSecurityGroupNames gets a reference to the given []string and assigns it to the OutboundRuleSecurityGroupNames field.
func (o *FiltersSecurityGroup) SetOutboundRuleSecurityGroupNames(v []string) {
	o.OutboundRuleSecurityGroupNames = &v
}

// GetOutboundRuleToPortRanges returns the OutboundRuleToPortRanges field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetOutboundRuleToPortRanges() []int32 {
	if o == nil || o.OutboundRuleToPortRanges == nil {
		var ret []int32
		return ret
	}
	return *o.OutboundRuleToPortRanges
}

// GetOutboundRuleToPortRangesOk returns a tuple with the OutboundRuleToPortRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetOutboundRuleToPortRangesOk() (*[]int32, bool) {
	if o == nil || o.OutboundRuleToPortRanges == nil {
		return nil, false
	}
	return o.OutboundRuleToPortRanges, true
}

// HasOutboundRuleToPortRanges returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasOutboundRuleToPortRanges() bool {
	if o != nil && o.OutboundRuleToPortRanges != nil {
		return true
	}

	return false
}

// SetOutboundRuleToPortRanges gets a reference to the given []int32 and assigns it to the OutboundRuleToPortRanges field.
func (o *FiltersSecurityGroup) SetOutboundRuleToPortRanges(v []int32) {
	o.OutboundRuleToPortRanges = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *FiltersSecurityGroup) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSecurityGroupNames returns the SecurityGroupNames field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetSecurityGroupNames() []string {
	if o == nil || o.SecurityGroupNames == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupNames
}

// GetSecurityGroupNamesOk returns a tuple with the SecurityGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetSecurityGroupNamesOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupNames == nil {
		return nil, false
	}
	return o.SecurityGroupNames, true
}

// HasSecurityGroupNames returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasSecurityGroupNames() bool {
	if o != nil && o.SecurityGroupNames != nil {
		return true
	}

	return false
}

// SetSecurityGroupNames gets a reference to the given []string and assigns it to the SecurityGroupNames field.
func (o *FiltersSecurityGroup) SetSecurityGroupNames(v []string) {
	o.SecurityGroupNames = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersSecurityGroup) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersSecurityGroup) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersSecurityGroup) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSecurityGroup) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersSecurityGroup) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersSecurityGroup) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersSecurityGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["AccountIds"] = o.AccountIds
	}
	if o.Descriptions != nil {
		toSerialize["Descriptions"] = o.Descriptions
	}
	if o.InboundRuleAccountIds != nil {
		toSerialize["InboundRuleAccountIds"] = o.InboundRuleAccountIds
	}
	if o.InboundRuleFromPortRanges != nil {
		toSerialize["InboundRuleFromPortRanges"] = o.InboundRuleFromPortRanges
	}
	if o.InboundRuleIpRanges != nil {
		toSerialize["InboundRuleIpRanges"] = o.InboundRuleIpRanges
	}
	if o.InboundRuleProtocols != nil {
		toSerialize["InboundRuleProtocols"] = o.InboundRuleProtocols
	}
	if o.InboundRuleSecurityGroupIds != nil {
		toSerialize["InboundRuleSecurityGroupIds"] = o.InboundRuleSecurityGroupIds
	}
	if o.InboundRuleSecurityGroupNames != nil {
		toSerialize["InboundRuleSecurityGroupNames"] = o.InboundRuleSecurityGroupNames
	}
	if o.InboundRuleToPortRanges != nil {
		toSerialize["InboundRuleToPortRanges"] = o.InboundRuleToPortRanges
	}
	if o.NetIds != nil {
		toSerialize["NetIds"] = o.NetIds
	}
	if o.OutboundRuleAccountIds != nil {
		toSerialize["OutboundRuleAccountIds"] = o.OutboundRuleAccountIds
	}
	if o.OutboundRuleFromPortRanges != nil {
		toSerialize["OutboundRuleFromPortRanges"] = o.OutboundRuleFromPortRanges
	}
	if o.OutboundRuleIpRanges != nil {
		toSerialize["OutboundRuleIpRanges"] = o.OutboundRuleIpRanges
	}
	if o.OutboundRuleProtocols != nil {
		toSerialize["OutboundRuleProtocols"] = o.OutboundRuleProtocols
	}
	if o.OutboundRuleSecurityGroupIds != nil {
		toSerialize["OutboundRuleSecurityGroupIds"] = o.OutboundRuleSecurityGroupIds
	}
	if o.OutboundRuleSecurityGroupNames != nil {
		toSerialize["OutboundRuleSecurityGroupNames"] = o.OutboundRuleSecurityGroupNames
	}
	if o.OutboundRuleToPortRanges != nil {
		toSerialize["OutboundRuleToPortRanges"] = o.OutboundRuleToPortRanges
	}
	if o.SecurityGroupIds != nil {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	if o.SecurityGroupNames != nil {
		toSerialize["SecurityGroupNames"] = o.SecurityGroupNames
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

type NullableFiltersSecurityGroup struct {
	value *FiltersSecurityGroup
	isSet bool
}

func (v NullableFiltersSecurityGroup) Get() *FiltersSecurityGroup {
	return v.value
}

func (v *NullableFiltersSecurityGroup) Set(val *FiltersSecurityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersSecurityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersSecurityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersSecurityGroup(val *FiltersSecurityGroup) *NullableFiltersSecurityGroup {
	return &NullableFiltersSecurityGroup{value: val, isSet: true}
}

func (v NullableFiltersSecurityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersSecurityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
