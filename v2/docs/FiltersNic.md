# FiltersNic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptions** | Pointer to **[]string** | The descriptions of the NICs. | [optional] 
**IsSourceDestCheck** | Pointer to **bool** | Whether the source/destination checking is enabled (true) or disabled (false). | [optional] 
**LinkNicDeleteOnVmDeletion** | Pointer to **bool** | Whether the NICs are deleted when the VMs they are attached to are terminated. | [optional] 
**LinkNicDeviceNumbers** | Pointer to **[]int32** | The device numbers the NICs are attached to. | [optional] 
**LinkNicLinkNicIds** | Pointer to **[]string** | The attachment IDs of the NICs. | [optional] 
**LinkNicStates** | Pointer to **[]string** | The states of the attachments. | [optional] 
**LinkNicVmAccountIds** | Pointer to **[]string** | The account IDs of the owners of the VMs the NICs are attached to. | [optional] 
**LinkNicVmIds** | Pointer to **[]string** | The IDs of the VMs the NICs are attached to. | [optional] 
**LinkPublicIpAccountIds** | Pointer to **[]string** | The account IDs of the owners of the EIPs associated with the NICs. | [optional] 
**LinkPublicIpLinkPublicIpIds** | Pointer to **[]string** | The association IDs returned when the EIPs were associated with the NICs. | [optional] 
**LinkPublicIpPublicIpIds** | Pointer to **[]string** | The allocation IDs returned when the EIPs were allocated to their accounts. | [optional] 
**LinkPublicIpPublicIps** | Pointer to **[]string** | The EIPs associated with the NICs. | [optional] 
**MacAddresses** | Pointer to **[]string** | The Media Access Control (MAC) addresses of the NICs. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets where the NICs are located. | [optional] 
**NicIds** | Pointer to **[]string** | The IDs of the NICs. | [optional] 
**PrivateDnsNames** | Pointer to **[]string** | The private DNS names associated with the primary private IP addresses. | [optional] 
**PrivateIpsLinkPublicIpAccountIds** | Pointer to **[]string** | The account IDs of the owner of the EIPs associated with the private IP addresses. | [optional] 
**PrivateIpsLinkPublicIpPublicIps** | Pointer to **[]string** | The EIPs associated with the private IP addresses. | [optional] 
**PrivateIpsPrimaryIp** | Pointer to **bool** | The primary private IP addresses of the NICs. | [optional] 
**PrivateIpsPrivateIps** | Pointer to **[]string** | The private IP addresses of the NICs. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups associated with the NICs. | [optional] 
**SecurityGroupNames** | Pointer to **[]string** | The names of the security groups associated with the NICs. | [optional] 
**States** | Pointer to **[]string** | The states of the NICs. | [optional] 
**SubnetIds** | Pointer to **[]string** | The IDs of the Subnets for the NICs. | [optional] 
**SubregionNames** | Pointer to **[]string** | The Subregions where the NICs are located. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the NICs. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the NICs. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the NICs, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

## Methods

### NewFiltersNic

`func NewFiltersNic() *FiltersNic`

NewFiltersNic instantiates a new FiltersNic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersNicWithDefaults

`func NewFiltersNicWithDefaults() *FiltersNic`

NewFiltersNicWithDefaults instantiates a new FiltersNic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptions

`func (o *FiltersNic) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersNic) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *FiltersNic) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *FiltersNic) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetIsSourceDestCheck

`func (o *FiltersNic) GetIsSourceDestCheck() bool`

GetIsSourceDestCheck returns the IsSourceDestCheck field if non-nil, zero value otherwise.

### GetIsSourceDestCheckOk

`func (o *FiltersNic) GetIsSourceDestCheckOk() (*bool, bool)`

GetIsSourceDestCheckOk returns a tuple with the IsSourceDestCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceDestCheck

`func (o *FiltersNic) SetIsSourceDestCheck(v bool)`

SetIsSourceDestCheck sets IsSourceDestCheck field to given value.

### HasIsSourceDestCheck

`func (o *FiltersNic) HasIsSourceDestCheck() bool`

HasIsSourceDestCheck returns a boolean if a field has been set.

### GetLinkNicDeleteOnVmDeletion

`func (o *FiltersNic) GetLinkNicDeleteOnVmDeletion() bool`

GetLinkNicDeleteOnVmDeletion returns the LinkNicDeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetLinkNicDeleteOnVmDeletionOk

`func (o *FiltersNic) GetLinkNicDeleteOnVmDeletionOk() (*bool, bool)`

GetLinkNicDeleteOnVmDeletionOk returns a tuple with the LinkNicDeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicDeleteOnVmDeletion

`func (o *FiltersNic) SetLinkNicDeleteOnVmDeletion(v bool)`

SetLinkNicDeleteOnVmDeletion sets LinkNicDeleteOnVmDeletion field to given value.

### HasLinkNicDeleteOnVmDeletion

`func (o *FiltersNic) HasLinkNicDeleteOnVmDeletion() bool`

HasLinkNicDeleteOnVmDeletion returns a boolean if a field has been set.

### GetLinkNicDeviceNumbers

`func (o *FiltersNic) GetLinkNicDeviceNumbers() []int32`

GetLinkNicDeviceNumbers returns the LinkNicDeviceNumbers field if non-nil, zero value otherwise.

### GetLinkNicDeviceNumbersOk

`func (o *FiltersNic) GetLinkNicDeviceNumbersOk() (*[]int32, bool)`

GetLinkNicDeviceNumbersOk returns a tuple with the LinkNicDeviceNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicDeviceNumbers

`func (o *FiltersNic) SetLinkNicDeviceNumbers(v []int32)`

SetLinkNicDeviceNumbers sets LinkNicDeviceNumbers field to given value.

### HasLinkNicDeviceNumbers

`func (o *FiltersNic) HasLinkNicDeviceNumbers() bool`

HasLinkNicDeviceNumbers returns a boolean if a field has been set.

### GetLinkNicLinkNicIds

`func (o *FiltersNic) GetLinkNicLinkNicIds() []string`

GetLinkNicLinkNicIds returns the LinkNicLinkNicIds field if non-nil, zero value otherwise.

### GetLinkNicLinkNicIdsOk

`func (o *FiltersNic) GetLinkNicLinkNicIdsOk() (*[]string, bool)`

GetLinkNicLinkNicIdsOk returns a tuple with the LinkNicLinkNicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicLinkNicIds

`func (o *FiltersNic) SetLinkNicLinkNicIds(v []string)`

SetLinkNicLinkNicIds sets LinkNicLinkNicIds field to given value.

### HasLinkNicLinkNicIds

`func (o *FiltersNic) HasLinkNicLinkNicIds() bool`

HasLinkNicLinkNicIds returns a boolean if a field has been set.

### GetLinkNicStates

`func (o *FiltersNic) GetLinkNicStates() []string`

GetLinkNicStates returns the LinkNicStates field if non-nil, zero value otherwise.

### GetLinkNicStatesOk

`func (o *FiltersNic) GetLinkNicStatesOk() (*[]string, bool)`

GetLinkNicStatesOk returns a tuple with the LinkNicStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicStates

`func (o *FiltersNic) SetLinkNicStates(v []string)`

SetLinkNicStates sets LinkNicStates field to given value.

### HasLinkNicStates

`func (o *FiltersNic) HasLinkNicStates() bool`

HasLinkNicStates returns a boolean if a field has been set.

### GetLinkNicVmAccountIds

`func (o *FiltersNic) GetLinkNicVmAccountIds() []string`

GetLinkNicVmAccountIds returns the LinkNicVmAccountIds field if non-nil, zero value otherwise.

### GetLinkNicVmAccountIdsOk

`func (o *FiltersNic) GetLinkNicVmAccountIdsOk() (*[]string, bool)`

GetLinkNicVmAccountIdsOk returns a tuple with the LinkNicVmAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicVmAccountIds

`func (o *FiltersNic) SetLinkNicVmAccountIds(v []string)`

SetLinkNicVmAccountIds sets LinkNicVmAccountIds field to given value.

### HasLinkNicVmAccountIds

`func (o *FiltersNic) HasLinkNicVmAccountIds() bool`

HasLinkNicVmAccountIds returns a boolean if a field has been set.

### GetLinkNicVmIds

`func (o *FiltersNic) GetLinkNicVmIds() []string`

GetLinkNicVmIds returns the LinkNicVmIds field if non-nil, zero value otherwise.

### GetLinkNicVmIdsOk

`func (o *FiltersNic) GetLinkNicVmIdsOk() (*[]string, bool)`

GetLinkNicVmIdsOk returns a tuple with the LinkNicVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicVmIds

`func (o *FiltersNic) SetLinkNicVmIds(v []string)`

SetLinkNicVmIds sets LinkNicVmIds field to given value.

### HasLinkNicVmIds

`func (o *FiltersNic) HasLinkNicVmIds() bool`

HasLinkNicVmIds returns a boolean if a field has been set.

### GetLinkPublicIpAccountIds

`func (o *FiltersNic) GetLinkPublicIpAccountIds() []string`

GetLinkPublicIpAccountIds returns the LinkPublicIpAccountIds field if non-nil, zero value otherwise.

### GetLinkPublicIpAccountIdsOk

`func (o *FiltersNic) GetLinkPublicIpAccountIdsOk() (*[]string, bool)`

GetLinkPublicIpAccountIdsOk returns a tuple with the LinkPublicIpAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIpAccountIds

`func (o *FiltersNic) SetLinkPublicIpAccountIds(v []string)`

SetLinkPublicIpAccountIds sets LinkPublicIpAccountIds field to given value.

### HasLinkPublicIpAccountIds

`func (o *FiltersNic) HasLinkPublicIpAccountIds() bool`

HasLinkPublicIpAccountIds returns a boolean if a field has been set.

### GetLinkPublicIpLinkPublicIpIds

`func (o *FiltersNic) GetLinkPublicIpLinkPublicIpIds() []string`

GetLinkPublicIpLinkPublicIpIds returns the LinkPublicIpLinkPublicIpIds field if non-nil, zero value otherwise.

### GetLinkPublicIpLinkPublicIpIdsOk

`func (o *FiltersNic) GetLinkPublicIpLinkPublicIpIdsOk() (*[]string, bool)`

GetLinkPublicIpLinkPublicIpIdsOk returns a tuple with the LinkPublicIpLinkPublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIpLinkPublicIpIds

`func (o *FiltersNic) SetLinkPublicIpLinkPublicIpIds(v []string)`

SetLinkPublicIpLinkPublicIpIds sets LinkPublicIpLinkPublicIpIds field to given value.

### HasLinkPublicIpLinkPublicIpIds

`func (o *FiltersNic) HasLinkPublicIpLinkPublicIpIds() bool`

HasLinkPublicIpLinkPublicIpIds returns a boolean if a field has been set.

### GetLinkPublicIpPublicIpIds

`func (o *FiltersNic) GetLinkPublicIpPublicIpIds() []string`

GetLinkPublicIpPublicIpIds returns the LinkPublicIpPublicIpIds field if non-nil, zero value otherwise.

### GetLinkPublicIpPublicIpIdsOk

`func (o *FiltersNic) GetLinkPublicIpPublicIpIdsOk() (*[]string, bool)`

GetLinkPublicIpPublicIpIdsOk returns a tuple with the LinkPublicIpPublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIpPublicIpIds

`func (o *FiltersNic) SetLinkPublicIpPublicIpIds(v []string)`

SetLinkPublicIpPublicIpIds sets LinkPublicIpPublicIpIds field to given value.

### HasLinkPublicIpPublicIpIds

`func (o *FiltersNic) HasLinkPublicIpPublicIpIds() bool`

HasLinkPublicIpPublicIpIds returns a boolean if a field has been set.

### GetLinkPublicIpPublicIps

`func (o *FiltersNic) GetLinkPublicIpPublicIps() []string`

GetLinkPublicIpPublicIps returns the LinkPublicIpPublicIps field if non-nil, zero value otherwise.

### GetLinkPublicIpPublicIpsOk

`func (o *FiltersNic) GetLinkPublicIpPublicIpsOk() (*[]string, bool)`

GetLinkPublicIpPublicIpsOk returns a tuple with the LinkPublicIpPublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIpPublicIps

`func (o *FiltersNic) SetLinkPublicIpPublicIps(v []string)`

SetLinkPublicIpPublicIps sets LinkPublicIpPublicIps field to given value.

### HasLinkPublicIpPublicIps

`func (o *FiltersNic) HasLinkPublicIpPublicIps() bool`

HasLinkPublicIpPublicIps returns a boolean if a field has been set.

### GetMacAddresses

`func (o *FiltersNic) GetMacAddresses() []string`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *FiltersNic) GetMacAddressesOk() (*[]string, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *FiltersNic) SetMacAddresses(v []string)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *FiltersNic) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersNic) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersNic) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersNic) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersNic) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetNicIds

`func (o *FiltersNic) GetNicIds() []string`

GetNicIds returns the NicIds field if non-nil, zero value otherwise.

### GetNicIdsOk

`func (o *FiltersNic) GetNicIdsOk() (*[]string, bool)`

GetNicIdsOk returns a tuple with the NicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicIds

`func (o *FiltersNic) SetNicIds(v []string)`

SetNicIds sets NicIds field to given value.

### HasNicIds

`func (o *FiltersNic) HasNicIds() bool`

HasNicIds returns a boolean if a field has been set.

### GetPrivateDnsNames

`func (o *FiltersNic) GetPrivateDnsNames() []string`

GetPrivateDnsNames returns the PrivateDnsNames field if non-nil, zero value otherwise.

### GetPrivateDnsNamesOk

`func (o *FiltersNic) GetPrivateDnsNamesOk() (*[]string, bool)`

GetPrivateDnsNamesOk returns a tuple with the PrivateDnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsNames

`func (o *FiltersNic) SetPrivateDnsNames(v []string)`

SetPrivateDnsNames sets PrivateDnsNames field to given value.

### HasPrivateDnsNames

`func (o *FiltersNic) HasPrivateDnsNames() bool`

HasPrivateDnsNames returns a boolean if a field has been set.

### GetPrivateIpsLinkPublicIpAccountIds

`func (o *FiltersNic) GetPrivateIpsLinkPublicIpAccountIds() []string`

GetPrivateIpsLinkPublicIpAccountIds returns the PrivateIpsLinkPublicIpAccountIds field if non-nil, zero value otherwise.

### GetPrivateIpsLinkPublicIpAccountIdsOk

`func (o *FiltersNic) GetPrivateIpsLinkPublicIpAccountIdsOk() (*[]string, bool)`

GetPrivateIpsLinkPublicIpAccountIdsOk returns a tuple with the PrivateIpsLinkPublicIpAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpsLinkPublicIpAccountIds

`func (o *FiltersNic) SetPrivateIpsLinkPublicIpAccountIds(v []string)`

SetPrivateIpsLinkPublicIpAccountIds sets PrivateIpsLinkPublicIpAccountIds field to given value.

### HasPrivateIpsLinkPublicIpAccountIds

`func (o *FiltersNic) HasPrivateIpsLinkPublicIpAccountIds() bool`

HasPrivateIpsLinkPublicIpAccountIds returns a boolean if a field has been set.

### GetPrivateIpsLinkPublicIpPublicIps

`func (o *FiltersNic) GetPrivateIpsLinkPublicIpPublicIps() []string`

GetPrivateIpsLinkPublicIpPublicIps returns the PrivateIpsLinkPublicIpPublicIps field if non-nil, zero value otherwise.

### GetPrivateIpsLinkPublicIpPublicIpsOk

`func (o *FiltersNic) GetPrivateIpsLinkPublicIpPublicIpsOk() (*[]string, bool)`

GetPrivateIpsLinkPublicIpPublicIpsOk returns a tuple with the PrivateIpsLinkPublicIpPublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpsLinkPublicIpPublicIps

`func (o *FiltersNic) SetPrivateIpsLinkPublicIpPublicIps(v []string)`

SetPrivateIpsLinkPublicIpPublicIps sets PrivateIpsLinkPublicIpPublicIps field to given value.

### HasPrivateIpsLinkPublicIpPublicIps

`func (o *FiltersNic) HasPrivateIpsLinkPublicIpPublicIps() bool`

HasPrivateIpsLinkPublicIpPublicIps returns a boolean if a field has been set.

### GetPrivateIpsPrimaryIp

`func (o *FiltersNic) GetPrivateIpsPrimaryIp() bool`

GetPrivateIpsPrimaryIp returns the PrivateIpsPrimaryIp field if non-nil, zero value otherwise.

### GetPrivateIpsPrimaryIpOk

`func (o *FiltersNic) GetPrivateIpsPrimaryIpOk() (*bool, bool)`

GetPrivateIpsPrimaryIpOk returns a tuple with the PrivateIpsPrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpsPrimaryIp

`func (o *FiltersNic) SetPrivateIpsPrimaryIp(v bool)`

SetPrivateIpsPrimaryIp sets PrivateIpsPrimaryIp field to given value.

### HasPrivateIpsPrimaryIp

`func (o *FiltersNic) HasPrivateIpsPrimaryIp() bool`

HasPrivateIpsPrimaryIp returns a boolean if a field has been set.

### GetPrivateIpsPrivateIps

`func (o *FiltersNic) GetPrivateIpsPrivateIps() []string`

GetPrivateIpsPrivateIps returns the PrivateIpsPrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsPrivateIpsOk

`func (o *FiltersNic) GetPrivateIpsPrivateIpsOk() (*[]string, bool)`

GetPrivateIpsPrivateIpsOk returns a tuple with the PrivateIpsPrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpsPrivateIps

`func (o *FiltersNic) SetPrivateIpsPrivateIps(v []string)`

SetPrivateIpsPrivateIps sets PrivateIpsPrivateIps field to given value.

### HasPrivateIpsPrivateIps

`func (o *FiltersNic) HasPrivateIpsPrivateIps() bool`

HasPrivateIpsPrivateIps returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *FiltersNic) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *FiltersNic) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *FiltersNic) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *FiltersNic) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSecurityGroupNames

`func (o *FiltersNic) GetSecurityGroupNames() []string`

GetSecurityGroupNames returns the SecurityGroupNames field if non-nil, zero value otherwise.

### GetSecurityGroupNamesOk

`func (o *FiltersNic) GetSecurityGroupNamesOk() (*[]string, bool)`

GetSecurityGroupNamesOk returns a tuple with the SecurityGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupNames

`func (o *FiltersNic) SetSecurityGroupNames(v []string)`

SetSecurityGroupNames sets SecurityGroupNames field to given value.

### HasSecurityGroupNames

`func (o *FiltersNic) HasSecurityGroupNames() bool`

HasSecurityGroupNames returns a boolean if a field has been set.

### GetStates

`func (o *FiltersNic) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersNic) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersNic) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersNic) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetSubnetIds

`func (o *FiltersNic) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *FiltersNic) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *FiltersNic) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *FiltersNic) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### GetSubregionNames

`func (o *FiltersNic) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersNic) GetSubregionNamesOk() (*[]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionNames

`func (o *FiltersNic) SetSubregionNames(v []string)`

SetSubregionNames sets SubregionNames field to given value.

### HasSubregionNames

`func (o *FiltersNic) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersNic) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersNic) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersNic) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersNic) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersNic) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersNic) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersNic) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersNic) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersNic) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersNic) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersNic) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersNic) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


