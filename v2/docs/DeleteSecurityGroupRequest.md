# DeleteSecurityGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**SecurityGroupId** | Pointer to **string** | The ID of the security group you want to delete. | [optional] 
**SecurityGroupName** | Pointer to **string** | The name of the security group. | [optional] 

## Methods

### NewDeleteSecurityGroupRequest

`func NewDeleteSecurityGroupRequest() *DeleteSecurityGroupRequest`

NewDeleteSecurityGroupRequest instantiates a new DeleteSecurityGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSecurityGroupRequestWithDefaults

`func NewDeleteSecurityGroupRequestWithDefaults() *DeleteSecurityGroupRequest`

NewDeleteSecurityGroupRequestWithDefaults instantiates a new DeleteSecurityGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteSecurityGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteSecurityGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteSecurityGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteSecurityGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *DeleteSecurityGroupRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *DeleteSecurityGroupRequest) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *DeleteSecurityGroupRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *DeleteSecurityGroupRequest) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetSecurityGroupName

`func (o *DeleteSecurityGroupRequest) GetSecurityGroupName() string`

GetSecurityGroupName returns the SecurityGroupName field if non-nil, zero value otherwise.

### GetSecurityGroupNameOk

`func (o *DeleteSecurityGroupRequest) GetSecurityGroupNameOk() (*string, bool)`

GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupName

`func (o *DeleteSecurityGroupRequest) SetSecurityGroupName(v string)`

SetSecurityGroupName sets SecurityGroupName field to given value.

### HasSecurityGroupName

`func (o *DeleteSecurityGroupRequest) HasSecurityGroupName() bool`

HasSecurityGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


