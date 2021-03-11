# FiltersNetAccessPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAccessPointIds** | Pointer to **[]string** | The IDs of the Net access points. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets. | [optional] 
**ServiceNames** | Pointer to **[]string** | The names of the services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices). | [optional] 
**States** | Pointer to **[]string** | The states of the Net access points (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the Net access points. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the Net access points. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the Net access points, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

## Methods

### NewFiltersNetAccessPoint

`func NewFiltersNetAccessPoint() *FiltersNetAccessPoint`

NewFiltersNetAccessPoint instantiates a new FiltersNetAccessPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersNetAccessPointWithDefaults

`func NewFiltersNetAccessPointWithDefaults() *FiltersNetAccessPoint`

NewFiltersNetAccessPointWithDefaults instantiates a new FiltersNetAccessPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetAccessPointIds

`func (o *FiltersNetAccessPoint) GetNetAccessPointIds() []string`

GetNetAccessPointIds returns the NetAccessPointIds field if non-nil, zero value otherwise.

### GetNetAccessPointIdsOk

`func (o *FiltersNetAccessPoint) GetNetAccessPointIdsOk() (*[]string, bool)`

GetNetAccessPointIdsOk returns a tuple with the NetAccessPointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAccessPointIds

`func (o *FiltersNetAccessPoint) SetNetAccessPointIds(v []string)`

SetNetAccessPointIds sets NetAccessPointIds field to given value.

### HasNetAccessPointIds

`func (o *FiltersNetAccessPoint) HasNetAccessPointIds() bool`

HasNetAccessPointIds returns a boolean if a field has been set.

### GetNetIds

`func (o *FiltersNetAccessPoint) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersNetAccessPoint) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *FiltersNetAccessPoint) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *FiltersNetAccessPoint) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetServiceNames

`func (o *FiltersNetAccessPoint) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *FiltersNetAccessPoint) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *FiltersNetAccessPoint) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *FiltersNetAccessPoint) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.

### GetStates

`func (o *FiltersNetAccessPoint) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersNetAccessPoint) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *FiltersNetAccessPoint) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *FiltersNetAccessPoint) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersNetAccessPoint) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersNetAccessPoint) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersNetAccessPoint) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersNetAccessPoint) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersNetAccessPoint) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersNetAccessPoint) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersNetAccessPoint) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersNetAccessPoint) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersNetAccessPoint) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersNetAccessPoint) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersNetAccessPoint) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersNetAccessPoint) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


