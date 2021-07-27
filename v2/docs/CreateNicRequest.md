# CreateNicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description for the NIC. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PrivateIps** | Pointer to [**[]PrivateIpLight**](PrivateIpLight.md) | The primary private IP address for the NIC.&lt;br /&gt; This IP address must be within the IP address range of the Subnet that you specify with the &#x60;SubnetId&#x60; attribute.&lt;br /&gt; If you do not specify this attribute, a random private IP address is selected within the IP address range of the Subnet. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security groups for the NIC. | [optional] 
**SubnetId** | **string** | The ID of the Subnet in which you want to create the NIC. | 

## Methods

### NewCreateNicRequest

`func NewCreateNicRequest(subnetId string, ) *CreateNicRequest`

NewCreateNicRequest instantiates a new CreateNicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNicRequestWithDefaults

`func NewCreateNicRequestWithDefaults() *CreateNicRequest`

NewCreateNicRequestWithDefaults instantiates a new CreateNicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateNicRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNicRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNicRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNicRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateNicRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateNicRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateNicRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateNicRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPrivateIps

`func (o *CreateNicRequest) GetPrivateIps() []PrivateIpLight`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *CreateNicRequest) GetPrivateIpsOk() (*[]PrivateIpLight, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *CreateNicRequest) SetPrivateIps(v []PrivateIpLight)`

SetPrivateIps sets PrivateIps field to given value.

### HasPrivateIps

`func (o *CreateNicRequest) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *CreateNicRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *CreateNicRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *CreateNicRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *CreateNicRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSubnetId

`func (o *CreateNicRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateNicRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateNicRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


