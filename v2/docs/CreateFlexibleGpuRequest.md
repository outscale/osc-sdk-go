# CreateFlexibleGpuRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If true, the fGPU is deleted when the VM is terminated. | [optional] [default to false]
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Generation** | Pointer to **string** | The processor generation that the fGPU must be compatible with. If not specified, the oldest possible processor generation is selected (as provided by [ReadFlexibleGpuCatalog](#readflexiblegpucatalog) for the specified model of fGPU). | [optional] 
**ModelName** | **string** | The model of fGPU you want to allocate. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html). | 
**SubregionName** | **string** | The Subregion in which you want to create the fGPU. | 

## Methods

### NewCreateFlexibleGpuRequest

`func NewCreateFlexibleGpuRequest(modelName string, subregionName string, ) *CreateFlexibleGpuRequest`

NewCreateFlexibleGpuRequest instantiates a new CreateFlexibleGpuRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFlexibleGpuRequestWithDefaults

`func NewCreateFlexibleGpuRequestWithDefaults() *CreateFlexibleGpuRequest`

NewCreateFlexibleGpuRequestWithDefaults instantiates a new CreateFlexibleGpuRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnVmDeletion

`func (o *CreateFlexibleGpuRequest) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *CreateFlexibleGpuRequest) GetDeleteOnVmDeletionOk() (*bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnVmDeletion

`func (o *CreateFlexibleGpuRequest) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion sets DeleteOnVmDeletion field to given value.

### HasDeleteOnVmDeletion

`func (o *CreateFlexibleGpuRequest) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateFlexibleGpuRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateFlexibleGpuRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateFlexibleGpuRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateFlexibleGpuRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetGeneration

`func (o *CreateFlexibleGpuRequest) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *CreateFlexibleGpuRequest) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *CreateFlexibleGpuRequest) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *CreateFlexibleGpuRequest) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetModelName

`func (o *CreateFlexibleGpuRequest) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *CreateFlexibleGpuRequest) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *CreateFlexibleGpuRequest) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetSubregionName

`func (o *CreateFlexibleGpuRequest) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *CreateFlexibleGpuRequest) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *CreateFlexibleGpuRequest) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


