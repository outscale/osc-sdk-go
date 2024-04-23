# UpdateUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NewPath** | Pointer to **string** | A new path for the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 
**NewUserGroupName** | Pointer to **string** | A new name for the user group. | [optional] 
**Path** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 
**UserGroupName** | **string** | The name of the group you want to update. | 

## Methods

### NewUpdateUserGroupRequest

`func NewUpdateUserGroupRequest(userGroupName string, ) *UpdateUserGroupRequest`

NewUpdateUserGroupRequest instantiates a new UpdateUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserGroupRequestWithDefaults

`func NewUpdateUserGroupRequestWithDefaults() *UpdateUserGroupRequest`

NewUpdateUserGroupRequestWithDefaults instantiates a new UpdateUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateUserGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateUserGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateUserGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateUserGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNewPath

`func (o *UpdateUserGroupRequest) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *UpdateUserGroupRequest) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *UpdateUserGroupRequest) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *UpdateUserGroupRequest) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### GetNewUserGroupName

`func (o *UpdateUserGroupRequest) GetNewUserGroupName() string`

GetNewUserGroupName returns the NewUserGroupName field if non-nil, zero value otherwise.

### GetNewUserGroupNameOk

`func (o *UpdateUserGroupRequest) GetNewUserGroupNameOk() (*string, bool)`

GetNewUserGroupNameOk returns a tuple with the NewUserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewUserGroupName

`func (o *UpdateUserGroupRequest) SetNewUserGroupName(v string)`

SetNewUserGroupName sets NewUserGroupName field to given value.

### HasNewUserGroupName

`func (o *UpdateUserGroupRequest) HasNewUserGroupName() bool`

HasNewUserGroupName returns a boolean if a field has been set.

### GetPath

`func (o *UpdateUserGroupRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateUserGroupRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateUserGroupRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateUserGroupRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserGroupName

`func (o *UpdateUserGroupRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *UpdateUserGroupRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *UpdateUserGroupRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


