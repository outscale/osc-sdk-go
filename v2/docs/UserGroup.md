# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **string** | The date and time (UTC) of creation of the user group. | [optional] 
**LastModificationDate** | Pointer to **string** | The date and time (UTC) of the last modification of the user group. | [optional] 
**Name** | Pointer to **string** | The name of the user group. | [optional] 
**Orn** | Pointer to **string** | The Outscale Resource Name (ORN) of the user group. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | [optional] 
**Path** | Pointer to **string** | The path to the user group. | [optional] 
**UserGroupId** | Pointer to **string** | The ID of the user group. | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup() *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *UserGroup) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UserGroup) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UserGroup) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *UserGroup) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModificationDate

`func (o *UserGroup) GetLastModificationDate() string`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *UserGroup) GetLastModificationDateOk() (*string, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *UserGroup) SetLastModificationDate(v string)`

SetLastModificationDate sets LastModificationDate field to given value.

### HasLastModificationDate

`func (o *UserGroup) HasLastModificationDate() bool`

HasLastModificationDate returns a boolean if a field has been set.

### GetName

`func (o *UserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrn

`func (o *UserGroup) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *UserGroup) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *UserGroup) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *UserGroup) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetPath

`func (o *UserGroup) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UserGroup) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UserGroup) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UserGroup) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserGroupId

`func (o *UserGroup) GetUserGroupId() string`

GetUserGroupId returns the UserGroupId field if non-nil, zero value otherwise.

### GetUserGroupIdOk

`func (o *UserGroup) GetUserGroupIdOk() (*string, bool)`

GetUserGroupIdOk returns a tuple with the UserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupId

`func (o *UserGroup) SetUserGroupId(v string)`

SetUserGroupId sets UserGroupId field to given value.

### HasUserGroupId

`func (o *UserGroup) HasUserGroupId() bool`

HasUserGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


