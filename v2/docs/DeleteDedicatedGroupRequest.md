# DeleteDedicatedGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedGroupId** | **string** | The ID of the dedicated group you want to delete. | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Force** | Pointer to **bool** | If true, forces the deletion of the dedicated group and all its dependencies. | [optional] 

## Methods

### NewDeleteDedicatedGroupRequest

`func NewDeleteDedicatedGroupRequest(dedicatedGroupId string, ) *DeleteDedicatedGroupRequest`

NewDeleteDedicatedGroupRequest instantiates a new DeleteDedicatedGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDedicatedGroupRequestWithDefaults

`func NewDeleteDedicatedGroupRequestWithDefaults() *DeleteDedicatedGroupRequest`

NewDeleteDedicatedGroupRequestWithDefaults instantiates a new DeleteDedicatedGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedGroupId

`func (o *DeleteDedicatedGroupRequest) GetDedicatedGroupId() string`

GetDedicatedGroupId returns the DedicatedGroupId field if non-nil, zero value otherwise.

### GetDedicatedGroupIdOk

`func (o *DeleteDedicatedGroupRequest) GetDedicatedGroupIdOk() (*string, bool)`

GetDedicatedGroupIdOk returns a tuple with the DedicatedGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedGroupId

`func (o *DeleteDedicatedGroupRequest) SetDedicatedGroupId(v string)`

SetDedicatedGroupId sets DedicatedGroupId field to given value.


### GetDryRun

`func (o *DeleteDedicatedGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteDedicatedGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteDedicatedGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteDedicatedGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetForce

`func (o *DeleteDedicatedGroupRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *DeleteDedicatedGroupRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *DeleteDedicatedGroupRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *DeleteDedicatedGroupRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


