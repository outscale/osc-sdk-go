# ReadUserGroupsPerUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**UserName** | **string** | The name of the user. | 
**UserPath** | Pointer to **string** | The path to the user (by default, &#x60;/&#x60;). | [optional] 

## Methods

### NewReadUserGroupsPerUserRequest

`func NewReadUserGroupsPerUserRequest(userName string, ) *ReadUserGroupsPerUserRequest`

NewReadUserGroupsPerUserRequest instantiates a new ReadUserGroupsPerUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUserGroupsPerUserRequestWithDefaults

`func NewReadUserGroupsPerUserRequestWithDefaults() *ReadUserGroupsPerUserRequest`

NewReadUserGroupsPerUserRequestWithDefaults instantiates a new ReadUserGroupsPerUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadUserGroupsPerUserRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadUserGroupsPerUserRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadUserGroupsPerUserRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadUserGroupsPerUserRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetUserName

`func (o *ReadUserGroupsPerUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ReadUserGroupsPerUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ReadUserGroupsPerUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserPath

`func (o *ReadUserGroupsPerUserRequest) GetUserPath() string`

GetUserPath returns the UserPath field if non-nil, zero value otherwise.

### GetUserPathOk

`func (o *ReadUserGroupsPerUserRequest) GetUserPathOk() (*string, bool)`

GetUserPathOk returns a tuple with the UserPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPath

`func (o *ReadUserGroupsPerUserRequest) SetUserPath(v string)`

SetUserPath sets UserPath field to given value.

### HasUserPath

`func (o *ReadUserGroupsPerUserRequest) HasUserPath() bool`

HasUserPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


