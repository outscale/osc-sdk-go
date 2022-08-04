# FiltersNet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptionsSetIds** | Pointer to **[]string** | The IDs of the DHCP options sets. | [optional] 
**IpRanges** | Pointer to **[]string** | The IP ranges for the Nets, in CIDR notation (for example, &#x60;10.0.0.0/16&#x60;). | [optional] 
**IsDefault** | Pointer to **bool** | If true, the Net used is the default one. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets. | [optional] 
**States** | Pointer to **[]string** | The states of the Nets (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the Nets. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the Nets. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the Nets, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

## Methods

### NewFiltersNet

`func NewFiltersNet() *FiltersNet`

NewFiltersNet instantiates a new FiltersNet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersNetWithDefaults

`func NewFiltersNetWithDefaults() *FiltersNet`

NewFiltersNetWithDefaults instantiates a new FiltersNet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpOptionsSetIds

`func (o *FiltersNet) GetDhcpOptionsSetIds() []string`

GetDhcpOptionsSetIds returns the DhcpOptionsSetIds field if non-nil, zero value otherwise.

### GetDhcpOptionsSetIdsOk

`func (o *FiltersNet) GetDhcpOptionsSetIdsOk() (*[]string, bool)`

GetDhcpOptionsSetIdsOk returns a tuple with the DhcpOptionsSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsSetIds

`func (o *FiltersNet) SetDhcpOptionsSetIds(v []string)`

SetDhcpOptionsSetIds sets DhcpOptionsSetIds field to given value.

### HasDhcpOptionsSetIds

`func (o *FiltersNet) HasDhcpOptionsSetIds() bool`

HasDhcpOptionsSetIds returns a boolean if a field has been set.

### GetIpRanges

`func (o *FiltersNet) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *FiltersNet) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *FiltersNet) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *FiltersNet) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetIsDefault

`func (o *FiltersNet) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *FiltersNet) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *FiltersNet) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *FiltersNet) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersNet) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersNet) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersNet) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersNet) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetStates

`func (o *FiltersNet) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersNet) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersNet) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersNet) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersNet) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersNet) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersNet) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersNet) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersNet) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersNet) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersNet) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersNet) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersNet) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersNet) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersNet) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersNet) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


