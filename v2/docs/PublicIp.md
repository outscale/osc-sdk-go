# PublicIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkPublicIpId** | Pointer to **string** | (Required in a Net) The ID representing the association of the public IP with the VM or the NIC. | [optional] 
**NicAccountId** | Pointer to **string** | The account ID of the owner of the NIC. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC the public IP is associated with (if any). | [optional] 
**PrivateIp** | Pointer to **string** | The private IP address associated with the public IP. | [optional] 
**PublicIp** | Pointer to **string** | The public IP. | [optional] 
**PublicIpId** | Pointer to **string** | The allocation ID of the public IP. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the public IP. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM the public IP is associated with (if any). | [optional] 

## Methods

### NewPublicIp

`func NewPublicIp() *PublicIp`

NewPublicIp instantiates a new PublicIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIpWithDefaults

`func NewPublicIpWithDefaults() *PublicIp`

NewPublicIpWithDefaults instantiates a new PublicIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkPublicIpId

`func (o *PublicIp) GetLinkPublicIpId() string`

GetLinkPublicIpId returns the LinkPublicIpId field if non-nil, zero value otherwise.

### GetLinkPublicIpIdOk

`func (o *PublicIp) GetLinkPublicIpIdOk() (*string, bool)`

GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIpId

`func (o *PublicIp) SetLinkPublicIpId(v string)`

SetLinkPublicIpId sets LinkPublicIpId field to given value.

### HasLinkPublicIpId

`func (o *PublicIp) HasLinkPublicIpId() bool`

HasLinkPublicIpId returns a boolean if a field has been set.

### GetNicAccountId

`func (o *PublicIp) GetNicAccountId() string`

GetNicAccountId returns the NicAccountId field if non-nil, zero value otherwise.

### GetNicAccountIdOk

`func (o *PublicIp) GetNicAccountIdOk() (*string, bool)`

GetNicAccountIdOk returns a tuple with the NicAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicAccountId

`func (o *PublicIp) SetNicAccountId(v string)`

SetNicAccountId sets NicAccountId field to given value.

### HasNicAccountId

`func (o *PublicIp) HasNicAccountId() bool`

HasNicAccountId returns a boolean if a field has been set.

### GetNicId

`func (o *PublicIp) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *PublicIp) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *PublicIp) SetNicId(v string)`

SetNicId sets NicId field to given value.

### HasNicId

`func (o *PublicIp) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### GetPrivateIp

`func (o *PublicIp) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *PublicIp) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *PublicIp) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *PublicIp) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *PublicIp) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *PublicIp) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *PublicIp) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *PublicIp) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicIpId

`func (o *PublicIp) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *PublicIp) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *PublicIp) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.

### HasPublicIpId

`func (o *PublicIp) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### GetTags

`func (o *PublicIp) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PublicIp) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PublicIp) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PublicIp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmId

`func (o *PublicIp) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *PublicIp) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *PublicIp) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *PublicIp) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


