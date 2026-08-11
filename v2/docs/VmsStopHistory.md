# VmsStopHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateReason** | Pointer to **string** | The reason explaining why the VM stopped. For more information, see [Creating VMs &gt; VM State Reference](https://docs.outscale.com/en/userguide/Creating-VMs.html#_vm_state_reference_statereason_2). | [optional] 
**StopDate** | Pointer to **string** | The date and time (UTC) of the stop event. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 

## Methods

### NewVmsStopHistory

`func NewVmsStopHistory() *VmsStopHistory`

NewVmsStopHistory instantiates a new VmsStopHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmsStopHistoryWithDefaults

`func NewVmsStopHistoryWithDefaults() *VmsStopHistory`

NewVmsStopHistoryWithDefaults instantiates a new VmsStopHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateReason

`func (o *VmsStopHistory) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *VmsStopHistory) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *VmsStopHistory) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *VmsStopHistory) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStopDate

`func (o *VmsStopHistory) GetStopDate() string`

GetStopDate returns the StopDate field if non-nil, zero value otherwise.

### GetStopDateOk

`func (o *VmsStopHistory) GetStopDateOk() (*string, bool)`

GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDate

`func (o *VmsStopHistory) SetStopDate(v string)`

SetStopDate sets StopDate field to given value.

### HasStopDate

`func (o *VmsStopHistory) HasStopDate() bool`

HasStopDate returns a boolean if a field has been set.

### GetVmId

`func (o *VmsStopHistory) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VmsStopHistory) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VmsStopHistory) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *VmsStopHistory) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


