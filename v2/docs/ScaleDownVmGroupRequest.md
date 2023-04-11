# ScaleDownVmGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VmGroupId** | **string** | The ID of the VM group you want to scale down. | 
**VmSubtraction** | **int32** | The number of VMs you want to delete from the VM group. | 

## Methods

### NewScaleDownVmGroupRequest

`func NewScaleDownVmGroupRequest(vmGroupId string, vmSubtraction int32, ) *ScaleDownVmGroupRequest`

NewScaleDownVmGroupRequest instantiates a new ScaleDownVmGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaleDownVmGroupRequestWithDefaults

`func NewScaleDownVmGroupRequestWithDefaults() *ScaleDownVmGroupRequest`

NewScaleDownVmGroupRequestWithDefaults instantiates a new ScaleDownVmGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ScaleDownVmGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ScaleDownVmGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ScaleDownVmGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ScaleDownVmGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVmGroupId

`func (o *ScaleDownVmGroupRequest) GetVmGroupId() string`

GetVmGroupId returns the VmGroupId field if non-nil, zero value otherwise.

### GetVmGroupIdOk

`func (o *ScaleDownVmGroupRequest) GetVmGroupIdOk() (*string, bool)`

GetVmGroupIdOk returns a tuple with the VmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroupId

`func (o *ScaleDownVmGroupRequest) SetVmGroupId(v string)`

SetVmGroupId sets VmGroupId field to given value.


### GetVmSubtraction

`func (o *ScaleDownVmGroupRequest) GetVmSubtraction() int32`

GetVmSubtraction returns the VmSubtraction field if non-nil, zero value otherwise.

### GetVmSubtractionOk

`func (o *ScaleDownVmGroupRequest) GetVmSubtractionOk() (*int32, bool)`

GetVmSubtractionOk returns a tuple with the VmSubtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSubtraction

`func (o *ScaleDownVmGroupRequest) SetVmSubtraction(v int32)`

SetVmSubtraction sets VmSubtraction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


