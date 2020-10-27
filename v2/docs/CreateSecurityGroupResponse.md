# CreateSecurityGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**SecurityGroup** | Pointer to [**SecurityGroup**](SecurityGroup.md) |  | [optional] 

## Methods

### NewCreateSecurityGroupResponse

`func NewCreateSecurityGroupResponse() *CreateSecurityGroupResponse`

NewCreateSecurityGroupResponse instantiates a new CreateSecurityGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupResponseWithDefaults

`func NewCreateSecurityGroupResponseWithDefaults() *CreateSecurityGroupResponse`

NewCreateSecurityGroupResponseWithDefaults instantiates a new CreateSecurityGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateSecurityGroupResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateSecurityGroupResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateSecurityGroupResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateSecurityGroupResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *CreateSecurityGroupResponse) GetSecurityGroup() SecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *CreateSecurityGroupResponse) GetSecurityGroupOk() (*SecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *CreateSecurityGroupResponse) SetSecurityGroup(v SecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *CreateSecurityGroupResponse) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


