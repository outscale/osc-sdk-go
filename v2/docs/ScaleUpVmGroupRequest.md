# ScaleUpVmGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VmAddition** | **int32** | The number of VMs you want to add to the VM group. | 
**VmGroupId** | **string** | The ID of the VM group you want to scale up. | 

## Methods

### NewScaleUpVmGroupRequest

`func NewScaleUpVmGroupRequest(vmAddition int32, vmGroupId string, ) *ScaleUpVmGroupRequest`

NewScaleUpVmGroupRequest instantiates a new ScaleUpVmGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaleUpVmGroupRequestWithDefaults

`func NewScaleUpVmGroupRequestWithDefaults() *ScaleUpVmGroupRequest`

NewScaleUpVmGroupRequestWithDefaults instantiates a new ScaleUpVmGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ScaleUpVmGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ScaleUpVmGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ScaleUpVmGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ScaleUpVmGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVmAddition

`func (o *ScaleUpVmGroupRequest) GetVmAddition() int32`

GetVmAddition returns the VmAddition field if non-nil, zero value otherwise.

### GetVmAdditionOk

`func (o *ScaleUpVmGroupRequest) GetVmAdditionOk() (*int32, bool)`

GetVmAdditionOk returns a tuple with the VmAddition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmAddition

`func (o *ScaleUpVmGroupRequest) SetVmAddition(v int32)`

SetVmAddition sets VmAddition field to given value.


### GetVmGroupId

`func (o *ScaleUpVmGroupRequest) GetVmGroupId() string`

GetVmGroupId returns the VmGroupId field if non-nil, zero value otherwise.

### GetVmGroupIdOk

`func (o *ScaleUpVmGroupRequest) GetVmGroupIdOk() (*string, bool)`

GetVmGroupIdOk returns a tuple with the VmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroupId

`func (o *ScaleUpVmGroupRequest) SetVmGroupId(v string)`

SetVmGroupId sets VmGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


