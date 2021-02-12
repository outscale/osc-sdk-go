# FlexibleGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If true, the fGPU is deleted when the VM is terminated. | [optional] 
**FlexibleGpuId** | Pointer to **string** | The ID of the fGPU. | [optional] 
**Generation** | Pointer to **string** | The compatible processor generation. | [optional] 
**ModelName** | Pointer to **string** | The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs). | [optional] 
**State** | Pointer to **string** | The state of the fGPU (&#x60;allocated&#x60; \\| &#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60;). | [optional] 
**SubregionName** | Pointer to **string** | The Subregion where the fGPU is located. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM the fGPU is attached to, if any. | [optional] 

## Methods

### NewFlexibleGpu

`func NewFlexibleGpu() *FlexibleGpu`

NewFlexibleGpu instantiates a new FlexibleGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexibleGpuWithDefaults

`func NewFlexibleGpuWithDefaults() *FlexibleGpu`

NewFlexibleGpuWithDefaults instantiates a new FlexibleGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnVmDeletion

`func (o *FlexibleGpu) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *FlexibleGpu) GetDeleteOnVmDeletionOk() (*bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnVmDeletion

`func (o *FlexibleGpu) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion sets DeleteOnVmDeletion field to given value.

### HasDeleteOnVmDeletion

`func (o *FlexibleGpu) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### GetFlexibleGpuId

`func (o *FlexibleGpu) GetFlexibleGpuId() string`

GetFlexibleGpuId returns the FlexibleGpuId field if non-nil, zero value otherwise.

### GetFlexibleGpuIdOk

`func (o *FlexibleGpu) GetFlexibleGpuIdOk() (*string, bool)`

GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleGpuId

`func (o *FlexibleGpu) SetFlexibleGpuId(v string)`

SetFlexibleGpuId sets FlexibleGpuId field to given value.

### HasFlexibleGpuId

`func (o *FlexibleGpu) HasFlexibleGpuId() bool`

HasFlexibleGpuId returns a boolean if a field has been set.

### GetGeneration

`func (o *FlexibleGpu) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *FlexibleGpu) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *FlexibleGpu) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *FlexibleGpu) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetModelName

`func (o *FlexibleGpu) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *FlexibleGpu) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *FlexibleGpu) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *FlexibleGpu) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetState

`func (o *FlexibleGpu) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlexibleGpu) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FlexibleGpu) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FlexibleGpu) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubregionName

`func (o *FlexibleGpu) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *FlexibleGpu) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *FlexibleGpu) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *FlexibleGpu) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### GetVmId

`func (o *FlexibleGpu) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *FlexibleGpu) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *FlexibleGpu) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *FlexibleGpu) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


