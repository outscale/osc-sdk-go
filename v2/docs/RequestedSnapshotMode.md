# RequestedSnapshotMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The requested mode (&#x60;application-consistent&#x60; \\| &#x60;crash-consistent&#x60;). | [optional] 
**OnFailure** | Pointer to **string** | The fallback behavior if snapshot creation in &#x60;application-consistent&#x60; mode fails (&#x60;continue&#x60; \\| &#x60;abort&#x60;). By default or if set to &#x60;continue&#x60;, the snapshot is created in crash-consistent mode. If set to &#x60;abort&#x60;, the task is canceled. | [optional] 

## Methods

### NewRequestedSnapshotMode

`func NewRequestedSnapshotMode() *RequestedSnapshotMode`

NewRequestedSnapshotMode instantiates a new RequestedSnapshotMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedSnapshotModeWithDefaults

`func NewRequestedSnapshotModeWithDefaults() *RequestedSnapshotMode`

NewRequestedSnapshotModeWithDefaults instantiates a new RequestedSnapshotMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RequestedSnapshotMode) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RequestedSnapshotMode) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RequestedSnapshotMode) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RequestedSnapshotMode) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOnFailure

`func (o *RequestedSnapshotMode) GetOnFailure() string`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *RequestedSnapshotMode) GetOnFailureOk() (*string, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *RequestedSnapshotMode) SetOnFailure(v string)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *RequestedSnapshotMode) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


