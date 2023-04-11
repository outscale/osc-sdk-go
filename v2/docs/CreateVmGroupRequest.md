# CreateVmGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description for the VM group. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PositioningStrategy** | Pointer to **string** | The positioning strategy of VMs on hypervisors. By default, or if set to &#x60;no-strategy&#x60; our orchestrator determines the most adequate position for your VMs. If set to &#x60;attract&#x60;, your VMs are deployed on the same hypervisor, which improves network performance. If set to &#x60;repulse&#x60;, your VMs are deployed on a different hypervisor, which improves fault tolerance. | [optional] [default to "no-strategy"]
**SecurityGroupIds** | **[]string** | One or more IDs of security groups for the VM group. | 
**SubnetId** | **string** | The ID of the Subnet in which you want to create the VM group. | 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags to add to the VM group. | [optional] 
**VmCount** | **int32** | The number of VMs deployed in the VM group. | 
**VmGroupName** | **string** | The name of the VM group. | 
**VmTemplateId** | **string** | The ID of the VM template used to launch VMs in the VM group. | 

## Methods

### NewCreateVmGroupRequest

`func NewCreateVmGroupRequest(securityGroupIds []string, subnetId string, vmCount int32, vmGroupName string, vmTemplateId string, ) *CreateVmGroupRequest`

NewCreateVmGroupRequest instantiates a new CreateVmGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVmGroupRequestWithDefaults

`func NewCreateVmGroupRequestWithDefaults() *CreateVmGroupRequest`

NewCreateVmGroupRequestWithDefaults instantiates a new CreateVmGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateVmGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVmGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVmGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVmGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateVmGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVmGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateVmGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateVmGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPositioningStrategy

`func (o *CreateVmGroupRequest) GetPositioningStrategy() string`

GetPositioningStrategy returns the PositioningStrategy field if non-nil, zero value otherwise.

### GetPositioningStrategyOk

`func (o *CreateVmGroupRequest) GetPositioningStrategyOk() (*string, bool)`

GetPositioningStrategyOk returns a tuple with the PositioningStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningStrategy

`func (o *CreateVmGroupRequest) SetPositioningStrategy(v string)`

SetPositioningStrategy sets PositioningStrategy field to given value.

### HasPositioningStrategy

`func (o *CreateVmGroupRequest) HasPositioningStrategy() bool`

HasPositioningStrategy returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *CreateVmGroupRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *CreateVmGroupRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *CreateVmGroupRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.


### GetSubnetId

`func (o *CreateVmGroupRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateVmGroupRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateVmGroupRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *CreateVmGroupRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateVmGroupRequest) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateVmGroupRequest) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateVmGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmCount

`func (o *CreateVmGroupRequest) GetVmCount() int32`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *CreateVmGroupRequest) GetVmCountOk() (*int32, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *CreateVmGroupRequest) SetVmCount(v int32)`

SetVmCount sets VmCount field to given value.


### GetVmGroupName

`func (o *CreateVmGroupRequest) GetVmGroupName() string`

GetVmGroupName returns the VmGroupName field if non-nil, zero value otherwise.

### GetVmGroupNameOk

`func (o *CreateVmGroupRequest) GetVmGroupNameOk() (*string, bool)`

GetVmGroupNameOk returns a tuple with the VmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroupName

`func (o *CreateVmGroupRequest) SetVmGroupName(v string)`

SetVmGroupName sets VmGroupName field to given value.


### GetVmTemplateId

`func (o *CreateVmGroupRequest) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *CreateVmGroupRequest) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *CreateVmGroupRequest) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


