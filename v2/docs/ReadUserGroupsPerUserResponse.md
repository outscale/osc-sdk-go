# ReadUserGroupsPerUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**UserGroups** | Pointer to [**[]UserGroup**](UserGroup.md) | A list of user groups. | [optional] 

## Methods

### NewReadUserGroupsPerUserResponse

`func NewReadUserGroupsPerUserResponse() *ReadUserGroupsPerUserResponse`

NewReadUserGroupsPerUserResponse instantiates a new ReadUserGroupsPerUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserGroupsPerUserResponseWithDefaults

`func NewReadUserGroupsPerUserResponseWithDefaults() *ReadUserGroupsPerUserResponse`

NewReadUserGroupsPerUserResponseWithDefaults instantiates a new ReadUserGroupsPerUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadUserGroupsPerUserResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadUserGroupsPerUserResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadUserGroupsPerUserResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadUserGroupsPerUserResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetUserGroups

`func (o *ReadUserGroupsPerUserResponse) GetUserGroups() []UserGroup`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *ReadUserGroupsPerUserResponse) GetUserGroupsOk() (*[]UserGroup, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *ReadUserGroupsPerUserResponse) SetUserGroups(v []UserGroup)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *ReadUserGroupsPerUserResponse) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


