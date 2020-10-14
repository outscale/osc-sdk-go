# NicLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the NIC. | [optional] 
**Description** | Pointer to **string** | The description of the NIC. | [optional] 
**IsSourceDestChecked** | Pointer to **bool** | (Net only) If &#x60;true&#x60;, the source/destination check is enabled. If &#x60;false&#x60;, it is disabled. This value must be &#x60;false&#x60; for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**LinkNic** | Pointer to [**LinkNicLight**](LinkNicLight.md) |  | [optional] 
**LinkPublicIp** | Pointer to [**LinkPublicIpLightForVm**](LinkPublicIpLightForVm.md) |  | [optional] 
**MacAddress** | Pointer to **string** | The Media Access Control (MAC) address of the NIC. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the NIC. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC. | [optional] 
**PrivateDnsName** | Pointer to **string** | The name of the private DNS. | [optional] 
**PrivateIps** | Pointer to [**[]PrivateIpLightForVm**](PrivateIpLightForVm.md) | The private IP address or addresses of the NIC. | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupLight**](SecurityGroupLight.md) | One or more IDs of security groups for the NIC. | [optional] 
**State** | Pointer to **string** | The state of the NIC (&#x60;available&#x60; \\| &#x60;attaching&#x60; \\| &#x60;in-use&#x60; \\| &#x60;detaching&#x60;). | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet for the NIC. | [optional] 

## Methods

### NewNicLight

`func NewNicLight() *NicLight`

NewNicLight instantiates a new NicLight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicLightWithDefaults

`func NewNicLightWithDefaults() *NicLight`

NewNicLightWithDefaults instantiates a new NicLight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NicLight) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NicLight) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NicLight) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *NicLight) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDescription

`func (o *NicLight) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NicLight) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NicLight) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NicLight) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsSourceDestChecked

`func (o *NicLight) GetIsSourceDestChecked() bool`

GetIsSourceDestChecked returns the IsSourceDestChecked field if non-nil, zero value otherwise.

### GetIsSourceDestCheckedOk

`func (o *NicLight) GetIsSourceDestCheckedOk() (*bool, bool)`

GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceDestChecked

`func (o *NicLight) SetIsSourceDestChecked(v bool)`

SetIsSourceDestChecked sets IsSourceDestChecked field to given value.

### HasIsSourceDestChecked

`func (o *NicLight) HasIsSourceDestChecked() bool`

HasIsSourceDestChecked returns a boolean if a field has been set.

### GetLinkNic

`func (o *NicLight) GetLinkNic() LinkNicLight`

GetLinkNic returns the LinkNic field if non-nil, zero value otherwise.

### GetLinkNicOk

`func (o *NicLight) GetLinkNicOk() (*LinkNicLight, bool)`

GetLinkNicOk returns a tuple with the LinkNic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNic

`func (o *NicLight) SetLinkNic(v LinkNicLight)`

SetLinkNic sets LinkNic field to given value.

### HasLinkNic

`func (o *NicLight) HasLinkNic() bool`

HasLinkNic returns a boolean if a field has been set.

### GetLinkPublicIp

`func (o *NicLight) GetLinkPublicIp() LinkPublicIpLightForVm`

GetLinkPublicIp returns the LinkPublicIp field if non-nil, zero value otherwise.

### GetLinkPublicIpOk

`func (o *NicLight) GetLinkPublicIpOk() (*LinkPublicIpLightForVm, bool)`

GetLinkPublicIpOk returns a tuple with the LinkPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIp

`func (o *NicLight) SetLinkPublicIp(v LinkPublicIpLightForVm)`

SetLinkPublicIp sets LinkPublicIp field to given value.

### HasLinkPublicIp

`func (o *NicLight) HasLinkPublicIp() bool`

HasLinkPublicIp returns a boolean if a field has been set.

### GetMacAddress

`func (o *NicLight) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NicLight) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NicLight) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NicLight) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetNetId

`func (o *NicLight) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *NicLight) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *NicLight) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *NicLight) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetNicId

`func (o *NicLight) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *NicLight) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *NicLight) SetNicId(v string)`

SetNicId sets NicId field to given value.

### HasNicId

`func (o *NicLight) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### GetPrivateDnsName

`func (o *NicLight) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *NicLight) GetPrivateDnsNameOk() (*string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsName

`func (o *NicLight) SetPrivateDnsName(v string)`

SetPrivateDnsName sets PrivateDnsName field to given value.

### HasPrivateDnsName

`func (o *NicLight) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### GetPrivateIps

`func (o *NicLight) GetPrivateIps() []PrivateIpLightForVm`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *NicLight) GetPrivateIpsOk() (*[]PrivateIpLightForVm, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *NicLight) SetPrivateIps(v []PrivateIpLightForVm)`

SetPrivateIps sets PrivateIps field to given value.

### HasPrivateIps

`func (o *NicLight) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *NicLight) GetSecurityGroups() []SecurityGroupLight`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *NicLight) GetSecurityGroupsOk() (*[]SecurityGroupLight, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *NicLight) SetSecurityGroups(v []SecurityGroupLight)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *NicLight) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetState

`func (o *NicLight) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NicLight) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NicLight) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NicLight) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubnetId

`func (o *NicLight) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *NicLight) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *NicLight) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *NicLight) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


