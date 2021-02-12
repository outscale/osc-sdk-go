# DeleteVmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VmIds** | **[]string** | One or more IDs of VMs. | 

## Methods

### NewDeleteVmsRequest

`func NewDeleteVmsRequest(vmIds []string, ) *DeleteVmsRequest`

NewDeleteVmsRequest instantiates a new DeleteVmsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVmsRequestWithDefaults

`func NewDeleteVmsRequestWithDefaults() *DeleteVmsRequest`

NewDeleteVmsRequestWithDefaults instantiates a new DeleteVmsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteVmsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteVmsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteVmsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteVmsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVmIds

`func (o *DeleteVmsRequest) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *DeleteVmsRequest) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *DeleteVmsRequest) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


