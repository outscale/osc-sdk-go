# FiltersSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptions** | Pointer to **[]string** | The descriptions of the security groups. | [optional] 
**InboundRuleAccountIds** | Pointer to **[]string** | The account IDs that have been granted permissions. | [optional] 
**InboundRuleFromPortRanges** | Pointer to **[]int32** | The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers. | [optional] 
**InboundRuleIpRanges** | Pointer to **[]string** | The IP ranges that have been granted permissions, in CIDR notation (for example, &#x60;10.0.0.0/24&#x60;). | [optional] 
**InboundRuleProtocols** | Pointer to **[]string** | The IP protocols for the permissions (&#x60;tcp&#x60; \\| &#x60;udp&#x60; \\| &#x60;icmp&#x60;, or a protocol number, or &#x60;-1&#x60; for all protocols). | [optional] 
**InboundRuleSecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups that have been granted permissions. | [optional] 
**InboundRuleSecurityGroupNames** | Pointer to **[]string** | The names of the security groups that have been granted permissions. | [optional] 
**InboundRuleToPortRanges** | Pointer to **[]int32** | The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets specified when the security groups were created. | [optional] 
**OutboundRuleAccountIds** | Pointer to **[]string** | The account IDs that have been granted permissions. | [optional] 
**OutboundRuleFromPortRanges** | Pointer to **[]int32** | The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers. | [optional] 
**OutboundRuleIpRanges** | Pointer to **[]string** | The IP ranges that have been granted permissions, in CIDR notation (for example, &#x60;10.0.0.0/24&#x60;). | [optional] 
**OutboundRuleProtocols** | Pointer to **[]string** | The IP protocols for the permissions (&#x60;tcp&#x60; \\| &#x60;udp&#x60; \\| &#x60;icmp&#x60;, or a protocol number, or &#x60;-1&#x60; for all protocols). | [optional] 
**OutboundRuleSecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups that have been granted permissions. | [optional] 
**OutboundRuleSecurityGroupNames** | Pointer to **[]string** | The names of the security groups that have been granted permissions. | [optional] 
**OutboundRuleToPortRanges** | Pointer to **[]int32** | The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups. | [optional] 
**SecurityGroupNames** | Pointer to **[]string** | The names of the security groups. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the security groups. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the security groups. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the security groups, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

## Methods

### NewFiltersSecurityGroup

`func NewFiltersSecurityGroup() *FiltersSecurityGroup`

NewFiltersSecurityGroup instantiates a new FiltersSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersSecurityGroupWithDefaults

`func NewFiltersSecurityGroupWithDefaults() *FiltersSecurityGroup`

NewFiltersSecurityGroupWithDefaults instantiates a new FiltersSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptions

`func (o *FiltersSecurityGroup) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersSecurityGroup) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *FiltersSecurityGroup) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *FiltersSecurityGroup) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetInboundRuleAccountIds

`func (o *FiltersSecurityGroup) GetInboundRuleAccountIds() []string`

GetInboundRuleAccountIds returns the InboundRuleAccountIds field if non-nil, zero value otherwise.

### GetInboundRuleAccountIdsOk

`func (o *FiltersSecurityGroup) GetInboundRuleAccountIdsOk() (*[]string, bool)`

GetInboundRuleAccountIdsOk returns a tuple with the InboundRuleAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRuleAccountIds

`func (o *FiltersSecurityGroup) SetInboundRuleAccountIds(v []string)`

SetInboundRuleAccountIds sets InboundRuleAccountIds field to given value.

### HasInboundRuleAccountIds

`func (o *FiltersSecurityGroup) HasInboundRuleAccountIds() bool`

HasInboundRuleAccountIds returns a boolean if a field has been set.

### GetInboundRuleFromPortRanges

`func (o *FiltersSecurityGroup) GetInboundRuleFromPortRanges() []int32`

GetInboundRuleFromPortRanges returns the InboundRuleFromPortRanges field if non-nil, zero value otherwise.

### GetInboundRuleFromPortRangesOk

`func (o *FiltersSecurityGroup) GetInboundRuleFromPortRangesOk() (*[]int32, bool)`

GetInboundRuleFromPortRangesOk returns a tuple with the InboundRuleFromPortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRuleFromPortRanges

`func (o *FiltersSecurityGroup) SetInboundRuleFromPortRanges(v []int32)`

SetInboundRuleFromPortRanges sets InboundRuleFromPortRanges field to given value.

### HasInboundRuleFromPortRanges

`func (o *FiltersSecurityGroup) HasInboundRuleFromPortRanges() bool`

HasInboundRuleFromPortRanges returns a boolean if a field has been set.

### GetInboundRuleIpRanges

`func (o *FiltersSecurityGroup) GetInboundRuleIpRanges() []string`

GetInboundRuleIpRanges returns the InboundRuleIpRanges field if non-nil, zero value otherwise.

### GetInboundRuleIpRangesOk

`func (o *FiltersSecurityGroup) GetInboundRuleIpRangesOk() (*[]string, bool)`

GetInboundRuleIpRangesOk returns a tuple with the InboundRuleIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRuleIpRanges

`func (o *FiltersSecurityGroup) SetInboundRuleIpRanges(v []string)`

SetInboundRuleIpRanges sets InboundRuleIpRanges field to given value.

### HasInboundRuleIpRanges

`func (o *FiltersSecurityGroup) HasInboundRuleIpRanges() bool`

HasInboundRuleIpRanges returns a boolean if a field has been set.

### GetInboundRuleProtocols

`func (o *FiltersSecurityGroup) GetInboundRuleProtocols() []string`

GetInboundRuleProtocols returns the InboundRuleProtocols field if non-nil, zero value otherwise.

### GetInboundRuleProtocolsOk

`func (o *FiltersSecurityGroup) GetInboundRuleProtocolsOk() (*[]string, bool)`

GetInboundRuleProtocolsOk returns a tuple with the InboundRuleProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRuleProtocols

`func (o *FiltersSecurityGroup) SetInboundRuleProtocols(v []string)`

SetInboundRuleProtocols sets InboundRuleProtocols field to given value.

### HasInboundRuleProtocols

`func (o *FiltersSecurityGroup) HasInboundRuleProtocols() bool`

HasInboundRuleProtocols returns a boolean if a field has been set.

### GetInboundRuleSecurityGroupIds

`func (o *FiltersSecurityGroup) GetInboundRuleSecurityGroupIds() []string`

GetInboundRuleSecurityGroupIds returns the InboundRuleSecurityGroupIds field if non-nil, zero value otherwise.

### GetInboundRuleSecurityGroupIdsOk

`func (o *FiltersSecurityGroup) GetInboundRuleSecurityGroupIdsOk() (*[]string, bool)`

GetInboundRuleSecurityGroupIdsOk returns a tuple with the InboundRuleSecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRuleSecurityGroupIds

`func (o *FiltersSecurityGroup) SetInboundRuleSecurityGroupIds(v []string)`

SetInboundRuleSecurityGroupIds sets InboundRuleSecurityGroupIds field to given value.

### HasInboundRuleSecurityGroupIds

`func (o *FiltersSecurityGroup) HasInboundRuleSecurityGroupIds() bool`

HasInboundRuleSecurityGroupIds returns a boolean if a field has been set.

### GetInboundRuleSecurityGroupNames

`func (o *FiltersSecurityGroup) GetInboundRuleSecurityGroupNames() []string`

GetInboundRuleSecurityGroupNames returns the InboundRuleSecurityGroupNames field if non-nil, zero value otherwise.

### GetInboundRuleSecurityGroupNamesOk

`func (o *FiltersSecurityGroup) GetInboundRuleSecurityGroupNamesOk() (*[]string, bool)`

GetInboundRuleSecurityGroupNamesOk returns a tuple with the InboundRuleSecurityGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRuleSecurityGroupNames

`func (o *FiltersSecurityGroup) SetInboundRuleSecurityGroupNames(v []string)`

SetInboundRuleSecurityGroupNames sets InboundRuleSecurityGroupNames field to given value.

### HasInboundRuleSecurityGroupNames

`func (o *FiltersSecurityGroup) HasInboundRuleSecurityGroupNames() bool`

HasInboundRuleSecurityGroupNames returns a boolean if a field has been set.

### GetInboundRuleToPortRanges

`func (o *FiltersSecurityGroup) GetInboundRuleToPortRanges() []int32`

GetInboundRuleToPortRanges returns the InboundRuleToPortRanges field if non-nil, zero value otherwise.

### GetInboundRuleToPortRangesOk

`func (o *FiltersSecurityGroup) GetInboundRuleToPortRangesOk() (*[]int32, bool)`

GetInboundRuleToPortRangesOk returns a tuple with the InboundRuleToPortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRuleToPortRanges

`func (o *FiltersSecurityGroup) SetInboundRuleToPortRanges(v []int32)`

SetInboundRuleToPortRanges sets InboundRuleToPortRanges field to given value.

### HasInboundRuleToPortRanges

`func (o *FiltersSecurityGroup) HasInboundRuleToPortRanges() bool`

HasInboundRuleToPortRanges returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersSecurityGroup) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersSecurityGroup) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersSecurityGroup) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersSecurityGroup) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetOutboundRuleAccountIds

`func (o *FiltersSecurityGroup) GetOutboundRuleAccountIds() []string`

GetOutboundRuleAccountIds returns the OutboundRuleAccountIds field if non-nil, zero value otherwise.

### GetOutboundRuleAccountIdsOk

`func (o *FiltersSecurityGroup) GetOutboundRuleAccountIdsOk() (*[]string, bool)`

GetOutboundRuleAccountIdsOk returns a tuple with the OutboundRuleAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundRuleAccountIds

`func (o *FiltersSecurityGroup) SetOutboundRuleAccountIds(v []string)`

SetOutboundRuleAccountIds sets OutboundRuleAccountIds field to given value.

### HasOutboundRuleAccountIds

`func (o *FiltersSecurityGroup) HasOutboundRuleAccountIds() bool`

HasOutboundRuleAccountIds returns a boolean if a field has been set.

### GetOutboundRuleFromPortRanges

`func (o *FiltersSecurityGroup) GetOutboundRuleFromPortRanges() []int32`

GetOutboundRuleFromPortRanges returns the OutboundRuleFromPortRanges field if non-nil, zero value otherwise.

### GetOutboundRuleFromPortRangesOk

`func (o *FiltersSecurityGroup) GetOutboundRuleFromPortRangesOk() (*[]int32, bool)`

GetOutboundRuleFromPortRangesOk returns a tuple with the OutboundRuleFromPortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundRuleFromPortRanges

`func (o *FiltersSecurityGroup) SetOutboundRuleFromPortRanges(v []int32)`

SetOutboundRuleFromPortRanges sets OutboundRuleFromPortRanges field to given value.

### HasOutboundRuleFromPortRanges

`func (o *FiltersSecurityGroup) HasOutboundRuleFromPortRanges() bool`

HasOutboundRuleFromPortRanges returns a boolean if a field has been set.

### GetOutboundRuleIpRanges

`func (o *FiltersSecurityGroup) GetOutboundRuleIpRanges() []string`

GetOutboundRuleIpRanges returns the OutboundRuleIpRanges field if non-nil, zero value otherwise.

### GetOutboundRuleIpRangesOk

`func (o *FiltersSecurityGroup) GetOutboundRuleIpRangesOk() (*[]string, bool)`

GetOutboundRuleIpRangesOk returns a tuple with the OutboundRuleIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundRuleIpRanges

`func (o *FiltersSecurityGroup) SetOutboundRuleIpRanges(v []string)`

SetOutboundRuleIpRanges sets OutboundRuleIpRanges field to given value.

### HasOutboundRuleIpRanges

`func (o *FiltersSecurityGroup) HasOutboundRuleIpRanges() bool`

HasOutboundRuleIpRanges returns a boolean if a field has been set.

### GetOutboundRuleProtocols

`func (o *FiltersSecurityGroup) GetOutboundRuleProtocols() []string`

GetOutboundRuleProtocols returns the OutboundRuleProtocols field if non-nil, zero value otherwise.

### GetOutboundRuleProtocolsOk

`func (o *FiltersSecurityGroup) GetOutboundRuleProtocolsOk() (*[]string, bool)`

GetOutboundRuleProtocolsOk returns a tuple with the OutboundRuleProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundRuleProtocols

`func (o *FiltersSecurityGroup) SetOutboundRuleProtocols(v []string)`

SetOutboundRuleProtocols sets OutboundRuleProtocols field to given value.

### HasOutboundRuleProtocols

`func (o *FiltersSecurityGroup) HasOutboundRuleProtocols() bool`

HasOutboundRuleProtocols returns a boolean if a field has been set.

### GetOutboundRuleSecurityGroupIds

`func (o *FiltersSecurityGroup) GetOutboundRuleSecurityGroupIds() []string`

GetOutboundRuleSecurityGroupIds returns the OutboundRuleSecurityGroupIds field if non-nil, zero value otherwise.

### GetOutboundRuleSecurityGroupIdsOk

`func (o *FiltersSecurityGroup) GetOutboundRuleSecurityGroupIdsOk() (*[]string, bool)`

GetOutboundRuleSecurityGroupIdsOk returns a tuple with the OutboundRuleSecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundRuleSecurityGroupIds

`func (o *FiltersSecurityGroup) SetOutboundRuleSecurityGroupIds(v []string)`

SetOutboundRuleSecurityGroupIds sets OutboundRuleSecurityGroupIds field to given value.

### HasOutboundRuleSecurityGroupIds

`func (o *FiltersSecurityGroup) HasOutboundRuleSecurityGroupIds() bool`

HasOutboundRuleSecurityGroupIds returns a boolean if a field has been set.

### GetOutboundRuleSecurityGroupNames

`func (o *FiltersSecurityGroup) GetOutboundRuleSecurityGroupNames() []string`

GetOutboundRuleSecurityGroupNames returns the OutboundRuleSecurityGroupNames field if non-nil, zero value otherwise.

### GetOutboundRuleSecurityGroupNamesOk

`func (o *FiltersSecurityGroup) GetOutboundRuleSecurityGroupNamesOk() (*[]string, bool)`

GetOutboundRuleSecurityGroupNamesOk returns a tuple with the OutboundRuleSecurityGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundRuleSecurityGroupNames

`func (o *FiltersSecurityGroup) SetOutboundRuleSecurityGroupNames(v []string)`

SetOutboundRuleSecurityGroupNames sets OutboundRuleSecurityGroupNames field to given value.

### HasOutboundRuleSecurityGroupNames

`func (o *FiltersSecurityGroup) HasOutboundRuleSecurityGroupNames() bool`

HasOutboundRuleSecurityGroupNames returns a boolean if a field has been set.

### GetOutboundRuleToPortRanges

`func (o *FiltersSecurityGroup) GetOutboundRuleToPortRanges() []int32`

GetOutboundRuleToPortRanges returns the OutboundRuleToPortRanges field if non-nil, zero value otherwise.

### GetOutboundRuleToPortRangesOk

`func (o *FiltersSecurityGroup) GetOutboundRuleToPortRangesOk() (*[]int32, bool)`

GetOutboundRuleToPortRangesOk returns a tuple with the OutboundRuleToPortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundRuleToPortRanges

`func (o *FiltersSecurityGroup) SetOutboundRuleToPortRanges(v []int32)`

SetOutboundRuleToPortRanges sets OutboundRuleToPortRanges field to given value.

### HasOutboundRuleToPortRanges

`func (o *FiltersSecurityGroup) HasOutboundRuleToPortRanges() bool`

HasOutboundRuleToPortRanges returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *FiltersSecurityGroup) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *FiltersSecurityGroup) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *FiltersSecurityGroup) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *FiltersSecurityGroup) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSecurityGroupNames

`func (o *FiltersSecurityGroup) GetSecurityGroupNames() []string`

GetSecurityGroupNames returns the SecurityGroupNames field if non-nil, zero value otherwise.

### GetSecurityGroupNamesOk

`func (o *FiltersSecurityGroup) GetSecurityGroupNamesOk() (*[]string, bool)`

GetSecurityGroupNamesOk returns a tuple with the SecurityGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupNames

`func (o *FiltersSecurityGroup) SetSecurityGroupNames(v []string)`

SetSecurityGroupNames sets SecurityGroupNames field to given value.

### HasSecurityGroupNames

`func (o *FiltersSecurityGroup) HasSecurityGroupNames() bool`

HasSecurityGroupNames returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersSecurityGroup) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersSecurityGroup) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersSecurityGroup) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersSecurityGroup) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersSecurityGroup) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersSecurityGroup) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersSecurityGroup) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersSecurityGroup) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersSecurityGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersSecurityGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersSecurityGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersSecurityGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


