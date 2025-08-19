# EnableOutscaleLoginPerUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**UserNames** | **[]string** | The usernames of the EIM users you want to enable the Outscale login for. | 

## Methods

### NewEnableOutscaleLoginPerUsersRequest

`func NewEnableOutscaleLoginPerUsersRequest(userNames []string, ) *EnableOutscaleLoginPerUsersRequest`

NewEnableOutscaleLoginPerUsersRequest instantiates a new EnableOutscaleLoginPerUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableOutscaleLoginPerUsersRequestWithDefaults

`func NewEnableOutscaleLoginPerUsersRequestWithDefaults() *EnableOutscaleLoginPerUsersRequest`

NewEnableOutscaleLoginPerUsersRequestWithDefaults instantiates a new EnableOutscaleLoginPerUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *EnableOutscaleLoginPerUsersRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *EnableOutscaleLoginPerUsersRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *EnableOutscaleLoginPerUsersRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *EnableOutscaleLoginPerUsersRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetUserNames

`func (o *EnableOutscaleLoginPerUsersRequest) GetUserNames() []string`

GetUserNames returns the UserNames field if non-nil, zero value otherwise.

### GetUserNamesOk

`func (o *EnableOutscaleLoginPerUsersRequest) GetUserNamesOk() (*[]string, bool)`

GetUserNamesOk returns a tuple with the UserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNames

`func (o *EnableOutscaleLoginPerUsersRequest) SetUserNames(v []string)`

SetUserNames sets UserNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


