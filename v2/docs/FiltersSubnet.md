# FiltersSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableIpsCounts** | Pointer to **[]int32** | The number of available IPs. | [optional] 
**IpRanges** | Pointer to **[]string** | The IP ranges in the Subnets, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets in which the Subnets are. | [optional] 
**States** | Pointer to **[]string** | The states of the Subnets (&#x60;pending&#x60; \\| &#x60;available&#x60;). | [optional] 
**SubnetIds** | Pointer to **[]string** | The IDs of the Subnets. | [optional] 
**SubregionNames** | Pointer to **[]string** | The names of the Subregions in which the Subnets are located. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the Subnets. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the Subnets. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the Subnets, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 

## Methods

### NewFiltersSubnet

`func NewFiltersSubnet() *FiltersSubnet`

NewFiltersSubnet instantiates a new FiltersSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersSubnetWithDefaults

`func NewFiltersSubnetWithDefaults() *FiltersSubnet`

NewFiltersSubnetWithDefaults instantiates a new FiltersSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableIpsCounts

`func (o *FiltersSubnet) GetAvailableIpsCounts() []int32`

GetAvailableIpsCounts returns the AvailableIpsCounts field if non-nil, zero value otherwise.

### GetAvailableIpsCountsOk

`func (o *FiltersSubnet) GetAvailableIpsCountsOk() (*[]int32, bool)`

GetAvailableIpsCountsOk returns a tuple with the AvailableIpsCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIpsCounts

`func (o *FiltersSubnet) SetAvailableIpsCounts(v []int32)`

SetAvailableIpsCounts sets AvailableIpsCounts field to given value.

### HasAvailableIpsCounts

`func (o *FiltersSubnet) HasAvailableIpsCounts() bool`

HasAvailableIpsCounts returns a boolean if a field has been set.

### GetIpRanges

`func (o *FiltersSubnet) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *FiltersSubnet) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *FiltersSubnet) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *FiltersSubnet) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersSubnet) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersSubnet) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersSubnet) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersSubnet) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetStates

`func (o *FiltersSubnet) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersSubnet) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersSubnet) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersSubnet) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetSubnetIds

`func (o *FiltersSubnet) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *FiltersSubnet) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *FiltersSubnet) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *FiltersSubnet) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### GetSubregionNames

`func (o *FiltersSubnet) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersSubnet) GetSubregionNamesOk() (*[]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionNames

`func (o *FiltersSubnet) SetSubregionNames(v []string)`

SetSubregionNames sets SubregionNames field to given value.

### HasSubregionNames

`func (o *FiltersSubnet) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersSubnet) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersSubnet) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersSubnet) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersSubnet) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersSubnet) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersSubnet) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersSubnet) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersSubnet) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersSubnet) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersSubnet) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersSubnet) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersSubnet) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


