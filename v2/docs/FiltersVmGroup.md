# FiltersVmGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptions** | Pointer to **[]string** | The descriptions of the VM groups. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | The IDs of the security groups. | [optional] 
**SubnetIds** | Pointer to **[]string** | The IDs of the Subnets. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the VM groups. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the VM groups. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the VMs, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 
**VmCounts** | Pointer to **[]int32** | The number of VMs in the VM group. | [optional] 
**VmGroupIds** | Pointer to **[]string** | The IDs of the VM groups. | [optional] 
**VmGroupNames** | Pointer to **[]string** | The names of the VM groups. | [optional] 
**VmTemplateIds** | Pointer to **[]string** | The IDs of the VM templates. | [optional] 

## Methods

### NewFiltersVmGroup

`func NewFiltersVmGroup() *FiltersVmGroup`

NewFiltersVmGroup instantiates a new FiltersVmGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersVmGroupWithDefaults

`func NewFiltersVmGroupWithDefaults() *FiltersVmGroup`

NewFiltersVmGroupWithDefaults instantiates a new FiltersVmGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptions

`func (o *FiltersVmGroup) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersVmGroup) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *FiltersVmGroup) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *FiltersVmGroup) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *FiltersVmGroup) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *FiltersVmGroup) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *FiltersVmGroup) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *FiltersVmGroup) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSubnetIds

`func (o *FiltersVmGroup) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *FiltersVmGroup) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *FiltersVmGroup) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *FiltersVmGroup) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersVmGroup) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersVmGroup) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersVmGroup) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersVmGroup) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersVmGroup) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersVmGroup) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersVmGroup) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersVmGroup) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersVmGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersVmGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersVmGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersVmGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmCounts

`func (o *FiltersVmGroup) GetVmCounts() []int32`

GetVmCounts returns the VmCounts field if non-nil, zero value otherwise.

### GetVmCountsOk

`func (o *FiltersVmGroup) GetVmCountsOk() (*[]int32, bool)`

GetVmCountsOk returns a tuple with the VmCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCounts

`func (o *FiltersVmGroup) SetVmCounts(v []int32)`

SetVmCounts sets VmCounts field to given value.

### HasVmCounts

`func (o *FiltersVmGroup) HasVmCounts() bool`

HasVmCounts returns a boolean if a field has been set.

### GetVmGroupIds

`func (o *FiltersVmGroup) GetVmGroupIds() []string`

GetVmGroupIds returns the VmGroupIds field if non-nil, zero value otherwise.

### GetVmGroupIdsOk

`func (o *FiltersVmGroup) GetVmGroupIdsOk() (*[]string, bool)`

GetVmGroupIdsOk returns a tuple with the VmGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroupIds

`func (o *FiltersVmGroup) SetVmGroupIds(v []string)`

SetVmGroupIds sets VmGroupIds field to given value.

### HasVmGroupIds

`func (o *FiltersVmGroup) HasVmGroupIds() bool`

HasVmGroupIds returns a boolean if a field has been set.

### GetVmGroupNames

`func (o *FiltersVmGroup) GetVmGroupNames() []string`

GetVmGroupNames returns the VmGroupNames field if non-nil, zero value otherwise.

### GetVmGroupNamesOk

`func (o *FiltersVmGroup) GetVmGroupNamesOk() (*[]string, bool)`

GetVmGroupNamesOk returns a tuple with the VmGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroupNames

`func (o *FiltersVmGroup) SetVmGroupNames(v []string)`

SetVmGroupNames sets VmGroupNames field to given value.

### HasVmGroupNames

`func (o *FiltersVmGroup) HasVmGroupNames() bool`

HasVmGroupNames returns a boolean if a field has been set.

### GetVmTemplateIds

`func (o *FiltersVmGroup) GetVmTemplateIds() []string`

GetVmTemplateIds returns the VmTemplateIds field if non-nil, zero value otherwise.

### GetVmTemplateIdsOk

`func (o *FiltersVmGroup) GetVmTemplateIdsOk() (*[]string, bool)`

GetVmTemplateIdsOk returns a tuple with the VmTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateIds

`func (o *FiltersVmGroup) SetVmTemplateIds(v []string)`

SetVmTemplateIds sets VmTemplateIds field to given value.

### HasVmTemplateIds

`func (o *FiltersVmGroup) HasVmTemplateIds() bool`

HasVmTemplateIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


