# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableIpsCount** | Pointer to **int32** | The number of available IPs in the Subnets. | [optional] 
**IpRange** | Pointer to **string** | The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**MapPublicIpOnLaunch** | Pointer to **bool** | If true, a public IP is assigned to the network interface cards (NICs) created in the specified Subnet. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net in which the Subnet is. | [optional] 
**State** | Pointer to **string** | The state of the Subnet (&#x60;pending&#x60; \\| &#x60;available&#x60;). | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet. | [optional] 
**SubregionName** | Pointer to **string** | The name of the Subregion in which the Subnet is located. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the Subnet. | [optional] 

## Methods

### NewSubnet

`func NewSubnet() *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableIpsCount

`func (o *Subnet) GetAvailableIpsCount() int32`

GetAvailableIpsCount returns the AvailableIpsCount field if non-nil, zero value otherwise.

### GetAvailableIpsCountOk

`func (o *Subnet) GetAvailableIpsCountOk() (*int32, bool)`

GetAvailableIpsCountOk returns a tuple with the AvailableIpsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIpsCount

`func (o *Subnet) SetAvailableIpsCount(v int32)`

SetAvailableIpsCount sets AvailableIpsCount field to given value.

### HasAvailableIpsCount

`func (o *Subnet) HasAvailableIpsCount() bool`

HasAvailableIpsCount returns a boolean if a field has been set.

### GetIpRange

`func (o *Subnet) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *Subnet) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *Subnet) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *Subnet) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetMapPublicIpOnLaunch

`func (o *Subnet) GetMapPublicIpOnLaunch() bool`

GetMapPublicIpOnLaunch returns the MapPublicIpOnLaunch field if non-nil, zero value otherwise.

### GetMapPublicIpOnLaunchOk

`func (o *Subnet) GetMapPublicIpOnLaunchOk() (*bool, bool)`

GetMapPublicIpOnLaunchOk returns a tuple with the MapPublicIpOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapPublicIpOnLaunch

`func (o *Subnet) SetMapPublicIpOnLaunch(v bool)`

SetMapPublicIpOnLaunch sets MapPublicIpOnLaunch field to given value.

### HasMapPublicIpOnLaunch

`func (o *Subnet) HasMapPublicIpOnLaunch() bool`

HasMapPublicIpOnLaunch returns a boolean if a field has been set.

### GetNetId

`func (o *Subnet) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *Subnet) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *Subnet) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *Subnet) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetState

`func (o *Subnet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Subnet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Subnet) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Subnet) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubnetId

`func (o *Subnet) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Subnet) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Subnet) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *Subnet) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetSubregionName

`func (o *Subnet) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *Subnet) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *Subnet) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *Subnet) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### GetTags

`func (o *Subnet) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Subnet) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Subnet) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Subnet) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


