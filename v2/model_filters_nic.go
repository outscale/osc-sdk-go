/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersNic One or more filters.
type FiltersNic struct {
	// The descriptions of the NICs.
	Descriptions *[]string `json:"Descriptions,omitempty"`
	// Whether the source/destination checking is enabled (true) or disabled (false).
	IsSourceDestCheck *bool `json:"IsSourceDestCheck,omitempty"`
	// Whether the NICs are deleted when the VMs they are attached to are terminated.
	LinkNicDeleteOnVmDeletion *bool `json:"LinkNicDeleteOnVmDeletion,omitempty"`
	// The device numbers the NICs are attached to.
	LinkNicDeviceNumbers *[]int32 `json:"LinkNicDeviceNumbers,omitempty"`
	// The attachment IDs of the NICs.
	LinkNicLinkNicIds *[]string `json:"LinkNicLinkNicIds,omitempty"`
	// The states of the attachments.
	LinkNicStates *[]string `json:"LinkNicStates,omitempty"`
	// The account IDs of the owners of the VMs the NICs are attached to.
	LinkNicVmAccountIds *[]string `json:"LinkNicVmAccountIds,omitempty"`
	// The IDs of the VMs the NICs are attached to.
	LinkNicVmIds *[]string `json:"LinkNicVmIds,omitempty"`
	// The account IDs of the owners of the EIPs associated with the NICs.
	LinkPublicIpAccountIds *[]string `json:"LinkPublicIpAccountIds,omitempty"`
	// The association IDs returned when the EIPs were associated with the NICs.
	LinkPublicIpLinkPublicIpIds *[]string `json:"LinkPublicIpLinkPublicIpIds,omitempty"`
	// The allocation IDs returned when the EIPs were allocated to their accounts.
	LinkPublicIpPublicIpIds *[]string `json:"LinkPublicIpPublicIpIds,omitempty"`
	// The EIPs associated with the NICs.
	LinkPublicIpPublicIps *[]string `json:"LinkPublicIpPublicIps,omitempty"`
	// The Media Access Control (MAC) addresses of the NICs.
	MacAddresses *[]string `json:"MacAddresses,omitempty"`
	// The IDs of the Nets where the NICs are located.
	NetIds *[]string `json:"NetIds,omitempty"`
	// The IDs of the NICs.
	NicIds *[]string `json:"NicIds,omitempty"`
	// The private DNS names associated with the primary private IP addresses.
	PrivateDnsNames *[]string `json:"PrivateDnsNames,omitempty"`
	// The account IDs of the owner of the EIPs associated with the private IP addresses.
	PrivateIpsLinkPublicIpAccountIds *[]string `json:"PrivateIpsLinkPublicIpAccountIds,omitempty"`
	// The EIPs associated with the private IP addresses.
	PrivateIpsLinkPublicIpPublicIps *[]string `json:"PrivateIpsLinkPublicIpPublicIps,omitempty"`
	// The primary private IP addresses of the NICs.
	PrivateIpsPrimaryIp *bool `json:"PrivateIpsPrimaryIp,omitempty"`
	// The private IP addresses of the NICs.
	PrivateIpsPrivateIps *[]string `json:"PrivateIpsPrivateIps,omitempty"`
	// The IDs of the security groups associated with the NICs.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The names of the security groups associated with the NICs.
	SecurityGroupNames *[]string `json:"SecurityGroupNames,omitempty"`
	// The states of the NICs.
	States *[]string `json:"States,omitempty"`
	// The IDs of the Subnets for the NICs.
	SubnetIds *[]string `json:"SubnetIds,omitempty"`
	// The Subregions where the NICs are located.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
	// The keys of the tags associated with the NICs.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the NICs.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the NICs, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersNic instantiates a new FiltersNic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersNic() *FiltersNic {
	this := FiltersNic{}
	return &this
}

// NewFiltersNicWithDefaults instantiates a new FiltersNic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersNicWithDefaults() *FiltersNic {
	this := FiltersNic{}
	return &this
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *FiltersNic) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetDescriptionsOk() (*[]string, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *FiltersNic) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *FiltersNic) SetDescriptions(v []string) {
	o.Descriptions = &v
}

// GetIsSourceDestCheck returns the IsSourceDestCheck field value if set, zero value otherwise.
func (o *FiltersNic) GetIsSourceDestCheck() bool {
	if o == nil || o.IsSourceDestCheck == nil {
		var ret bool
		return ret
	}
	return *o.IsSourceDestCheck
}

// GetIsSourceDestCheckOk returns a tuple with the IsSourceDestCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetIsSourceDestCheckOk() (*bool, bool) {
	if o == nil || o.IsSourceDestCheck == nil {
		return nil, false
	}
	return o.IsSourceDestCheck, true
}

// HasIsSourceDestCheck returns a boolean if a field has been set.
func (o *FiltersNic) HasIsSourceDestCheck() bool {
	if o != nil && o.IsSourceDestCheck != nil {
		return true
	}

	return false
}

// SetIsSourceDestCheck gets a reference to the given bool and assigns it to the IsSourceDestCheck field.
func (o *FiltersNic) SetIsSourceDestCheck(v bool) {
	o.IsSourceDestCheck = &v
}

// GetLinkNicDeleteOnVmDeletion returns the LinkNicDeleteOnVmDeletion field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkNicDeleteOnVmDeletion() bool {
	if o == nil || o.LinkNicDeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.LinkNicDeleteOnVmDeletion
}

// GetLinkNicDeleteOnVmDeletionOk returns a tuple with the LinkNicDeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkNicDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.LinkNicDeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.LinkNicDeleteOnVmDeletion, true
}

// HasLinkNicDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkNicDeleteOnVmDeletion() bool {
	if o != nil && o.LinkNicDeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetLinkNicDeleteOnVmDeletion gets a reference to the given bool and assigns it to the LinkNicDeleteOnVmDeletion field.
func (o *FiltersNic) SetLinkNicDeleteOnVmDeletion(v bool) {
	o.LinkNicDeleteOnVmDeletion = &v
}

// GetLinkNicDeviceNumbers returns the LinkNicDeviceNumbers field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkNicDeviceNumbers() []int32 {
	if o == nil || o.LinkNicDeviceNumbers == nil {
		var ret []int32
		return ret
	}
	return *o.LinkNicDeviceNumbers
}

// GetLinkNicDeviceNumbersOk returns a tuple with the LinkNicDeviceNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkNicDeviceNumbersOk() (*[]int32, bool) {
	if o == nil || o.LinkNicDeviceNumbers == nil {
		return nil, false
	}
	return o.LinkNicDeviceNumbers, true
}

// HasLinkNicDeviceNumbers returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkNicDeviceNumbers() bool {
	if o != nil && o.LinkNicDeviceNumbers != nil {
		return true
	}

	return false
}

// SetLinkNicDeviceNumbers gets a reference to the given []int32 and assigns it to the LinkNicDeviceNumbers field.
func (o *FiltersNic) SetLinkNicDeviceNumbers(v []int32) {
	o.LinkNicDeviceNumbers = &v
}

// GetLinkNicLinkNicIds returns the LinkNicLinkNicIds field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkNicLinkNicIds() []string {
	if o == nil || o.LinkNicLinkNicIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkNicLinkNicIds
}

// GetLinkNicLinkNicIdsOk returns a tuple with the LinkNicLinkNicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkNicLinkNicIdsOk() (*[]string, bool) {
	if o == nil || o.LinkNicLinkNicIds == nil {
		return nil, false
	}
	return o.LinkNicLinkNicIds, true
}

// HasLinkNicLinkNicIds returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkNicLinkNicIds() bool {
	if o != nil && o.LinkNicLinkNicIds != nil {
		return true
	}

	return false
}

// SetLinkNicLinkNicIds gets a reference to the given []string and assigns it to the LinkNicLinkNicIds field.
func (o *FiltersNic) SetLinkNicLinkNicIds(v []string) {
	o.LinkNicLinkNicIds = &v
}

// GetLinkNicStates returns the LinkNicStates field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkNicStates() []string {
	if o == nil || o.LinkNicStates == nil {
		var ret []string
		return ret
	}
	return *o.LinkNicStates
}

// GetLinkNicStatesOk returns a tuple with the LinkNicStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkNicStatesOk() (*[]string, bool) {
	if o == nil || o.LinkNicStates == nil {
		return nil, false
	}
	return o.LinkNicStates, true
}

// HasLinkNicStates returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkNicStates() bool {
	if o != nil && o.LinkNicStates != nil {
		return true
	}

	return false
}

// SetLinkNicStates gets a reference to the given []string and assigns it to the LinkNicStates field.
func (o *FiltersNic) SetLinkNicStates(v []string) {
	o.LinkNicStates = &v
}

// GetLinkNicVmAccountIds returns the LinkNicVmAccountIds field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkNicVmAccountIds() []string {
	if o == nil || o.LinkNicVmAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkNicVmAccountIds
}

// GetLinkNicVmAccountIdsOk returns a tuple with the LinkNicVmAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkNicVmAccountIdsOk() (*[]string, bool) {
	if o == nil || o.LinkNicVmAccountIds == nil {
		return nil, false
	}
	return o.LinkNicVmAccountIds, true
}

// HasLinkNicVmAccountIds returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkNicVmAccountIds() bool {
	if o != nil && o.LinkNicVmAccountIds != nil {
		return true
	}

	return false
}

// SetLinkNicVmAccountIds gets a reference to the given []string and assigns it to the LinkNicVmAccountIds field.
func (o *FiltersNic) SetLinkNicVmAccountIds(v []string) {
	o.LinkNicVmAccountIds = &v
}

// GetLinkNicVmIds returns the LinkNicVmIds field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkNicVmIds() []string {
	if o == nil || o.LinkNicVmIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkNicVmIds
}

// GetLinkNicVmIdsOk returns a tuple with the LinkNicVmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkNicVmIdsOk() (*[]string, bool) {
	if o == nil || o.LinkNicVmIds == nil {
		return nil, false
	}
	return o.LinkNicVmIds, true
}

// HasLinkNicVmIds returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkNicVmIds() bool {
	if o != nil && o.LinkNicVmIds != nil {
		return true
	}

	return false
}

// SetLinkNicVmIds gets a reference to the given []string and assigns it to the LinkNicVmIds field.
func (o *FiltersNic) SetLinkNicVmIds(v []string) {
	o.LinkNicVmIds = &v
}

// GetLinkPublicIpAccountIds returns the LinkPublicIpAccountIds field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkPublicIpAccountIds() []string {
	if o == nil || o.LinkPublicIpAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkPublicIpAccountIds
}

// GetLinkPublicIpAccountIdsOk returns a tuple with the LinkPublicIpAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkPublicIpAccountIdsOk() (*[]string, bool) {
	if o == nil || o.LinkPublicIpAccountIds == nil {
		return nil, false
	}
	return o.LinkPublicIpAccountIds, true
}

// HasLinkPublicIpAccountIds returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkPublicIpAccountIds() bool {
	if o != nil && o.LinkPublicIpAccountIds != nil {
		return true
	}

	return false
}

// SetLinkPublicIpAccountIds gets a reference to the given []string and assigns it to the LinkPublicIpAccountIds field.
func (o *FiltersNic) SetLinkPublicIpAccountIds(v []string) {
	o.LinkPublicIpAccountIds = &v
}

// GetLinkPublicIpLinkPublicIpIds returns the LinkPublicIpLinkPublicIpIds field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkPublicIpLinkPublicIpIds() []string {
	if o == nil || o.LinkPublicIpLinkPublicIpIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkPublicIpLinkPublicIpIds
}

// GetLinkPublicIpLinkPublicIpIdsOk returns a tuple with the LinkPublicIpLinkPublicIpIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkPublicIpLinkPublicIpIdsOk() (*[]string, bool) {
	if o == nil || o.LinkPublicIpLinkPublicIpIds == nil {
		return nil, false
	}
	return o.LinkPublicIpLinkPublicIpIds, true
}

// HasLinkPublicIpLinkPublicIpIds returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkPublicIpLinkPublicIpIds() bool {
	if o != nil && o.LinkPublicIpLinkPublicIpIds != nil {
		return true
	}

	return false
}

// SetLinkPublicIpLinkPublicIpIds gets a reference to the given []string and assigns it to the LinkPublicIpLinkPublicIpIds field.
func (o *FiltersNic) SetLinkPublicIpLinkPublicIpIds(v []string) {
	o.LinkPublicIpLinkPublicIpIds = &v
}

// GetLinkPublicIpPublicIpIds returns the LinkPublicIpPublicIpIds field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkPublicIpPublicIpIds() []string {
	if o == nil || o.LinkPublicIpPublicIpIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkPublicIpPublicIpIds
}

// GetLinkPublicIpPublicIpIdsOk returns a tuple with the LinkPublicIpPublicIpIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkPublicIpPublicIpIdsOk() (*[]string, bool) {
	if o == nil || o.LinkPublicIpPublicIpIds == nil {
		return nil, false
	}
	return o.LinkPublicIpPublicIpIds, true
}

// HasLinkPublicIpPublicIpIds returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkPublicIpPublicIpIds() bool {
	if o != nil && o.LinkPublicIpPublicIpIds != nil {
		return true
	}

	return false
}

// SetLinkPublicIpPublicIpIds gets a reference to the given []string and assigns it to the LinkPublicIpPublicIpIds field.
func (o *FiltersNic) SetLinkPublicIpPublicIpIds(v []string) {
	o.LinkPublicIpPublicIpIds = &v
}

// GetLinkPublicIpPublicIps returns the LinkPublicIpPublicIps field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkPublicIpPublicIps() []string {
	if o == nil || o.LinkPublicIpPublicIps == nil {
		var ret []string
		return ret
	}
	return *o.LinkPublicIpPublicIps
}

// GetLinkPublicIpPublicIpsOk returns a tuple with the LinkPublicIpPublicIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkPublicIpPublicIpsOk() (*[]string, bool) {
	if o == nil || o.LinkPublicIpPublicIps == nil {
		return nil, false
	}
	return o.LinkPublicIpPublicIps, true
}

// HasLinkPublicIpPublicIps returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkPublicIpPublicIps() bool {
	if o != nil && o.LinkPublicIpPublicIps != nil {
		return true
	}

	return false
}

// SetLinkPublicIpPublicIps gets a reference to the given []string and assigns it to the LinkPublicIpPublicIps field.
func (o *FiltersNic) SetLinkPublicIpPublicIps(v []string) {
	o.LinkPublicIpPublicIps = &v
}

// GetMacAddresses returns the MacAddresses field value if set, zero value otherwise.
func (o *FiltersNic) GetMacAddresses() []string {
	if o == nil || o.MacAddresses == nil {
		var ret []string
		return ret
	}
	return *o.MacAddresses
}

// GetMacAddressesOk returns a tuple with the MacAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetMacAddressesOk() (*[]string, bool) {
	if o == nil || o.MacAddresses == nil {
		return nil, false
	}
	return o.MacAddresses, true
}

// HasMacAddresses returns a boolean if a field has been set.
func (o *FiltersNic) HasMacAddresses() bool {
	if o != nil && o.MacAddresses != nil {
		return true
	}

	return false
}

// SetMacAddresses gets a reference to the given []string and assigns it to the MacAddresses field.
func (o *FiltersNic) SetMacAddresses(v []string) {
	o.MacAddresses = &v
}

// GetNetIds returns the NetIds field value if set, zero value otherwise.
func (o *FiltersNic) GetNetIds() []string {
	if o == nil || o.NetIds == nil {
		var ret []string
		return ret
	}
	return *o.NetIds
}

// GetNetIdsOk returns a tuple with the NetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetNetIdsOk() (*[]string, bool) {
	if o == nil || o.NetIds == nil {
		return nil, false
	}
	return o.NetIds, true
}

// HasNetIds returns a boolean if a field has been set.
func (o *FiltersNic) HasNetIds() bool {
	if o != nil && o.NetIds != nil {
		return true
	}

	return false
}

// SetNetIds gets a reference to the given []string and assigns it to the NetIds field.
func (o *FiltersNic) SetNetIds(v []string) {
	o.NetIds = &v
}

// GetNicIds returns the NicIds field value if set, zero value otherwise.
func (o *FiltersNic) GetNicIds() []string {
	if o == nil || o.NicIds == nil {
		var ret []string
		return ret
	}
	return *o.NicIds
}

// GetNicIdsOk returns a tuple with the NicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetNicIdsOk() (*[]string, bool) {
	if o == nil || o.NicIds == nil {
		return nil, false
	}
	return o.NicIds, true
}

// HasNicIds returns a boolean if a field has been set.
func (o *FiltersNic) HasNicIds() bool {
	if o != nil && o.NicIds != nil {
		return true
	}

	return false
}

// SetNicIds gets a reference to the given []string and assigns it to the NicIds field.
func (o *FiltersNic) SetNicIds(v []string) {
	o.NicIds = &v
}

// GetPrivateDnsNames returns the PrivateDnsNames field value if set, zero value otherwise.
func (o *FiltersNic) GetPrivateDnsNames() []string {
	if o == nil || o.PrivateDnsNames == nil {
		var ret []string
		return ret
	}
	return *o.PrivateDnsNames
}

// GetPrivateDnsNamesOk returns a tuple with the PrivateDnsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetPrivateDnsNamesOk() (*[]string, bool) {
	if o == nil || o.PrivateDnsNames == nil {
		return nil, false
	}
	return o.PrivateDnsNames, true
}

// HasPrivateDnsNames returns a boolean if a field has been set.
func (o *FiltersNic) HasPrivateDnsNames() bool {
	if o != nil && o.PrivateDnsNames != nil {
		return true
	}

	return false
}

// SetPrivateDnsNames gets a reference to the given []string and assigns it to the PrivateDnsNames field.
func (o *FiltersNic) SetPrivateDnsNames(v []string) {
	o.PrivateDnsNames = &v
}

// GetPrivateIpsLinkPublicIpAccountIds returns the PrivateIpsLinkPublicIpAccountIds field value if set, zero value otherwise.
func (o *FiltersNic) GetPrivateIpsLinkPublicIpAccountIds() []string {
	if o == nil || o.PrivateIpsLinkPublicIpAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIpsLinkPublicIpAccountIds
}

// GetPrivateIpsLinkPublicIpAccountIdsOk returns a tuple with the PrivateIpsLinkPublicIpAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetPrivateIpsLinkPublicIpAccountIdsOk() (*[]string, bool) {
	if o == nil || o.PrivateIpsLinkPublicIpAccountIds == nil {
		return nil, false
	}
	return o.PrivateIpsLinkPublicIpAccountIds, true
}

// HasPrivateIpsLinkPublicIpAccountIds returns a boolean if a field has been set.
func (o *FiltersNic) HasPrivateIpsLinkPublicIpAccountIds() bool {
	if o != nil && o.PrivateIpsLinkPublicIpAccountIds != nil {
		return true
	}

	return false
}

// SetPrivateIpsLinkPublicIpAccountIds gets a reference to the given []string and assigns it to the PrivateIpsLinkPublicIpAccountIds field.
func (o *FiltersNic) SetPrivateIpsLinkPublicIpAccountIds(v []string) {
	o.PrivateIpsLinkPublicIpAccountIds = &v
}

// GetPrivateIpsLinkPublicIpPublicIps returns the PrivateIpsLinkPublicIpPublicIps field value if set, zero value otherwise.
func (o *FiltersNic) GetPrivateIpsLinkPublicIpPublicIps() []string {
	if o == nil || o.PrivateIpsLinkPublicIpPublicIps == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIpsLinkPublicIpPublicIps
}

// GetPrivateIpsLinkPublicIpPublicIpsOk returns a tuple with the PrivateIpsLinkPublicIpPublicIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetPrivateIpsLinkPublicIpPublicIpsOk() (*[]string, bool) {
	if o == nil || o.PrivateIpsLinkPublicIpPublicIps == nil {
		return nil, false
	}
	return o.PrivateIpsLinkPublicIpPublicIps, true
}

// HasPrivateIpsLinkPublicIpPublicIps returns a boolean if a field has been set.
func (o *FiltersNic) HasPrivateIpsLinkPublicIpPublicIps() bool {
	if o != nil && o.PrivateIpsLinkPublicIpPublicIps != nil {
		return true
	}

	return false
}

// SetPrivateIpsLinkPublicIpPublicIps gets a reference to the given []string and assigns it to the PrivateIpsLinkPublicIpPublicIps field.
func (o *FiltersNic) SetPrivateIpsLinkPublicIpPublicIps(v []string) {
	o.PrivateIpsLinkPublicIpPublicIps = &v
}

// GetPrivateIpsPrimaryIp returns the PrivateIpsPrimaryIp field value if set, zero value otherwise.
func (o *FiltersNic) GetPrivateIpsPrimaryIp() bool {
	if o == nil || o.PrivateIpsPrimaryIp == nil {
		var ret bool
		return ret
	}
	return *o.PrivateIpsPrimaryIp
}

// GetPrivateIpsPrimaryIpOk returns a tuple with the PrivateIpsPrimaryIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetPrivateIpsPrimaryIpOk() (*bool, bool) {
	if o == nil || o.PrivateIpsPrimaryIp == nil {
		return nil, false
	}
	return o.PrivateIpsPrimaryIp, true
}

// HasPrivateIpsPrimaryIp returns a boolean if a field has been set.
func (o *FiltersNic) HasPrivateIpsPrimaryIp() bool {
	if o != nil && o.PrivateIpsPrimaryIp != nil {
		return true
	}

	return false
}

// SetPrivateIpsPrimaryIp gets a reference to the given bool and assigns it to the PrivateIpsPrimaryIp field.
func (o *FiltersNic) SetPrivateIpsPrimaryIp(v bool) {
	o.PrivateIpsPrimaryIp = &v
}

// GetPrivateIpsPrivateIps returns the PrivateIpsPrivateIps field value if set, zero value otherwise.
func (o *FiltersNic) GetPrivateIpsPrivateIps() []string {
	if o == nil || o.PrivateIpsPrivateIps == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIpsPrivateIps
}

// GetPrivateIpsPrivateIpsOk returns a tuple with the PrivateIpsPrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetPrivateIpsPrivateIpsOk() (*[]string, bool) {
	if o == nil || o.PrivateIpsPrivateIps == nil {
		return nil, false
	}
	return o.PrivateIpsPrivateIps, true
}

// HasPrivateIpsPrivateIps returns a boolean if a field has been set.
func (o *FiltersNic) HasPrivateIpsPrivateIps() bool {
	if o != nil && o.PrivateIpsPrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIpsPrivateIps gets a reference to the given []string and assigns it to the PrivateIpsPrivateIps field.
func (o *FiltersNic) SetPrivateIpsPrivateIps(v []string) {
	o.PrivateIpsPrivateIps = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *FiltersNic) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *FiltersNic) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *FiltersNic) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSecurityGroupNames returns the SecurityGroupNames field value if set, zero value otherwise.
func (o *FiltersNic) GetSecurityGroupNames() []string {
	if o == nil || o.SecurityGroupNames == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupNames
}

// GetSecurityGroupNamesOk returns a tuple with the SecurityGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetSecurityGroupNamesOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupNames == nil {
		return nil, false
	}
	return o.SecurityGroupNames, true
}

// HasSecurityGroupNames returns a boolean if a field has been set.
func (o *FiltersNic) HasSecurityGroupNames() bool {
	if o != nil && o.SecurityGroupNames != nil {
		return true
	}

	return false
}

// SetSecurityGroupNames gets a reference to the given []string and assigns it to the SecurityGroupNames field.
func (o *FiltersNic) SetSecurityGroupNames(v []string) {
	o.SecurityGroupNames = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersNic) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetStatesOk() (*[]string, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersNic) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersNic) SetStates(v []string) {
	o.States = &v
}

// GetSubnetIds returns the SubnetIds field value if set, zero value otherwise.
func (o *FiltersNic) GetSubnetIds() []string {
	if o == nil || o.SubnetIds == nil {
		var ret []string
		return ret
	}
	return *o.SubnetIds
}

// GetSubnetIdsOk returns a tuple with the SubnetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetSubnetIdsOk() (*[]string, bool) {
	if o == nil || o.SubnetIds == nil {
		return nil, false
	}
	return o.SubnetIds, true
}

// HasSubnetIds returns a boolean if a field has been set.
func (o *FiltersNic) HasSubnetIds() bool {
	if o != nil && o.SubnetIds != nil {
		return true
	}

	return false
}

// SetSubnetIds gets a reference to the given []string and assigns it to the SubnetIds field.
func (o *FiltersNic) SetSubnetIds(v []string) {
	o.SubnetIds = &v
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *FiltersNic) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetSubregionNamesOk() (*[]string, bool) {
	if o == nil || o.SubregionNames == nil {
		return nil, false
	}
	return o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *FiltersNic) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *FiltersNic) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersNic) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersNic) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersNic) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersNic) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersNic) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersNic) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersNic) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersNic) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersNic) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersNic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Descriptions != nil {
		toSerialize["Descriptions"] = o.Descriptions
	}
	if o.IsSourceDestCheck != nil {
		toSerialize["IsSourceDestCheck"] = o.IsSourceDestCheck
	}
	if o.LinkNicDeleteOnVmDeletion != nil {
		toSerialize["LinkNicDeleteOnVmDeletion"] = o.LinkNicDeleteOnVmDeletion
	}
	if o.LinkNicDeviceNumbers != nil {
		toSerialize["LinkNicDeviceNumbers"] = o.LinkNicDeviceNumbers
	}
	if o.LinkNicLinkNicIds != nil {
		toSerialize["LinkNicLinkNicIds"] = o.LinkNicLinkNicIds
	}
	if o.LinkNicStates != nil {
		toSerialize["LinkNicStates"] = o.LinkNicStates
	}
	if o.LinkNicVmAccountIds != nil {
		toSerialize["LinkNicVmAccountIds"] = o.LinkNicVmAccountIds
	}
	if o.LinkNicVmIds != nil {
		toSerialize["LinkNicVmIds"] = o.LinkNicVmIds
	}
	if o.LinkPublicIpAccountIds != nil {
		toSerialize["LinkPublicIpAccountIds"] = o.LinkPublicIpAccountIds
	}
	if o.LinkPublicIpLinkPublicIpIds != nil {
		toSerialize["LinkPublicIpLinkPublicIpIds"] = o.LinkPublicIpLinkPublicIpIds
	}
	if o.LinkPublicIpPublicIpIds != nil {
		toSerialize["LinkPublicIpPublicIpIds"] = o.LinkPublicIpPublicIpIds
	}
	if o.LinkPublicIpPublicIps != nil {
		toSerialize["LinkPublicIpPublicIps"] = o.LinkPublicIpPublicIps
	}
	if o.MacAddresses != nil {
		toSerialize["MacAddresses"] = o.MacAddresses
	}
	if o.NetIds != nil {
		toSerialize["NetIds"] = o.NetIds
	}
	if o.NicIds != nil {
		toSerialize["NicIds"] = o.NicIds
	}
	if o.PrivateDnsNames != nil {
		toSerialize["PrivateDnsNames"] = o.PrivateDnsNames
	}
	if o.PrivateIpsLinkPublicIpAccountIds != nil {
		toSerialize["PrivateIpsLinkPublicIpAccountIds"] = o.PrivateIpsLinkPublicIpAccountIds
	}
	if o.PrivateIpsLinkPublicIpPublicIps != nil {
		toSerialize["PrivateIpsLinkPublicIpPublicIps"] = o.PrivateIpsLinkPublicIpPublicIps
	}
	if o.PrivateIpsPrimaryIp != nil {
		toSerialize["PrivateIpsPrimaryIp"] = o.PrivateIpsPrimaryIp
	}
	if o.PrivateIpsPrivateIps != nil {
		toSerialize["PrivateIpsPrivateIps"] = o.PrivateIpsPrivateIps
	}
	if o.SecurityGroupIds != nil {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	if o.SecurityGroupNames != nil {
		toSerialize["SecurityGroupNames"] = o.SecurityGroupNames
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

type NullableFiltersNic struct {
	value *FiltersNic
	isSet bool
}

func (v NullableFiltersNic) Get() *FiltersNic {
	return v.value
}

func (v *NullableFiltersNic) Set(val *FiltersNic) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersNic) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersNic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersNic(val *FiltersNic) *NullableFiltersNic {
	return &NullableFiltersNic{value: val, isSet: true}
}

func (v NullableFiltersNic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersNic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
