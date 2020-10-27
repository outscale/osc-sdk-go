# BackendVmHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the state of the back-end VM. | [optional] 
**State** | Pointer to **string** | The state of the back-end VM (&#x60;InService&#x60; \\| &#x60;OutOfService&#x60; \\| &#x60;Unknown&#x60;). | [optional] 
**StateReason** | Pointer to **string** | Information about the cause of &#x60;OutOfService&#x60; VMs.&lt;br /&gt; Specifically, whether the cause is Elastic Load Balancing or the VM (&#x60;ELB&#x60; \\| &#x60;Instance&#x60; \\| &#x60;N/A&#x60;). | [optional] 
**VmId** | Pointer to **string** | The ID of the back-end VM. | [optional] 

## Methods

### NewBackendVmHealth

`func NewBackendVmHealth() *BackendVmHealth`

NewBackendVmHealth instantiates a new BackendVmHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackendVmHealthWithDefaults

`func NewBackendVmHealthWithDefaults() *BackendVmHealth`

NewBackendVmHealthWithDefaults instantiates a new BackendVmHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BackendVmHealth) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackendVmHealth) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackendVmHealth) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackendVmHealth) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *BackendVmHealth) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackendVmHealth) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackendVmHealth) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BackendVmHealth) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *BackendVmHealth) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *BackendVmHealth) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *BackendVmHealth) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *BackendVmHealth) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetVmId

`func (o *BackendVmHealth) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *BackendVmHealth) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *BackendVmHealth) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *BackendVmHealth) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


