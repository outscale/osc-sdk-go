# VmStates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceEvents** | Pointer to [**[]MaintenanceEvent**](MaintenanceEvent.md) | One or more scheduled events associated with the VM. | [optional] 
**SubregionName** | Pointer to **string** | The name of the Subregion of the VM. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 
**VmState** | Pointer to **string** | The state of the VM (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 

## Methods

### NewVmStates

`func NewVmStates() *VmStates`

NewVmStates instantiates a new VmStates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmStatesWithDefaults

`func NewVmStatesWithDefaults() *VmStates`

NewVmStatesWithDefaults instantiates a new VmStates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceEvents

`func (o *VmStates) GetMaintenanceEvents() []MaintenanceEvent`

GetMaintenanceEvents returns the MaintenanceEvents field if non-nil, zero value otherwise.

### GetMaintenanceEventsOk

`func (o *VmStates) GetMaintenanceEventsOk() (*[]MaintenanceEvent, bool)`

GetMaintenanceEventsOk returns a tuple with the MaintenanceEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceEvents

`func (o *VmStates) SetMaintenanceEvents(v []MaintenanceEvent)`

SetMaintenanceEvents sets MaintenanceEvents field to given value.

### HasMaintenanceEvents

`func (o *VmStates) HasMaintenanceEvents() bool`

HasMaintenanceEvents returns a boolean if a field has been set.

### GetSubregionName

`func (o *VmStates) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *VmStates) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *VmStates) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *VmStates) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### GetVmId

`func (o *VmStates) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VmStates) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VmStates) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *VmStates) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetVmState

`func (o *VmStates) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *VmStates) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *VmStates) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *VmStates) HasVmState() bool`

HasVmState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


