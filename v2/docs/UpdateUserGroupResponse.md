# UpdateUserGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**UserGroup** | Pointer to [**UserGroup**](UserGroup.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) | A list of EIM users. | [optional] 

## Methods

### NewUpdateUserGroupResponse

`func NewUpdateUserGroupResponse() *UpdateUserGroupResponse`

NewUpdateUserGroupResponse instantiates a new UpdateUserGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserGroupResponseWithDefaults

`func NewUpdateUserGroupResponseWithDefaults() *UpdateUserGroupResponse`

NewUpdateUserGroupResponseWithDefaults instantiates a new UpdateUserGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *UpdateUserGroupResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UpdateUserGroupResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UpdateUserGroupResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UpdateUserGroupResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetUserGroup

`func (o *UpdateUserGroupResponse) GetUserGroup() UserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *UpdateUserGroupResponse) GetUserGroupOk() (*UserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *UpdateUserGroupResponse) SetUserGroup(v UserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *UpdateUserGroupResponse) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetUsers

`func (o *UpdateUserGroupResponse) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateUserGroupResponse) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateUserGroupResponse) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateUserGroupResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


