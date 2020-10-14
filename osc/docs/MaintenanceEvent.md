# MaintenanceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The code of the event (&#x60;system-reboot&#x60; \\| &#x60;system-maintenance&#x60;). | [optional] 
**Description** | Pointer to **string** | The description of the event. | [optional] 
**NotAfter** | Pointer to **string** | The latest scheduled end time for the event. | [optional] 
**NotBefore** | Pointer to **string** | The earliest scheduled start time for the event. | [optional] 

## Methods

### NewMaintenanceEvent

`func NewMaintenanceEvent() *MaintenanceEvent`

NewMaintenanceEvent instantiates a new MaintenanceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceEventWithDefaults

`func NewMaintenanceEventWithDefaults() *MaintenanceEvent`

NewMaintenanceEventWithDefaults instantiates a new MaintenanceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MaintenanceEvent) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MaintenanceEvent) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MaintenanceEvent) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MaintenanceEvent) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *MaintenanceEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenanceEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenanceEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MaintenanceEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotAfter

`func (o *MaintenanceEvent) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *MaintenanceEvent) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *MaintenanceEvent) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *MaintenanceEvent) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *MaintenanceEvent) GetNotBefore() string`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *MaintenanceEvent) GetNotBeforeOk() (*string, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *MaintenanceEvent) SetNotBefore(v string)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *MaintenanceEvent) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


