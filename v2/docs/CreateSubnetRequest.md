# CreateSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**IpRange** | **string** | The IP range in the Subnet, in CIDR notation (for example, &#x60;10.0.0.0/16&#x60;).&lt;br /&gt; The IP range of the Subnet can be either the same as the Net one if you create only a single Subnet in this Net, or a subset of the Net one. In case of several Subnets in a Net, their IP ranges must not overlap. The smallest Subnet you can create uses a /29 netmask (eight IPs). For more information, see [About Nets](https://docs.outscale.com/en/userguide/About-Nets.html). | 
**NetId** | **string** | The ID of the Net for which you want to create a Subnet. | 
**SubregionName** | Pointer to **string** | The name of the Subregion in which you want to create the Subnet. | [optional] 

## Methods

### NewCreateSubnetRequest

`func NewCreateSubnetRequest(ipRange string, netId string, ) *CreateSubnetRequest`

NewCreateSubnetRequest instantiates a new CreateSubnetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnetRequestWithDefaults

`func NewCreateSubnetRequestWithDefaults() *CreateSubnetRequest`

NewCreateSubnetRequestWithDefaults instantiates a new CreateSubnetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateSubnetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSubnetRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateSubnetRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateSubnetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetIpRange

`func (o *CreateSubnetRequest) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *CreateSubnetRequest) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *CreateSubnetRequest) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.


### GetNetId

`func (o *CreateSubnetRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *CreateSubnetRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *CreateSubnetRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.


### GetSubregionName

`func (o *CreateSubnetRequest) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *CreateSubnetRequest) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *CreateSubnetRequest) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *CreateSubnetRequest) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


