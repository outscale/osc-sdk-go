# UpdateDedicatedGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedGroupId** | **string** | The ID of the dedicated group you want to update. | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Name** | **string** | The new name of the dedicated group. | 

## Methods

### NewUpdateDedicatedGroupRequest

`func NewUpdateDedicatedGroupRequest(dedicatedGroupId string, name string, ) *UpdateDedicatedGroupRequest`

NewUpdateDedicatedGroupRequest instantiates a new UpdateDedicatedGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDedicatedGroupRequestWithDefaults

`func NewUpdateDedicatedGroupRequestWithDefaults() *UpdateDedicatedGroupRequest`

NewUpdateDedicatedGroupRequestWithDefaults instantiates a new UpdateDedicatedGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedGroupId

`func (o *UpdateDedicatedGroupRequest) GetDedicatedGroupId() string`

GetDedicatedGroupId returns the DedicatedGroupId field if non-nil, zero value otherwise.

### GetDedicatedGroupIdOk

`func (o *UpdateDedicatedGroupRequest) GetDedicatedGroupIdOk() (*string, bool)`

GetDedicatedGroupIdOk returns a tuple with the DedicatedGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedGroupId

`func (o *UpdateDedicatedGroupRequest) SetDedicatedGroupId(v string)`

SetDedicatedGroupId sets DedicatedGroupId field to given value.


### GetDryRun

`func (o *UpdateDedicatedGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateDedicatedGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateDedicatedGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateDedicatedGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetName

`func (o *UpdateDedicatedGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDedicatedGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDedicatedGroupRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


