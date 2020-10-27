# SourceSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupAccountId** | Pointer to **string** | The account ID of the owner of the security group. | [optional] 
**SecurityGroupName** | Pointer to **string** | The name of the security group. | [optional] 

## Methods

### NewSourceSecurityGroup

`func NewSourceSecurityGroup() *SourceSecurityGroup`

NewSourceSecurityGroup instantiates a new SourceSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceSecurityGroupWithDefaults

`func NewSourceSecurityGroupWithDefaults() *SourceSecurityGroup`

NewSourceSecurityGroupWithDefaults instantiates a new SourceSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupAccountId

`func (o *SourceSecurityGroup) GetSecurityGroupAccountId() string`

GetSecurityGroupAccountId returns the SecurityGroupAccountId field if non-nil, zero value otherwise.

### GetSecurityGroupAccountIdOk

`func (o *SourceSecurityGroup) GetSecurityGroupAccountIdOk() (*string, bool)`

GetSecurityGroupAccountIdOk returns a tuple with the SecurityGroupAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupAccountId

`func (o *SourceSecurityGroup) SetSecurityGroupAccountId(v string)`

SetSecurityGroupAccountId sets SecurityGroupAccountId field to given value.

### HasSecurityGroupAccountId

`func (o *SourceSecurityGroup) HasSecurityGroupAccountId() bool`

HasSecurityGroupAccountId returns a boolean if a field has been set.

### GetSecurityGroupName

`func (o *SourceSecurityGroup) GetSecurityGroupName() string`

GetSecurityGroupName returns the SecurityGroupName field if non-nil, zero value otherwise.

### GetSecurityGroupNameOk

`func (o *SourceSecurityGroup) GetSecurityGroupNameOk() (*string, bool)`

GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupName

`func (o *SourceSecurityGroup) SetSecurityGroupName(v string)`

SetSecurityGroupName sets SecurityGroupName field to given value.

### HasSecurityGroupName

`func (o *SourceSecurityGroup) HasSecurityGroupName() bool`

HasSecurityGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


