# Nic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the NIC. | [optional] 
**Description** | Pointer to **string** | The description of the NIC. | [optional] 
**IsSourceDestChecked** | Pointer to **bool** | (Net only) If &#x60;true&#x60;, the source/destination check is enabled. If &#x60;false&#x60;, it is disabled. This value must be &#x60;false&#x60; for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**LinkNic** | Pointer to [**LinkNic**](LinkNic.md) |  | [optional] 
**LinkPublicIp** | Pointer to [**LinkPublicIp**](LinkPublicIp.md) |  | [optional] 
**MacAddress** | Pointer to **string** | The Media Access Control (MAC) address of the NIC. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the NIC. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC. | [optional] 
**PrivateDnsName** | Pointer to **string** | The name of the private DNS. | [optional] 
**PrivateIps** | Pointer to [**[]PrivateIp**](PrivateIp.md) | The private IP addresses of the NIC. | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupLight**](SecurityGroupLight.md) | One or more IDs of security groups for the NIC. | [optional] 
**State** | Pointer to **string** | The state of the NIC (&#x60;available&#x60; \\| &#x60;attaching&#x60; \\| &#x60;in-use&#x60; \\| &#x60;detaching&#x60;). | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet. | [optional] 
**SubregionName** | Pointer to **string** | The Subregion in which the NIC is located. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the NIC. | [optional] 

## Methods

### NewNic

`func NewNic() *Nic`

NewNic instantiates a new Nic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicWithDefaults

`func NewNicWithDefaults() *Nic`

NewNicWithDefaults instantiates a new Nic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Nic) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Nic) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Nic) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Nic) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDescription

`func (o *Nic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Nic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Nic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Nic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsSourceDestChecked

`func (o *Nic) GetIsSourceDestChecked() bool`

GetIsSourceDestChecked returns the IsSourceDestChecked field if non-nil, zero value otherwise.

### GetIsSourceDestCheckedOk

`func (o *Nic) GetIsSourceDestCheckedOk() (*bool, bool)`

GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceDestChecked

`func (o *Nic) SetIsSourceDestChecked(v bool)`

SetIsSourceDestChecked sets IsSourceDestChecked field to given value.

### HasIsSourceDestChecked

`func (o *Nic) HasIsSourceDestChecked() bool`

HasIsSourceDestChecked returns a boolean if a field has been set.

### GetLinkNic

`func (o *Nic) GetLinkNic() LinkNic`

GetLinkNic returns the LinkNic field if non-nil, zero value otherwise.

### GetLinkNicOk

`func (o *Nic) GetLinkNicOk() (*LinkNic, bool)`

GetLinkNicOk returns a tuple with the LinkNic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNic

`func (o *Nic) SetLinkNic(v LinkNic)`

SetLinkNic sets LinkNic field to given value.

### HasLinkNic

`func (o *Nic) HasLinkNic() bool`

HasLinkNic returns a boolean if a field has been set.

### GetLinkPublicIp

`func (o *Nic) GetLinkPublicIp() LinkPublicIp`

GetLinkPublicIp returns the LinkPublicIp field if non-nil, zero value otherwise.

### GetLinkPublicIpOk

`func (o *Nic) GetLinkPublicIpOk() (*LinkPublicIp, bool)`

GetLinkPublicIpOk returns a tuple with the LinkPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIp

`func (o *Nic) SetLinkPublicIp(v LinkPublicIp)`

SetLinkPublicIp sets LinkPublicIp field to given value.

### HasLinkPublicIp

`func (o *Nic) HasLinkPublicIp() bool`

HasLinkPublicIp returns a boolean if a field has been set.

### GetMacAddress

`func (o *Nic) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Nic) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Nic) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Nic) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetNetId

`func (o *Nic) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *Nic) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *Nic) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *Nic) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetNicId

`func (o *Nic) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *Nic) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *Nic) SetNicId(v string)`

SetNicId sets NicId field to given value.

### HasNicId

`func (o *Nic) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### GetPrivateDnsName

`func (o *Nic) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *Nic) GetPrivateDnsNameOk() (*string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsName

`func (o *Nic) SetPrivateDnsName(v string)`

SetPrivateDnsName sets PrivateDnsName field to given value.

### HasPrivateDnsName

`func (o *Nic) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### GetPrivateIps

`func (o *Nic) GetPrivateIps() []PrivateIp`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *Nic) GetPrivateIpsOk() (*[]PrivateIp, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *Nic) SetPrivateIps(v []PrivateIp)`

SetPrivateIps sets PrivateIps field to given value.

### HasPrivateIps

`func (o *Nic) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *Nic) GetSecurityGroups() []SecurityGroupLight`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *Nic) GetSecurityGroupsOk() (*[]SecurityGroupLight, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *Nic) SetSecurityGroups(v []SecurityGroupLight)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *Nic) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetState

`func (o *Nic) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Nic) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Nic) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Nic) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubnetId

`func (o *Nic) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Nic) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Nic) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *Nic) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetSubregionName

`func (o *Nic) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *Nic) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *Nic) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *Nic) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### GetTags

`func (o *Nic) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Nic) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Nic) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Nic) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


