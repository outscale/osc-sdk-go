# FiltersDedicatedGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuGenerations** | Pointer to **[]int32** | The processor generation for the VMs in the dedicated group (for example, &#x60;4&#x60;). | [optional] 
**DedicatedGroupIds** | Pointer to **[]string** | The IDs of the dedicated groups. | [optional] 
**Names** | Pointer to **[]string** | The names of the dedicated groups. | [optional] 
**SubregionNames** | Pointer to **[]string** | The names of the Subregions in which the dedicated groups are located. | [optional] 

## Methods

### NewFiltersDedicatedGroup

`func NewFiltersDedicatedGroup() *FiltersDedicatedGroup`

NewFiltersDedicatedGroup instantiates a new FiltersDedicatedGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersDedicatedGroupWithDefaults

`func NewFiltersDedicatedGroupWithDefaults() *FiltersDedicatedGroup`

NewFiltersDedicatedGroupWithDefaults instantiates a new FiltersDedicatedGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuGenerations

`func (o *FiltersDedicatedGroup) GetCpuGenerations() []int32`

GetCpuGenerations returns the CpuGenerations field if non-nil, zero value otherwise.

### GetCpuGenerationsOk

`func (o *FiltersDedicatedGroup) GetCpuGenerationsOk() (*[]int32, bool)`

GetCpuGenerationsOk returns a tuple with the CpuGenerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuGenerations

`func (o *FiltersDedicatedGroup) SetCpuGenerations(v []int32)`

SetCpuGenerations sets CpuGenerations field to given value.

### HasCpuGenerations

`func (o *FiltersDedicatedGroup) HasCpuGenerations() bool`

HasCpuGenerations returns a boolean if a field has been set.

### GetDedicatedGroupIds

`func (o *FiltersDedicatedGroup) GetDedicatedGroupIds() []string`

GetDedicatedGroupIds returns the DedicatedGroupIds field if non-nil, zero value otherwise.

### GetDedicatedGroupIdsOk

`func (o *FiltersDedicatedGroup) GetDedicatedGroupIdsOk() (*[]string, bool)`

GetDedicatedGroupIdsOk returns a tuple with the DedicatedGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedGroupIds

`func (o *FiltersDedicatedGroup) SetDedicatedGroupIds(v []string)`

SetDedicatedGroupIds sets DedicatedGroupIds field to given value.

### HasDedicatedGroupIds

`func (o *FiltersDedicatedGroup) HasDedicatedGroupIds() bool`

HasDedicatedGroupIds returns a boolean if a field has been set.

### GetNames

`func (o *FiltersDedicatedGroup) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *FiltersDedicatedGroup) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *FiltersDedicatedGroup) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *FiltersDedicatedGroup) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSubregionNames

`func (o *FiltersDedicatedGroup) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersDedicatedGroup) GetSubregionNamesOk() (*[]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionNames

`func (o *FiltersDedicatedGroup) SetSubregionNames(v []string)`

SetSubregionNames sets SubregionNames field to given value.

### HasSubregionNames

`func (o *FiltersDedicatedGroup) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


