# CreateSecurityGroupRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**SecurityGroup** | Pointer to [**SecurityGroup**](SecurityGroup.md) |  | [optional] 

## Methods

### NewCreateSecurityGroupRuleResponse

`func NewCreateSecurityGroupRuleResponse() *CreateSecurityGroupRuleResponse`

NewCreateSecurityGroupRuleResponse instantiates a new CreateSecurityGroupRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRuleResponseWithDefaults

`func NewCreateSecurityGroupRuleResponseWithDefaults() *CreateSecurityGroupRuleResponse`

NewCreateSecurityGroupRuleResponseWithDefaults instantiates a new CreateSecurityGroupRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateSecurityGroupRuleResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateSecurityGroupRuleResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateSecurityGroupRuleResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateSecurityGroupRuleResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *CreateSecurityGroupRuleResponse) GetSecurityGroup() SecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *CreateSecurityGroupRuleResponse) GetSecurityGroupOk() (*SecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *CreateSecurityGroupRuleResponse) SetSecurityGroup(v SecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *CreateSecurityGroupRuleResponse) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


