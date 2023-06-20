# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NewPath** | Pointer to **string** | A new path for the EIM user. | [optional] 
**NewUserName** | Pointer to **string** | A new name for the EIM user. | [optional] 
**UserName** | **string** | The name of the EIM user you want to modify. | 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest(userName string, ) *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateUserRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateUserRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateUserRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateUserRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNewPath

`func (o *UpdateUserRequest) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *UpdateUserRequest) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *UpdateUserRequest) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *UpdateUserRequest) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### GetNewUserName

`func (o *UpdateUserRequest) GetNewUserName() string`

GetNewUserName returns the NewUserName field if non-nil, zero value otherwise.

### GetNewUserNameOk

`func (o *UpdateUserRequest) GetNewUserNameOk() (*string, bool)`

GetNewUserNameOk returns a tuple with the NewUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewUserName

`func (o *UpdateUserRequest) SetNewUserName(v string)`

SetNewUserName sets NewUserName field to given value.

### HasNewUserName

`func (o *UpdateUserRequest) HasNewUserName() bool`

HasNewUserName returns a boolean if a field has been set.

### GetUserName

`func (o *UpdateUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UpdateUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UpdateUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


