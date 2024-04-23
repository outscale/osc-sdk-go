# RemoveUserFromUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**UserGroupName** | **string** | The name of the group you want to remove the user from. | 
**UserGroupPath** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 
**UserName** | **string** | The name of the user you want to remove from the group. | 
**UserPath** | Pointer to **string** | The path to the user (by default, &#x60;/&#x60;). | [optional] 

## Methods

### NewRemoveUserFromUserGroupRequest

`func NewRemoveUserFromUserGroupRequest(userGroupName string, userName string, ) *RemoveUserFromUserGroupRequest`

NewRemoveUserFromUserGroupRequest instantiates a new RemoveUserFromUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveUserFromUserGroupRequestWithDefaults

`func NewRemoveUserFromUserGroupRequestWithDefaults() *RemoveUserFromUserGroupRequest`

NewRemoveUserFromUserGroupRequestWithDefaults instantiates a new RemoveUserFromUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *RemoveUserFromUserGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *RemoveUserFromUserGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *RemoveUserFromUserGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *RemoveUserFromUserGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetUserGroupName

`func (o *RemoveUserFromUserGroupRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *RemoveUserFromUserGroupRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *RemoveUserFromUserGroupRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.


### GetUserGroupPath

`func (o *RemoveUserFromUserGroupRequest) GetUserGroupPath() string`

GetUserGroupPath returns the UserGroupPath field if non-nil, zero value otherwise.

### GetUserGroupPathOk

`func (o *RemoveUserFromUserGroupRequest) GetUserGroupPathOk() (*string, bool)`

GetUserGroupPathOk returns a tuple with the UserGroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupPath

`func (o *RemoveUserFromUserGroupRequest) SetUserGroupPath(v string)`

SetUserGroupPath sets UserGroupPath field to given value.

### HasUserGroupPath

`func (o *RemoveUserFromUserGroupRequest) HasUserGroupPath() bool`

HasUserGroupPath returns a boolean if a field has been set.

### GetUserName

`func (o *RemoveUserFromUserGroupRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RemoveUserFromUserGroupRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RemoveUserFromUserGroupRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserPath

`func (o *RemoveUserFromUserGroupRequest) GetUserPath() string`

GetUserPath returns the UserPath field if non-nil, zero value otherwise.

### GetUserPathOk

`func (o *RemoveUserFromUserGroupRequest) GetUserPathOk() (*string, bool)`

GetUserPathOk returns a tuple with the UserPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPath

`func (o *RemoveUserFromUserGroupRequest) SetUserPath(v string)`

SetUserPath sets UserPath field to given value.

### HasUserPath

`func (o *RemoveUserFromUserGroupRequest) HasUserPath() bool`

HasUserPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


