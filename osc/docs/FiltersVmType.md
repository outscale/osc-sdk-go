# FiltersVmType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BsuOptimized** | Pointer to **bool** | Indicates whether the VM is optimized for BSU I/O. | [optional] 
**MemorySizes** | Pointer to **[]float32** | The amounts of memory, in gibibytes (GiB). | [optional] 
**VcoreCounts** | Pointer to **[]int32** | The numbers of vCores. | [optional] 
**VmTypeNames** | Pointer to **[]string** | The names of the VM types. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). | [optional] 
**VolumeCounts** | Pointer to **[]int32** | The maximum number of ephemeral storage disks. | [optional] 
**VolumeSizes** | Pointer to **[]int32** | The size of one ephemeral storage disk, in gibibytes (GiB). | [optional] 

## Methods

### NewFiltersVmType

`func NewFiltersVmType() *FiltersVmType`

NewFiltersVmType instantiates a new FiltersVmType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersVmTypeWithDefaults

`func NewFiltersVmTypeWithDefaults() *FiltersVmType`

NewFiltersVmTypeWithDefaults instantiates a new FiltersVmType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsuOptimized

`func (o *FiltersVmType) GetBsuOptimized() bool`

GetBsuOptimized returns the BsuOptimized field if non-nil, zero value otherwise.

### GetBsuOptimizedOk

`func (o *FiltersVmType) GetBsuOptimizedOk() (*bool, bool)`

GetBsuOptimizedOk returns a tuple with the BsuOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsuOptimized

`func (o *FiltersVmType) SetBsuOptimized(v bool)`

SetBsuOptimized sets BsuOptimized field to given value.

### HasBsuOptimized

`func (o *FiltersVmType) HasBsuOptimized() bool`

HasBsuOptimized returns a boolean if a field has been set.

### GetMemorySizes

`func (o *FiltersVmType) GetMemorySizes() []float32`

GetMemorySizes returns the MemorySizes field if non-nil, zero value otherwise.

### GetMemorySizesOk

`func (o *FiltersVmType) GetMemorySizesOk() (*[]float32, bool)`

GetMemorySizesOk returns a tuple with the MemorySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizes

`func (o *FiltersVmType) SetMemorySizes(v []float32)`

SetMemorySizes sets MemorySizes field to given value.

### HasMemorySizes

`func (o *FiltersVmType) HasMemorySizes() bool`

HasMemorySizes returns a boolean if a field has been set.

### GetVcoreCounts

`func (o *FiltersVmType) GetVcoreCounts() []int32`

GetVcoreCounts returns the VcoreCounts field if non-nil, zero value otherwise.

### GetVcoreCountsOk

`func (o *FiltersVmType) GetVcoreCountsOk() (*[]int32, bool)`

GetVcoreCountsOk returns a tuple with the VcoreCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcoreCounts

`func (o *FiltersVmType) SetVcoreCounts(v []int32)`

SetVcoreCounts sets VcoreCounts field to given value.

### HasVcoreCounts

`func (o *FiltersVmType) HasVcoreCounts() bool`

HasVcoreCounts returns a boolean if a field has been set.

### GetVmTypeNames

`func (o *FiltersVmType) GetVmTypeNames() []string`

GetVmTypeNames returns the VmTypeNames field if non-nil, zero value otherwise.

### GetVmTypeNamesOk

`func (o *FiltersVmType) GetVmTypeNamesOk() (*[]string, bool)`

GetVmTypeNamesOk returns a tuple with the VmTypeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTypeNames

`func (o *FiltersVmType) SetVmTypeNames(v []string)`

SetVmTypeNames sets VmTypeNames field to given value.

### HasVmTypeNames

`func (o *FiltersVmType) HasVmTypeNames() bool`

HasVmTypeNames returns a boolean if a field has been set.

### GetVolumeCounts

`func (o *FiltersVmType) GetVolumeCounts() []int32`

GetVolumeCounts returns the VolumeCounts field if non-nil, zero value otherwise.

### GetVolumeCountsOk

`func (o *FiltersVmType) GetVolumeCountsOk() (*[]int32, bool)`

GetVolumeCountsOk returns a tuple with the VolumeCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCounts

`func (o *FiltersVmType) SetVolumeCounts(v []int32)`

SetVolumeCounts sets VolumeCounts field to given value.

### HasVolumeCounts

`func (o *FiltersVmType) HasVolumeCounts() bool`

HasVolumeCounts returns a boolean if a field has been set.

### GetVolumeSizes

`func (o *FiltersVmType) GetVolumeSizes() []int32`

GetVolumeSizes returns the VolumeSizes field if non-nil, zero value otherwise.

### GetVolumeSizesOk

`func (o *FiltersVmType) GetVolumeSizesOk() (*[]int32, bool)`

GetVolumeSizesOk returns a tuple with the VolumeSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSizes

`func (o *FiltersVmType) SetVolumeSizes(v []int32)`

SetVolumeSizes sets VolumeSizes field to given value.

### HasVolumeSizes

`func (o *FiltersVmType) HasVolumeSizes() bool`

HasVolumeSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


