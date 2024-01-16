# VmType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BsuOptimized** | Pointer to **bool** | This parameter is not available. It is present in our API for the sake of historical compatibility with AWS. | [optional] 
**EphemeralsType** | Pointer to **string** | The type of ephemeral storage disk. | [optional] 
**Eth** | Pointer to **int32** | The number of Ethernet interface available. | [optional] 
**Gpu** | Pointer to **int32** | The number of GPU available. | [optional] 
**MaxPrivateIps** | Pointer to **int32** | The maximum number of private IPs per network interface card (NIC). | [optional] 
**MemorySize** | Pointer to **float32** | The amount of memory, in gibibytes. | [optional] 
**VcoreCount** | Pointer to **int32** | The number of vCores. | [optional] 
**VmTypeName** | Pointer to **string** | The name of the VM type. | [optional] 
**VolumeCount** | Pointer to **int32** | The maximum number of ephemeral storage disks. | [optional] 
**VolumeSize** | Pointer to **int32** | The size of one ephemeral storage disk, in gibibytes (GiB). | [optional] 

## Methods

### NewVmType

`func NewVmType() *VmType`

NewVmType instantiates a new VmType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmTypeWithDefaults

`func NewVmTypeWithDefaults() *VmType`

NewVmTypeWithDefaults instantiates a new VmType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsuOptimized

`func (o *VmType) GetBsuOptimized() bool`

GetBsuOptimized returns the BsuOptimized field if non-nil, zero value otherwise.

### GetBsuOptimizedOk

`func (o *VmType) GetBsuOptimizedOk() (*bool, bool)`

GetBsuOptimizedOk returns a tuple with the BsuOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsuOptimized

`func (o *VmType) SetBsuOptimized(v bool)`

SetBsuOptimized sets BsuOptimized field to given value.

### HasBsuOptimized

`func (o *VmType) HasBsuOptimized() bool`

HasBsuOptimized returns a boolean if a field has been set.

### GetEphemeralsType

`func (o *VmType) GetEphemeralsType() string`

GetEphemeralsType returns the EphemeralsType field if non-nil, zero value otherwise.

### GetEphemeralsTypeOk

`func (o *VmType) GetEphemeralsTypeOk() (*string, bool)`

GetEphemeralsTypeOk returns a tuple with the EphemeralsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralsType

`func (o *VmType) SetEphemeralsType(v string)`

SetEphemeralsType sets EphemeralsType field to given value.

### HasEphemeralsType

`func (o *VmType) HasEphemeralsType() bool`

HasEphemeralsType returns a boolean if a field has been set.

### GetEth

`func (o *VmType) GetEth() int32`

GetEth returns the Eth field if non-nil, zero value otherwise.

### GetEthOk

`func (o *VmType) GetEthOk() (*int32, bool)`

GetEthOk returns a tuple with the Eth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEth

`func (o *VmType) SetEth(v int32)`

SetEth sets Eth field to given value.

### HasEth

`func (o *VmType) HasEth() bool`

HasEth returns a boolean if a field has been set.

### GetGpu

`func (o *VmType) GetGpu() int32`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *VmType) GetGpuOk() (*int32, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *VmType) SetGpu(v int32)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *VmType) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetMaxPrivateIps

`func (o *VmType) GetMaxPrivateIps() int32`

GetMaxPrivateIps returns the MaxPrivateIps field if non-nil, zero value otherwise.

### GetMaxPrivateIpsOk

`func (o *VmType) GetMaxPrivateIpsOk() (*int32, bool)`

GetMaxPrivateIpsOk returns a tuple with the MaxPrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrivateIps

`func (o *VmType) SetMaxPrivateIps(v int32)`

SetMaxPrivateIps sets MaxPrivateIps field to given value.

### HasMaxPrivateIps

`func (o *VmType) HasMaxPrivateIps() bool`

HasMaxPrivateIps returns a boolean if a field has been set.

### GetMemorySize

`func (o *VmType) GetMemorySize() float32`

GetMemorySize returns the MemorySize field if non-nil, zero value otherwise.

### GetMemorySizeOk

`func (o *VmType) GetMemorySizeOk() (*float32, bool)`

GetMemorySizeOk returns a tuple with the MemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySize

`func (o *VmType) SetMemorySize(v float32)`

SetMemorySize sets MemorySize field to given value.

### HasMemorySize

`func (o *VmType) HasMemorySize() bool`

HasMemorySize returns a boolean if a field has been set.

### GetVcoreCount

`func (o *VmType) GetVcoreCount() int32`

GetVcoreCount returns the VcoreCount field if non-nil, zero value otherwise.

### GetVcoreCountOk

`func (o *VmType) GetVcoreCountOk() (*int32, bool)`

GetVcoreCountOk returns a tuple with the VcoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcoreCount

`func (o *VmType) SetVcoreCount(v int32)`

SetVcoreCount sets VcoreCount field to given value.

### HasVcoreCount

`func (o *VmType) HasVcoreCount() bool`

HasVcoreCount returns a boolean if a field has been set.

### GetVmTypeName

`func (o *VmType) GetVmTypeName() string`

GetVmTypeName returns the VmTypeName field if non-nil, zero value otherwise.

### GetVmTypeNameOk

`func (o *VmType) GetVmTypeNameOk() (*string, bool)`

GetVmTypeNameOk returns a tuple with the VmTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTypeName

`func (o *VmType) SetVmTypeName(v string)`

SetVmTypeName sets VmTypeName field to given value.

### HasVmTypeName

`func (o *VmType) HasVmTypeName() bool`

HasVmTypeName returns a boolean if a field has been set.

### GetVolumeCount

`func (o *VmType) GetVolumeCount() int32`

GetVolumeCount returns the VolumeCount field if non-nil, zero value otherwise.

### GetVolumeCountOk

`func (o *VmType) GetVolumeCountOk() (*int32, bool)`

GetVolumeCountOk returns a tuple with the VolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCount

`func (o *VmType) SetVolumeCount(v int32)`

SetVolumeCount sets VolumeCount field to given value.

### HasVolumeCount

`func (o *VmType) HasVolumeCount() bool`

HasVolumeCount returns a boolean if a field has been set.

### GetVolumeSize

`func (o *VmType) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *VmType) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *VmType) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *VmType) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


