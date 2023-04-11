# DeleteVmGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VmGroupId** | **string** | The ID of the VM group you want to delete. | 

## Methods

### NewDeleteVmGroupRequest

`func NewDeleteVmGroupRequest(vmGroupId string, ) *DeleteVmGroupRequest`

NewDeleteVmGroupRequest instantiates a new DeleteVmGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVmGroupRequestWithDefaults

`func NewDeleteVmGroupRequestWithDefaults() *DeleteVmGroupRequest`

NewDeleteVmGroupRequestWithDefaults instantiates a new DeleteVmGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteVmGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteVmGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteVmGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteVmGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVmGroupId

`func (o *DeleteVmGroupRequest) GetVmGroupId() string`

GetVmGroupId returns the VmGroupId field if non-nil, zero value otherwise.

### GetVmGroupIdOk

`func (o *DeleteVmGroupRequest) GetVmGroupIdOk() (*string, bool)`

GetVmGroupIdOk returns a tuple with the VmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmGroupId

`func (o *DeleteVmGroupRequest) SetVmGroupId(v string)`

SetVmGroupId sets VmGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


