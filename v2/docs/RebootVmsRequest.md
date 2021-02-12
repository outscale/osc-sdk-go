# RebootVmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**VmIds** | **[]string** | One or more IDs of the VMs you want to reboot. | 

## Methods

### NewRebootVmsRequest

`func NewRebootVmsRequest(vmIds []string, ) *RebootVmsRequest`

NewRebootVmsRequest instantiates a new RebootVmsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootVmsRequestWithDefaults

`func NewRebootVmsRequestWithDefaults() *RebootVmsRequest`

NewRebootVmsRequestWithDefaults instantiates a new RebootVmsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *RebootVmsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *RebootVmsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *RebootVmsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *RebootVmsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVmIds

`func (o *RebootVmsRequest) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *RebootVmsRequest) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *RebootVmsRequest) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


