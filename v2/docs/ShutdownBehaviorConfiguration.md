# ShutdownBehaviorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestAction** | Pointer to **string** | The action performed by the orchestrator when the VM is shut down from the guest operating system. By default, &#x60;stop&#x60;. | [optional] 
**HostAction** | Pointer to **string** | The action performed by the orchestrator when the VM is shut down due to a host infrastructure failure. By default, &#x60;restart&#x60;. | [optional] 

## Methods

### NewShutdownBehaviorConfiguration

`func NewShutdownBehaviorConfiguration() *ShutdownBehaviorConfiguration`

NewShutdownBehaviorConfiguration instantiates a new ShutdownBehaviorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutdownBehaviorConfigurationWithDefaults

`func NewShutdownBehaviorConfigurationWithDefaults() *ShutdownBehaviorConfiguration`

NewShutdownBehaviorConfigurationWithDefaults instantiates a new ShutdownBehaviorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestAction

`func (o *ShutdownBehaviorConfiguration) GetGuestAction() string`

GetGuestAction returns the GuestAction field if non-nil, zero value otherwise.

### GetGuestActionOk

`func (o *ShutdownBehaviorConfiguration) GetGuestActionOk() (*string, bool)`

GetGuestActionOk returns a tuple with the GuestAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestAction

`func (o *ShutdownBehaviorConfiguration) SetGuestAction(v string)`

SetGuestAction sets GuestAction field to given value.

### HasGuestAction

`func (o *ShutdownBehaviorConfiguration) HasGuestAction() bool`

HasGuestAction returns a boolean if a field has been set.

### GetHostAction

`func (o *ShutdownBehaviorConfiguration) GetHostAction() string`

GetHostAction returns the HostAction field if non-nil, zero value otherwise.

### GetHostActionOk

`func (o *ShutdownBehaviorConfiguration) GetHostActionOk() (*string, bool)`

GetHostActionOk returns a tuple with the HostAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAction

`func (o *ShutdownBehaviorConfiguration) SetHostAction(v string)`

SetHostAction sets HostAction field to given value.

### HasHostAction

`func (o *ShutdownBehaviorConfiguration) HasHostAction() bool`

HasHostAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


