# FiltersNic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkNicVmIds** | Pointer to **[]string** | The IDs of the VMs the NICs are attached to. | [optional] 
**NicIds** | Pointer to **[]string** | The IDs of the NICs. | [optional] 
**PrivateIpsPrivateIps** | Pointer to **[]string** | The private IP addresses of the NICs. | [optional] 
**SubnetIds** | Pointer to **[]string** | The IDs of the Subnets for the NICs. | [optional] 

## Methods

### NewFiltersNic

`func NewFiltersNic() *FiltersNic`

NewFiltersNic instantiates a new FiltersNic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersNicWithDefaults

`func NewFiltersNicWithDefaults() *FiltersNic`

NewFiltersNicWithDefaults instantiates a new FiltersNic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkNicVmIds

`func (o *FiltersNic) GetLinkNicVmIds() []string`

GetLinkNicVmIds returns the LinkNicVmIds field if non-nil, zero value otherwise.

### GetLinkNicVmIdsOk

`func (o *FiltersNic) GetLinkNicVmIdsOk() (*[]string, bool)`

GetLinkNicVmIdsOk returns a tuple with the LinkNicVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicVmIds

`func (o *FiltersNic) SetLinkNicVmIds(v []string)`

SetLinkNicVmIds sets LinkNicVmIds field to given value.

### HasLinkNicVmIds

`func (o *FiltersNic) HasLinkNicVmIds() bool`

HasLinkNicVmIds returns a boolean if a field has been set.

### GetNicIds

`func (o *FiltersNic) GetNicIds() []string`

GetNicIds returns the NicIds field if non-nil, zero value otherwise.

### GetNicIdsOk

`func (o *FiltersNic) GetNicIdsOk() (*[]string, bool)`

GetNicIdsOk returns a tuple with the NicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicIds

`func (o *FiltersNic) SetNicIds(v []string)`

SetNicIds sets NicIds field to given value.

### HasNicIds

`func (o *FiltersNic) HasNicIds() bool`

HasNicIds returns a boolean if a field has been set.

### GetPrivateIpsPrivateIps

`func (o *FiltersNic) GetPrivateIpsPrivateIps() []string`

GetPrivateIpsPrivateIps returns the PrivateIpsPrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsPrivateIpsOk

`func (o *FiltersNic) GetPrivateIpsPrivateIpsOk() (*[]string, bool)`

GetPrivateIpsPrivateIpsOk returns a tuple with the PrivateIpsPrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpsPrivateIps

`func (o *FiltersNic) SetPrivateIpsPrivateIps(v []string)`

SetPrivateIpsPrivateIps sets PrivateIpsPrivateIps field to given value.

### HasPrivateIpsPrivateIps

`func (o *FiltersNic) HasPrivateIpsPrivateIps() bool`

HasPrivateIpsPrivateIps returns a boolean if a field has been set.

### GetSubnetIds

`func (o *FiltersNic) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *FiltersNic) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *FiltersNic) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *FiltersNic) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


