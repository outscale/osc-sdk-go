# DeleteUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Force** | Pointer to **bool** | If true, forces the deletion of the user group even if it is not empty. | [optional] 
**Path** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 
**UserGroupName** | **string** | The name of the group you want to delete. | 

## Methods

### NewDeleteUserGroupRequest

`func NewDeleteUserGroupRequest(userGroupName string, ) *DeleteUserGroupRequest`

NewDeleteUserGroupRequest instantiates a new DeleteUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserGroupRequestWithDefaults

`func NewDeleteUserGroupRequestWithDefaults() *DeleteUserGroupRequest`

NewDeleteUserGroupRequestWithDefaults instantiates a new DeleteUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteUserGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteUserGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteUserGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteUserGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetForce

`func (o *DeleteUserGroupRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *DeleteUserGroupRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *DeleteUserGroupRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *DeleteUserGroupRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetPath

`func (o *DeleteUserGroupRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeleteUserGroupRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeleteUserGroupRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeleteUserGroupRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserGroupName

`func (o *DeleteUserGroupRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *DeleteUserGroupRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *DeleteUserGroupRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


