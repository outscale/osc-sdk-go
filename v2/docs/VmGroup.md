# VmGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** | The date and time of creation of the VM group. | [optional] 
**Description** | Pointer to **string** | The description of the VM group. | [optional] 
**PositioningStrategy** | Pointer to **string** | The positioning strategy of the VMs on hypervisors. By default, or if set to &#x60;no-strategy&#x60;, TINA determines the most adequate position for the VMs. If set to &#x60;attract&#x60;, the VMs are deployed on the same hypervisor, which improves network performance. If set to &#x60;repulse&#x60;, the VMs are deployed on a different hypervisor, which improves fault tolerance. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security groups for the VM group. | [optional] 
**State** | Pointer to **string** | The state of the VM group (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;scaling up&#x60; \\| &#x60;scaling down&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet for the VM group. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the VM group. | [optional] 
**VmCount** | Pointer to **int32** | The number of VMs in the VM group. | [optional] 
**VmGroupId** | Pointer to **string** | The ID of the VM group. | [optional] 
**VmGroupName** | Pointer to **string** | The name of the VM group. | [optional] 
**VmIds** | Pointer to **[]string** | The IDs of the VMs in the VM group. | [optional] 
**VmTemplateId** | Pointer to **string** | The ID of the VM template used by the VM group. | [optional] 

## Methods

### NewVmGroup

`func NewVmGroup() *VmGroup`

NewVmGroup instantiates a new VmGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmGroupWithDefaults

`func NewVmGroupWithDefaults() *VmGroup`

NewVmGroupWithDefaults instantiates a new VmGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *VmGroup) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *VmGroup) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *VmGroup) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *VmGroup) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *VmGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VmGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VmGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VmGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPositioningStrategy

`func (o *VmGroup) GetPositioningStrategy() string`

GetPositioningStrategy returns the PositioningStrategy field if non-nil, zero value otherwise.

### GetPositioningStrategyOk

`func (o *VmGroup) GetPositioningStrategyOk() (*string, bool)`

GetPositioningStrategyOk returns a tuple with the PositioningStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningStrategy

`func (o *VmGroup) SetPositioningStrategy(v string)`

SetPositioningStrategy sets PositioningStrategy field to given value.

### HasPositioningStrategy

`func (o *VmGroup) HasPositioningStrategy() bool`

HasPositioningStrategy returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *VmGroup) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *VmGroup) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *VmGroup) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *VmGroup) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetState

`func (o *VmGroup) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VmGroup) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VmGroup) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VmGroup) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubnetId

`func (o *VmGroup) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *VmGroup) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *VmGroup) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *VmGroup) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTags

`func (o *VmGroup) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VmGroup) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VmGroup) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VmGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmCount

`func (o *VmGroup) GetVmCount() int32`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *VmGroup) GetVmCountOk() (*int32, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *VmGroup) SetVmCount(v int32)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *VmGroup) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetVmGroupId

`func (o *VmGroup) GetVmGroupId() string`

GetVmGroupId returns the VmGroupId field if non-nil, zero value otherwise.

### GetVmGroupIdOk

`func (o *VmGroup) GetVmGroupIdOk() (*string, bool)`

GetVmGroupIdOk returns a tuple with the VmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroupId

`func (o *VmGroup) SetVmGroupId(v string)`

SetVmGroupId sets VmGroupId field to given value.

### HasVmGroupId

`func (o *VmGroup) HasVmGroupId() bool`

HasVmGroupId returns a boolean if a field has been set.

### GetVmGroupName

`func (o *VmGroup) GetVmGroupName() string`

GetVmGroupName returns the VmGroupName field if non-nil, zero value otherwise.

### GetVmGroupNameOk

`func (o *VmGroup) GetVmGroupNameOk() (*string, bool)`

GetVmGroupNameOk returns a tuple with the VmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroupName

`func (o *VmGroup) SetVmGroupName(v string)`

SetVmGroupName sets VmGroupName field to given value.

### HasVmGroupName

`func (o *VmGroup) HasVmGroupName() bool`

HasVmGroupName returns a boolean if a field has been set.

### GetVmIds

`func (o *VmGroup) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *VmGroup) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *VmGroup) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *VmGroup) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### GetVmTemplateId

`func (o *VmGroup) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *VmGroup) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *VmGroup) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.

### HasVmTemplateId

`func (o *VmGroup) HasVmTemplateId() bool`

HasVmTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


