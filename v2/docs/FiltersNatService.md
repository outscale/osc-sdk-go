# FiltersNatService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatServiceIds** | Pointer to **[]string** | The IDs of the NAT services. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets in which the NAT services are. | [optional] 
**States** | Pointer to **[]string** | The states of the NAT services (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**SubnetIds** | Pointer to **[]string** | The IDs of the Subnets in which the NAT services are. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the NAT services. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the NAT services. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the NAT services, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

## Methods

### NewFiltersNatService

`func NewFiltersNatService() *FiltersNatService`

NewFiltersNatService instantiates a new FiltersNatService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersNatServiceWithDefaults

`func NewFiltersNatServiceWithDefaults() *FiltersNatService`

NewFiltersNatServiceWithDefaults instantiates a new FiltersNatService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatServiceIds

`func (o *FiltersNatService) GetNatServiceIds() []string`

GetNatServiceIds returns the NatServiceIds field if non-nil, zero value otherwise.

### GetNatServiceIdsOk

`func (o *FiltersNatService) GetNatServiceIdsOk() (*[]string, bool)`

GetNatServiceIdsOk returns a tuple with the NatServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatServiceIds

`func (o *FiltersNatService) SetNatServiceIds(v []string)`

SetNatServiceIds sets NatServiceIds field to given value.

### HasNatServiceIds

`func (o *FiltersNatService) HasNatServiceIds() bool`

HasNatServiceIds returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersNatService) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersNatService) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersNatService) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersNatService) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetStates

`func (o *FiltersNatService) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersNatService) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersNatService) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersNatService) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetSubnetIds

`func (o *FiltersNatService) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *FiltersNatService) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *FiltersNatService) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *FiltersNatService) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersNatService) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersNatService) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersNatService) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersNatService) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersNatService) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersNatService) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersNatService) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersNatService) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersNatService) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersNatService) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersNatService) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersNatService) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


