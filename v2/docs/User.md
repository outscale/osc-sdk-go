# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **string** | The date and time (UTC) of creation of the EIM user. | [optional] 
**LastModificationDate** | Pointer to **string** | The date and time (UTC) of the last modification of the EIM user. | [optional] 
**OutscaleLoginAllowed** | Pointer to **bool** | Whether the user is allowed to log in to Cockpit v2 using its Outscale credentials when identity federation is activated. | [optional] 
**Path** | Pointer to **string** | The path to the EIM user. | [optional] 
**UserEmail** | Pointer to **string** | The email address of the EIM user. | [optional] 
**UserId** | Pointer to **string** | The ID of the EIM user. | [optional] 
**UserName** | Pointer to **string** | The name of the EIM user. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *User) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *User) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *User) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *User) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModificationDate

`func (o *User) GetLastModificationDate() string`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *User) GetLastModificationDateOk() (*string, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *User) SetLastModificationDate(v string)`

SetLastModificationDate sets LastModificationDate field to given value.

### HasLastModificationDate

`func (o *User) HasLastModificationDate() bool`

HasLastModificationDate returns a boolean if a field has been set.

### GetOutscaleLoginAllowed

`func (o *User) GetOutscaleLoginAllowed() bool`

GetOutscaleLoginAllowed returns the OutscaleLoginAllowed field if non-nil, zero value otherwise.

### GetOutscaleLoginAllowedOk

`func (o *User) GetOutscaleLoginAllowedOk() (*bool, bool)`

GetOutscaleLoginAllowedOk returns a tuple with the OutscaleLoginAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutscaleLoginAllowed

`func (o *User) SetOutscaleLoginAllowed(v bool)`

SetOutscaleLoginAllowed sets OutscaleLoginAllowed field to given value.

### HasOutscaleLoginAllowed

`func (o *User) HasOutscaleLoginAllowed() bool`

HasOutscaleLoginAllowed returns a boolean if a field has been set.

### GetPath

`func (o *User) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *User) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *User) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *User) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserEmail

`func (o *User) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *User) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *User) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *User) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserId

`func (o *User) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *User) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *User) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *User) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *User) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *User) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *User) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *User) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


