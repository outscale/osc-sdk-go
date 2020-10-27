# DeleteSecurityGroupRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**SecurityGroup** | Pointer to [**SecurityGroup**](SecurityGroup.md) |  | [optional] 

## Methods

### NewDeleteSecurityGroupRuleResponse

`func NewDeleteSecurityGroupRuleResponse() *DeleteSecurityGroupRuleResponse`

NewDeleteSecurityGroupRuleResponse instantiates a new DeleteSecurityGroupRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSecurityGroupRuleResponseWithDefaults

`func NewDeleteSecurityGroupRuleResponseWithDefaults() *DeleteSecurityGroupRuleResponse`

NewDeleteSecurityGroupRuleResponseWithDefaults instantiates a new DeleteSecurityGroupRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *DeleteSecurityGroupRuleResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *DeleteSecurityGroupRuleResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *DeleteSecurityGroupRuleResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *DeleteSecurityGroupRuleResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *DeleteSecurityGroupRuleResponse) GetSecurityGroup() SecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *DeleteSecurityGroupRuleResponse) GetSecurityGroupOk() (*SecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *DeleteSecurityGroupRuleResponse) SetSecurityGroup(v SecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *DeleteSecurityGroupRuleResponse) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


