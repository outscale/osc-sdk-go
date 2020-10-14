# UpdateNicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A new description for the NIC. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LinkNic** | Pointer to [**LinkNicToUpdate**](LinkNicToUpdate.md) |  | [optional] 
**NicId** | **string** | The ID of the NIC you want to modify. | 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security groups for the NIC.&lt;br /&gt; You must specify at least one group, even if you use the default security group in the Net. | [optional] 

## Methods

### NewUpdateNicRequest

`func NewUpdateNicRequest(nicId string, ) *UpdateNicRequest`

NewUpdateNicRequest instantiates a new UpdateNicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNicRequestWithDefaults

`func NewUpdateNicRequestWithDefaults() *UpdateNicRequest`

NewUpdateNicRequestWithDefaults instantiates a new UpdateNicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateNicRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNicRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNicRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNicRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateNicRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateNicRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateNicRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateNicRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLinkNic

`func (o *UpdateNicRequest) GetLinkNic() LinkNicToUpdate`

GetLinkNic returns the LinkNic field if non-nil, zero value otherwise.

### GetLinkNicOk

`func (o *UpdateNicRequest) GetLinkNicOk() (*LinkNicToUpdate, bool)`

GetLinkNicOk returns a tuple with the LinkNic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNic

`func (o *UpdateNicRequest) SetLinkNic(v LinkNicToUpdate)`

SetLinkNic sets LinkNic field to given value.

### HasLinkNic

`func (o *UpdateNicRequest) HasLinkNic() bool`

HasLinkNic returns a boolean if a field has been set.

### GetNicId

`func (o *UpdateNicRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *UpdateNicRequest) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *UpdateNicRequest) SetNicId(v string)`

SetNicId sets NicId field to given value.


### GetSecurityGroupIds

`func (o *UpdateNicRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *UpdateNicRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *UpdateNicRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *UpdateNicRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


