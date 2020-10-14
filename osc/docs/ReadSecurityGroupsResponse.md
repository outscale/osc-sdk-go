# ReadSecurityGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroup**](SecurityGroup.md) | Information about one or more security groups. | [optional] 

## Methods

### NewReadSecurityGroupsResponse

`func NewReadSecurityGroupsResponse() *ReadSecurityGroupsResponse`

NewReadSecurityGroupsResponse instantiates a new ReadSecurityGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSecurityGroupsResponseWithDefaults

`func NewReadSecurityGroupsResponseWithDefaults() *ReadSecurityGroupsResponse`

NewReadSecurityGroupsResponseWithDefaults instantiates a new ReadSecurityGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadSecurityGroupsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSecurityGroupsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadSecurityGroupsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadSecurityGroupsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ReadSecurityGroupsResponse) GetSecurityGroups() []SecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ReadSecurityGroupsResponse) GetSecurityGroupsOk() (*[]SecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ReadSecurityGroupsResponse) SetSecurityGroups(v []SecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ReadSecurityGroupsResponse) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


