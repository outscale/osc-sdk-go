# AddUserToUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**UserGroupName** | **string** | The name of the group you want to add a user to. | 
**UserGroupPath** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 
**UserName** | **string** | The name of the user you want to add to the group. | 
**UserPath** | Pointer to **string** | The path to the user. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 

## Methods

### NewAddUserToUserGroupRequest

`func NewAddUserToUserGroupRequest(userGroupName string, userName string, ) *AddUserToUserGroupRequest`

NewAddUserToUserGroupRequest instantiates a new AddUserToUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserToUserGroupRequestWithDefaults

`func NewAddUserToUserGroupRequestWithDefaults() *AddUserToUserGroupRequest`

NewAddUserToUserGroupRequestWithDefaults instantiates a new AddUserToUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *AddUserToUserGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AddUserToUserGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AddUserToUserGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AddUserToUserGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetUserGroupName

`func (o *AddUserToUserGroupRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *AddUserToUserGroupRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *AddUserToUserGroupRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.


### GetUserGroupPath

`func (o *AddUserToUserGroupRequest) GetUserGroupPath() string`

GetUserGroupPath returns the UserGroupPath field if non-nil, zero value otherwise.

### GetUserGroupPathOk

`func (o *AddUserToUserGroupRequest) GetUserGroupPathOk() (*string, bool)`

GetUserGroupPathOk returns a tuple with the UserGroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupPath

`func (o *AddUserToUserGroupRequest) SetUserGroupPath(v string)`

SetUserGroupPath sets UserGroupPath field to given value.

### HasUserGroupPath

`func (o *AddUserToUserGroupRequest) HasUserGroupPath() bool`

HasUserGroupPath returns a boolean if a field has been set.

### GetUserName

`func (o *AddUserToUserGroupRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AddUserToUserGroupRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AddUserToUserGroupRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserPath

`func (o *AddUserToUserGroupRequest) GetUserPath() string`

GetUserPath returns the UserPath field if non-nil, zero value otherwise.

### GetUserPathOk

`func (o *AddUserToUserGroupRequest) GetUserPathOk() (*string, bool)`

GetUserPathOk returns a tuple with the UserPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPath

`func (o *AddUserToUserGroupRequest) SetUserPath(v string)`

SetUserPath sets UserPath field to given value.

### HasUserPath

`func (o *AddUserToUserGroupRequest) HasUserPath() bool`

HasUserPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


