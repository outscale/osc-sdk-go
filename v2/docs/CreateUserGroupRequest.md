# CreateUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Path** | Pointer to **string** | The path to the group. If not specified, it is set to a slash (&#x60;/&#x60;). | [optional] 
**UserGroupName** | **string** | The name of the group. | 

## Methods

### NewCreateUserGroupRequest

`func NewCreateUserGroupRequest(userGroupName string, ) *CreateUserGroupRequest`

NewCreateUserGroupRequest instantiates a new CreateUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserGroupRequestWithDefaults

`func NewCreateUserGroupRequestWithDefaults() *CreateUserGroupRequest`

NewCreateUserGroupRequestWithDefaults instantiates a new CreateUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateUserGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateUserGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateUserGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateUserGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPath

`func (o *CreateUserGroupRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateUserGroupRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateUserGroupRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateUserGroupRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserGroupName

`func (o *CreateUserGroupRequest) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *CreateUserGroupRequest) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *CreateUserGroupRequest) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


