# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Path** | Pointer to **string** | The path to the EIM user you want to create (by default, &#x60;/&#x60;). This path name must begin and end with a slash (&#x60;/&#x60;), and contain between 1 and 512 alphanumeric characters and/or slashes (&#x60;/&#x60;), or underscores (&#x60;_&#x60;). | [optional] 
**UserEmail** | Pointer to **string** | The email address of the EIM user. | [optional] 
**UserName** | **string** | The name of the EIM user. This user name must contain between 1 and 64 alphanumeric characters and/or pluses (&#x60;+&#x60;), equals (&#x60;&#x3D;&#x60;), commas (&#x60;,&#x60;), periods (&#x60;.&#x60;), at signs (&#x60;@&#x60;), dashes (&#x60;-&#x60;), or underscores (&#x60;_&#x60;). | 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(userName string, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateUserRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateUserRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateUserRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateUserRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPath

`func (o *CreateUserRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateUserRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateUserRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateUserRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserEmail

`func (o *CreateUserRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *CreateUserRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *CreateUserRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *CreateUserRequest) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserName

`func (o *CreateUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


