# SecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID that has been granted permission. | [optional] 
**Description** | Pointer to **string** | The description of the security group. | [optional] 
**InboundRules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | The inbound rules associated with the security group. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the security group. | [optional] 
**OutboundRules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | The outbound rules associated with the security group. | [optional] 
**SecurityGroupId** | Pointer to **string** | The ID of the security group. | [optional] 
**SecurityGroupName** | Pointer to **string** | The name of the security group. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the security group. | [optional] 

## Methods

### NewSecurityGroup

`func NewSecurityGroup() *SecurityGroup`

NewSecurityGroup instantiates a new SecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupWithDefaults

`func NewSecurityGroupWithDefaults() *SecurityGroup`

NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SecurityGroup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SecurityGroup) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SecurityGroup) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SecurityGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInboundRules

`func (o *SecurityGroup) GetInboundRules() []SecurityGroupRule`

GetInboundRules returns the InboundRules field if non-nil, zero value otherwise.

### GetInboundRulesOk

`func (o *SecurityGroup) GetInboundRulesOk() (*[]SecurityGroupRule, bool)`

GetInboundRulesOk returns a tuple with the InboundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRules

`func (o *SecurityGroup) SetInboundRules(v []SecurityGroupRule)`

SetInboundRules sets InboundRules field to given value.

### HasInboundRules

`func (o *SecurityGroup) HasInboundRules() bool`

HasInboundRules returns a boolean if a field has been set.

### GetNetId

`func (o *SecurityGroup) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *SecurityGroup) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *SecurityGroup) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *SecurityGroup) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetOutboundRules

`func (o *SecurityGroup) GetOutboundRules() []SecurityGroupRule`

GetOutboundRules returns the OutboundRules field if non-nil, zero value otherwise.

### GetOutboundRulesOk

`func (o *SecurityGroup) GetOutboundRulesOk() (*[]SecurityGroupRule, bool)`

GetOutboundRulesOk returns a tuple with the OutboundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundRules

`func (o *SecurityGroup) SetOutboundRules(v []SecurityGroupRule)`

SetOutboundRules sets OutboundRules field to given value.

### HasOutboundRules

`func (o *SecurityGroup) HasOutboundRules() bool`

HasOutboundRules returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *SecurityGroup) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *SecurityGroup) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *SecurityGroup) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *SecurityGroup) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetSecurityGroupName

`func (o *SecurityGroup) GetSecurityGroupName() string`

GetSecurityGroupName returns the SecurityGroupName field if non-nil, zero value otherwise.

### GetSecurityGroupNameOk

`func (o *SecurityGroup) GetSecurityGroupNameOk() (*string, bool)`

GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupName

`func (o *SecurityGroup) SetSecurityGroupName(v string)`

SetSecurityGroupName sets SecurityGroupName field to given value.

### HasSecurityGroupName

`func (o *SecurityGroup) HasSecurityGroupName() bool`

HasSecurityGroupName returns a boolean if a field has been set.

### GetTags

`func (o *SecurityGroup) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityGroup) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityGroup) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecurityGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


