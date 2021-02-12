# CreateSecurityGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description for the security group, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters). | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the security group. | [optional] 
**SecurityGroupName** | **string** | The name of the security group.&lt;br /&gt; This name must not start with &#x60;sg-&#x60;.&lt;/br&gt; This name must be unique and contain between 1 and 255 ASCII characters. Accented letters are not allowed. | 

## Methods

### NewCreateSecurityGroupRequest

`func NewCreateSecurityGroupRequest(description string, securityGroupName string, ) *CreateSecurityGroupRequest`

NewCreateSecurityGroupRequest instantiates a new CreateSecurityGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRequestWithDefaults

`func NewCreateSecurityGroupRequestWithDefaults() *CreateSecurityGroupRequest`

NewCreateSecurityGroupRequestWithDefaults instantiates a new CreateSecurityGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateSecurityGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecurityGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSecurityGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDryRun

`func (o *CreateSecurityGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSecurityGroupRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateSecurityGroupRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateSecurityGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetId

`func (o *CreateSecurityGroupRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *CreateSecurityGroupRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *CreateSecurityGroupRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *CreateSecurityGroupRequest) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetSecurityGroupName

`func (o *CreateSecurityGroupRequest) GetSecurityGroupName() string`

GetSecurityGroupName returns the SecurityGroupName field if non-nil, zero value otherwise.

### GetSecurityGroupNameOk

`func (o *CreateSecurityGroupRequest) GetSecurityGroupNameOk() (*string, bool)`

GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupName

`func (o *CreateSecurityGroupRequest) SetSecurityGroupName(v string)`

SetSecurityGroupName sets SecurityGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


