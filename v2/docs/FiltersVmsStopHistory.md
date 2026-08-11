# FiltersVmsStopHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateReasons** | Pointer to **[]string** | The reason explaining why the VM stopped. You can filter by reason code or reason prefix (for example, &#x60;Client.ApiGracefulShutdown&#x60; or &#x60;Client.*&#x60;). For the list of reason codes, see [Creating VMs &gt; VM State Reference](https://docs.outscale.com/en/userguide/Creating-VMs). | [optional] 
**StopDateAfter** | Pointer to **string** | The date and time (UTC), or the date, after which you want to retrieve VM stops, in ISO 8601 format (for example, &#x60;2026-06-14T00:00:00.000Z&#x60; or &#x60;2026-06-14&#x60;). | [optional] 
**StopDateBefore** | Pointer to **string** | The date and time (UTC), or the date, before which you want to retrieve VM stops, in ISO 8601 format (for example, &#x60;2026-06-14T00:00:00.000Z&#x60; or &#x60;2026-06-14&#x60;). | [optional] 
**VmIds** | Pointer to **[]string** | The IDs of the stopped VM(s). | [optional] 

## Methods

### NewFiltersVmsStopHistory

`func NewFiltersVmsStopHistory() *FiltersVmsStopHistory`

NewFiltersVmsStopHistory instantiates a new FiltersVmsStopHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersVmsStopHistoryWithDefaults

`func NewFiltersVmsStopHistoryWithDefaults() *FiltersVmsStopHistory`

NewFiltersVmsStopHistoryWithDefaults instantiates a new FiltersVmsStopHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateReasons

`func (o *FiltersVmsStopHistory) GetStateReasons() []string`

GetStateReasons returns the StateReasons field if non-nil, zero value otherwise.

### GetStateReasonsOk

`func (o *FiltersVmsStopHistory) GetStateReasonsOk() (*[]string, bool)`

GetStateReasonsOk returns a tuple with the StateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReasons

`func (o *FiltersVmsStopHistory) SetStateReasons(v []string)`

SetStateReasons sets StateReasons field to given value.

### HasStateReasons

`func (o *FiltersVmsStopHistory) HasStateReasons() bool`

HasStateReasons returns a boolean if a field has been set.

### GetStopDateAfter

`func (o *FiltersVmsStopHistory) GetStopDateAfter() string`

GetStopDateAfter returns the StopDateAfter field if non-nil, zero value otherwise.

### GetStopDateAfterOk

`func (o *FiltersVmsStopHistory) GetStopDateAfterOk() (*string, bool)`

GetStopDateAfterOk returns a tuple with the StopDateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDateAfter

`func (o *FiltersVmsStopHistory) SetStopDateAfter(v string)`

SetStopDateAfter sets StopDateAfter field to given value.

### HasStopDateAfter

`func (o *FiltersVmsStopHistory) HasStopDateAfter() bool`

HasStopDateAfter returns a boolean if a field has been set.

### GetStopDateBefore

`func (o *FiltersVmsStopHistory) GetStopDateBefore() string`

GetStopDateBefore returns the StopDateBefore field if non-nil, zero value otherwise.

### GetStopDateBeforeOk

`func (o *FiltersVmsStopHistory) GetStopDateBeforeOk() (*string, bool)`

GetStopDateBeforeOk returns a tuple with the StopDateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDateBefore

`func (o *FiltersVmsStopHistory) SetStopDateBefore(v string)`

SetStopDateBefore sets StopDateBefore field to given value.

### HasStopDateBefore

`func (o *FiltersVmsStopHistory) HasStopDateBefore() bool`

HasStopDateBefore returns a boolean if a field has been set.

### GetVmIds

`func (o *FiltersVmsStopHistory) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *FiltersVmsStopHistory) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *FiltersVmsStopHistory) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *FiltersVmsStopHistory) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


