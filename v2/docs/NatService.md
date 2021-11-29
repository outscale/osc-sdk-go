# NatService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatServiceId** | Pointer to **string** | The ID of the NAT service. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net in which the NAT service is. | [optional] 
**PublicIps** | Pointer to [**[]PublicIpLight**](PublicIpLight.md) | Information about the public IP or IPs associated with the NAT service. | [optional] 
**State** | Pointer to **string** | The state of the NAT service (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet in which the NAT service is. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the NAT service. | [optional] 

## Methods

### NewNatService

`func NewNatService() *NatService`

NewNatService instantiates a new NatService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatServiceWithDefaults

`func NewNatServiceWithDefaults() *NatService`

NewNatServiceWithDefaults instantiates a new NatService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatServiceId

`func (o *NatService) GetNatServiceId() string`

GetNatServiceId returns the NatServiceId field if non-nil, zero value otherwise.

### GetNatServiceIdOk

`func (o *NatService) GetNatServiceIdOk() (*string, bool)`

GetNatServiceIdOk returns a tuple with the NatServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatServiceId

`func (o *NatService) SetNatServiceId(v string)`

SetNatServiceId sets NatServiceId field to given value.

### HasNatServiceId

`func (o *NatService) HasNatServiceId() bool`

HasNatServiceId returns a boolean if a field has been set.

### GetNetId

`func (o *NatService) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *NatService) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *NatService) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *NatService) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetPublicIps

`func (o *NatService) GetPublicIps() []PublicIpLight`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *NatService) GetPublicIpsOk() (*[]PublicIpLight, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *NatService) SetPublicIps(v []PublicIpLight)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *NatService) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### GetState

`func (o *NatService) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NatService) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NatService) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NatService) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubnetId

`func (o *NatService) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *NatService) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *NatService) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *NatService) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTags

`func (o *NatService) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NatService) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NatService) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NatService) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


