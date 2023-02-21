# SecurityGroupsMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID that owns the source or destination security group. | [optional] 
**SecurityGroupId** | Pointer to **string** | The ID of a source or destination security group that you want to link to the security group of the rule. | [optional] 
**SecurityGroupName** | Pointer to **string** | (Public Cloud only) The name of a source or destination security group that you want to link to the security group of the rule. | [optional] 

## Methods

### NewSecurityGroupsMember

`func NewSecurityGroupsMember() *SecurityGroupsMember`

NewSecurityGroupsMember instantiates a new SecurityGroupsMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupsMemberWithDefaults

`func NewSecurityGroupsMemberWithDefaults() *SecurityGroupsMember`

NewSecurityGroupsMemberWithDefaults instantiates a new SecurityGroupsMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SecurityGroupsMember) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SecurityGroupsMember) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SecurityGroupsMember) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SecurityGroupsMember) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *SecurityGroupsMember) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *SecurityGroupsMember) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *SecurityGroupsMember) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *SecurityGroupsMember) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetSecurityGroupName

`func (o *SecurityGroupsMember) GetSecurityGroupName() string`

GetSecurityGroupName returns the SecurityGroupName field if non-nil, zero value otherwise.

### GetSecurityGroupNameOk

`func (o *SecurityGroupsMember) GetSecurityGroupNameOk() (*string, bool)`

GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupName

`func (o *SecurityGroupsMember) SetSecurityGroupName(v string)`

SetSecurityGroupName sets SecurityGroupName field to given value.

### HasSecurityGroupName

`func (o *SecurityGroupsMember) HasSecurityGroupName() bool`

HasSecurityGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


