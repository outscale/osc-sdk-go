# FiltersVmsState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceEventCodes** | Pointer to **[]string** | The code for the scheduled event (&#x60;system-reboot&#x60; | &#x60;system-maintenance&#x60;). | [optional] 
**MaintenanceEventDescriptions** | Pointer to **[]string** | The description of the scheduled event. | [optional] 
**MaintenanceEventsNotAfter** | Pointer to **[]string** | The latest time the event can end. | [optional] 
**MaintenanceEventsNotBefore** | Pointer to **[]string** | The earliest time the event can start. | [optional] 
**SubregionNames** | Pointer to **[]string** | The names of the Subregions of the VMs. | [optional] 
**VmIds** | Pointer to **[]string** | One or more IDs of VMs. | [optional] 
**VmStates** | Pointer to **[]string** | The states of the VMs (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 

## Methods

### NewFiltersVmsState

`func NewFiltersVmsState() *FiltersVmsState`

NewFiltersVmsState instantiates a new FiltersVmsState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersVmsStateWithDefaults

`func NewFiltersVmsStateWithDefaults() *FiltersVmsState`

NewFiltersVmsStateWithDefaults instantiates a new FiltersVmsState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceEventCodes

`func (o *FiltersVmsState) GetMaintenanceEventCodes() []string`

GetMaintenanceEventCodes returns the MaintenanceEventCodes field if non-nil, zero value otherwise.

### GetMaintenanceEventCodesOk

`func (o *FiltersVmsState) GetMaintenanceEventCodesOk() (*[]string, bool)`

GetMaintenanceEventCodesOk returns a tuple with the MaintenanceEventCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceEventCodes

`func (o *FiltersVmsState) SetMaintenanceEventCodes(v []string)`

SetMaintenanceEventCodes sets MaintenanceEventCodes field to given value.

### HasMaintenanceEventCodes

`func (o *FiltersVmsState) HasMaintenanceEventCodes() bool`

HasMaintenanceEventCodes returns a boolean if a field has been set.

### GetMaintenanceEventDescriptions

`func (o *FiltersVmsState) GetMaintenanceEventDescriptions() []string`

GetMaintenanceEventDescriptions returns the MaintenanceEventDescriptions field if non-nil, zero value otherwise.

### GetMaintenanceEventDescriptionsOk

`func (o *FiltersVmsState) GetMaintenanceEventDescriptionsOk() (*[]string, bool)`

GetMaintenanceEventDescriptionsOk returns a tuple with the MaintenanceEventDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceEventDescriptions

`func (o *FiltersVmsState) SetMaintenanceEventDescriptions(v []string)`

SetMaintenanceEventDescriptions sets MaintenanceEventDescriptions field to given value.

### HasMaintenanceEventDescriptions

`func (o *FiltersVmsState) HasMaintenanceEventDescriptions() bool`

HasMaintenanceEventDescriptions returns a boolean if a field has been set.

### GetMaintenanceEventsNotAfter

`func (o *FiltersVmsState) GetMaintenanceEventsNotAfter() []string`

GetMaintenanceEventsNotAfter returns the MaintenanceEventsNotAfter field if non-nil, zero value otherwise.

### GetMaintenanceEventsNotAfterOk

`func (o *FiltersVmsState) GetMaintenanceEventsNotAfterOk() (*[]string, bool)`

GetMaintenanceEventsNotAfterOk returns a tuple with the MaintenanceEventsNotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceEventsNotAfter

`func (o *FiltersVmsState) SetMaintenanceEventsNotAfter(v []string)`

SetMaintenanceEventsNotAfter sets MaintenanceEventsNotAfter field to given value.

### HasMaintenanceEventsNotAfter

`func (o *FiltersVmsState) HasMaintenanceEventsNotAfter() bool`

HasMaintenanceEventsNotAfter returns a boolean if a field has been set.

### GetMaintenanceEventsNotBefore

`func (o *FiltersVmsState) GetMaintenanceEventsNotBefore() []string`

GetMaintenanceEventsNotBefore returns the MaintenanceEventsNotBefore field if non-nil, zero value otherwise.

### GetMaintenanceEventsNotBeforeOk

`func (o *FiltersVmsState) GetMaintenanceEventsNotBeforeOk() (*[]string, bool)`

GetMaintenanceEventsNotBeforeOk returns a tuple with the MaintenanceEventsNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceEventsNotBefore

`func (o *FiltersVmsState) SetMaintenanceEventsNotBefore(v []string)`

SetMaintenanceEventsNotBefore sets MaintenanceEventsNotBefore field to given value.

### HasMaintenanceEventsNotBefore

`func (o *FiltersVmsState) HasMaintenanceEventsNotBefore() bool`

HasMaintenanceEventsNotBefore returns a boolean if a field has been set.

### GetSubregionNames

`func (o *FiltersVmsState) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersVmsState) GetSubregionNamesOk() (*[]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionNames

`func (o *FiltersVmsState) SetSubregionNames(v []string)`

SetSubregionNames sets SubregionNames field to given value.

### HasSubregionNames

`func (o *FiltersVmsState) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### GetVmIds

`func (o *FiltersVmsState) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *FiltersVmsState) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *FiltersVmsState) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *FiltersVmsState) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### GetVmStates

`func (o *FiltersVmsState) GetVmStates() []string`

GetVmStates returns the VmStates field if non-nil, zero value otherwise.

### GetVmStatesOk

`func (o *FiltersVmsState) GetVmStatesOk() (*[]string, bool)`

GetVmStatesOk returns a tuple with the VmStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmStates

`func (o *FiltersVmsState) SetVmStates(v []string)`

SetVmStates sets VmStates field to given value.

### HasVmStates

`func (o *FiltersVmsState) HasVmStates() bool`

HasVmStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


