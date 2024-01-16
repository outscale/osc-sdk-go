# DedicatedGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owners of the dedicated group. | [optional] 
**CpuGeneration** | Pointer to **int32** | The processor generation. | [optional] 
**DedicatedGroupId** | Pointer to **string** | The ID of the dedicated group. | [optional] 
**Name** | Pointer to **string** | The name of the dedicated group. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets in the dedicated group. | [optional] 
**SubregionName** | Pointer to **string** | The name of the Subregion in which the dedicated group is located. | [optional] 
**VmIds** | Pointer to **[]string** | The IDs of the VMs in the dedicated group. | [optional] 

## Methods

### NewDedicatedGroup

`func NewDedicatedGroup() *DedicatedGroup`

NewDedicatedGroup instantiates a new DedicatedGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedGroupWithDefaults

`func NewDedicatedGroupWithDefaults() *DedicatedGroup`

NewDedicatedGroupWithDefaults instantiates a new DedicatedGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DedicatedGroup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DedicatedGroup) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DedicatedGroup) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *DedicatedGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCpuGeneration

`func (o *DedicatedGroup) GetCpuGeneration() int32`

GetCpuGeneration returns the CpuGeneration field if non-nil, zero value otherwise.

### GetCpuGenerationOk

`func (o *DedicatedGroup) GetCpuGenerationOk() (*int32, bool)`

GetCpuGenerationOk returns a tuple with the CpuGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuGeneration

`func (o *DedicatedGroup) SetCpuGeneration(v int32)`

SetCpuGeneration sets CpuGeneration field to given value.

### HasCpuGeneration

`func (o *DedicatedGroup) HasCpuGeneration() bool`

HasCpuGeneration returns a boolean if a field has been set.

### GetDedicatedGroupId

`func (o *DedicatedGroup) GetDedicatedGroupId() string`

GetDedicatedGroupId returns the DedicatedGroupId field if non-nil, zero value otherwise.

### GetDedicatedGroupIdOk

`func (o *DedicatedGroup) GetDedicatedGroupIdOk() (*string, bool)`

GetDedicatedGroupIdOk returns a tuple with the DedicatedGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedGroupId

`func (o *DedicatedGroup) SetDedicatedGroupId(v string)`

SetDedicatedGroupId sets DedicatedGroupId field to given value.

### HasDedicatedGroupId

`func (o *DedicatedGroup) HasDedicatedGroupId() bool`

HasDedicatedGroupId returns a boolean if a field has been set.

### GetName

`func (o *DedicatedGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DedicatedGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DedicatedGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DedicatedGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetIds

`func (o *DedicatedGroup) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *DedicatedGroup) GetNetIdsOk() (*[]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIds

`func (o *DedicatedGroup) SetNetIds(v []string)`

SetNetIds sets NetIds field to given value.

### HasNetIds

`func (o *DedicatedGroup) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### GetSubregionName

`func (o *DedicatedGroup) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *DedicatedGroup) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *DedicatedGroup) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *DedicatedGroup) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### GetVmIds

`func (o *DedicatedGroup) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *DedicatedGroup) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *DedicatedGroup) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *DedicatedGroup) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


