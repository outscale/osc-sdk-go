# CreateVmTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCores** | **int32** | The number of vCores to use for each VM. | 
**CpuGeneration** | **string** | The processor generation to use for each VM (for example, &#x60;v4&#x60;). | 
**CpuPerformance** | Pointer to **string** | The performance of the VMs (&#x60;medium&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;).  | [optional] [default to "high"]
**Description** | Pointer to **string** | A description for the VM template. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | **string** | The ID of the OMI to use for each VM. You can find a list of OMIs by calling the [ReadImages](#readimages) method. | 
**KeypairName** | Pointer to **string** | The name of the keypair to use for each VM. | [optional] 
**Ram** | **int32** | The amount of RAM to use for each VM. | 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags to add to the VM template. | [optional] 
**VmTemplateName** | **string** | The name of the VM template. | 

## Methods

### NewCreateVmTemplateRequest

`func NewCreateVmTemplateRequest(cpuCores int32, cpuGeneration string, imageId string, ram int32, vmTemplateName string, ) *CreateVmTemplateRequest`

NewCreateVmTemplateRequest instantiates a new CreateVmTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVmTemplateRequestWithDefaults

`func NewCreateVmTemplateRequestWithDefaults() *CreateVmTemplateRequest`

NewCreateVmTemplateRequestWithDefaults instantiates a new CreateVmTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCores

`func (o *CreateVmTemplateRequest) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *CreateVmTemplateRequest) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *CreateVmTemplateRequest) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.


### GetCpuGeneration

`func (o *CreateVmTemplateRequest) GetCpuGeneration() string`

GetCpuGeneration returns the CpuGeneration field if non-nil, zero value otherwise.

### GetCpuGenerationOk

`func (o *CreateVmTemplateRequest) GetCpuGenerationOk() (*string, bool)`

GetCpuGenerationOk returns a tuple with the CpuGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuGeneration

`func (o *CreateVmTemplateRequest) SetCpuGeneration(v string)`

SetCpuGeneration sets CpuGeneration field to given value.


### GetCpuPerformance

`func (o *CreateVmTemplateRequest) GetCpuPerformance() string`

GetCpuPerformance returns the CpuPerformance field if non-nil, zero value otherwise.

### GetCpuPerformanceOk

`func (o *CreateVmTemplateRequest) GetCpuPerformanceOk() (*string, bool)`

GetCpuPerformanceOk returns a tuple with the CpuPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPerformance

`func (o *CreateVmTemplateRequest) SetCpuPerformance(v string)`

SetCpuPerformance sets CpuPerformance field to given value.

### HasCpuPerformance

`func (o *CreateVmTemplateRequest) HasCpuPerformance() bool`

HasCpuPerformance returns a boolean if a field has been set.

### GetDescription

`func (o *CreateVmTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVmTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVmTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVmTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateVmTemplateRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVmTemplateRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateVmTemplateRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateVmTemplateRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetImageId

`func (o *CreateVmTemplateRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateVmTemplateRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateVmTemplateRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetKeypairName

`func (o *CreateVmTemplateRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *CreateVmTemplateRequest) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *CreateVmTemplateRequest) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *CreateVmTemplateRequest) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetRam

`func (o *CreateVmTemplateRequest) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *CreateVmTemplateRequest) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *CreateVmTemplateRequest) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetTags

`func (o *CreateVmTemplateRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateVmTemplateRequest) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateVmTemplateRequest) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateVmTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmTemplateName

`func (o *CreateVmTemplateRequest) GetVmTemplateName() string`

GetVmTemplateName returns the VmTemplateName field if non-nil, zero value otherwise.

### GetVmTemplateNameOk

`func (o *CreateVmTemplateRequest) GetVmTemplateNameOk() (*string, bool)`

GetVmTemplateNameOk returns a tuple with the VmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateName

`func (o *CreateVmTemplateRequest) SetVmTemplateName(v string)`

SetVmTemplateName sets VmTemplateName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


