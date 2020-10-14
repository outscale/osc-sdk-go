# FlexibleGpuCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generations** | Pointer to **[]string** | The generations of VMs that the fGPU is compatible with. | [optional] 
**MaxCpu** | Pointer to **int32** | The maximum number of VM vCores that the fGPU is compatible with. | [optional] 
**MaxRam** | Pointer to **int32** | The maximum amount of VM memory that the fGPU is compatible with. | [optional] 
**ModelName** | Pointer to **string** | The model of fGPU. | [optional] 
**VRam** | Pointer to **int32** | The amount of video RAM (VRAM) of the fGPU. | [optional] 

## Methods

### NewFlexibleGpuCatalog

`func NewFlexibleGpuCatalog() *FlexibleGpuCatalog`

NewFlexibleGpuCatalog instantiates a new FlexibleGpuCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexibleGpuCatalogWithDefaults

`func NewFlexibleGpuCatalogWithDefaults() *FlexibleGpuCatalog`

NewFlexibleGpuCatalogWithDefaults instantiates a new FlexibleGpuCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerations

`func (o *FlexibleGpuCatalog) GetGenerations() []string`

GetGenerations returns the Generations field if non-nil, zero value otherwise.

### GetGenerationsOk

`func (o *FlexibleGpuCatalog) GetGenerationsOk() (*[]string, bool)`

GetGenerationsOk returns a tuple with the Generations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerations

`func (o *FlexibleGpuCatalog) SetGenerations(v []string)`

SetGenerations sets Generations field to given value.

### HasGenerations

`func (o *FlexibleGpuCatalog) HasGenerations() bool`

HasGenerations returns a boolean if a field has been set.

### GetMaxCpu

`func (o *FlexibleGpuCatalog) GetMaxCpu() int32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *FlexibleGpuCatalog) GetMaxCpuOk() (*int32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *FlexibleGpuCatalog) SetMaxCpu(v int32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *FlexibleGpuCatalog) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxRam

`func (o *FlexibleGpuCatalog) GetMaxRam() int32`

GetMaxRam returns the MaxRam field if non-nil, zero value otherwise.

### GetMaxRamOk

`func (o *FlexibleGpuCatalog) GetMaxRamOk() (*int32, bool)`

GetMaxRamOk returns a tuple with the MaxRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRam

`func (o *FlexibleGpuCatalog) SetMaxRam(v int32)`

SetMaxRam sets MaxRam field to given value.

### HasMaxRam

`func (o *FlexibleGpuCatalog) HasMaxRam() bool`

HasMaxRam returns a boolean if a field has been set.

### GetModelName

`func (o *FlexibleGpuCatalog) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *FlexibleGpuCatalog) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *FlexibleGpuCatalog) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *FlexibleGpuCatalog) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetVRam

`func (o *FlexibleGpuCatalog) GetVRam() int32`

GetVRam returns the VRam field if non-nil, zero value otherwise.

### GetVRamOk

`func (o *FlexibleGpuCatalog) GetVRamOk() (*int32, bool)`

GetVRamOk returns a tuple with the VRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVRam

`func (o *FlexibleGpuCatalog) SetVRam(v int32)`

SetVRam sets VRam field to given value.

### HasVRam

`func (o *FlexibleGpuCatalog) HasVRam() bool`

HasVRam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


